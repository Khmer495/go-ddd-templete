//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"
	"fmt"

	"github.com/Khmer495/go-templete/internal/pkg/config"
	"github.com/Khmer495/go-templete/internal/pkg/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util"
	"go.uber.org/zap"
)

var MysqlClient *ent.Client

func NewMysqlClient() {
	env := config.Env()
	if util.IsPrd(env) {
		newPrdMysqlClient()
	} else {
		newDevMysqlClient()
	}
}

func newPrdMysqlClient() {
	var err error
	MysqlClient, err = ent.Open("mysql", fmt.Sprintf("%s?parseTime=True", config.MysqlDSN()))
	if err != nil {
		zap.S().Panicf("ent.Open: %+v", err)
	}
	if err := MysqlClient.Schema.Create(context.Background()); err != nil {
		zap.S().Panicf("mysqlClient.Schema.Create: %+v", err)
	}
}

func newDevMysqlClient() {
	var err error
	MysqlClient, err = ent.Open("mysql", fmt.Sprintf("%s?parseTime=True", config.MysqlDSN()), ent.Debug())
	if err != nil {
		zap.S().Panicf("ent.Open: %+v", err)
	}
	if err := MysqlClient.Schema.Create(context.Background()); err != nil {
		zap.S().Panicf("mysqlClient.Schema.Create: %+v", err)
	}
}
