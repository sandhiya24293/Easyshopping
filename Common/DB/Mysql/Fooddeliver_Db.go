package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"

	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	_ "github.com/go-sql-driver/mysql"
)

func Foodinsert_DB(Instancedata Model.Fooddelivery) string {
	var stringprod []string

	row, err := OpenConnection["Rentmatics"].Exec("insert into fooddeliver (Hotelname,DeliverAddress,Delivernumber,Status,Loginid) values (?,?,?,?,?)", Instancedata.Hotelname, Instancedata.Address, Instancedata.Mobilenumber, "Order", Instancedata.Loginid)
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	Orderid, _ := row.LastInsertId()

	for _, Food := range Instancedata.Food {
		var getstring string
		rows, err := OpenConnection["Rentmatics"].Exec("insert into fooddelivername (foodid,foodname,foodqty) values (?,?,?)", Orderid, Food.Foodname, Food.Quantity)
		if err != nil {
			log.Println("Error -DB: Executive insert picture", err, rows)
		}
		getstring = "<tr style='font-style:sans-serif'><td>" + Food.Foodname + "</td><td>" + fmt.Sprintf("%v", Food.Quantity) + "</td></tr>"
		stringprod = append(stringprod, getstring)
	}

	result := strings.Join(stringprod, "")
	fmt.Println(stringprod)

	sendkey := os.Getenv("SENDGRID_API_KEYGO")

	from := mail.NewEmail("E3 Shopping", "e3easyshopping@gmail.com")
	subject := "E3 NOTIFICATION - New Order Recieved!"
	to := mail.NewEmail("E3 Admin", "e3easyshopping@gmail.com.com")
	plainTextContent := "e3easyshopping@gmail.com"
	htmlContent := "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
		"<div style='font-style:sans-serif'>LOGINID - " + Instancedata.Loginid +
		"<div style='font-style:sans-serif'>ADDRESS - " + Instancedata.Address +
		"</div><div style='font-style:sans-serif'>NUMBER -" + Instancedata.Mobilenumber +
		"</div><div style='font-style:sans-serif'>HOTEL NAME -" + Instancedata.Hotelname +

		"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >FOOD </th><th>RATE</th><th>QUANTITY</th></tr>" +
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

func UpdateOwnDelivered_DB(Orderid int) string {
	Status := "Delivered"
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

func Getallfooddelivery_DB() (Response []Model.Foodtrackreesponse) {
	Querystring := "SELECT * FROM fooddeliver"
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
func Getsingle_DB(foodid int) (Response Model.Foodtrackreesponse1) {

	var Data Model.Foodtrack
	rows, err := OpenConnection["Rentmatics"].Query("SELECT * FROM fooddeliver WHERE Foodid=?", foodid)
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
		row, _ := OpenConnection["Rentmatics"].Query("select fooddelid,foodname,foodqty from fooddelivername where foodid=?", Data.Foodid)

		var Productlist []Model.Foodcount1

		for row.Next() {
			var Productlist1 Model.Foodcount1

			row.Scan(
				&Productlist1.Foodid,
				&Productlist1.Foodname,
				&Productlist1.Quantity,
			)

			Productlist = append(Productlist, Productlist1)

		}
		Response.Details = Data
		Response.Fooddetails = Productlist

	}

	return

}

//OwnFood Function

func FoodOwn_DB(Instancedata Model.Ownfood) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into owndeliver (dishname,Rate,platecount,image,status) values (?,?,?,?,?)", Instancedata.Dishname, Instancedata.Rate, Instancedata.Platecount, Instancedata.Image, "show")
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	return "Success"

}

func GetFood_DB() (res Model.SendOwnfood) {

	var Data Model.Responseownfood
	var Foodres []Model.Responseownfood
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
			&Data.Image,
			&Data.Status,
		)
		Foodres = append(Foodres, Data)
		res.Data = Foodres

	}

	return

}

