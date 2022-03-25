package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/api/nft", get_nfts)
	http.HandleFunc("/api/news", get_news)
	http.HandleFunc("/", running)
	http.HandleFunc("/api/topmonth", top_nft_this_month)
	http.HandleFunc("/api/artist", get_nft_artist)

	fmt.Printf("Server running at http://localhost:8888/api/topmonth ")
	http.ListenAndServe(":8888", nil)
}
