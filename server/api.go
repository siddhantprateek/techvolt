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
	fmt.Fprintf(w, "%v", "server is running")
}

type Reponse struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Page_Size int `json:"page_size"`
	Result    []struct {
		TokenId             int    `json:"token_id"`
		TokenAddress        string `json:"token_address"`
		TokenUri            string `json:"token_uri"`
		MetaData            string `json:"metadata"`
		Contract_Type       string `json:"contract_Type"`
		Token_hash          string `json:"token_hash"`
		Minter_address      string `json:"minter_address"`
		Block_number_minted int    `json:"block_number_minted"`
		Transaction_minted  string `json:"transaction_minted"`
		Synced_at           string `json:"synced_at"`
		Created_at          string `json:"created_at"`
	}
}

func get_nfts(w http.ResponseWriter, req *http.Request) {
	eror := godotenv.Load("./.env")
	if eror != nil {
		log.Fatal("Error loading .env file")
	}
	url := "https://nft-explorer.p.rapidapi.com/search?chain=eth&filter=name&offset=0&q=ape"
	rapidApiKey := os.Getenv("RAPID_API_KEY")

	req, _err := http.NewRequest("GET", url, nil)

	if _err != nil {
		log.Fatalln(_err)
	}

	req.Header.Add("x-rapidapi-host", "nft-explorer.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", rapidApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Fprintf(w, "%v", res)
	fmt.Fprintf(w, string(body))
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

	fmt.Fprintf(w, string(body))
}

func top_nft_this_month(w http.ResponseWriter, req *http.Request) {
	eror := godotenv.Load("./.env")
	if eror != nil {
		log.Fatal("Error Loading .env file")
	}
	rapidApiKey := os.Getenv("RAPID_API_KEY")
	url := "https://top-nft-sales.p.rapidapi.com/sales/30d"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("X-RapidAPI-Host", "top-nft-sales.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", rapidApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func get_nft_artist(w http.ResponseWriter, req *http.Request) {
	eror := godotenv.Load("./.env")
	if eror != nil {
		log.Fatal("Error Loading .env file")
	}

	rapidApiKey := os.Getenv("RAPID_API_KEY")
	url := "https://niftygateway-data-scraper.p.rapidapi.com/artists"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error requesting")
	}
	req.Header.Add("X-RapidAPI-Host", "niftygateway-data-scraper.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", rapidApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, string(body))

}
