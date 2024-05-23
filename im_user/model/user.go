package model

import (
	"im_server/common/model"
	"time"
)

type User struct {
	model.PrimaryModel
	Password string `json:"password"` // 密码
	// 基本信息
	Username string    `json:"username"` // 用户名
	Nickname string    `json:"nickname"` // 昵称
	Gender   int8      `json:"gender"`   // 性别
	Email    string    `json:"email"`    // 电子邮箱
	Phone    string    `json:"phone"`    // 手机号
	IP       string    `json:"ip"`       // ip地址
	Addr     string    `json:"addr"`     // 地址
	Birthday time.Time `json:"birthday"` // 出生日期
	// 个性化信息
	Signature   string `json:"signature"`   // 个性签名
	Description string `json:"description"` // 简介
	Avatar      string `json:"avatar"`      // 头像
}
