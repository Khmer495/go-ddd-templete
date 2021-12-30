package di

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/firebase"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func newPanic(err error) {
	zap.S().Panicf("di: %+v", err)
}

func provideService(d *dig.Container) {
	if err := d.Provide(service.NewUserService); err != nil {
		newPanic(err)
	}
	if err := d.Provide(service.NewTeamService); err != nil {
		newPanic(err)
	}
}

func provideRepository(d *dig.Container) {
	if err := d.Provide(repository.NewTeamRepository); err != nil {
		newPanic(err)
	}
}

func provideUsecase(d *dig.Container) {
}

func provideInfrastructure(d *dig.Container) {
	if err := d.Provide(firebase.NewFirebaseClient); err != nil {
		newPanic(err)
	}
}

func NewDig() *dig.Container {
	d := dig.New()

	provideService(d)
	provideRepository(d)
	provideUsecase(d)
	provideInfrastructure(d)

	return d
}
