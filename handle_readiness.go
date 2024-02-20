package main

import (
	"net/http"
)

func handleReadines( w http.ResponseWriter , r *http.Request ){
	respondWithJson(w,200, struct{}{})

}