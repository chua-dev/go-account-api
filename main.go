package main

import (
	"errors"
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

func accountById(c *gin.Context) {
	id := c.Param("id")
	account, error := getAccountById(id)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, account)
}

// Helper Function
func getAccountById(id string) (*account, error) {
	for index, account := range accounts {
		if account.ID == id {
			return &accounts[index], nil
		}
	}

	return nil, errors.New("account not found")
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

func addMember(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id param from query"})
		return
	}

	account, err := getAccountById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		return
	}

	if account.User >= 50 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Account has maximum user"})
		return
	}
	account.User += 1
	c.IndentedJSON(http.StatusOK, account)
	return
}

func removeMember(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing Id Param"})
		return
	}

	account, err := getAccountById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		return
	}

	if account.User <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No more use for this account"})
		return
	}

	account.User -= 1
	c.IndentedJSON(http.StatusOK, account)
}

func main() {
	var pop string
	fmt.Println(&pop)

	router := gin.Default()
	router.GET("/accounts", getAccounts)
	router.GET("/accounts/:id", accountById) // Path Parameters
	router.POST("/accounts", createAccount)
	router.PATCH("/addmember", addMember)
	router.PATCH("/removemember", removeMember)
	router.Run("localhost:8080")
}
