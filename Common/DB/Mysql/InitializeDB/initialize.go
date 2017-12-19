package InitializeDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var OpenConnection = make(map[string]*sql.DB)

func init() {
	Rentmatics, err := sql.Open("mysql", "motoskia_rmlist:RENTMATICS2017@tcp(143.95.248.120:3306)/motoskia_easyshopping?charset=utf8")
	//Rentmatics, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/easyshopping?charset=utf8")

	if err != nil {
		fmt.Println("error", err)
	}

	OpenConnection["Rentmatics"] = Rentmatics

}
func Ret() map[string]*sql.DB {
	return OpenConnection
}
