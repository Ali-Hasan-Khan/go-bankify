package api

import (
	"github.com/Ali-Hasan-Khan/go-bankify/api/handlers"
	db "github.com/Ali-Hasan-Khan/go-bankify/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store          *db.Store
	router         *gin.Engine
	accountHandler *handlers.AccountHandler
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store, accountHandler: handlers.NewAccountHandler(store)}

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/accounts", server.accountHandler.CreateAccount)
		v1.GET("/accounts/:id", server.accountHandler.GetAccount)
		v1.GET("/accounts", server.accountHandler.ListAccounts)
		v1.PUT("/accounts/:id", server.accountHandler.UpdateAccount)
	}

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
