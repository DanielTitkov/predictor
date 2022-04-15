package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID           uuid.UUID
		Name         string
		Email        string
		Admin        bool
		Picture      string
		Password     string
		PasswordHash string
		Locale       string
		Meta         map[string]interface{}
	}

	UserSummary struct {
		UserID               uuid.UUID
		CorrectPredictions   int
		IncorrectPredictions int
		UnknownPredictions   int
	}

	UserSession struct {
		ID           int    // probably uuid not needed here, sessions are temporary anyways
		SID          string // code to identify the session
		UserID       uuid.UUID
		IP           string
		UserAgent    string
		CreateTime   time.Time
		UpdateTime   time.Time
		LastActivity time.Time
		Meta         map[string]interface{}
		Active       bool
	}

	Badge struct {
		ID     int // probably not needed
		UserID uuid.UUID
		Type   string
		Active bool
		Meta   map[string]interface{}
	}

	Challenge struct {
		ID             uuid.UUID
		Type           string
		Content        string
		Description    string
		Outcome        *bool
		StartTime      time.Time
		EndTime        time.Time
		Predictions    []*Prediction
		UserPrediction *Prediction
	}

	Prediction struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		ChallengeID uuid.UUID
		Prognosis   bool
		Meta        map[string]interface{}
	}

	SystemSymmary struct {
		ID                 int
		Users              int
		Challenges         int
		OngoingChallenges  int
		FinishedChallenges int
		Predictions        int
		CreateTime         time.Time
	}
)
