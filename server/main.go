package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func running(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v", "server")
}

func get_news(w http.ResponseWriter, req *http.Request) {
	eror := godotenv.Load("./.env")
	if eror != nil {
		log.Fatal("Error loading .env file")
	}

	rapidApiKey := os.Getenv("RAPID_API_KEY")

	url := "https://crypto-news-live3.p.rapidapi.com/news"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("x-rapidapi-host", "crypto-news-live3.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", rapidApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Fprintf(w, "%v", res)
	fmt.Fprintf(w, string(body))
}

func main() {
	fmt.Printf("Server running at http://localhost:8888/api ")
	http.HandleFunc("/api", get_news)
	http.ListenAndServe(":8888", nil)
}
