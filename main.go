package main

import (
	_ "github.com/lib/pq"
	"movie_review_apis/cmd"
)

// Main ...
func main() {
	cmd.Execute()
}
