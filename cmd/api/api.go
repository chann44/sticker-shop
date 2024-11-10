package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/chann44/go-shop/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	userRouter := user.NewHandler(userStore)
	userRouter.RegisterRoutes(subRouter)

	log.Println("Listening on port ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
