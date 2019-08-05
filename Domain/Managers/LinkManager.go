package Managers

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/alistairfink/Link-Shortener/Config"
	"github.com/alistairfink/Link-Shortener/Data/Commands"
	"github.com/alistairfink/Link-Shortener/Data/DataModels"
	"github.com/alistairfink/Link-Shortener/Domain/ViewModels"
	"github.com/go-pg/pg"
	"strings"
	"time"
)

type LinkManager struct {
	DB          *pg.DB
	Config      *Config.Config
	LinkCommand *Commands.LinkCommand
}

func NewLinkManager(db *pg.DB, config *Config.Config) *LinkManager {
	linkManager := &LinkManager{
		DB:          db,
		Config:      config,
		LinkCommand: &Commands.LinkCommand{DB: db},
	}

	return linkManager
}

func (this *LinkManager) GetLink(shortenedLink string) *DataModels.LinkDataModel {
	result := this.LinkCommand.GetLink(shortenedLink)
	if result == nil {
		return nil
	}

	return result
}

func (this *LinkManager) GetAllLinks(filter int) *[]DataModels.LinkDataModel {
	results := this.LinkCommand.GetAllLinks()
	filter = len(*results) - filter
	if filter < 0 {
		filter = 0
	}

	filteredResults := (*results)[filter:len(*results)]
	return &filteredResults
}

func (this *LinkManager) CreateLink(link *ViewModels.LinkRequestModel) *DataModels.LinkDataModel {
	dataModel := this.generateLinkId(link.Link)
	return dataModel
}

func (this *LinkManager) generateLinkId(link string) *DataModels.LinkDataModel {
	var upsertResult *DataModels.LinkDataModel
	upsertResult = nil
	for upsertResult == nil {
		currTime := time.Now()
		hash := md5.New()
		hash.Write([]byte(currTime.String() + link))
		hashedContent := hash.Sum(nil)
		encodedHash := base64.StdEncoding.EncodeToString(hashedContent)
		encodedSubstring := encodedHash[0:6]
		encodedSubstring = strings.Replace(encodedSubstring, "/", "-", -1)
		encodedSubstring = strings.Replace(encodedSubstring, "\\", "_", -1)

		linkModel := &DataModels.LinkDataModel{
			Id:   encodedSubstring,
			Link: link,
		}
		upsertResult = this.LinkCommand.CreateLink(linkModel)
	}

	return upsertResult
}
