package handlers

import (
	"database/sql"
)

type Server struct {
	DB    *sql.DB
	Token string
}

func NewServer(db *sql.DB, token string) *Server {
	return &Server{DB: db, Token: token}
}
