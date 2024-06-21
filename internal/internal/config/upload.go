package config

type Upload struct {
	Enabled *bool  `default:"${ENABLED=true}" json:"enabled,omitempty"`
	Name    string `default:"${NAME=未命名}" json:"name,omitempty"`

	File  string `default:"${FILE}" json:"file,omitempty" validate:"required_without=Files"`
	Files string `default:"${FILES}" json:"files,omitempty" validate:"required_without=File"`
}

func (u *Upload) Checked() bool {
	return "" != u.File || 0 != len(u.Files)
}
