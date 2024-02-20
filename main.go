package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	portString:= os.Getenv("PORT")
	
	router:= chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		 // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
		
	}) )

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz",handleReadines)
	
	v1Router.HandleFunc("/err",handleErr)
	router.Mount("/v1",v1Router)


	srv:= &http.Server{
		
		Addr: ":"+portString,
		Handler: router,
	
	}
	
	fmt.Printf("Server starting at port : %v",portString)
	err:= srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

	
	

}
