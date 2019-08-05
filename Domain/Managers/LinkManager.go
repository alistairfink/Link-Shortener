package Managers

import (
	"github.com/go-pg/pg"
)

type LinkManager struct {
	DB *pg.DB
}

func NewLinkManager(db *pg.DB) *LinkManager {
	linkManager := &LinkManager{
		DB: db,
	}

	return linkManager
}

func (this *LinkManager) GetLink(shortenedLink string) {

}

func (this *LinkManager) CreateLink(link string) {

}
