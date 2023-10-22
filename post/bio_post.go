package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
    Id     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserId int    `json:"userId"`
}


func Postuser() {
	url := "https://jsonplaceholder.typicode.com/posts"

	body := []byte(`{
        "title": "Post title",''
        "body": "Post description",
        "userId": 1
    }`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err :=  client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	post := &Post{}
	decode := json.NewDecoder(res.Body).Decode(post)
	if decode != nil {
		panic(res.Status)
	}

	fmt.Println("Id:", post.Id)
    fmt.Println("Title:", post.Title)
    fmt.Println("Body:", post.Body)
    fmt.Println("UserId:", post.UserId)

}