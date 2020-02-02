package models

type Audience2Banner struct {
	TableName  struct{} `sql:"audience2banner"`
	Id         uint64
	AudienceId uint64 `sql:"audience_id,notnull,unique"`
	BannerId   uint64 `sql:"banner_id,notnull,unique"`
}
