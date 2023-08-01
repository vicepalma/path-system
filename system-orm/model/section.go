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


Table: section
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
[ 1] id_stage                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] title                                          VARCHAR(250)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 250     default: []
[ 3] description                                    TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 97,    "id_stage": 17,    "title": "VGAapPGvaIvHxbnBSGSWsfcHM",    "description": "PMnZRhaSICpnvogEZoeHocAAT"}



*/

// Section struct is a row record of the section table in the path-system database
type Section struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] id_stage                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	IDStage sql.NullInt64 `gorm:"column:id_stage;type:INT4;" json:"id_stage"`
	//[ 2] title                                          VARCHAR(250)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 250     default: []
	Title string `gorm:"column:title;type:VARCHAR;size:250;" json:"title"`
	//[ 3] description                                    TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Description string `gorm:"column:description;type:TEXT;" json:"description"`
}

var sectionTableInfo = &TableInfo{
	Name: "section",
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
			Name:               "id_stage",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "IDStage",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "id_stage",
			ProtobufFieldName:  "id_stage",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(250)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       250,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Section) TableName() string {
	return "section"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Section) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Section) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Section) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Section) TableInfo() *TableInfo {
	return sectionTableInfo
}
