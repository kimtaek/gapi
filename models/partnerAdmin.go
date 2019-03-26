package models

import (
	"gapi/lib"
)

type PartnerAdmin struct {
	Id        int `json:"id"`
	HotelId   int
	Name      *string `json:"name"`
	Email     string  `json:"email"`
	Mobile    *string `json:"mobile"`
	Password  string  `json:"password"`
	CreatedBy int
	UpdatedBy int
	DeletedBy int
	lib.CudAtModel
}

func (_ *PartnerAdmin) Index() []PartnerAdmin {
	var data []PartnerAdmin
	lib.Slave().Select("*").Order("id desc").Find(&data)
	return data
}

func (_ *PartnerAdmin) WhereIndex(where PartnerAdmin) []PartnerAdmin {
	var data []PartnerAdmin
	lib.Slave().Where(&where).Order("id desc").Find(&data)
	return data
}

func (data *PartnerAdmin) Find(id int) *PartnerAdmin {
	lib.Slave().Where(&PartnerAdmin{Id: id}).First(data)
	return data
}

func (data *PartnerAdmin) FindByEmail(email string) *PartnerAdmin {
	lib.Slave().Where(&PartnerAdmin{Email: email}).First(data)
	return data
}

func (data *PartnerAdmin) WhereFind(where PartnerAdmin) *PartnerAdmin {
	lib.Slave().Where(&where).First(data)
	return data
}

func (_ *PartnerAdmin) Create(data PartnerAdmin) int {
	lib.Master().Create(data)
	return data.Id
}

func (data *PartnerAdmin) Update(id int, update PartnerAdmin) int64 {
	result := data.Find(id)
	if result.Id == 0 {
		return 0
	}
	rowsAffected := lib.Master().Model(result).Updates(update).RowsAffected
	return rowsAffected
}

func (data *PartnerAdmin) Delete(id int) int64 {
	hotel := data.Find(id)
	if hotel.Id == 0 {
		return 0
	}
	rowsAffected := lib.Master().Delete(hotel).RowsAffected
	return rowsAffected
}
