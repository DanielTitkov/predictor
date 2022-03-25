package entgo

import (
	"context"

	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/challenge"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreateOrUpdateChallengeByContent(ctx context.Context, ch *domain.Challenge) (*domain.Challenge, error) {
	// query challenge by content
	c, err := r.client.Challenge.
		Query().
		Where(challenge.ContentEQ(ch.Content)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create challenge
		chQuery := r.client.Challenge.
			Create().
			SetContent(ch.Content).
			SetStartTime(ch.StartTime).
			SetEndTime(ch.EndTime).
			SetDescription(ch.Description)

		c, err = chQuery.Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainChallenge(c), nil
	}

	// update challenge
	chUpdateQuery := c.Update().
		SetStartTime(ch.StartTime).
		SetEndTime(ch.EndTime).
		SetDescription(ch.Description)

	c, err = chUpdateQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c), nil
}

func entToDomainChallenge(ch *ent.Challenge) *domain.Challenge {
	return &domain.Challenge{
		ID:          ch.ID,
		Content:     ch.Content,
		Description: ch.Description,
		StartTime:   ch.StartTime,
		EndTime:     ch.EndTime,
	}
}
