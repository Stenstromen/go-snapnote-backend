module github.com/stenstromen/go-snapnote-backend

go 1.20

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gorilla/mux v1.8.0
	github.com/rs/cors v1.9.0
)

replace github.com/stenstromen/go-snapnote-backend => /
replace github.com/stenstromen/go-snapnote-backend/models => /models
replace github.com/stenstromen/go-snapnote-backend/controller => /controller
replace github.com/stenstromen/go-snapnote-backend/service => /service