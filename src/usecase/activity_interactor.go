package usecase

import (
    "domain"
)

type ActivityInteractor struct {
    ActivityRepository ActivityRepository
}

func (interactor *ActivityInteractor) FindBy(userId int) (activities domain.Activity, err error) {
    activities, err = interactor.ActivityRepository.FindBy(userId)
    return
}