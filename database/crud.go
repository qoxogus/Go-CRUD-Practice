package database

// User - 유저 관리 테이블
type Crud struct {
	Idx       uint   `gorm:"primary_key; auto_increment:true" json:"idx"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	ID        string `gorm:"type:varchar(255);not null" json:"id"`
	PW        string `gorm:"type:varchar(255);not null" json:"pw"`
	IsManager bool   `gorm:"not null" sql:"DEFAULT:false" json:"ismanager"`
}
