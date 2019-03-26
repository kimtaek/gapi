package models

import "gapi/lib"

type PartnerAdminService struct {
	PartnerAdminId int
	Service        string
	ServiceId      int
	CreatedBy      int
	UpdatedBy      int
	DeletedBy      int
	lib.CudAtModel
}

func (_ *PartnerAdminService) WhereIndex(where PartnerAdminService) []PartnerAdminService {
	var data []PartnerAdminService
	lib.Slave().Where(&where).Find(&data)
	return data
}
