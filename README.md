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
curl -X PUT localhost:4000/v1/books/125

#delete item
curl -X DELETE localhost:4000/v1/books/125
```
