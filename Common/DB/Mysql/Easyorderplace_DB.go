package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Orderplaced_DB(Order Model.Orderplaced, Billid string) (Orderres Model.Placeorderresponse) {

	var Orderuser Model.Orderdata
	var Productdata Model.Product
	rows, err := OpenConnection["Rentmatics"].Query("select userid,uniqueid,emailid,phonenumber	 from  easylogin where Loginid=?", Order.Loginid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Orderuser.Userid,
			&Orderuser.Uniqueid,
			&Orderuser.Emailid,
			&Orderuser.Phonenumber,
		)

	}

	row, err := OpenConnection["Rentmatics"].Exec("insert into easyorderplaced (Billid,Uniqueid,userid,loginid,Phonenumber,Emailid,Totalproducts,TotalAmount,Date) values (?,?,?,?,?,?,?,?,?)", Billid, Orderuser.Uniqueid, Orderuser.Userid, Order.Loginid, Orderuser.Phonenumber, Orderuser.Emailid, Order.Noofproducts, Order.TotalAmount, Order.Date)
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	Orderid, _ := row.LastInsertId()

	for _, Productdata = range Order.Products {
		rows, err := OpenConnection["Rentmatics"].Exec("insert into productlist (orderplaceid,Productname,productrate,weight) values (?,?,?,?)", Orderid, Productdata.Product, Productdata.Rate, Productdata.Weight)
		if err != nil {
			log.Println("Error -DB: Executive insert picture", err, rows)
		}
	}

	row1, err := OpenConnection["Rentmatics"].Exec("insert into ordertracking (orderid,Orderplacedid,loginid,status) values (?,?,?,?)", Orderid, Billid, Order.Loginid, "Dispatched")
	if err != nil {
		log.Println("Error -DB: User", err, row1)
	}
	fmt.Println("Orderid", Orderid)
	row2, err := OpenConnection["Rentmatics"].Query("select Orderid,Billid  from  easyorderplaced where Orderid=?", Orderid)
	if err != nil {
		log.Println("Error -DB: Get User", err, row2)
	}
	for row2.Next() {

		row2.Scan(
			&Orderres.Orderid,
			&Orderres.Billid,
		)

	}

	return Orderres
}

func OrderTracking_DB(orderid int) (Trackingres Model.Trackingresponse) {

	rows, err := OpenConnection["Rentmatics"].Query("select *  from  ordertracking where orderid=?", orderid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Trackingres.Ordertrackingid,
			&Trackingres.Orderid,
			&Trackingres.Orderplacedid,
			&Trackingres.Loginid,
			&Trackingres.Status,
		)

	}
	return

}
func GetTracking_DB() (Tracking []Model.Trackingresponses) {

	var Trackingres Model.Trackingresponses

	var Userid int

	rows, err := OpenConnection["Rentmatics"].Query("select *  from  ordertracking WHERE status IN ('Order','Dispatched')")
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Trackingres.Ordertrackingid,
			&Trackingres.Orderid,
			&Trackingres.Orderplacedid,
			&Trackingres.Loginid,
			&Trackingres.Status,
		)

		row1, err := OpenConnection["Rentmatics"].Query("select userid,Phonenumber  from  easylogin WHERE Loginid=?", Trackingres.Loginid)
		if err != nil {
			log.Println("Error -DB: Get User", err)
		}
		for row1.Next() {

			row1.Scan(
				&Userid,
				&Trackingres.Phonenumber,
			)
		}
		row2, err := OpenConnection["Rentmatics"].Query("select Adress  from  easyprofile WHERE easyuserid=?", Userid)
		if err != nil {
			log.Println("Error -DB: Get User", err)
		}
		for row2.Next() {

			row2.Scan(
				&Trackingres.Address,
			)
			Tracking = append(Tracking, Trackingres)
		}

	}

	return

}

func GetTrackupdates_DB(Trackingres Model.Trackupdate) {

	Queryupdate := "UPDATE ordertracking SET status='" + Trackingres.Orderstatus + "'where Ordertrackingid=" + fmt.Sprintf("%v", Trackingres.Orderid)

	rows, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: Get User", err, rows)
	}

}

func OrderHistory_DB(Login string) (Response []Model.Orderhistory) {
	var historyresp Model.Orderhistory
	var Orderhis Model.OrderPlacedhistory
	rows, err := OpenConnection["Rentmatics"].Query("select Orderid,Billid,Totalproducts,TotalAmount,Date  from  easyorderplaced where loginid=?", Login)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Orderhis.Orderid,
			&Orderhis.Billid,
			&Orderhis.Totalproduct,
			&Orderhis.Totalamount,
			&Orderhis.Date,
		)

		row, err := OpenConnection["Rentmatics"].Query("select Productname,productrate,weight from productlist where orderplaceid=?", Orderhis.Orderid)
		fmt.Println("err", err)
		var Productlist []Model.Product

		for row.Next() {
			var Productlist1 Model.Product

			row.Scan(

				&Productlist1.Product,
				&Productlist1.Rate,
				&Productlist1.Weight)

			Productlist = append(Productlist, Productlist1)
		}
		historyresp.Orderdetails = Orderhis
		historyresp.Productdetails = Productlist
		Response = append(Response, historyresp)
	}

	return

}
