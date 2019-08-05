package Commands

import (
	"github.com/alistairfink/Link-Shortener/Data/DataModels"
	"github.com/go-pg/pg"
)

type LinkCommand struct {
	DB *pg.DB
}

func (this *LinkCommand) GetLink(id string) *DataModels.LinkDataModel {
	if !this.Exists(id) {
		return nil
	}

	var models []DataModels.LinkDataModel
	err := this.DB.Model(&models).Where("id = ?", id).Select()
	if err != nil {
		panic(err.Error())
	}

	return &models[0]
}

func (this *LinkCommand) GetAllLinks() *[]DataModels.LinkDataModel {
	var models []DataModels.LinkDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err.Error())
	}

	return &models
}

func (this *LinkCommand) CreateLink(link *DataModels.LinkDataModel) *DataModels.LinkDataModel {
	if this.Exists(link.Id) {
		return nil
	}

	err := this.DB.Insert(link)
	if err != nil {
		panic(err.Error())
	}

	return this.GetLink(link.Id)
}

func (this *LinkCommand) Exists(id string) bool {
	var models []DataModels.LinkDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", id).Exists()
	if err != nil {
		panic(err.Error())
	}

	return exists
}