func Orderfooddeliver_DB(getorder Model.Ownorderplaced) string {
	var stringprod []string
	var Userid int
	var Number string
	var Address string

	row, err := OpenConnection["Rentmatics"].Exec("insert into ownfoodorder (Loginid,Itemscount,Totamount,Date,Status) values (?,?,?,?,?)", getorder.Loginid, getorder.Noofitems, getorder.TotalAmount, getorder.Date, "Order")
	if err != nil {
		log.Println("Error -DB: User", err, row)
	}

	Orderid, _ := row.LastInsertId()

	for _, Food := range getorder.Food {
		var getstring string
		rows, err := OpenConnection["Rentmatics"].Exec("insert into ownorderlist (Orderid,Dishname,Dishrate,platecount) values (?,?,?,?)", Orderid, Food.Dishname, Food.Rate, Food.Platecount)
		if err != nil {
			log.Println("Error -DB: Executive insert picture", err, rows)
		}
		getstring = "<tr style='font-style:sans-serif'><td>" + Food.Dishname + "</td><td>" + fmt.Sprintf("%v", Food.Rate) + "</td><td>" + fmt.Sprintf("%v", Food.Platecount) + "</td></tr>"
		stringprod = append(stringprod, getstring)
	}

	rows, err := OpenConnection["Rentmatics"].Query("select userid,phonenumber	 from  easylogin where Loginid=?", getorder.Loginid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Userid,
			&Number,
		)

	}
	rows1, err := OpenConnection["Rentmatics"].Query("select Adress	 from  easyprofile where easyuserid=?", Userid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows1.Next() {

		rows1.Scan(
			&Address,
		)

	}

	result := strings.Join(stringprod, "")
	fmt.Println(stringprod)

	sendkey := os.Getenv("SENDGRID_API_KEYGO")

	from := mail.NewEmail("E3 Shopping", "e3easyshopping@gmail.com")
	subject := "E3 NOTIFICATION - New Order Recieved!"
	to := mail.NewEmail("E3 Admin", "e3easyshopping@gmail.com.com")
	plainTextContent := "e3easyshopping@gmail.com"
	htmlContent := "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
		"<div style='font-style:sans-serif'>LOGINID - " + getorder.Loginid +
		"<div style='font-style:sans-serif'>ADDRESS - " + Address +
		"</div><div style='font-style:sans-serif'>NUMBER -" + Number +
		"</div><div style='font-style:sans-serif'>NO Of ITEMS -" + fmt.Sprintf("%v", getorder.Noofitems) +
		"</div><div style='font-style:sans-serif'>TOTAL AMOUNT -" + fmt.Sprintf("%v", getorder.TotalAmount) +

		"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >DISH </th><th>RATE</th><th>PLATECOUNT</th></tr>" +
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

	return "success"

}

func Getallundeliverfood_DB() (FinalResponse []Model.ResOwnorderplacedAll) {
	Querystring := "SELECT * FROM ownfoodorder "
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
		rows1, err := OpenConnection["Rentmatics"].Query("SELECT emailid,phonenumber FROM easylogin WHERE Loginid=?", Data.Loginid)
		if err != nil {
			log.Println("Error -Db:Activity", err)
		}
		for rows1.Next() {

			rows1.Scan(
				&Response.Emailid,
				&Response.Phonenumber,
			)
		}

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

func Getownundeliverfood_DB() (FinalResponse []Model.ResOwnorderplacedAll) {
	Querystring := "SELECT * FROM ownfoodorder WHERE Status IN ( 'Order','Dispatched')"
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
		rows1, err := OpenConnection["Rentmatics"].Query("SELECT emailid,phonenumber FROM easylogin WHERE Loginid=?", Data.Loginid)
		if err != nil {
			log.Println("Error -Db:Activity", err)
		}
		for rows1.Next() {

			rows1.Scan(
				&Response.Emailid,
				&Response.Phonenumber,
			)
		}

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

func Getownsinglefood_DB(Productid int) (Data Model.Responseownfood) {

	rows, err := OpenConnection["Rentmatics"].Query("select * from owndeliver where id=?", Productid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(

			&Data.ID,
			&Data.Dishname,
			&Data.Rate,
			&Data.Platecount,
			&Data.Image,
			&Data.Status,
		)

	}

	return
}
func OrderOwnsinglefood_DB(ownfoodid int) (FinalResponse Model.ResOwnorderplacedAll1) {
	var userid int
	var Data Model.ResOwnorderplaced
	rows, err := OpenConnection["Rentmatics"].Query("SELECT * FROM ownfoodorder WHERE Id=?", ownfoodid)
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
		rowss, err := OpenConnection["Rentmatics"].Query("SELECT userid,emailid,phonenumber FROM easylogin WHERE Loginid=?", Data.Loginid)
		if err != nil {
			log.Println("Error -Db:Activity", err)
		}
		for rowss.Next() {

			rowss.Scan(
				&userid,
				&FinalResponse.Emailid,
				&FinalResponse.Phonenumber,
			)
		}
		rows, err := OpenConnection["Rentmatics"].Query("SELECT Adress FROM easyprofile WHERE easyuserid=?", userid)
		if err != nil {
			log.Println("Error -Db:Activity", err)
		}
		for rows.Next() {

			rows.Scan(
				&FinalResponse.Address,
			)
		}
		row, _ := OpenConnection["Rentmatics"].Query("select Id,Dishname,Dishrate,platecount from ownorderlist where Orderid=?", Data.Id)

		var Productlist []Model.Ownfood1

		for row.Next() {
			var Productlist1 Model.Ownfood1

			row.Scan(
				&Productlist1.Id,
				&Productlist1.Dishname,
				&Productlist1.Rate,
				&Productlist1.Platecount,
			)

			Productlist = append(Productlist, Productlist1)

		}

		FinalResponse.Res = Data
		FinalResponse.Food = Productlist

	}

	return

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
