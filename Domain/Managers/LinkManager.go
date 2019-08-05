package Managers

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/alistairfink/Link-Shortener/Config"
	"github.com/alistairfink/Link-Shortener/Domain/ViewModels"
	"github.com/go-pg/pg"
	"time"
)

type LinkManager struct {
	DB     *pg.DB
	Config *Config.Config
}

func NewLinkManager(db *pg.DB, config *Config.Config) *LinkManager {
	linkManager := &LinkManager{
		DB:     db,
		Config: config,
	}

	return linkManager
}

func (this *LinkManager) GetLink(shortenedLink string) {

}

func (this *LinkManager) CreateLink(link *ViewModels.LinkRequestModel) {
	println(link.Link)
	println(this.Config.BaseLink + this.generateLinkId(link.Link))
}

func (this *LinkManager) generateLinkId(link string) string {
	currTime := time.Now()
	hash := md5.New()
	hash.Write([]byte(currTime.String() + link))
	hashedContent := hash.Sum(nil)
	encodedHash := base64.StdEncoding.EncodeToString(hashedContent)
	return encodedHash[0:6]
}
