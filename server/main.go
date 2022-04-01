package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// http.HandleFunc("/api/nft", get_nfts)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/news", get_news)
	mux.HandleFunc("/", running)
	mux.HandleFunc("/api/topmonth", top_nft_this_month)
	mux.HandleFunc("/api/artist", get_nft_artist)
	handler := cors.Default().Handler(mux)
	fmt.Printf("Server running at \n http://localhost:4000/api/ \n http://localhost:4000/api/artist \n http://localhost:4000/api/news")
	http.ListenAndServe(":4000", handler)
}
