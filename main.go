package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sturct add json behind to enable conversion to json
type account struct {
	ID    string `json:"id"`
	Name  string `json:"title"`
	Email string `json:"email"`
	User  int    `json:"quantity"`
}

// Slice of account
var accounts = []account{
	{ID: "1", Name: "Apex Lock", Email: "apex101@gmail.com", User: 15},
	{ID: "2", Name: "Best shop", Email: "shop102@gmail.com", User: 10},
	{ID: "3", Name: "Cashless", Email: "cashless103@gmail.com", User: 25},
}

func getAccounts(c *gin.Context) {
	// IndentedJSON convert accounts to json and return status ok
	c.IndentedJSON(http.StatusOK, accounts)
}

func createAccount(c *gin.Context) {
	var newAccount account

	// Binding Input Json into newAccount pointer, check if there is error return
	if err := c.BindJSON(&newAccount); err != nil {
		// If Any Error, BindJSON will return fail status code
		return
	}

	accounts = append(accounts, newAccount)
	c.IndentedJSON(http.StatusCreated, newAccount)

}

func main() {
	var pop string
	fmt.Println(&pop)

	router := gin.Default()
	router.GET("/accounts", getAccounts)
	router.POST("/accounts", createAccount)
	router.Run("localhost:8080")
}
