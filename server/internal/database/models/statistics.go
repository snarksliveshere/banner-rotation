package models

type Statistics struct {
	TableName  struct{} `sql:"statistics"`
	Id         uint64
	AudienceId string `sql:"audience_id,notnull" pg:"unique:public_statistics__uidx"`
	BannerId   string `sql:"banner_id,notnull" pg:"unique:public_statistics_uidx"`
	SlotId     string `sql:"slot_id,notnull" pg:"unique:public_statistics_uidx"`
	Clicks     uint64 `sql:"clicks,use_zero,notnull"`
	Shows      uint64 `sql:"shows,use_zero,notnull"`
}
