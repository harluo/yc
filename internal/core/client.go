package core

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/yc"
	"github.com/harluo/yc/internal/config"
)

type Client = yc.Client

func newClient(secret *config.Secret, logger log.Logger, http *http.Client) *Client {
	builder := yc.New(secret.Id, secret.Key)
	builder.Logger(logger)
	builder.Http(http)

	return builder.Build()
}
