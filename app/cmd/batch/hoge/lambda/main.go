package main

import (
	"context"

	"github.com/Khmer495/go-templete/internal/app/batch/hoge/handler"
	"github.com/Khmer495/go-templete/internal/app/batch/hoge/util/di"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/util/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func lambdaHandler(ctx context.Context, event events.CloudWatchEvent) {
	d := di.NewDig()
	targetUTCTime := event.Time
	err := d.Invoke(
		func() {
			handler.Handler(targetUTCTime)
		},
	)
	if err != nil {
		zap.S().Fatalf("di: %+v", err)
	}
}

func init() {
	logger.NewLogger()
	repository.NewMysqlClient()
}

func main() {
	zap.L().Info("Initialize.")
	lambda.Start(lambdaHandler)
	defer repository.MysqlClient.Close()
}
