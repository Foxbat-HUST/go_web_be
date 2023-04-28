package gen

import "gorm.io/gorm"

//******************************************************//
//**This file is generated by `make gen_model` command**//
//************Do not edit this file manually************//
//******************************************************//

type Job struct {

	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`

	//[ 1] name                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Name *string `gorm:"column:name;type:varchar;size:30;"`

	//[ 2] description                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Description *string `gorm:"column:description;type:varchar;size:30;"`
}

var jobsTableInfo = &TableInfo{
	Name: "jobs",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
		},

		{
			Index:              1,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Name",
			GoFieldType:        "*string",
		},

		{
			Index:              2,
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Description",
			GoFieldType:        "*string",
		},

		{
			Index:              3,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "*time.Time",
		},

		{
			Index:              4,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "*time.Time",
		},

		{
			Index:              5,
			Name:               "deleted_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "DeletedAt",
			GoFieldType:        "*time.Time",
		},
	},
}

// TableName sets the insert table name for this struct type
func (j *Job) TableName() string {
	return "jobs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (j *Job) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (j *Job) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (j *Job) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (j *Job) TableInfo() *TableInfo {
	return jobsTableInfo
}

// Columns return map of column
func (j Job) Columns() map[string]bool {
	var result = map[string]bool{
		"id":          true,
		"name":        true,
		"description": true,
		"created_at":  true,
		"updated_at":  true,
		"deleted_at":  true,
	}
	return result
}
