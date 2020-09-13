package store

var (
	insertUserOauthToken = `UPDATE Users SET Token = $1 WHERE ID = $2`
)

func (db *DB) AddUserOauthToken(id string, token string) error {
	stmt, _ := db.Prepare(insertUserOauthToken)
	_, err := stmt.Exec(id, token)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteAuthToken(id string) error {
	stmt, _ := db.Prepare(insertUserOauthToken)
	_, err := stmt.Exec(id, nil)

	if err != nil {
		return err
	}

	return nil
}
