package main

import (
	"context"

	"github.com/Khmer495/go-templete/internal/app/batch/samplebatch/handler"
	"github.com/Khmer495/go-templete/internal/app/batch/samplebatch/util/di"
	"github.com/Khmer495/go-templete/internal/pkg/domain/usecase"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/entmysql"
	"github.com/Khmer495/go-templete/internal/pkg/util/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func lambdaHandler(ctx context.Context, event events.CloudWatchEvent) {
	d := di.NewDig()
	triggeredUTCTime := event.Time
	err := d.Invoke(
		func(tu usecase.ITeamUsecase) {
			handler.Handler(ctx, tu, triggeredUTCTime)
		},
	)
	if err != nil {
		zap.S().Fatalf("di: %+v", err)
	}
}

func init() {
	logger.NewLogger()
	entmysql.NewMysqlClient()
}

func main() {
	zap.L().Info("Initialize.")
	lambda.Start(lambdaHandler)
	defer entmysql.MysqlClient.Close()
}
