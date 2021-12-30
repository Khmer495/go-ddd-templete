package main

import (
	"github.com/Khmer495/go-templete/internal/app/api/v1/util/di"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/util/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func init() {
	logger.NewLogger()
	repository.NewMysqlClient()
}

func main() {
	zap.L().Info("Start Server.")
	di.NewDig()
	defer repository.MysqlClient.Close()
}
