package di

import (
	"github.com/Khmer495/go-templete/internal/app/api/v1/handler"
	"github.com/Khmer495/go-templete/internal/app/api/v1/infrastructure/echo"
	"github.com/Khmer495/go-templete/internal/pkg/config"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/domain/usecase"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/firebase"
	"github.com/Khmer495/go-templete/internal/pkg/util"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func newPanic(err error) {
	zap.S().Panicf("di: %+v", err)
}

func provideFirebaseAuth(d *dig.Container) {
	if err := d.Provide(echo.NewFirebaseAuth); err != nil {
		newPanic(err)
	}
}

func provideSimpleAuth(d *dig.Container) {
	if err := d.Provide(echo.NewSimpleAuth); err != nil {
		newPanic(err)
	}
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
	if err := d.Provide(repository.NewAuthRepository); err != nil {
		newPanic(err)
	}
	if err := d.Provide(repository.NewUserRepository); err != nil {
		newPanic(err)
	}
	if err := d.Provide(repository.NewTeamRepository); err != nil {
		newPanic(err)
	}
}

func provideUsecase(d *dig.Container) {
	if err := d.Provide(usecase.NewUserUsecase); err != nil {
		newPanic(err)
	}
	if err := d.Provide(usecase.NewTeamUsecase); err != nil {
		newPanic(err)
	}
}

func provideHandler(d *dig.Container) {
	if err := d.Provide(handler.NewUserHandler); err != nil {
		newPanic(err)
	}
	if err := d.Provide(handler.NewTeamHandler); err != nil {
		newPanic(err)
	}
	if err := d.Provide(handler.NewHandler); err != nil {
		newPanic(err)
	}
}

func provideInfrastructure(d *dig.Container) {
	if err := d.Provide(firebase.NewFirebaseClient); err != nil {
		newPanic(err)
	}
}

func invokeEchoServer(d *dig.Container) {
	if err := d.Invoke(echo.NewEchoServer); err != nil {
		newPanic(err)
	}
}

func NewDig() *dig.Container {
	d := dig.New()
	env := config.Env()
	if util.IsPrd(env) || util.IsStg(env) || util.IsDev(env) {
		provideFirebaseAuth(d)
	} else {
		provideSimpleAuth(d)
	}
	provideHandler(d)
	provideService(d)
	provideRepository(d)
	provideUsecase(d)
	provideInfrastructure(d)
	invokeEchoServer(d)
	return d
}
