package models

type Audience2Banner struct {
	TableName  struct{} `sql:"audience2banner"`
	Id         uint64
	AudienceFK uint64 `sql:"audience_fk,notnull,unique"`
	BannerFK   uint64 `sql:"banner_fk,notnull,unique"`
}
