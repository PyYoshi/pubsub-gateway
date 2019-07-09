package config

import (
	"encoding/base64"
)

// Base64String Base64文字列
type Base64String string

// Decode Base64Stringをenvconfig.Processでデコードできるようにする
func (b *Base64String) Decode(value string) error {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	*b = Base64String(string(data))
	return nil
}
