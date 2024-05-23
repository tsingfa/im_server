package model

import "im_server/common/model"

type UserConf struct {
	model.PrimaryModel
	UserId int64 `json:"user_id"` // 用户 id
	// 基本设置
	OnlineState           int8    `json:"online_state"`          // 在线状态
	WithdrawMessagePrompt *string `json:"withdrawMessagePrompt"` // 撤回消息的提示信息
	FriendOnline          bool    `json:"friendOnline"`          // 好友上线提醒
	Mute                  bool    `json:"mute"`                  // 静音
	// 安全设置
	SecureLink   bool `json:"secureLink"`   // 安全链接
	SavePassword bool `json:"savePassword"` // 记住密码
	// 权限设置
	SearchPermission      int8                  `json:"searchPermission"`   // 允许被搜索到的方式
	FriendVerification    int8                  `json:"friendVerification"` // 好友验证
	VerificationQuestions VerificationQuestions `json:"verificationQuestions"`
}

type VerificationQuestions struct {
	Question1 *string `json:"question1"`
	Question2 *string `json:"question2"`
	Question3 *string `json:"question3"`
	Answer1   *string `json:"answer1"`
	Answer2   *string `json:"answer2"`
	Answer3   *string `json:"answer3"`
}
