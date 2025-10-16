package service

import (
	"context"
	"strings"
	"time"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/repository/prospect"
)

type ProspectService interface {
	NewProspect(prospect entity.Prospect, ctx context.Context) error
}
type prospectService struct {
	prospectRepo prospect.ProspectRepo
}

func NewProspectService(prospectRepo prospect.ProspectRepo) ProspectService {
	return &prospectService{
		prospectRepo: prospectRepo,
	}
}
func (p *prospectService) NewProspect(prospect entity.Prospect, ctx context.Context) error {
	prospect.CompanyName = strings.TrimSpace(strings.ToLower(prospect.CompanyName))
	prospect.ContactName = strings.TrimSpace(strings.ToLower(prospect.ContactName))
	prospect.Email = strings.TrimSpace(strings.ToLower(prospect.Email))
	prospect.Phone = strings.TrimSpace(strings.ToLower(prospect.Phone))
	prospect.CreatedAt = time.Now()

	if prospect.CompanyName == "" || prospect.ContactName == "" || prospect.Email == "" || prospect.Phone == "" {
		return entity.ErrMissingField
	}

	err := p.prospectRepo.InsertOne(prospect, ctx)
	if err != nil {
		return err
	}
	return nil
}
