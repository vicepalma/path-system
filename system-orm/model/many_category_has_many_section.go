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


Table: many_category_has_many_section
[ 0] id_category                                    INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] id_section                                     INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id_category": 47,    "id_section": 66}



*/

// ManyCategoryHasManySection struct is a row record of the many_category_has_many_section table in the path-system database
type ManyCategoryHasManySection struct {
	//[ 0] id_category                                    INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	IDCategory int32 `gorm:"primary_key;column:id_category;type:INT4;" json:"id_category"`
	//[ 1] id_section                                     INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	IDSection int32 `gorm:"primary_key;column:id_section;type:INT4;" json:"id_section"`
}

var many_category_has_many_sectionTableInfo = &TableInfo{
	Name: "many_category_has_many_section",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id_category",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "IDCategory",
			GoFieldType:        "int32",
			JSONFieldName:      "id_category",
			ProtobufFieldName:  "id_category",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "id_section",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "IDSection",
			GoFieldType:        "int32",
			JSONFieldName:      "id_section",
			ProtobufFieldName:  "id_section",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *ManyCategoryHasManySection) TableName() string {
	return "many_category_has_many_section"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *ManyCategoryHasManySection) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *ManyCategoryHasManySection) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *ManyCategoryHasManySection) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *ManyCategoryHasManySection) TableInfo() *TableInfo {
	return many_category_has_many_sectionTableInfo
}
