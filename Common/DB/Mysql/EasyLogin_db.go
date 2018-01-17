package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	_ "github.com/go-sql-driver/mysql"
)

func InsertnewUser(Userdata Model.EasyRegister, randomno string) (Userinfo Model.UserResponse) {
	var Count int

	query := "SELECT COUNT(*) FROM easylogin WHERE easylogin.Loginid ='" + Userdata.Loginid + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		row, err := OpenConnection["Rentmatics"].Exec("insert into easylogin (uniqueid,username,Loginid,password,phonenumber,emailid,logintype) values (?,?,?,?,?,?,?)", randomno, Userdata.Username, Userdata.Loginid, Userdata.Password, Userdata.Phonenumber, Userdata.Emailid, Userdata.Logintype)
		if err != nil {
			log.Println("Error -DB: User", err, row)
		}
		rows, err := OpenConnection["Rentmatics"].Query("select username,Loginid,phonenumber from easylogin where Loginid=?", Userdata.Loginid)
		if err != nil {
			log.Println("Error -DB: User", err)
		}
		for rows.Next() {

			rows.Scan(

				&Userinfo.Username,
				&Userinfo.Loginid,
				&Userinfo.Phonenumber,
			)

		}
		Userinfo.Status = "Success"
		return Userinfo
	} else {

		rows, _ := OpenConnection["Rentmatics"].Query("select username,Loginid,phonenumber from easylogin where Loginid=?", Userdata.Loginid)
		for rows.Next() {

			rows.Scan(
				&Userinfo.Username,
				&Userinfo.Loginid,
				&Userinfo.Phonenumber,
			)

		}

		Userinfo.Status = "Already_Exist"
		return Userinfo

	}
}

func LoginUser(User1 Model.EasyLogin) (Userinfo Model.UserResponse) {
	var Getuser Model.Easyloginverify
	rows, err := OpenConnection["Rentmatics"].Query("select username,Loginid,password,phonenumber from  easylogin where Loginid=?", User1.Loginid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Getuser.Username,
			&Getuser.Loginid,
			&Getuser.Password,
			&Getuser.Phonenumber,
		)

	}

	if User1.Loginid == "" || User1.Password == "" {
		Userinfo.Status = "Failure"
		return

	} else if User1.Password == Getuser.Password {

		Userinfo.Username = Getuser.Username
		Userinfo.Loginid = Getuser.Loginid
		Userinfo.Phonenumber = Getuser.Phonenumber
		Userinfo.Status = "Success"
		return

	} else {
		Userinfo.Status = "Failure"
		return
	}

}

func EasyChangepassword(User1 Model.Changepass) (Userinfo string) {
	var Getuser string
	rows, err := OpenConnection["Rentmatics"].Query("select password from easylogin where Loginid=?", User1.Loginid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(

			&Getuser,
		)

	}
	fmt.Println(Getuser)
	fmt.Println(User1.Oldpassword)
	if User1.Oldpassword == Getuser {

		Queryupdate := "UPDATE  easylogin SET password='" + User1.Newpassword + "'where Loginid='" + User1.Loginid + "'"

		row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
		if err != nil {
			log.Println("Error -DB: update Profile", err, row)
		}
		return "success"

	} else {

		return "Failure"
	}

}

