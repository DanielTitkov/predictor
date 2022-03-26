package entgo

import (
	"context"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetUserCount(ctx context.Context) (int, error) {
	return r.client.User.Query().Count(ctx)
}

func (r *EntgoRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	user, err := r.client.User.
		Create().
		SetName(u.Name).
		SetEmail(u.Email).
		SetPasswordHash(u.PasswordHash).
		// TODO: not setting admin here
		SetMeta(u.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func entToDomainUser(user *ent.User) *domain.User {
	return &domain.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Meta:         user.Meta,
		Admin:        user.Admin,
	}
}
