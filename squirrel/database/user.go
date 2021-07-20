package database

import (
	"db_connect/model"
	"git.aimap.io/go/logs"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"time"
)

func (mgr *Manager) InsertUsers(users []*model.User) error {
	// 将 users 参数转化为 [][]interface{} 的数据加在 sqlStr.Values 里
	values := getInsertUserValues(users)

	// PlaceholderFormat(sq.Dollar) 用来将默认的 ? 占位符替换成 $ ，是 pg 的格式
	sqlStr := sq.StatementBuilder.
		RunWith(mgr.db).
		Insert(UserTableName).
		Columns(
			UserInsertColumns...,
		).
		PlaceholderFormat(sq.Dollar)

	for _, v := range values {
		sqlStr = sqlStr.Values(v...)
	}

	logs.Debug(sqlStr.ToSql())
	_, err := sqlStr.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (mgr *Manager) QueryUsers(query *model.UserQuery) ([]*model.User, error) {
	sqlStr := sq.StatementBuilder.
		RunWith(mgr.db).
		Select(UserSelectColumns...).
		From(UserTableName).
		PlaceholderFormat(sq.Dollar)

	// 批量处理查询参数
	sqlStr = getQueryParams(query, sqlStr)

	logs.Debug(sqlStr.ToSql())

	sql, args, _ := sqlStr.ToSql()

	// 这里用 sqlx 查询会直接用 sq 查询方便
	var users []*model.User
	err := mgr.db.Select(&users, sql, args...)
	if err != nil {
		logs.Error(err)
	}

	return users, nil
}

func (mgr *Manager) UpdateUserByName(user *model.User) error {
	// 得到更新需要的 map 参数
	m := getUpdateUserMap(user)

	sqlStr := sq.StatementBuilder.
		RunWith(mgr.db).
		Update(UserTableName).
		PlaceholderFormat(sq.Dollar)

	sqlStr = sqlStr.SetMap(m).Where(sq.Eq{UserTableColName: m[UserTableColName]})

	logs.Debug(sqlStr.ToSql())
	_, err := sqlStr.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (mgr *Manager) DeleteUserByName(name string) error {
	m := make(map[string]interface{})
	m[UserTableColDeleted] = true

	sqlStr := sq.StatementBuilder.
		RunWith(mgr.db).
		Update(UserTableName).
		PlaceholderFormat(sq.Dollar)

	sqlStr = sqlStr.SetMap(m).Where(sq.Eq{UserTableColName: name})

	logs.Debug(sqlStr.ToSql())
	_, err := sqlStr.Exec()
	if err != nil {
		return err
	}
	return nil
}

func getQueryParams(query *model.UserQuery, sqlStr sq.SelectBuilder) sq.SelectBuilder {
	if len(query.Names) > 0 {
		args := getMultiParams(query.Names)
		sqlStr = sqlStr.Where(sq.Eq{UserTableColName: args})
	}
	if len(query.Emails) > 0 {
		args := getMultiParams(query.Emails)
		sqlStr = sqlStr.Where(sq.Eq{UserTableColEmail: args})
	}
	if len(query.PhoneNumbers) > 0 {
		args := getMultiParams(query.PhoneNumbers)
		sqlStr = sqlStr.Where(sq.Eq{UserTableColPhoneNumber: args})
	}
	if query.Age > 0 {
		sqlStr = sqlStr.Where(sq.GtOrEq{UserTableColAge: query.Age})
	}
	if len(query.MemberNumbers) > 0 {
		args := getMultiParams(query.MemberNumbers)
		sqlStr = sqlStr.Where(sq.Eq{UserTableColMemberNumber: args})
	}
	if !query.BirthDay.IsZero() {
		sqlStr = sqlStr.Where(sq.GtOrEq{UserTableColBirthday: query.BirthDay})
	}
	if !query.ActivatedAt.IsZero() {
		sqlStr = sqlStr.Where(sq.GtOrEq{UserTableColActivatedAt: query.ActivatedAt})
	}
	if query.Deleted {
		sqlStr = sqlStr.Where(sq.Eq{UserTableColDeleted: query.Deleted})
	}
	return sqlStr
}

func getMultiParams(items []string) (args []interface{}) {
	for i := range items {
		args = append(args, items[i])
	}
	return
}

func getInsertUserValues(users []*model.User) [][]interface{} {
	var res [][]interface{}

	for _, user := range users {
		var values []interface{}
		for _, column := range UserInsertColumns {
			switch column {
			case UserTableColId:
				if user.Id == "" {
					user.Id = uuid.NewString()
				}
				values = append(values, user.Id)
			case UserTableColName:
				values = append(values, user.Name)
			case UserTableColEmail:
				values = append(values, user.Email)
			case UserTableColPhoneNumber:
				values = append(values, user.PhoneNumber)
			case UserTableColAge:
				values = append(values, user.Age)
			case UserTableColMemberNumber:
				values = append(values, user.MemberNumber)
			case UserTableColBirthday:
				if user.BirthDay.IsZero() {
					user.BirthDay = time.Now()
				}
				values = append(values, user.BirthDay)
			case UserTableColActivatedAt:
				if user.ActivatedAt.IsZero() {
					user.ActivatedAt = time.Now()
				}
				values = append(values, user.ActivatedAt)
			}
		}
		res = append(res, values)
	}

	return res
}

func getUpdateUserMap(user *model.User) map[string]interface{} {
	updateMap := make(map[string]interface{})

	if user.Name != "" {
		updateMap[UserTableColName] = user.Name
	}
	if user.Email != "" {
		updateMap[UserTableColEmail] = user.Email
	}
	if user.PhoneNumber != "" {
		updateMap[UserTableColPhoneNumber] = user.PhoneNumber
	}
	if user.Age != 0 {
		updateMap[UserTableColAge] = user.Age
	}
	if user.MemberNumber != "" {
		updateMap[UserTableColMemberNumber] = user.MemberNumber
	}
	if !user.BirthDay.IsZero() {
		updateMap[UserTableColBirthday] = user.BirthDay
	}
	if !user.ActivatedAt.IsZero() {
		updateMap[UserTableColActivatedAt] = user.ActivatedAt
	}
	return updateMap
}
