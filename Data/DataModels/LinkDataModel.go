package DataModels

import ()

type LinkDataModel struct {
	tableName struct{} `sql:"link"`
	Id        string   `sql:"id, pk, notnull"`
	Link      string   `sql:"link, notnull"`
}
