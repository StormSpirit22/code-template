package database

import (
	"db_connect/model"
	"testing"
	"time"
)

var manager *Manager

func init() {
	var err error
	manager, err = NewManager()
	if err != nil {
		panic(err)
	}
}

func TestManager_InsertUser(t *testing.T) {
	var users []*model.User
	user := &model.User{
		Name:         "test3",
		Email:        "123@qq.com",
		PhoneNumber:  "123456",
		Age:          18,
		BirthDay:     time.Now(),
		MemberNumber: "abc",
		Deleted:      false,
	}
	user2 := &model.User{
		Name:         "test5",
		Email:        "123@qq.com",
		PhoneNumber:  "123456",
		Age:          18,
		BirthDay:     time.Now(),
		MemberNumber: "abc",
		Deleted:      false,
	}
	users = append(users, user, user2)
	err := manager.InsertUsers(users)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestManager_QueryUsers(t *testing.T) {
	userQuery := &model.UserQuery{
		Names:         []string{"test", "test5"},
		Emails:        []string{"123@qq.com", "123456@qq.com"},
		PhoneNumbers:  []string{"123456"},
		Age:           18,
		BirthDay:      time.Now().AddDate(0, 0, -1),
		MemberNumbers: nil,
		ActivatedAt:   time.Now().AddDate(0, 0, -1),
		Deleted:       false,
	}
	users, err := manager.QueryUsers(userQuery)
	if err != nil {
		t.Log(err.Error())
	}
	for _, user := range users {
		t.Log(user)
	}
}

func TestManager_UpdateUserByUserName(t *testing.T) {
	user := &model.User{
		Name:         "test3",
		Email:        "123456@qq.com",
		PhoneNumber:  "123456",
		Age:          18,
		BirthDay:     time.Now(),
		MemberNumber: "abc",
		Deleted:      false,
	}

	err := manager.UpdateUserByName(user)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestManager_DeleteUserByName(t *testing.T) {
	err := manager.DeleteUserByName("test3")
	if err != nil {
		t.Log(err.Error())
	}
}