func EasyForgotpassword(User1 string) string {
	var Username string
	var Password string
	var Email string
	var Count1 int

	query := "SELECT COUNT(*) FROM easylogin WHERE easylogin.Loginid ='" + User1 + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count1,
	)
	if Count1 != 0 {

		rows, err := OpenConnection["Rentmatics"].Query("select Loginid,password,emailid from easylogin where Loginid=?", User1)
		if err != nil {
			log.Println("Error -DB: Get User", err)
		}
		for rows.Next() {

			rows.Scan(

				&Username,
				&Password,
				&Email,
			)

		}
		Tostring := "Username:" + Username + "<br>Password:" + Password

		sendkey := os.Getenv("SENDGRID_API_KEYGO")
		fmt.Println(sendkey)

		from := mail.NewEmail("Rentmatics User", "sandhiyabalakrishnan6@gmail.com")
		subject := "EASY E3 SHOPPING NOTIFICATION - FORGOT PASSWORD!"
		to := mail.NewEmail("Example User", Email)
		plainTextContent := Tostring
		htmlContent := "<strong>EASY SHOPPING NOTIFICATION - FORGOT PASSWORD <br> <br></strong><p>" + Tostring + "</p><br><b >Thank you for Contacting Easy E3 shopping</b>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		fmt.Println(Email)
		client := sendgrid.NewSendClient(sendkey)
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)

		}

		return "Success"
	} else {
		return "Failure"

	}

}
func EasyRetriveprofile(Loginid string) (Profileresp Model.Profileresponse) {
	var Address string
	var Data Model.Profile
	rows, err := OpenConnection["Rentmatics"].Query("Select userid,uniqueid,username,loginid,emailid,phonenumber from  easylogin where Loginid=?", Loginid)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Userid,
			&Data.Uniueid,
			&Data.Username,
			&Data.Loginid,
			&Data.Phonenumber,
			&Data.Emailid,
		)

	}
	row1, err := OpenConnection["Rentmatics"].Query("select Adress from easyprofile where easyuserid=?", Data.Userid)
	if err != nil {
		log.Println("Error -DB: Tenant", err)
	}
	for row1.Next() {
		row1.Scan(

			&Address,
		)

	}
	Profileresp.Profileresponse = Data
	Profileresp.Address = Address

	return Profileresp

}

func Easyinsertprofile(LoginidAddress Model.InsertAddress) string {

	var Data string
	rows, err := OpenConnection["Rentmatics"].Query("Select userid from  easylogin where Loginid=?", LoginidAddress.Loginid)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data,
		)

	}

	row, err := OpenConnection["Rentmatics"].Exec("insert into easyprofile (easyuserid,Adress) values (?,?)", Data, LoginidAddress.Address)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}

	return "Success"

}
func Easyupdateprofile(LoginidAddress Model.UpdateProfile) string {

	var Data string
	rows, err := OpenConnection["Rentmatics"].Query("Select userid from  easylogin where Loginid=?", LoginidAddress.Loginid)
	if err != nil {
		log.Println("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data,
		)

	}

	Queryupdate := "UPDATE  easylogin SET username='" + LoginidAddress.Username + "',emailid='" + LoginidAddress.Emailid + "',phonenumber='" + LoginidAddress.Phonenumber + "'where Loginid='" + LoginidAddress.Loginid + "'"

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}

	var Count int

	query := "SELECT COUNT(*) FROM easyprofile WHERE easyprofile.easyuserid =" + Data
	rowcount := OpenConnection["Rentmatics"].QueryRow(query)
	rowcount.Scan(
		&Count,
	)
	if Count == 0 {

		rowinsert, err := OpenConnection["Rentmatics"].Exec("insert into easyprofile (easyuserid,Adress) values (?,?)", Data, LoginidAddress.Address)
		if err != nil {
			log.Println("Error -DB: User", err, rowinsert)
		}

	} else {

		Queryupdate1 := "UPDATE  easyprofile SET Adress=' " + LoginidAddress.Address + " '  where easyuserid=" + Data

		row1, err := OpenConnection["Rentmatics"].Exec(Queryupdate1)
		if err != nil {
			log.Println("Error -DB: update Profile", err, row1)
		}
	}
	return "Success"

}

func Changepassword_DB(User1 Model.Changepassword) (Userinfo string) {
	var Getuser string
	rows, err := OpenConnection["Rentmatics"].Query("select password from easyadmin where Loginid=?", User1.Loginid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(

			&Getuser,
		)

	}
	fmt.Println(Getuser)
	fmt.Println(User1.Oldpass)
	if User1.Oldpass == Getuser {

		Queryupdate := "UPDATE  easyadmin SET password='" + User1.Newpass + "'where Loginid='" + User1.Loginid + "'"

		row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
		if err != nil {
			log.Println("Error -DB: update Profile", err, row)
		}
		return "success"

	} else {

		return "Failure"
	}

}

func Adminlogin_DB(User1 Model.AdminLogin) string {

	var Username string
	var password string
	rows, err := OpenConnection["Rentmatics"].Query("select Loginid,password from  easyadmin where Loginid=?", User1.User)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Username,
			&password,
		)

	}
	if User1.Pass == password {

		return "Success"

	} else {

		return "Failure"
	}

}
