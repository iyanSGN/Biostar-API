package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

func PostUser() {
	url := "https://example.com/api/users" // Replace with your actual API endpoint
	method := "POST"

	// Create a JSON payload using a struct and marshaling it
	payload := []byte(`{
		"User": {
			"user_group_id": {
				"id": "1"
			},
			"start_datetime": "2001-01-01T00:00:00.00Z",
			"expiry_datetime": "2030-12-31T23:59:00.00Z",
			"name": "Rita",
			"email": "Makayla15@yahoo.com",
			"password": "P96E2a16QMXTZIY",
			"user_ip": "72.162.118.63"
		}
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type
	req.Header.Set("bs-session-id", "0ee8b546c303427794d7921d0193c1ba")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
