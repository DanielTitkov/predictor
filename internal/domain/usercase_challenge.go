package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func (ch *Challenge) Finished() bool {
	return ch.EndTime.Before(time.Now())
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

func (a *FilterChallengesArgs) Validate() error {
	if a.Ongoing && a.Finished {
		return errors.New("invalid request: cannot query 'ongoing' and 'finished' at the same time")
	}

	if a.Unvoted && a.UserID == uuid.Nil {
		return errors.New("invalid query: to query unvoted user id is required")
	}

	return nil
}
