package entity

type Token struct {
	Id         int    `gorm:"not null;primary_key:id;AUTO_INCREMENT" json:"-"`
	Uid        int    `gorm:"uid" json:"-"`
	Token      string `gorm:"token" json:"token"`
	Content    string `gorm:"content" json:"-"`
	Expire     int    `gorm:"expire" json:"expire"`
	UpdateTime int    `gorm:"update_time" json:"-"`
	CreateTime int    `gorm:"create_time" json:"-"`
}

func (Token) TableName() string {
	return "mnt_token"
}
