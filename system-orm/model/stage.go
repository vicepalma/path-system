package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
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


Table: stage
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] title                                          VARCHAR(250)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 250     default: []


JSON Sample
-------------------------------------
{    "id": 92,    "title": "vGHGVaPYBlrePlLxmnELFGpWe"}



*/

// Stage struct is a row record of the stage table in the path-system database
type Stage struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] title                                          VARCHAR(250)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 250     default: []
	Title sql.NullString `gorm:"column:title;type:VARCHAR;size:250;" json:"title"`
}

var stageTableInfo = &TableInfo{
	Name: "stage",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(250)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       250,
			GoFieldName:        "Title",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Stage) TableName() string {
	return "stage"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Stage) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Stage) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Stage) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Stage) TableInfo() *TableInfo {
	return stageTableInfo
}
