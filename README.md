# Preparation 
## initiating go project
```bash
go mod init web-test 
```

## install external database dependency
```bash
go get github.com/lib/pq
```

# Running
## starting server
```bash
#run local db server
docker start web-hello

#export db connection
export WEBHELLO_DB_DSN="postgres://webhello:pass@localhost/webhello?sslmode=disable"

#start web server
go run ./cmd/api
```

# Tests
## testing healthcheck
```bash
curl localhost:4000/v1/healthcheck
```

## testing healthcheck correct method
```bash
curl -X POST localhost:4000/v1/healthcheck
```

## testing books api
```bash
#get collection
curl localhost:4000/v1/books

#add new item
BODY='{"title":"The Black Soulstone","published":2001,"pages":107,"genres":["Fiction","Mystery"],"rating":3.5}'
curl -i -d "$BODY" -X POST localhost:4000/v1/books

#get item
curl localhost:4000/v1/books/125

#update item
BODY='{"title":"The Black Soulstone","published":2001,"pages":107,"genres":["Fiction","Mystery"],"rating":3.5}'
curl -i -d "$BODY" -X PUT localhost:4000/v1/books/125

#delete item
curl -X DELETE localhost:4000/v1/books/125
```


# Database
## preparing postgreSQL on local machine
```bash
#pull latest pastge image
docker pull postgres

#check if imgage is already there
docker images | grep postgres

#variable must be defined or db will not start
docker run --name web-hello -e POSTGRES_PASSWORD=mylocalpass -d -p 5432:5432 postgres

#check running proces
docker ps

psql -h localhost -p 5432 -U postgres
```

## cleanup docker
```bash
docker stop web-hello
docker rm web-hello
docker image rm -f postgres
```

## install psql client(ubuntu)
```bash
sudo apt-get install -y postgresql-client
```

## prepare db structure
```sql
CREATE DATABASE webhello;
CREATE ROLE webhello WITH LOGIN PASSWORD 'pass';
\c webhello


CREATE TABLE IF NOT EXISTS books (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  title text NOT NULL,
  published integer NOT NULL,
  pages integer NOT NULL,
  genres text[] NOT NULL,
  rating real NOT NULL,
  version integer NOT NULL DEFAULT 1
);

GRANT SELECT, INSERT, UPDATE, DELETE ON books TO webhello;
GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO webhello;
```

# Go Remote Debugging
## installing delve debugger for go
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

## start remote session
```bash
# wait for debugger
dlv debug ./cmd/api --headless --listen=:12345 --api-version=2 --accept-multiclient

# start the procedure --continue flag
dlv debug ./cmd/api --headless --listen=:12345 --api-version=2 --accept-multiclient --continue
```

## connecting vscode to remote session 
notes for vscode
[documentation vscode-go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging)
launch.json in folder ./.vscode/launch.json for attach to the session
```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Connect to external session",
      "type": "go",
      "debugAdapter": "dlv-dap",
      "request": "attach",
      "mode": "remote",
      "port": 12345
    }
  ]
}
```

## connecting neovim to remote session with nvim-dap
[documentation nvim-dap](https://github.com/mfussenegger/nvim-dap/wiki/Debug-Adapter-installation#go-using-delve-directly)
```lua
require("dap").adapters.delve = {
	type = "server",
	host = "127.0.0.1",
	port = "12345",
}

require("dap").configurations.go = {
	{
		type = "delve",
		name = "Attach remote v2",
		mode = "remote",
		request = "attach",
	},
}
```
