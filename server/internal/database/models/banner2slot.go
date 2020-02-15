package models

type Banner2Slot struct {
	TableName struct{} `sql:"banner2slot"`
	SlotFK    uint64   `sql:"slot_fk,notnull,unique" pg:"unique:public_banner2slot_slot_banner_uidx"`
	BannerFK  uint64   `sql:"banner_fk,notnull,unique" pg:"unique:public_banner2slot_slot_banner_uidx"`
}
