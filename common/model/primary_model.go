package model

import "time"

type PrimaryModel struct {
	Id         int64     `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
