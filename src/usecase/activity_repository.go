package usecase

import (
    "domain"
)

type ActivityRepository interface {
    FindBy(int) (domain.Activity, error)
}