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


Table: category
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] name                                           VARCHAR(250)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 250     default: []


JSON Sample
-------------------------------------
{    "id": 35,    "name": "mTwOPRxyUxrraeTsJcNUvkJSk"}



*/

// Category struct is a row record of the category table in the path-system database
type Category struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] name                                           VARCHAR(250)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 250     default: []
	Name sql.NullString `gorm:"column:name;type:VARCHAR;size:250;" json:"name"`
}

var categoryTableInfo = &TableInfo{
	Name: "category",
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
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "category"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Category) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Category) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Category) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Category) TableInfo() *TableInfo {
	return categoryTableInfo
}
