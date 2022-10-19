package views

import (
	"fmt"
	"net/http"
)

// HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Movie Review Website!")
}
