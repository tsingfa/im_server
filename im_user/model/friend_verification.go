package model

import "im_server/common/model"

// FriendVerification 好友验证表
type FriendVerification struct {
	model.PrimaryModel
	SendUserId            int64                  `json:"sendUserId"`            // 发送方 userId
	RecvUserId            int64                  `json:"recvUserId"`            // 接收方 userId
	AdditionalMessage     string                 `json:"additionalMessage"`     // 附加消息
	Status                int8                   `json:"status"`                // 验证请求的处理状态
	VerificationQuestions *VerificationQuestions `json:"verificationQuestions"` // 验证问题
}
