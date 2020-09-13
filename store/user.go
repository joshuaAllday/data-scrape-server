package store

var (
	insertUser = `INSERT INTO Users(email, password) VALUES($1, $2)`
)

func (db *DB) GetUser() {

}

func (db *DB) CreateUser(email string, password string) (*bool, error) {
	stmt, _ := db.Prepare(insertUser)
	_, err := stmt.Exec(email, password)
	valid := true
	if err != nil {
		return nil, err
	}

	return &valid, nil
}
