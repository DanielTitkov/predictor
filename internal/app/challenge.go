package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/DanielTitkov/predictor/internal/domain"
)

func (a *App) CreateTaskFromArgs(ctx context.Context, args domain.CreateChallengeArgs) (*domain.Challenge, error) {
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

			challenge, err := a.CreateTaskFromArgs(ctx, args)
			if err != nil {
				return err
			}

			a.log.Info("loaded challenge", fmt.Sprintf("%+v", challenge))
		}
	}

	return nil
}
