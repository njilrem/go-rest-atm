package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/njilrem/go-rest-atm/controllers"
  middleware "github.com/njilrem/go-rest-atm/middlewares"
)

// SetupRouter ... setting up router
func SetupRouter() *gin.Engine {
  r := gin.Default()
  accountGroup := r.Group("/accounts-api")
  accountGroup.Use(middleware.AuthorizeJWT())
  {
    accountGroup.GET("account", controllers.GetAccounts)
    accountGroup.POST("account", controllers.CreateAccount)
    accountGroup.GET("account/:id", controllers.GetAccountByID)
    accountGroup.PUT("account/:id", controllers.UpdateAccount)
    accountGroup.DELETE("account/:id", controllers.DeleteAccount)
  }
  cardGroup := r.Group("/cards-api")
  cardGroup.Use(middleware.AuthorizeJWT())
  {
    cardGroup.GET("card/:id", controllers.GetCardsByHolderID)
    cardGroup.PUT("card/:id", controllers.UpdateCard)
    cardGroup.POST("card", controllers.CreateCard)
    cardGroup.GET("balance/:id", controllers.GetBalance)
  }
  authGroup := r.Group("/auth")
  {
    authGroup.POST("login", controllers.AuthAccount)
    authGroup.POST("auth", controllers.AuthAdmin)
    authGroup.POST("pin", controllers.AuthAccountPIN)
  }
  transactionGroup := r.Group("/transactions-api")
  transactionGroup.Use(middleware.AuthorizeTransactionJWT())
  {
    transactionGroup.GET("transaction/:trId", controllers.GetTransactionById)

    transactionGroup.GET("transactions/:id", controllers.GetTransactionsById)
    transactionGroup.POST("transaction", controllers.CreateTransaction)
    transactionGroup.POST("refill", controllers.Refill)
  }
  return r
}

