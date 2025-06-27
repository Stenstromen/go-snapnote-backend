# Go-Snapnote-Backend

Backend server for Snapnote

```bash
DB_HOST:            String
DB_USERNAME:        String
DB_PASSWORD:        String
DB_DATABASE:        String
AUTHORIZATION:      String
ALLOWED_ORIGINS     String (Comma separated)
```

## Docker

### Run

```bash
docker run -d --rm \
--name snapnote \
-p 8080:8080 \
-e DB_HOST="STRING" \
-e DB_USERNAME="STRING" \
-e DB_PASSWORD="STRING" \
-e DB_DATABASE="STRING" \
-e AUTHORIZATION="STRING" \
-e ALLOWED_ORIGINS="http://localhost:8080,http://localhost:5173"
ghcr.io/stenstromen/go-snapnote-backend:latest
```
