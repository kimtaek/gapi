package models

import (
	"gapi/lib"
	"time"
)

type Admin struct {
	Id            int
	Name          string
	Email         string
	Password      string
	Status        string
	RememberToken string
	Sex           string
	Birthday      time.Time
	Mobile        string
	Phone         string
	Address       string
	Avatar        string
	CreatedBy     int
	UpdatedBy     int
	DeletedBy     int
	lib.CudAtModel
}

func (data *Admin) Find(id int) *Admin {
	lib.Slave().Where(&Admin{Id: id}).First(data)
	return data
}

func (data *Admin) FindByEmail(email string) *Admin {
	lib.Slave().Where(&Admin{Email: email}).First(data)
	return data
}

func (data *Admin) WhereFind(where Admin) *Admin {
	lib.Slave().Where(&where).First(data)
	return data
}

func (_ *Admin) Create(data Admin) int {
	lib.Master().Create(data)
	return data.Id
}

func (data *Admin) Update(id int, update Admin) int64 {
	result := data.Find(id)
	if result.Id == 0 {
		return 0
	}
	rowsAffected := lib.Master().Model(result).Updates(update).RowsAffected
	return rowsAffected
}
