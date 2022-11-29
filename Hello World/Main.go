package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Posts    []Post
	Comments []Comment
	Profile  Profile
}
type Post struct {
	Id    int
	Title string
}

type Profile struct {
	Name string
}
type Comment struct {
	Id     int
	Body   string
	PostId int
}

func getData() {
	resp, err := http.Get("https://my-json-server.typicode.com/typicode/demo/db")
	if err != nil {
		print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	var response Response
	json.Unmarshal(body, &response)
	println("Name Of Profile:", response.Profile.Name+"\n")
	println("Response Part")

	for _, s := range response.Posts {
		println("Id ", s.Id, "has Title: ", s.Title)
	}
	println("Comment Part")
	for _, m := range response.Comments {
		println("id:", m.Id, ", body: ", m.Body, ", PostId: ", m.PostId)
	}
}

func main() {
	getData()
}
