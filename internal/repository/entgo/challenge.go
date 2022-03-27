package entgo

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/challenge"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
)

func (r *EntgoRepository) GetChallengeByContent(ctx context.Context, content string) (*domain.Challenge, error) {
	c, err := r.client.Challenge.
		Query().
		Where(challenge.ContentEQ(content)).
		WithPredictions().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c), nil
}

func (r *EntgoRepository) GetChallengeByID(ctx context.Context, id uuid.UUID) (*domain.Challenge, error) {
	c, err := r.client.Challenge.
		Query().
		Where(challenge.IDEQ(id)).
		WithPredictions().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c), nil
}

func (r *EntgoRepository) GetRandomFinishedChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.CreateTimeLT(time.Now()),
				challenge.EndTimeLT(time.Now()),
			),
		).
		WithPredictions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch))
	}

	return res, nil
}

func (r *EntgoRepository) GetRandomOngoingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.CreateTimeLT(time.Now()),
				challenge.EndTimeGT(time.Now()),
			),
		).
		WithPredictions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch))
	}

	return res, nil
}

// GetPopularOngoingChallenges returns ongoing challenges that have more predictions
func (r *EntgoRepository) GetPopularOngoingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	return nil, nil
}

// GetPopularFinishedChallenges returns finished challenges that have more predictions
func (r *EntgoRepository) GetPopularFinishedChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	return nil, nil
}

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
			SetNillableOutcome(ch.Outcome).
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
		SetNillableOutcome(ch.Outcome).
		SetDescription(ch.Description)

	c, err = chUpdateQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c), nil
}

func entToDomainChallenge(ch *ent.Challenge) *domain.Challenge {
	var predictions []*domain.Prediction
	if ch.Edges.Predictions != nil {
		for _, p := range ch.Edges.Predictions {
			predictions = append(predictions, entToDomainPrediction(p))
		}
	}

	return &domain.Challenge{
		ID:          ch.ID,
		Content:     ch.Content,
		Outcome:     ch.Outcome,
		Description: ch.Description,
		StartTime:   ch.StartTime,
		EndTime:     ch.EndTime,
		Predictions: predictions,
	}
}
