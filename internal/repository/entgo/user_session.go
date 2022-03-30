package entgo

import (
	"context"
	"errors"
	"time"

	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/user"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/usersession"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreateUserSession(ctx context.Context, s *domain.UserSession) (*domain.UserSession, error) {
	ses, err := r.client.UserSession.
		Create().
		SetSid(s.SID).
		SetUserID(s.UserID).
		SetIP(s.IP).
		SetUserAgent(s.UserAgent).
		SetMeta(s.Meta).
		SetLastActivity(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUserSession(ses), nil
}

// CreateOrUpdateUserSession must be used only with authentication like oauth
// because it updated ip and user agent and in other uses may be vulnerable
// to session steal
func (r *EntgoRepository) CreateOrUpdateUserSession(ctx context.Context, s *domain.UserSession) (*domain.UserSession, error) {
	if s.UserID == uuid.Nil {
		return nil, errors.New("user id required to create or update user session")
	}

	ses, err := r.client.UserSession.
		Query().
		Where(
			usersession.And(
				usersession.SidEQ(s.SID),
				usersession.HasUserWith(
					user.IDEQ(s.UserID),
				),
			),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create session
		ses, err := r.client.UserSession.
			Create().
			SetSid(s.SID).
			SetUserID(s.UserID).
			SetIP(s.IP).
			SetUserAgent(s.UserAgent).
			SetMeta(s.Meta).
			SetLastActivity(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainUserSession(ses), nil
	}

	// update session
	ses, err = ses.Update().
		SetIP(s.IP).
		SetUserAgent(s.UserAgent).
		SetLastActivity(time.Now()).
		SetMeta(s.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUserSession(ses), nil
}

func (r *EntgoRepository) IfSessionRegistered(ctx context.Context, s *domain.UserSession) (bool, error) {
	return r.client.UserSession.
		Query().
		Where(
			usersession.And(
				usersession.SidEQ(s.SID),
				usersession.IPEQ(s.IP),
				usersession.UserAgentEQ(s.UserAgent),
			),
		).
		Exist(ctx)
}

func (r *EntgoRepository) UpdateUserSessionLastActivityBySID(ctx context.Context, sid string) error {
	ses, err := r.client.UserSession.Query().Where(usersession.SidEQ(sid)).Only(ctx)
	if err != nil {
		return err
	}

	_, err = ses.Update().SetLastActivity(time.Now()).Save(ctx)
	return err
}

func (r *EntgoRepository) GetUserBySession(ctx context.Context, s *domain.UserSession) (*domain.User, error) {
	user, err := r.client.UserSession.
		Query().
		Where(
			usersession.And(
				usersession.SidEQ(s.SID),
				usersession.IPEQ(s.IP),
				usersession.UserAgentEQ(s.UserAgent),
			),
		).
		QueryUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func entToDomainUserSession(s *ent.UserSession) *domain.UserSession {
	var userID uuid.UUID
	if s.Edges.User != nil {
		userID = s.Edges.User.ID
	}

	return &domain.UserSession{
		ID:           s.ID,
		UserID:       userID,
		SID:          s.Sid,
		IP:           s.IP,
		UserAgent:    s.UserAgent,
		CreateTime:   s.CreateTime,
		UpdateTime:   s.UpdateTime,
		LastActivity: s.LastActivity,
		Meta:         s.Meta,
	}
}
