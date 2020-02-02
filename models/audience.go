package models

type Audience struct {
	TableName  struct{} `sql:"audience"`
	Id         uint64
	AudienceId string `sql:"audience_id,notnull,unique"`
}
