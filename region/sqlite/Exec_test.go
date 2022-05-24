package sqlite

import (
	"context"
	"testing"
)

func TestExec(t *testing.T) {

	//Exec("CREATE TABLE `userinfo` (`till_id` INTEGER PRIMARY KEY AUTOINCREMENT, `client_id` VARCHAR(64) NULL, `first_name` VARCHAR(255) NOT NULL, `last_name` VARCHAR(255) NOT NULL, `guid` VARCHAR(255) NULL, `dob` DATETIME NULL, `type` VARCHAR(1))")
	//rows, err := Query("SELECT * FROM userinfo")
	//fmt.Println(rows, err)
	Exec("INSERT INTO userinfo(client_id, first_name, last_name) values(1,'1','1')", "userinfo")
}

func TestBack(t *testing.T) {
	ctx := context.Background()
	BackupSQLite3(ctx)
}
