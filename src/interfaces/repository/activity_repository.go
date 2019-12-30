package repository

import (
    "domain"
)

type ActivityRepository struct {
    LogSearchEngineHandler
}

func (repo *ActivityRepository) FindBy(userId int) (activity domain.Activity, err error){
    logs, _ := repo.Query("TODO")
    activity.Actions = logs
    return
}