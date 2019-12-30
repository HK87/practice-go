package infrastructure

import (
    "net/http"

	"github.com/gorilla/mux"
	"interfaces/controllers"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	
	// REST Path 設定
	router.HandleFunc("/activities/{userId}", getActivities)
	
	// 起動
    http.ListenAndServe(":8080", router)
}

func getActivities(w http.ResponseWriter, r *http.Request) {
	controller := controllers.NewActivityController(NewElasticsearchHandler())
	controller.Get(w, mux.Vars(r))
}