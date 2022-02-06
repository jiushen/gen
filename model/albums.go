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
{    "album_id": 99,    "title": "SFPSmfjOlMbCivLsyspujKVQd",    "artist_id": 55}



*/

// Albums struct is a row record of the albums table in the gormtest database
type Albums struct {
	//[ 0] AlbumId                                        uint(64)             null: false  primary: true   isArray: false  auto: false  col: uint            len: 64      default: []
	AlbumID uint32 `gorm:"primary_key;column:AlbumId;type:uint;size:64;" json:"album_id"`
	//[ 1] Title                                          string(160)          null: true   primary: false  isArray: false  auto: false  col: string          len: 160     default: []
	Title string `gorm:"column:Title;type:string;size:160;" json:"title"`
	//[ 2] ArtistId                                       int(32)              null: false  primary: false  isArray: false  auto: false  col: int             len: 32      default: []
	ArtistID int32 `gorm:"column:ArtistId;type:int;size:32;" json:"artist_id"`
}

// TableName sets the insert table name for this struct type
func (a *Albums) TableName() string {
	return "albums"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Albums) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Albums) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Albums) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
//func (a *Albums) TableInfo() *TableInfo {
//	return albumsTableInfo
//}
