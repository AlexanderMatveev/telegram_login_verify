package telegram_login_verify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

var (
	ErrWrongLoginData = errors.New("wrong login data")
)

type LoginData struct {
	AuthDate  uint64
	FirstName string
	Id        uint64
	LastName  string
	PhotoUrl  string
	Username  string
	Hash      string
}

func Verify(d *LoginData, token string) error {
	keyHash := sha256.New()
	keyHash.Write([]byte(token))
	key := keyHash.Sum(nil)
	h := hmac.New(sha256.New, key)
	message := fmt.Sprintf("auth_date=%d", d.AuthDate)
	if d.FirstName != "" {
		message += fmt.Sprintf("\nfirst_name=%s", d.FirstName)
	}
	if d.Id != 0 {
		message += fmt.Sprintf("\nid=%d", d.Id)
	}
	if d.LastName != "" {
		message += fmt.Sprintf("\nlast_name=%s", d.LastName)
	}
	if d.PhotoUrl != "" {
		message += fmt.Sprintf("\nphoto_url=%s", d.PhotoUrl)
	}
	if d.Username != "" {
		message += fmt.Sprintf("\nusername=%s", d.Username)
	}
	h.Write([]byte(message))
	if hex.EncodeToString(h.Sum(nil)) != d.Hash {
		return ErrWrongLoginData
	}
	return nil
}
