package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Company struct {
	Name		string	`json:"name"`
	CatchPrase	string	`json:"catchPrase"`
	Bs			string	`json:"bs"`	
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type User struct {
	ID     int    	`json:"id"`
	Name   string 	`json:"name"`
	Email  string 	`json:"email"`
	Address Address `json:"address"`
}

type getUser struct {
	ID     int    	`json:"id"`
	Name   string 	`json:"name"`
	Email  string 	`json:"email"`
	Address Address `json:"address"`
}

func GetUser() ([]getUser, error) {
    apiUrl := "https://jsonplaceholder.typicode.com/users"

    response, err := http.Get(apiUrl)
    if err != nil {
        fmt.Println("Error making the request:", err)
    }
    defer response.Body.Close()

    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading the response:", err)
    }

    var users []User
    if err := json.Unmarshal(responseBody, &users); err != nil {
        fmt.Println("Error decoding JSON:", err)
    }

	getusers := make([]getUser, len(users))
    for i, user := range users {
		getusers[i] = getUser(user)
	}

    return getusers, nil
}

func HandleUser(c echo.Context) error {
	userID, err := GetUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data" : userID,
		"status_Code" : http.StatusOK,
	})
}