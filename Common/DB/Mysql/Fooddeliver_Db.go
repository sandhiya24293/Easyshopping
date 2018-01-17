package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Foodinsert_DB(Instancedata Model.Fooddelivery) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into fooddeliver (Hotelname,DeliverAddress,Delivernumber,Status,Loginid) values (?,?,?,?,?)", Instancedata.Hotelname, Instancedata.Address, Instancedata.Mobilenumber, "Order", Instancedata.Loginid)
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	Orderid, _ := row.LastInsertId()

	for _, Food := range Instancedata.Food {
		rows, err := OpenConnection["Rentmatics"].Exec("insert into fooddelivername (foodid,foodname,foodqty) values (?,?,?)", Orderid, Food.Foodname, Food.Quantity)
		if err != nil {
			log.Println("Error -DB: Executive insert picture", err, rows)
		}
	}
	return "Success"

}

//func EasyInstance() (Easyresponse []Model.Easyinstances) {

//	var Data Model.Easyinstances
//	rows, err := OpenConnection["Rentmatics"].Query("Select * from  easydelivery")
//	if err != nil {
//		log.Println("Error -Db:Activity", err)
//	}
//	for rows.Next() {

//		rows.Scan(
//			&Data.Productid,
//			&Data.Uniqueid,
//			&Data.Productname,
//			&Data.Dispactchername,
//			&Data.Dispatcheraddess,
//			&Data.Dispatchernumber,
//			&Data.Deliveryaddress,
//			&Data.Deliveryname,
//			&Data.Deliverynumber,
//			&Data.Message,
//			&Data.Status,
//		)
//		Easyresponse = append(Easyresponse, Data)

//	}

//	return

//}

//func Instantundelivered_DB() (Easyresponse []Model.Easyinstances) {

//	var Data Model.Easyinstances
//	rows, err := OpenConnection["Rentmatics"].Query("Select * from  easydelivery where Status =?", "Order")
//	if err != nil {
//		log.Println("Error -Db:Activity", err)
//	}
//	for rows.Next() {

//		rows.Scan(
//			&Data.Productid,
//			&Data.Uniqueid,
//			&Data.Productname,
//			&Data.Dispactchername,
//			&Data.Dispatcheraddess,
//			&Data.Dispatchernumber,
//			&Data.Deliveryaddress,
//			&Data.Deliveryname,
//			&Data.Deliverynumber,
//			&Data.Message,
//			&Data.Status,
//			&Data.Loginid,
//		)
//		Easyresponse = append(Easyresponse, Data)

//	}

//	return

//}
func UpdateDelivered_DB(Orderid int) string {
	Status := "Delivered"
	Queryupdate := "UPDATE fooddeliver SET Status ='" + Status + "' where Foodid= " + fmt.Sprintf("%v", Orderid)
	fmt.Println(Queryupdate)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
		return "failure"

	} else {
		return "success"

	}

}

func Updatecancel_DB(Orderid int) string {
	Status := "Cancelled"
	Queryupdate := "UPDATE fooddeliver SET Status ='" + Status + "' where Foodid= " + fmt.Sprintf("%v", Orderid)
	fmt.Println(Queryupdate)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
		return "failure"

	} else {
		return "success"

	}

}

func GetFooddeliver_DB(loginid string) (Response []Model.Foodtrackreesponse) {
	Querystring := "SELECT * FROM fooddeliver WHERE Status IN ( 'Order','Dispatched') && loginid ='" + loginid + "'"
	var tempres Model.Foodtrackreesponse

	var Data Model.Foodtrack
	rows, err := OpenConnection["Rentmatics"].Query(Querystring)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Foodid,
			&Data.Hotelname,
			&Data.Address,
			&Data.Mobilenumber,
			&Data.Status,
			&Data.Loginid,
		)
		row, _ := OpenConnection["Rentmatics"].Query("select foodname,foodqty from fooddelivername where foodid=?", Data.Foodid)

		var Productlist []Model.Foodcount

		for row.Next() {
			var Productlist1 Model.Foodcount

			row.Scan(

				&Productlist1.Foodname,
				&Productlist1.Quantity,
			)

			Productlist = append(Productlist, Productlist1)

		}
		tempres.Details = Data
		tempres.Fooddetails = Productlist
		Response = append(Response, tempres)

	}

	return

}
func Gethistory_DB(loginid string) (Response []Model.Foodtrackreesponse) {
	var tempres Model.Foodtrackreesponse

	var Data Model.Foodtrack
	rows, err := OpenConnection["Rentmatics"].Query("SELECT * FROM fooddeliver WHERE  loginid =?", loginid)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Foodid,
			&Data.Hotelname,
			&Data.Address,
			&Data.Mobilenumber,
			&Data.Status,
			&Data.Loginid,
		)
		row, _ := OpenConnection["Rentmatics"].Query("select foodname,foodqty from fooddelivername where foodid=?", Data.Foodid)

		var Productlist []Model.Foodcount

		for row.Next() {
			var Productlist1 Model.Foodcount

			row.Scan(

				&Productlist1.Foodname,
				&Productlist1.Quantity,
			)

			Productlist = append(Productlist, Productlist1)

		}
		tempres.Details = Data
		tempres.Fooddetails = Productlist
		Response = append(Response, tempres)

	}

	return

}
func Getundelivered() (Response []Model.Foodtrackreesponse) {
	Querystring := "SELECT * FROM fooddeliver WHERE Status IN ( 'Order','Dispatched')"
	var tempres Model.Foodtrackreesponse

	var Data Model.Foodtrack
	rows, err := OpenConnection["Rentmatics"].Query(Querystring)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Foodid,
			&Data.Hotelname,
			&Data.Address,
			&Data.Mobilenumber,
			&Data.Status,
			&Data.Loginid,
		)
		row, _ := OpenConnection["Rentmatics"].Query("select foodname,foodqty from fooddelivername where foodid=?", Data.Foodid)

		var Productlist []Model.Foodcount

		for row.Next() {
			var Productlist1 Model.Foodcount

			row.Scan(

				&Productlist1.Foodname,
				&Productlist1.Quantity,
			)

			Productlist = append(Productlist, Productlist1)

		}
		tempres.Details = Data
		tempres.Fooddetails = Productlist
		Response = append(Response, tempres)

	}

	return

}
