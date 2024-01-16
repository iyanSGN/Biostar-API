package get

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Rows struct {
	Name             string `json:"name"`
	Gender			string	`json:"gender"`
	LoginID			string	`json:"login_id"`
	
}

type UserCollection struct {
	Total		string	`json:"total"`
	Rows		[]Rows	`json:"rows"`
}


type APIResponse struct {
	UserCollection UserCollection `json:"UserCollection"`
}

type getApiResponse struct {
	UserCollection UserCollection `json:"UserCollection"`
}

func GetUser() ([]getApiResponse, error) {
	apiUrl := "https://192.168.88.79:8443/api/users"
	
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// req.Header.Set("Host", "192.168.88.79:8443")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("bs-session-id", "440ac596d7f743d2aac948e031bf3e7e")
	// req.Header.Set("Cookie", "bs-session-id=8a2efee39aa44023b89e03b48ab107e0")

	res,err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return nil, err
	}

	var users APIResponse
	if err := json.Unmarshal(body, &users);
	err != nil {
		fmt.Println("error decoding JSON:", err)
		return nil, err
	}

	getusers := []getApiResponse{{ UserCollection: users.UserCollection }}

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