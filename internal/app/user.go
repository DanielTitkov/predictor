package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/markbates/goth"
	"github.com/sethvargo/go-password/password"

	"github.com/DanielTitkov/predictor/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	if u.Password == "" {
		// TODO: add password strength checks
		return nil, errors.New("user password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	u.PasswordHash = string(hash)
	user, err := a.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticateGothUser creates new user or returns existsing one.
// It relies on goth authetication to verify user has access
// to that profile and thus doesn't check password.
func (a *App) AuthenticateGothUser(ctx context.Context, gu *goth.User) (*domain.User, error) {
	exists, err := a.repo.IfEmailRegistered(ctx, gu.Email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return a.CreateUserFromGoth(ctx, gu)
	}

	// TODO if user came from another provider add new data to meta
	return a.GetUserByEmail(ctx, gu.Email)
}

func (a *App) LogoutUser(ctx context.Context) error {
	return nil
}

func (a *App) CreateUserFromGoth(ctx context.Context, gu *goth.User) (*domain.User, error) {
	passw, err := password.Generate(16, 5, 0, false, true)
	if err != nil {
		return nil, err
	}

	meta := make(map[string]interface{})
	meta[gu.Provider] = *gu

	user := &domain.User{
		Name:     gu.NickName,
		Email:    gu.Email,
		Picture:  gu.AvatarURL,
		Password: passw,
		Admin:    false,
		Meta:     meta,
	}

	user, err = a.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Password = passw // TODO: this is to show or send password to the user

	return user, nil
}

func (a *App) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := a.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// we should not return password hash if this is not needed
	return &domain.User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
		Admin:   user.Admin,
		Meta:    user.Meta,
	}, nil
}

func (a *App) ValidateUserPassword(ctx context.Context, u *domain.User) (bool, error) {
	user, err := a.repo.GetUserByEmail(ctx, u.Email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (a *App) loadUserPresets() error {
	a.log.Info("loading user presets", fmt.Sprint(a.Cfg.Data.Presets.UserPresetsPaths))
	for _, path := range a.Cfg.Data.Presets.UserPresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var users []domain.User
		err = json.Unmarshal(data, &users)
		if err != nil {
			return err
		}

		for _, user := range users {
			ctx := context.Background()

			// its only for debug purposes so checking errors is not critical
			u, _ := a.GetUserByEmail(ctx, user.Email)
			if u != nil {
				a.log.Debug("user already exists", fmt.Sprintf("%+v", u))
				continue
			}

			u, err := a.CreateUser(ctx, &user)
			if err != nil {
				return err
			}

			a.log.Debug("loaded user", fmt.Sprintf("%+v", u))
		}
	}

	return nil
}
