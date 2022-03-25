package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID    uuid.UUID
		Name  string
		Email string
		Admin bool
		Meta  map[string]interface{}
	}

	Challenge struct {
		ID          uuid.UUID
		Type        string
		Content     string
		Description string
		StartTime   time.Time
		EndTime     time.Time
		Predictions []*Prediction
	}

	Prediction struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		ChallengeID uuid.UUID
		Prognosis   bool
		Meta        map[string]interface{}
	}

	SystemSymmary struct {
		ID         int
		CreateTime time.Time
	}
)
