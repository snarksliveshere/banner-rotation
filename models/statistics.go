package models

type Statistics struct {
	TableName  struct{} `sql:"statistics"`
	Id         uint64
	AudienceId uint64 `sql:"audience_id,notnull,unique" pg:"unique:public_statistics_audience_banner_uidx"`
	BannerId   uint64 `sql:"banner_id,notnull,unique" pg:"unique:public_statistics_audience_banner_uidx"`
	Clicks     uint64 `sql:"clicks,use_zero,notnull"`
	Shows      uint64 `sql:"shows,use_zero,notnull"`
}
