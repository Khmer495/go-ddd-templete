//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type ITeamService interface {
	IsExist(ctx context.Context, teamId entity.Id) (bool, error)
}

type teamService struct {
	tr repository.ITeamRepository
}

func NewTeamService(tr repository.ITeamRepository) ITeamService {
	return teamService{
		tr: tr,
	}
}

func (ts teamService) IsExist(ctx context.Context, teamId entity.Id) (bool, error) {
	isExist, err := ts.tr.IsExist(ctx, teamId)
	if err != nil {
		return false, xerrors.Errorf("os.or.IsExist: %w", err)
	}
	return isExist, err
}
