package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"
)

func (a *App) GetChallengeByContent(ctx context.Context, content string) (*domain.Challenge, error) {
	return a.repo.GetChallengeByContent(ctx, content)
}

func (a *App) GetChallengeByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Challenge, error) {
	return a.repo.GetChallengeByID(ctx, id, userID)
}

func (a *App) GetRandomFinishedChallenges(ctx context.Context) ([]*domain.Challenge, error) {
	return a.repo.GetRandomFinishedChallenges(ctx, a.Cfg.App.HomeChallengePageLimit)
}

func (a *App) GetRandomPendingChallenges(ctx context.Context) ([]*domain.Challenge, error) {
	return a.repo.GetRandomPendingChallenges(ctx, a.Cfg.App.HomeChallengePageLimit)
}

func (a *App) GetClosingChallenges(ctx context.Context) ([]*domain.Challenge, error) {
	return a.repo.GetClosingChallenges(ctx, a.Cfg.App.HomeChallengePageLimit)
}

func (a *App) GetRandomOngoingChallenges(ctx context.Context, userID uuid.UUID, limit int) ([]*domain.Challenge, error) {
	if limit == 0 || limit > a.Cfg.App.HomeChallengePageLimit {
		limit = a.Cfg.App.HomeChallengePageLimit
	}

	return a.repo.GetRandomOngoingChallenges(ctx, limit, userID)
}

func (a *App) GetRandomTrueChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	if limit == 0 || limit > a.Cfg.App.HomeChallengePageLimit {
		limit = a.Cfg.App.HomeChallengePageLimit
	}

	return a.repo.GetRandomTrueChallenges(ctx, limit)
}

func (a *App) GetRandomFalseChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	if limit == 0 || limit > a.Cfg.App.HomeChallengePageLimit {
		limit = a.Cfg.App.HomeChallengePageLimit
	}

	return a.repo.GetRandomFalseChallenges(ctx, limit)
}

func (a *App) GetUserChallenges(ctx context.Context, userID uuid.UUID) ([]*domain.Challenge, error) {
	return a.repo.GetUserChallenges(ctx, userID)
}

func (a *App) FilterChallenges(ctx context.Context, args *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error) {
	if err := args.Validate(); err != nil {
		return nil, 0, err
	}

	return a.repo.FilterChallenges(ctx, args)
}

func (a *App) CreateChallengeFromArgs(ctx context.Context, args domain.CreateChallengeArgs) (*domain.Challenge, error) {
	startTime, err := time.Parse(args.TimeLayout, args.StartTime)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse(args.TimeLayout, args.EndTime)
	if err != nil {
		return nil, err
	}

	challenge := &domain.Challenge{
		Type:        args.Type,
		Content:     args.Content,
		Description: args.Description,
		StartTime:   startTime,
		EndTime:     endTime,
		Outcome:     args.Outcome,
	}

	challenge, err = a.repo.CreateOrUpdateChallengeByContent(ctx, challenge)
	if err != nil {
		return nil, err
	}

	return challenge, nil

}

func (a *App) loadChallengePresets() error {
	a.log.Info("loading challenge presets", fmt.Sprint(a.Cfg.Data.Presets.ChallengePresetsPaths))
	for _, path := range a.Cfg.Data.Presets.ChallengePresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var presets []domain.CreateChallengeArgs
		err = json.Unmarshal(data, &presets)
		if err != nil {
			return err
		}

		for _, args := range presets {
			ctx := context.Background()

			challenge, err := a.CreateChallengeFromArgs(ctx, args)
			if err != nil {
				return err
			}

			a.log.Debug("loaded challenge", fmt.Sprintf("%+v", challenge))
		}
	}

	return nil
}
