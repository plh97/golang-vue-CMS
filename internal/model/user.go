package model

import "gorm.io/plugin/soft_delete"

type User struct {
	ID              int                   `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`                 // 主键ID
	Name            string                `gorm:"column:name;type:varchar(255);not null" json:"name"`                         // 用户姓名
	Gender          int                   `gorm:"column:gender;type:tinyint;not null;default 0" json:"gender"`                // 性别 1.男 2.女
	Udid            string                `gorm:"column:udid;type:varchar(255);not null" json:"udid"`                         // 设备唯一标识
	UserType        int                   `gorm:"column:user_type;type:tinyint;not null" json:"user_type"`                    // 用户类型 1访客用户 2注册用户
	IsSubmitProfile int                   `gorm:"column:is_submit_profile;type:tinyint;not null" json:"is_submit_profile"`    // 是否已经上传过个人资料 1.是 2.否
	Status          int                   `gorm:"column:status;type:tinyint;not null" json:"status"`                          // 状态 0未知状态 1正常 2禁用 3注销
	CountryID       int                   `gorm:"column:country_id;type:int;not null" json:"country_id"`                      // 国家ID
	RegisterIP      string                `gorm:"column:register_ip;type:varchar(45);not null" json:"register_ip"`            // 注册IP
	RegisterTime    int                   `gorm:"column:register_time;type:int;not null" json:"register_time"`                // 注册时间
	RegisterType    int                   `gorm:"column:register_type;type:tinyint;not null" json:"register_type"`            // 注册类型 0游客 1微信服务号 2微信小程序
	LastLoginIP     string                `gorm:"column:last_login_ip;type:varchar(45);not null" json:"last_login_ip"`        // 最后登录IP
	LastLoginTime   int                   `gorm:"column:last_login_time;type:int;not null" json:"last_login_time"`            // 最后登录时间
	LastLoginType   int                   `gorm:"column:last_login_type;type:tinyint;not null" json:"last_login_type"`        // 最后登录类型 0游客 1微信服务号 2微信小程序
	DeactivateTime  int                   `gorm:"column:deactivate_time;type:int;not null" json:"deactivate_time"`            // 注销时间
	IsFake          int                   `gorm:"column:is_fake;type:tinyint;not null;default 2" json:"is_fake"`              // 是否为假用户 1.是 2.否
	CreateTime      int                   `gorm:"column:create_time;type:int(11);not null;default 0;autoCreateTime" json:"-"` // 创建时间
	UpdateTime      int                   `gorm:"column:update_time;type:int(11);not null;default 0;autoUpdateTime" json:"-"` // 更新时间
	DeleteTime      soft_delete.DeletedAt `gorm:"column:delete_time;type:int(11);not null;default 0;" json:"-"`               // 删除时间
	UserId          string                `gorm:"column:user_id;unique;not null"`
	Nickname        string                `gorm:"column:Nickname;not null"`
	Password        string                `gorm:"column:Password;not null"`
	Email           string                `gorm:"column:Email;not null"`
}

func (u *User) TableName() string {
	return "users"
}
