package entity

type Account struct {
	Id         int    `gorm:"not null;primary_key:id;AUTO_INCREMENT" json:"id" form:"id"`
	Username   string `gorm:"column:username;not null;" json:"username" form:"username" binding:"required"`
	Password   string `gorm:"column:password;not null;" json:"-" form:"password"`
	Nickname   string `gorm:"column:nickname;not null;" json:"nickname" form:"nickname"`
	Avatar     string `gorm:"column:avatar;not null;" json:"avatar" form:"avatar"`
	Email      string `gorm:"column:email;not null;" json:"email" form:"email"`
	Ip         string `gorm:"column:ip;not null;" json:"ip" form:"ip"`
	Status     int    `gorm:"column:status;not null;" json:"status" form:"status"`
	UpdateTime int    `gorm:"column:update_time;" json:"-" form:"updateTime"`
	CreateTime int    `gorm:"column:create_time;not null;" json:"createTime" form:"createTime"`
}

func (Account) TableName() string {
	return "mnt_account"
}
