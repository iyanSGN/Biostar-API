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

// package main

// import (
// 	"main/routes"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()

// 	routes.Init(e)

// 	e.Logger.Fatal(e.Start(":8700"))

// }

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type Todo struct {
// 	UserID    int    `json:"userId"`
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// func main() {
// 	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
// 	jsonReq, err := json.Marshal(todo)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 	}

// 	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
// 	if err != nil {
// 		fmt.Println("error req: ", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)

// 	// Convert response body to Todo struct
// 	var todoStruct Todo
// 	json.Unmarshal(bodyBytes, &todoStruct)
// 	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
// }

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

  url := "https://192.168.88.79:8443/api/users?id=876"
  method := "DELETE"

  transport := &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

	client := &http.Client{Transport: transport}

  req, err := http.NewRequest(method, url, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  	req.Header.Set("Content-Type", "application/json") // Set the content type
	req.Header.Set("bs-session-id", "ee5117a427bb4f7f9afc135d7d5beb75")

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