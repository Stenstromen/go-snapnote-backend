module github.com/stenstromen/go-snapnote-backend

go 1.24.0

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/rs/cors v1.11.1
)

require filippo.io/edwards25519 v1.1.0 // indirect

replace github.com/stenstromen/go-snapnote-backend => /

replace github.com/stenstromen/go-snapnote-backend/models => /models

replace github.com/stenstromen/go-snapnote-backend/controller => /controller

replace github.com/stenstromen/go-snapnote-backend/service => /service
