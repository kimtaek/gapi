package models

import (
	"gapi/lib"
)

type Hotel struct {
	Id             int `gorm:"primary_key"`
	Brand          int
	City           int
	Zone           int
	Star           int
	Use            int
	CheckInPolicy  string
	CheckOutPolicy string
	Provider       int
	TitleChs       string
	TitleCht       string
	TitleEng       string
	TitleKor       string
	AddressChs     string
	AddressCht     string
	AddressEng     string
	AddressKor     string
	DescriptionChs string
	DescriptionCht string
	DescriptionEng string
	DescriptionKor string
	CreatedBy      uint64
	UpdatedBy      uint64
	DeletedBy      uint64
	lib.CudAtModel
}

func (_ *Hotel) Index() []Hotel {
	var data []Hotel
	lib.Slave().Select("*").Order("id desc").Find(&data)
	return data
}

func (_ *Hotel) WhereIndex(where Hotel) []Hotel {
	var data []Hotel
	lib.Slave().Where(&where).Order("id desc").Find(&data)
	return data
}

func (data *Hotel) Find(id int) *Hotel {
	lib.Slave().Where(&Hotel{Id: id}).First(data)
	return data
}

func (data *Hotel) WhereFind(where Hotel) *Hotel {
	lib.Slave().Where(&where).First(data)
	return data
}

func (_ *Hotel) Create(data Hotel) int {
	lib.Master().Create(data)
	return data.Id
}

func (data *Hotel) Update(id int, update Hotel) int64 {
	hotel := data.Find(id)
	if hotel.Id == 0 {
		return 0
	}
	rowsAffected := lib.Master().Model(hotel).Updates(update).RowsAffected
	return rowsAffected
}

func (data *Hotel) Delete(id int) int64 {
	hotel := data.Find(id)
	if hotel.Id == 0 {
		return 0
	}
	rowsAffected := lib.Master().Delete(hotel).RowsAffected
	return rowsAffected
}
