package model

import "time"

type User struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Age          int       `json:"age"`
	BirthDay     time.Time `json:"birthday"`
	MemberNumber string    `json:"member_number"`
	ActivatedAt  time.Time `json:"activated_at"`
	Deleted      bool      `json:"deleted"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

// 为了批量处理查询请求，需要在上一层根据参数进行构造
type UserQuery struct {
	Names         []string  `json:"names"`
	Emails        []string  `json:"emails"`
	PhoneNumbers  []string  `json:"phone_numbers"`
	Age           int       `json:"age"`
	BirthDay      time.Time `json:"birthday"`
	MemberNumbers []string  `json:"member_numbers"`
	ActivatedAt   time.Time `json:"activated_at"`
	Deleted       bool      `json:"deleted"`
}
