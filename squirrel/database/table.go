package database

const (
	UserTableName = "users"

	UserTableColId = "id"
	UserTableColName = "name"
	UserTableColEmail = "email"
	UserTableColPhoneNumber = "phone_number"
	UserTableColAge = "age"
	UserTableColBirthday = "birthday"
	UserTableColMemberNumber = "member_number"
	UserTableColActivatedAt = "activated_at"
	UserTableColDeleted = "deleted"
	UserTableColCreateTime = "create_time"
	UserTableColUpdateTime = "update_time"
)


var UserSelectColumns = []string {
	UserTableColId,
	UserTableColName,
	UserTableColEmail,
	UserTableColPhoneNumber,
	UserTableColAge,
	UserTableColBirthday,
	UserTableColMemberNumber,
	UserTableColActivatedAt,
	UserTableColDeleted,
	UserTableColCreateTime,
	UserTableColUpdateTime,
}

var UserInsertColumns = []string {
	UserTableColId,
	UserTableColName,
	UserTableColEmail,
	UserTableColPhoneNumber,
	UserTableColAge,
	UserTableColBirthday,
	UserTableColMemberNumber,
	UserTableColActivatedAt,
}


func (mgr *Manager) createUserTable() error {
	schema := `CREATE OR REPLACE FUNCTION update_timestamp() RETURNS 
		TRIGGER AS $$
		BEGIN
			new.update_time = current_timestamp;
			RETURN NEW;
		END
		$$ language plpgsql;

		CREATE TABLE IF NOT EXISTS ` + UserTableName + ` (
		` + UserTableColId + ` varchar(255) PRIMARY KEY  NOT NULL,
		` + UserTableColName + ` varchar(255) DEFAULT NULL,
		` + UserTableColEmail + ` varchar(255) DEFAULT NULL,
		` + UserTableColPhoneNumber + ` varchar(255) DEFAULT NULL,
		` + UserTableColAge + ` int NOT NULL,
		` + UserTableColBirthday + ` timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
		` + UserTableColMemberNumber + ` varchar(255) DEFAULT NULL,
		` + UserTableColActivatedAt + ` timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
		` + UserTableColDeleted + ` boolean DEFAULT false,
		` + UserTableColCreateTime + ` timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
		` + UserTableColUpdateTime + ` timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		DROP TRIGGER IF EXISTS ` + UserTableName + `_update_timestamp ON ` + UserTableName + `;
		CREATE TRIGGER ` + UserTableName + `_update_timestamp BEFORE UPDATE ON ` + UserTableName + ` for each row execute procedure update_timestamp();

		CREATE UNIQUE INDEX IF NOT EXISTS ` + UserTableName + `_` + UserTableColName + `_` + UserTableColDeleted + ` ON ` + UserTableName + `(` + UserTableColName + `,` + UserTableColDeleted + `);
		CREATE INDEX IF NOT EXISTS ` + UserTableName + `_` + UserTableColId + ` ON ` + UserTableName + `(` + UserTableColId + `);
		CREATE INDEX IF NOT EXISTS ` + UserTableName + `_` + UserTableColName + ` ON ` + UserTableName + `(` + UserTableColName + `);
		CREATE INDEX IF NOT EXISTS ` + UserTableName + `_` + UserTableColEmail + ` ON ` + UserTableName + `(` + UserTableColEmail + `);
		CREATE INDEX IF NOT EXISTS ` + UserTableName + `_` + UserTableColPhoneNumber + ` ON ` + UserTableName + `(` + UserTableColPhoneNumber + `);
		CREATE INDEX IF NOT EXISTS ` + UserTableName + `_` + UserTableColDeleted + ` ON ` + UserTableName + `(` + UserTableColDeleted + `);`

	_, err := mgr.db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
