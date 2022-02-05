package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


dummy ddl

JSON Sample
-------------------------------------
{    "artist_id": 69,    "name": "qANuKGlLkIeCNSLsSTKwgSXfJ"}



*/

// Artists struct is a row record of the artists table in the gormtest database
type Artists struct {
	//[ 0] ArtistId                                       int(32)              null: false  primary: true   isArray: false  auto: false  col: int             len: 32      default: []
	ArtistID int32 `gorm:"primary_key;column:ArtistId;type:int;size:32;" json:"artist_id"`
	//[ 1] Name                                           string               null: true   primary: false  isArray: false  auto: false  col: string          len: 0       default: []
	Name null.String `gorm:"column:Name;type:string;" json:"name"`
}

// TableName sets the insert table name for this struct type
func (a *Artists) TableName() string {
	return "artists"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Artists) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Artists) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Artists) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
//func (a *Artists) TableInfo() *TableInfo {
//	return artistsTableInfo
//}
