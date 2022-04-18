package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (ch *Challenge) Finished() bool {
	return ch.EndTime.Before(time.Now())
}

func (ch *Challenge) Started() bool {
	return ch.StartTime.Before(time.Now())
}

func (ch *Challenge) Votes() int {
	return len(ch.Predictions)
}

func (ch *Challenge) VotesTrue() int {
	var res int
	for _, pred := range ch.Predictions {
		if pred.Prognosis {
			res++
		}
	}
	return res
}

func (ch *Challenge) VotesFalse() int {
	var res int
	for _, pred := range ch.Predictions {
		if !pred.Prognosis {
			res++
		}
	}
	return res
}

func (ch *Challenge) PercTrue() int {
	return int(float64(ch.VotesTrue()) / float64(ch.Votes()) * 100.0)
}

func (ch *Challenge) PercFalse() int {
	return 100 - ch.PercTrue()
}

func (ch *Challenge) HasOutcome() bool {
	return ch.Outcome != nil
}

func (ch *Challenge) AllowOutcomeEdit() bool {
	if ch.HasOutcome() {
		return false
	}

	if !ch.Finished() {
		return false
	}

	return true
}

func (ch *Challenge) AllowDetailsEdit() bool {
	if ch.Started() {
		return false
	}

	// TODO: maybe allow to edit
	if ch.Published {
		return false
	}

	return true
}

func (ch *Challenge) HasOutcomeAndTrue() bool {
	// safety check
	if !ch.HasOutcome() {
		return false
	}
	return *ch.Outcome
}

func (ch *Challenge) StartStr() string {
	return ch.StartTime.Format(ChallengeTimeFormat)
}

func (ch *Challenge) EndStr() string {
	return ch.EndTime.Format(ChallengeTimeFormat)
}

func (a *CreateChallengeArgs) Validate() error {
	if a.Outcome != nil {
		return errors.New("not allowed to create with outcome")
	}

	if len(a.Content) > 140 {
		return errors.New("content must be less than 140 characters")
	}

	if len(a.Content) < 10 {
		return errors.New("content must be more than 10 characters")
	}

	if len(a.Description) < 10 {
		return errors.New("description must be more than 10 characters")
	}

	if len(a.Description) > 280 {
		return errors.New("description must be less than 280 characters")
	}

	startTime, err := time.Parse(a.TimeLayout, a.StartTime)
	if err != nil {
		return fmt.Errorf("failed to parse start time: %s", err)
	}

	endTime, err := time.Parse(a.TimeLayout, a.EndTime)
	if err != nil {
		return fmt.Errorf("failed to parse end time: %s", err)
	}

	if !startTime.Before(endTime) {
		return errors.New("start time must be before end time")
	}

	if endTime.Before(time.Now()) {
		return errors.New("end time must be in the future")
	}

	if startTime.Before(time.Now()) {
		return errors.New("start time must be in the future")
	}

	return nil
}

func (a *FilterChallengesArgs) Validate(requireUser bool) error {
	if requireUser {
		if a.UserID == uuid.Nil {
			return errors.New("user id required")
		}

		if a.Unvoted {
			return errors.New("unvoted is not supported with require user query")
		}
	}

	if a.Ongoing && a.Finished {
		return errors.New("invalid request: cannot query 'ongoing' and 'finished' at the same time")
	}

	if a.Unvoted && a.UserID == uuid.Nil {
		return errors.New("invalid query: to query unvoted user id is required")
	}

	return nil
}
