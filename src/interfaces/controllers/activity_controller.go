package controllers

import (
    "encoding/json"
    "interfaces/repository"
    "net/http"
    "strconv"
    "usecase"
)

type ActivityController struct {
    Interactor usecase.ActivityInteractor
}

func NewActivityController(handler repository.LogSearchEngineHandler) *ActivityController {
    return &ActivityController{
        Interactor: usecase.ActivityInteractor{
            ActivityRepository: &repository.ActivityRepository{
                LogSearchEngineHandler: handler,
            },
        },
    }
}

func responseAsJson(rw http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(code)
    rw.Write(response)
}

func (controller *ActivityController) Get(w http.ResponseWriter, vars map[string] string) {
    userId, _ := strconv.Atoi(vars["userId"])
    activity, _ := controller.Interactor.FindBy(userId)
    responseAsJson(w, http.StatusOK, activity)
}


func respondWithError(w http.ResponseWriter, code int, msg string) {
    responseAsJson(w, code, map[string]string{"error": msg})
}