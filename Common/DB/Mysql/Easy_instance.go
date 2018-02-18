package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Instance_DB(Instancedata Model.Easyinstance, Uniqueid string) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into easydelivery (uniqueid,Productname,Deliveryaddress,Deliveryname,Deliverynumber,Message,Status,loginid ) values (?,?,?,?,?,?,?,?)", Uniqueid, Instancedata.Productname, Instancedata.Deliveryaddress, Instancedata.Deliveryname, Instancedata.Deliverynumber, Instancedata.Message, "Order", Instancedata.Loginid)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
	Mail(Instancedata, "Dummyaddress", "dummynumber")

	return "Success"

}
func EasyInstance() (Easyresponse []Model.Easyinstances) {

	var Data Model.Easyinstances
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  easydelivery")
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Productid,
			&Data.Productname,
			&Data.Uniqueid,
			&Data.Deliveryaddress,
			&Data.Deliveryname,
			&Data.Deliverynumber,
			&Data.Message,
			&Data.Status,
			&Data.Loginid,
		)
		Easyresponse = append(Easyresponse, Data)

	}

	return

}

func Instantundelivered_DB() (Easyresponse []Model.Easyinstances) {

	var Data Model.Easyinstances
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  easydelivery where Status =?", "Order")
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Productid,
			&Data.Productname,
			&Data.Uniqueid,

			&Data.Deliveryaddress,
			&Data.Deliveryname,
			&Data.Deliverynumber,
			&Data.Message,
			&Data.Status,
			&Data.Loginid,
		)
		Easyresponse = append(Easyresponse, Data)

	}

	return

}
func Updateinstant_DB(Data Model.Updateinstant) string {

	Queryupdate := "UPDATE easydelivery SET Status ='" + Data.Status + "' where easydeliveryid= " + fmt.Sprintf("%v", Data.Easyid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
	return "success"
}

func Gettrackinginstance(loginid string) (Easyresponse []Model.Easyinstances) {

	var Data Model.Easyinstances
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  easydelivery where loginid =?", loginid)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Productid,
			&Data.Productname,
			&Data.Uniqueid,
			&Data.Deliveryaddress,
			&Data.Deliveryname,
			&Data.Deliverynumber,
			&Data.Message,
			&Data.Status,
			&Data.Loginid,
		)
		Easyresponse = append(Easyresponse, Data)

	}

	return

}
