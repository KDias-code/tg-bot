package storage

type Storage interface {
	Save(p *Page) error
	Remove(p *Page) error
	PickRandom(username string) (*Page, error)
	IsExist(p *Page) (bool, error)
}

type Page struct {
	URL      string
	Username string
}
