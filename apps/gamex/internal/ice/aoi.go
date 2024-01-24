package ice

type IAoi interface {
	SetAoiId(aoiId int64)
	GetAoiId() (aoiId int64)
	GetAoiIdStr() (aoiIdStr string)
}

type IAoiManager interface {
	NewAoi(id int64) (aoi IAoi)
	AddAoi(aoi IAoi)
	GetAoiByAoiId(aoiId int64) (aoi IAoi, err error)
	GetAoiByAoiIdStr(aoiIdStr string) (aoi IAoi, err error)
	RemoveAoi(aoi IAoi)
}
