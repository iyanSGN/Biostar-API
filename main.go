// package main

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "net/http"
// )

// type Post struct {
//     Id     int    `json:"id"`
//     Title  string `json:"title"`
//     Body   string `json:"body"`
//     UserId int    `json:"userId"`
// }

// func main() {
//     url := "https://jsonplaceholder.typicode.com/posts"

//     post := Post{
//         Title:  "Post title",
//         Body:   "Post description",
//         UserId: 1,
//     }

//     jsonBody, err := json.Marshal(post)
//     if err != nil {
//         panic(err)
//     }

//     req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
//     if err != nil {
//         panic(err)
//     }

//     req.Header.Add("Content-Type", "application/json")

//     client := &http.Client{}

//     res, err := client.Do(req)
//     if err != nil {
//         panic(err)
//     }

//     defer res.Body.Close()

//     if res.StatusCode != http.StatusCreated {
//         fmt.Println("Failed to create a post. Status code:", res.Status)
//         return
//     }

//     createdPost := &Post{}
//     decode := json.NewDecoder(res.Body).Decode(createdPost)
//     if decode != nil {
//         panic(decode)
//     }

//	    fmt.Println("Id:", createdPost.Id)
//	    fmt.Println("Title:", createdPost.Title)
//	    fmt.Println("Body:", createdPost.Body)
//	    fmt.Println("UserId:", createdPost.UserId)
//	}
package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://192.168.88.79:8443/api/users" // Replace with your actual API endpoint
	method := "POST"

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	// Create a JSON payload using a struct and marshaling it
	payload := []byte(`{
		"User": {
			"user_id": "820",
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

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
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
