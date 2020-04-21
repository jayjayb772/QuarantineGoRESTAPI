package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type BlogPost struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content`
}

type Posts []BlogPost

func allPosts(w http.ResponseWriter, r *http.Request){
	posts := Posts{
		BlogPost{Title:"Test",Desc:"Test",Content:"Hello World"},
	}

	fmt.Println("endpoint hit: All posts endpoint")
	json.NewEncoder(w).Encode(posts)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allposts",allPosts)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
}
