package routes

import (
	"ALURA/jcqueiroz2/api-go-rest/controllers"
	"ALURA/jcqueiroz2/api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnOnePersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteOnePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
