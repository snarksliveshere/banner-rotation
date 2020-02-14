package models

type Banner struct {
	TableName  struct{} `sql:"banner"`
	Id         uint64
	BannerId   string `sql:"banner_id,notnull,unique"`
	Audience   uint64
	AudienceId string
}
