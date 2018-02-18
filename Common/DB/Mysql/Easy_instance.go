package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Instance_DB(Instancedata Model.Easyinstance, Uniqueid string) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into easydelivery (uniqueid,Productname,Deliveryaddress,Deliveryname,Deliverynumber,Message,Status,loginid ) values (?,?,?,?,?,?,?,?)", Uniqueid, Instancedata.Productname, Instancedata.Deliveryaddress, Instancedata.Deliveryname, Instancedata.Deliverynumber, Instancedata.Message, "Order", Instancedata.Loginid)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}

	sendkey := os.Getenv("SENDGRID_API_KEYGO")

	from := mail.NewEmail("E3 Shopping", "e3easyshopping@gmail.com")
	subject := "E3 NOTIFICATION - New Instant Order Recieved!"
	to := mail.NewEmail("E3 Admin", "e3easyshopping@gmail.com.com")
	plainTextContent := "e3easyshopping@gmail.com"
	htmlContent := "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
		"<div style='font-style:sans-serif'>NAME  - " + Instancedata.Deliveryname +
		"<div style='font-style:sans-serif'>ADDRESS - " + Instancedata.Deliveryaddress +
		"</div><div style='font-style:sans-serif'>NUMBER -" + Instancedata.Deliverynumber +
		"</div><div style='font-style:sans-serif'>PRODUCT  -" + Instancedata.Productname +
		"</div><div style='font-style:sans-serif'>MESSAGE" + Instancedata.Message +
		"</div></tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

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
