package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Article is...
type Article struct {
	ArticleID string `json:"ArticleID"`
	Title     string `json:"Title"`
	SubTitle  string `json:"SubTitle"`
	Content   string `json:"content"`
	TimeStamp string `json:"TimeStamp"`
}

//Articles is...
var Articles []Article

//list all articles
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// Get an article using id
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ArticleID"]

	for _, article := range Articles {
		if article.ArticleID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

//Search for an article
func searchQuery(w http.ResponseWriter, r *http.Request) {
	paramTitle := r.URL.Query().Get("Title")
	if paramTitle != "" {
		vars := mux.Vars(r)
		key := vars["Title"]

		for _, article := range Articles {
			if article.Title == key {
				json.NewEncoder(w).Encode(article)
			}
		}
	}

	paramSubTitle := r.URL.Query().Get("SubTitle")
	if paramSubTitle != "" {
		vars := mux.Vars(r)
		key := vars["SubTitle"]

		for _, article := range Articles {
			if article.SubTitle == key {
				json.NewEncoder(w).Encode(article)
			}
		}
	}

	paramContent := r.URL.Query().Get("Content")
	if paramContent != "" {
		vars := mux.Vars(r)
		key := vars["Content"]

		for _, article := range Articles {
			if article.Content == key {
				json.NewEncoder(w).Encode(article)
			}
		}
	}
}

//Create an article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

//Home Page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{ArticleID}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles/search", returnAllArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	t := time.Now()
	Articles = []Article{
		Article{ArticleID: "1", Title: "Vodafone Network down", SubTitle: "Vodafone Idea network down in parts of Maharashtra as key site in Pune flooded", Content: "Vodafone Idea subscribers in Maharashtra are facing connectivity issues since Thursday morning as a key site in Pune got flooded following heavy rains. The company said a key site in Pune's Kalyani Nagar area got flooded and that its technical team is working to fully restore the services soon. Several users took to social media to complain of network disruption.", TimeStamp: t.Format("2006-01-02 15:04:05")},
		Article{ArticleID: "2", Title: "Red Light On, Gaadi Off", SubTitle: "Delhi CM launches 'Red Light On, Gaadi Off' campaign to curb pollution", Content: "Delhi CM Arvind Kejriwal has launched the 'Red Light On, Gaadi Off' campaign to tackle air pollution in the national capital. Kejriwal stated, 'There are one crore vehicles registered in Delhi. According to experts, even if 10 lakh vehicles turn off ignition at traffic signals, then 1.5 tonnes of PM10 will reduce in a year.'", TimeStamp: t.Format("2006-01-02 15:04:05")},
	}
	handleRequests()
}
