# initiating go project
```bash
go mod init web-test 
```

# starting server
```bash
go run ./cmd/api
```

# testing healthcheck
```bash
curl localhost:4000/v1/healthcheck
```

# testing healthcheck correct method
```bash
curl -X POST localhost:4000/v1/healthcheck
```

# testing books api
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


# preparing postgreSQL on local machine
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

# cleanup docker
```bash
docker stop web-hello
docker rm web-hello
docker image rm -f postgres
```

# install psql client(ubuntu)
```bash
sudo apt-get install -y postgresql-client
```

```postgreSQL
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
