package routes

import (
	"ApiRestGolang/controllers"
	"ApiRestGolang/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadesPorId).Methods("Get")
	r.HandleFunc("/api/register", controllers.RegisterPersonalidade).Methods("Post")
	r.HandleFunc("/api/edit/{id}", controllers.UpdatePersonalidade).Methods("Put")
	r.HandleFunc("/api/delete/{id}", controllers.DeletePersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
