package helper

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollBack := tx.Rollback()
		PanicIfError(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
