package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/example/federation/accounts/graph"
	"github.com/99designs/gqlgen/example/federation/products/graph/generated"
	"github.com/99designs/gqlgen/example/federation/reviews/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"honnef.co/go/tools/analysis/facts/generated"
)

func main() {
	// http.HandleFunc("/api/nft", get_nfts)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
		log.Printf("defaulting to port %s", port)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()
	mux.HandleFunc("/api/news", get_news)
	mux.HandleFunc("/", running)
	mux.HandleFunc("/api/topmonth", top_nft_this_month)
	mux.HandleFunc("/api/artist", get_nft_artist)
	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", srv)
	handler := cors.Default().Handler(mux)
	fmt.Printf("Server running at \n http://localhost:4000/api/ \n http://localhost:4000/api/artist \n http://localhost:4000/api/news")
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
