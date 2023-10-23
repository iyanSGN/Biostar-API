// package main

// import (
// 	"bytes"
// 	"crypto/tls"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	url := "https://192.168.88.79:8443/api/users" // Replace with your actual API endpoint
// 	method := "POST"

// 	transport := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}

// 	client := &http.Client{Transport: transport}

// 	// Create a JSON payload using a struct and marshaling it
// 	payload := []byte(`{
// 		"User": {
// 			"user_id": "820",
// 			"user_group_id": {
// 				"id": "1"
// 			},
// 			"start_datetime": "2001-01-01T00:00:00.00Z",
// 			"expiry_datetime": "2030-12-31T23:59:00.00Z",
// 			"name": "Rita",
// 			"email": "Makayla15@yahoo.com",
// 			"password": "P96E2a16QMXTZIY",
// 			"user_ip": "72.162.118.63"
// 		}
// 	}`)

// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("bs-session-id", "5e25ef700af24769ae58895db6b1cd04")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 		body, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println(string(body))
// 	}
package main

import (
	"main/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8700"))

}