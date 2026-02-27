package config

import (
	"github.com/harluo/config"
)

type Yaothink struct {
	Secret *Secret `default:"{}" json:"secret,omitempty" validate:"required"`
}

func newYaothink(config config.Getter) (yaothink *Yaothink, err error) {
	yaothink = new(Yaothink)
	err = config.Get(&struct {
		Yaothink *Yaothink `default:"{}" json:"yaothink,omitempty" validate:"required"`
	}{
		Yaothink: yaothink,
	})

	return
}
