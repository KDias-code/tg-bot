package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	err2 "github.com/KDias-code/lib/err"
	"io"
)

type Storage interface {
	Save(p *Page) error
	Remove(p *Page) error
	PickRandom(username string) (*Page, error)
	IsExist(p *Page) (bool, error)
}

var ErrNoSaved = errors.New("No Saved pages")

type Page struct {
	URL      string
	Username string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", err2.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.Username); err != nil {
		return "", err2.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
