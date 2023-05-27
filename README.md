# Go-Snapnote-Backend

Backend server for Snapnote

```
DB_HOST:            String
DB_USERNAME:        String
DB_PASSWORD:        String
DB_DATABASE:        String
AUTHORIZATION:      String
PAXSSWORD:          String
```

## Docker

### Build
```
docker build -t go-snapnote-backend .
```

### Run
```
docker run -d --rm \
--name snapnote \
-p 8080:8080 \
-e DB_HOST="STRING" \
-e DB_USERNAME="STRING" \
-e DB_PASSWORD="STRING" \
-e DB_DATABASE="STRING" \
-e AUTHORIZATION="STRING" \
ghcr.io/stenstromen/go-snapnote-backend:latest
```