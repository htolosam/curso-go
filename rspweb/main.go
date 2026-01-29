package main

import (
	"fmt"
	"log"
	"net/http"
	"rspweb/handlers"
)

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static/"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	fmt.Println("Servidor corriendo en el puerto", port)
	log.Fatal(http.ListenAndServe(port, router))
}
