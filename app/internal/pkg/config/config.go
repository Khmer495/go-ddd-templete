package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func Env() string {
	if val, ok := os.LookupEnv("ENV"); ok {
		return val
	}
	return "LOCAL"
}

func MysqlDSN() string {
	user, ok := os.LookupEnv("MYSQL_USER")
	if !ok {
		zap.L().Panic(`os.LookupEnv("MYSQL_USER")`)
	}
	password, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		zap.L().Panic(`os.LookupEnv("MYSQL_PASSWORD")`)
	}
	host, ok := os.LookupEnv("MYSQL_HOST")
	if !ok {
		zap.L().Panic(`os.LookupEnv("MYSQL_HOST")`)
	}
	port, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		zap.L().Panic(`os.LookupEnv("MYSQL_PORT")`)
	}
	db, ok := os.LookupEnv("MYSQL_DATABASE")
	if !ok {
		zap.L().Panic(`os.LookupEnv("MYSQL_DATABASE")`)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}

func AWSAccessKey() string {
	val, ok := os.LookupEnv("AWS_ACCESS_KEY")
	if !ok {
		zap.L().Panic(`os.LookupEnv("AWS_ACCESS_KEY")`)
	}
	return val
}

func AWSSecretKey() string {
	val, ok := os.LookupEnv("AWS_SECRET_KEY")
	if !ok {
		zap.L().Panic(`os.LookupEnv("AWS_SECRET_KEY")`)
	}
	return val
}

func AWSRegion() string {
	return "ap-northeast-1"
}
