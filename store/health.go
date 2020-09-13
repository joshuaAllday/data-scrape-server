package store

func (db *DB) GetHealth() error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
