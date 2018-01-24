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

//OwnFood Function

func FoodOwn_DB(Instancedata Model.Ownfood) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into owndeliver (dishname,Rate,platecount,status) values (?,?,?,?)", Instancedata.Dishname, Instancedata.Rate, Instancedata.Platecount, "show")
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	return "Success"

}

func GetFood_DB() (Foodres []Model.Responseownfood) {

	var Data Model.Responseownfood
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  owndeliver")
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.ID,
			&Data.Dishname,
			&Data.Rate,
			&Data.Platecount,
			&Data.Status,
		)
		Foodres = append(Foodres, Data)

	}

	return

}

func Orderfooddeliver_DB(getorder Model.Ownorderplaced) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into ownfoodorder (Loginid,Itemscount,Totamount,Date,Status) values (?,?,?,?,?)", getorder.Loginid, getorder.Noofitems, getorder.TotalAmount, getorder.Date, "Order")
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	Orderid, _ := row.LastInsertId()

	for _, Food := range getorder.Food {
		rows, err := OpenConnection["Rentmatics"].Exec("insert into ownorderlist (Orderid,Dishname,Dishrate,platecount) values (?,?,?,?)", Orderid, Food.Dishname, Food.Rate, Food.Platecount)
		if err != nil {
			log.Println("Error -DB: Executive insert picture", err, rows)
		}
	}

	return "success"

}

func TrackOwnfood_DB(loginid string) (FinalResponse []Model.ResOwnorderplacedAll) {
	Querystring := "SELECT * FROM ownfoodorder WHERE Status IN ( 'Order','Dispatched') && Loginid ='" + loginid + "'"
	var Response Model.ResOwnorderplacedAll
	var Data Model.ResOwnorderplaced
	rows, err := OpenConnection["Rentmatics"].Query(Querystring)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Id,
			&Data.Loginid,
			&Data.Noofitems,
			&Data.TotalAmount,
			&Data.Date,
			&Data.Status,
		)

		row, _ := OpenConnection["Rentmatics"].Query("select Dishname,Dishrate,platecount from ownorderlist where Orderid=?", Data.Id)

		var Productlist []Model.Ownfood

		for row.Next() {
			var Productlist1 Model.Ownfood

			row.Scan(

				&Productlist1.Dishname,
				&Productlist1.Rate,
				&Productlist1.Platecount,
			)

			Productlist = append(Productlist, Productlist1)

		}

		Response.Res = Data
		Response.Food = Productlist
		FinalResponse = append(FinalResponse, Response)

	}

	return

}
func CancelOwnFooddeliver_DB(Orderid int) string {
	Status := "Cancelled"
	Queryupdate := "UPDATE ownfoodorder SET Status ='" + Status + "' where Id= " + fmt.Sprintf("%v", Orderid)
	fmt.Println(Queryupdate)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
		return "failure"

	} else {
		return "success"

	}

}
