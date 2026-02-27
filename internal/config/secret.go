package config

type Secret struct {
	Id  string `json:"id,omitempty" validate:"required"`
	Key string `json:"key,omitempty" validate:"required"`
}

func newSecret(yaothink *Yaothink) *Secret {
	return yaothink.Secret
}
