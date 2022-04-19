package entgo

import (
	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
)

func entToDomainProof(p *ent.Proof) *domain.Proof {
	return &domain.Proof{
		ID:      p.ID,
		Content: p.Content,
		Link:    p.Link,
		Meta:    p.Meta,
	}
}
