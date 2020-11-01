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
  accountGroup.Use(middleware.AuthorizeAdminJWT())
  {
    accountGroup.GET("account", controllers.GetAccounts)
    accountGroup.POST("account", controllers.CreateAccount)
    accountGroup.GET("account/:id", controllers.GetAccountByID)
    accountGroup.PUT("account/:id", controllers.UpdateAccount)
    accountGroup.DELETE("account/:id", controllers.DeleteAccount)
  }
  cardGroup := r.Group("/cards-api")
  cardGroup.Use(middleware.AuthorizeAdminJWT())
  {
    cardGroup.GET("card/:id", controllers.GetCardsByHolderID)
    cardGroup.PUT("card/:id", controllers.UpdateCard)
    cardGroup.POST("card", controllers.CreateCard)
  }
  authGroup := r.Group("/auth")
  {
    authGroup.POST("pay", controllers.AuthAccount)
  }
  return r
}

