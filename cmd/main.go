package main

import (
	"fmt"
	"log"
	"net/http"

	"fmt"
	"log"
	"net/http"

	repository "github.com/SahilBheke25/ResourceSharingApplication/internal/Repository"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/login"
	User "github.com/SahilBheke25/ResourceSharingApplication/internal/app/user"
	_ "github.com/lib/pq"

	Equipment "github.com/SahilBheke25/ResourceSharingApplication/internal/app/equipment"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/login"

	repository "github.com/SahilBheke25/ResourceSharingApplication/internal/repository"

	_ "github.com/lib/pq"
)

func main() {

	// Creating DB connection
	repository.InitializeDatabase()
	defer repository.DB.Close()

	mux := http.DefaultServeMux
	mux.HandleFunc("POST /login", login.Verify)
	mux.HandleFunc("POST /register", login.Register)
	mux.HandleFunc("POST /user", User.GetUserByIdHandler)

	fmt.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

	// Creating DB connection
	repository.InitializeDatabase()
	defer repository.DB.Close()

	mux := http.DefaultServeMux
	mux.HandleFunc("POST /login", login.Verify)
	mux.HandleFunc("POST /register", login.Register)
	mux.HandleFunc("POST /equipment", Equipment.PostLendEquipmentHandler)
	mux.HandleFunc("GET /equipments", Equipment.GetAllEquipmentHandler)

	fmt.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
