package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newYaothink,

		newSecret,
	).Build().Apply()
}
