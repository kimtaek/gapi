package models

import (
	"gapi/lib"
)

type Code struct {
	Id       int
	Group    string
	Code     string
	TitleChs string
	TitleCht string
	TitleEng string
	TitleKor string
	Memo     string
	lib.CudAtModel
}

type CodeShow struct {
	Id       int `gorm:"primary_key"`
	Code     string
	TitleKor string
	TitleChs string
	TitleCht string
	TitleEng string
}

func (data *Code) Find(id int) *Code {
	lib.Slave().Where(&Code{Id: id}).First(data)
	return data
}

func (data *Code) WhereFind(where Code) *Code {
	lib.Slave().Where(&where).First(data)
	return data
}

func (_ *Code) WhereFindOptions(where Code) interface{} {
	type Code struct {
		lib.OptionModel
	}
	var data []Code
	lib.Slave().Where(&where).Select("id as `key`, title_chs as name, code as value").Find(&data)
	return data
}

func (_ *Code) Create(data Code) int {
	lib.Master().Create(data)
	return data.Id
}

func (data *Code) Update(id int, update Code) int64 {
	result := data.Find(id)
	if result.Id == 0 {
		return 0
	}
	rowsAffected := lib.Master().Model(result).Updates(update).RowsAffected
	return rowsAffected
}
