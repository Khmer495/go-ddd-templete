package appconfig

import (
	"os"

	"go.uber.org/zap"
)

var RestPort = func() string {
	val, ok := os.LookupEnv("REST_PORT")
	if !ok {
		val = "8080"
	}
	zap.S().Info("REST_PORT: %s", val)
	return val
}()
