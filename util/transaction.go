package util

import (
	"database/sql"
)


func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		SendPanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		SendPanicIfError(errCommit)
	}
}