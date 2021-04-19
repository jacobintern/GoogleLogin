package controllers

import (
	"fmt"
	"net/http"
	"time"
)

// HomePage is home page
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello jacob"+time.Now().String())
}
