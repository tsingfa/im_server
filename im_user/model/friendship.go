package model

import "im_server/common/model"

// Friendship 好友关系表
type Friendship struct {
	model.PrimaryModel
	UserId       int64  `json:"userId"`
	FriendUserId int64  `json:"friendUserId"`
	FriendRemark string `json:"friendRemark"` // 好友备注
	Status       int8   `json:"status"`       // 好友关系状态
}
