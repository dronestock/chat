package config

type Lark struct {
	Id     string `default:"${ID}" json:"id,omitempty"`
	Secret string `default:"${SECRET}" json:"secret,omitempty"`
}
