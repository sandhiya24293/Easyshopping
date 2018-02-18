package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
	"fmt"

	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Orderplaced_DB(Order Model.Orderplaced, Billid string) (Orderres Model.Placeorderresponse) {

	var Orderuser Model.Orderdata
	var Productdata Model.Product
	var Address string
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
	row3, err := OpenConnection["Rentmatics"].Query("select Adress  from  easyprofile where easyuserid=?", Orderuser.Userid)
	if err != nil {
		log.Println("Error -DB: Get User", err, row2)
	}
	for row3.Next() {

		row3.Scan(
			&Address,
		)

	}

	var stringprod []string

	for _, v := range Order.Products {
		var getstring string
		getstring = "<tr style='font-style:sans-serif'><td>" + v.Product + "</td><td>" + fmt.Sprintf("%v", v.Rate) + "</td><td>" + v.Weight + "</td></tr>"
		stringprod = append(stringprod, getstring)
	}

	result := strings.Join(stringprod, "")
	fmt.Println(stringprod)

	sendkey := os.Getenv("SENDGRID_API_KEYGO")

	from := mail.NewEmail("E3 Shopping", "e3easyshopping@gmail.com")
	subject := "E3 NOTIFICATION - New Order Recieved!"
	to := mail.NewEmail("E3 Admin", "e3easyshopping@gmail.com")
	plainTextContent := "e3easyshopping@gmail.com"
	htmlContent := "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
		"<div style='font-style:sans-serif'>LOGINID - " + Order.Loginid +
		"<div style='font-style:sans-serif'>ADDRESS - " + Address +
		"</div><div style='font-style:sans-serif'>DATE -" + Order.Date +
		"</div><div style='font-style:sans-serif'>TOTAL AMOUNT -" + fmt.Sprintf("%v", Order.TotalAmount) +
		"</div><div style='font-style:sans-serif'>NO OF PRODUCTS -" + fmt.Sprintf("%v", Order.Noofproducts) +
		"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >PRODUCT</th><th>RATE</th></tr>" +
		result + "</tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(sendkey)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)

	}

	return Orderres

}

func OrderTracking_DB(loginid string) (Trackingres1 []Model.Trackingresponse) {
	var Trackingres Model.Trackingresponse
	Querystring := "SELECT * FROM ordertracking WHERE status IN ('Dispatched', 'Order') && loginid ='" + loginid + "'"

	rows, err := OpenConnection["Rentmatics"].Query(Querystring)
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
		Trackingres1 = append(Trackingres1, Trackingres)

	}

	return

}

func OrderTrackingadmin_DB(orderid int) (Trackingres1 Model.ResponseOrderplaced) {
	var easyorderid int
	var userid int

	getrow, err := OpenConnection["Rentmatics"].Query("select orderid from ordertracking where Ordertrackingid=?", orderid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for getrow.Next() {

		getrow.Scan(
			&easyorderid,
		)
	}

	rows, err := OpenConnection["Rentmatics"].Query("select userid,Billid,Phonenumber,Totalproducts,TotalAmount,Date from easyorderplaced where Orderid=?", easyorderid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&userid,
			&Trackingres1.Billid,
			&Trackingres1.Phonenumber,
			&Trackingres1.Totalproduct,
			&Trackingres1.TotalAmount,
			&Trackingres1.Date,
		)

		Addr, err := OpenConnection["Rentmatics"].Query("select Adress from easyprofile where easyuserid=?", userid)
		fmt.Println("err", err)
		for Addr.Next() {

			Addr.Scan(
				&Trackingres1.Address,
			)
		}

		row, err := OpenConnection["Rentmatics"].Query("select Productid,Productname,productrate,weight from productlist where orderplaceid=?", orderid)
		fmt.Println("err", err)
		var Productlist []Model.Product1

		for row.Next() {
			var Productlist1 Model.Product1

			row.Scan(
				&Productlist1.Id,
				&Productlist1.Product,
				&Productlist1.Rate,
				&Productlist1.Weight)

			Productlist = append(Productlist, Productlist1)
		}

		Trackingres1.Products = Productlist

		return

	}

	return

}
func OrderCancel_DB(cancelorder Model.Ordercancel) string {
	Orderstatus := "Cancelled"
	Queryupdate := "UPDATE ordertracking SET status='" + Orderstatus + "'where orderid=" + fmt.Sprintf("%v", cancelorder.Orderid)

	rows, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: Get User", err, rows)

		return "invalid orderid"
	} else {
		return "Success"

	}

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

func GetTrackupdates_DB(Trackingres Model.Trackupdate) string {
	fmt.Println(Trackingres.Orderstatus, Trackingres.Orderid)

	Queryupdate := "UPDATE ordertracking SET status='" + Trackingres.Orderstatus + "'where Ordertrackingid=" + fmt.Sprintf("%v", Trackingres.Orderid)

	rows, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: Get User", err, rows)
	}
	return "Success"

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
