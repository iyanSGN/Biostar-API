package post

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {	
	UserID     		string 			`json:"user_id"`
	UserGroupID    struct {
		ID string `json:"id"`
	} `json:"user_group_id"`
	StartDatetime  	time.Time 		`json:"start_datetime"`
	ExpiryDatetime  time.Time 		`json:"expiry_datetime"`
	Name            string 			`json:"name"`
	Email           string 			`json:"email"`
}

type PostUsers struct {
	UserID     		string 			`json:"user_id"`
	UserGroupID    struct {
		ID string `json:"id"`
	} `json:"user_group_id"`
	StartDatetime   time.Time 		`json:"start_datetime"`
	ExpiryDatetime  time.Time 		`json:"expiry_datetime"`
	Name            string 			`json:"name"`
	Email           string 			`json:"email"`
}

// type Post struct {
//     Id     int    `json:"id"`
//     Title  string `json:"title"`
//     Body   string `json:"body"`
//     UserId int    `json:"userId"`
// }

// type PostData struct {
//     Title  string `json:"title"`
//     Body   string `json:"body"`
//     UserId int    `json:"userId"`
// }



func PostUser(post User) (User, error) {
	url := "https://192.168.88.79:8443/api/users" // Replace with your actual API endpoint
	method := "POST"
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	// Create a JSON payload using a struct and marshaling it
	payload := PostUsers{
		UserID: post.UserID,
		UserGroupID: post.UserGroupID,
		Name: post.Name,
		Email: post.Email,
		}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(payloadBytes))
	
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type
	req.Header.Set("bs-session-id", "5e25ef700af24769ae58895db6b1cd04")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Println("error req: %w", err))
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(fmt.Println("error body: %w", err))
	}

	fmt.Println("Response Body:", string(body))

	var users User
	if err := json.Unmarshal(body, &users);
	err != nil {
		fmt.Println("error decoding JSON: ", err.Error())
		return post, nil
	}

	return post, nil
}


func HandleUser(c echo.Context) error {
	var request User

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	createdUser, err := PostUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"status_code" : http.StatusOK,
		"data" : map[string]interface{}{
			"message" :  "register Successfull",
			"Name" : createdUser.Name,
		},
	}

	return c.JSON(http.StatusOK, response)
}