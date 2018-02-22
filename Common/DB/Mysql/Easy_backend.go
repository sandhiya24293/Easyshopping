package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Getuser_DB() (Userresponse []Model.Profile) {
	var Userresp Model.Profile

	rows, err := OpenConnection["Rentmatics"].Query("select userid,uniqueid,username,Loginid,phonenumber,emailid from  easylogin")
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Userresp.Userid,
			&Userresp.Uniueid,
			&Userresp.Username,
			&Userresp.Loginid,
			&Userresp.Phonenumber,
			&Userresp.Emailid,
		)

		Userresponse = append(Userresponse, Userresp)
	}
	return

}

func GetAllorder_DB() (Orderresult []Model.Orderdataresponse) {
	var Userresponse Model.Orderdataresponse
	var Orderdata Model.Orderresponse

	rows, err := OpenConnection["Rentmatics"].Query("select Orderid,Billid,Uniqueid,loginid,Totalproducts,TotalAmount,Date from  easyorderplaced")
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Orderdata.Orderid,
			&Orderdata.Billnumber,
			&Orderdata.Userid,
			&Orderdata.Loginid,
			&Orderdata.Totalproduct,
			&Orderdata.Totalamount,
			&Orderdata.Date,
		)
		row, err := OpenConnection["Rentmatics"].Query("select Productid,Productname,productrate,weight from  productlist where orderplaceid =?", Orderdata.Orderid)
		if err != nil {
			log.Println("Error -DB: Get User", err)
		}
		var Prodarray []Model.Productdata
		for row.Next() {
			var Prodarr Model.Productdata

			row.Scan(
				&Prodarr.Productid,
				&Prodarr.Productname,
				&Prodarr.Productrate,
				&Prodarr.Weight,
			)
			Prodarray = append(Prodarray, Prodarr)

		}

		Userresponse.Orderdetails = Orderdata
		Userresponse.Productsdetails = Prodarray
		Orderresult = append(Orderresult, Userresponse)
	}

	return

}

func Addproduct_DB(Adddata Model.Addproduct) {
	row, err := OpenConnection["Rentmatics"].Exec("insert into easyvegetables (Easyproduct,Type,Rate1KG,Rate500gm,Rate250gm,display,pictureurl ) values (?,?,?,?,?,?,?)", Adddata.Productname, Adddata.Categorylist, Adddata.Productrate1kg, Adddata.Productrate500kg, Adddata.Productrate250gm, "show", Adddata.Pictureurl)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}

}

func AddFood_DB(Adddata Model.Ownfood) {
	row, err := OpenConnection["Rentmatics"].Exec("insert into owndeliver (dishname,Rate,platecount,image,status ) values (?,?,?,?,?)", Adddata.Dishname, Adddata.Rate, Adddata.Platecount, Adddata.Image, "show")
	if err != nil {
		log.Println("Error -DB: update FOOD", err, row)
	}

}
func AddNonproduct_DB(Adddata Model.NonAddproduct) {
	row, err := OpenConnection["Rentmatics"].Exec("insert into easynonveg (Easyproduct,Type,Rate2KG,Rate1750gm,Rate1500gm,Rate1250gm,Rate1KG,Rate500gm,Rate250gm,display,pictureurl ) values (?,?,?,?,?,?,?,?,?,?,?)", Adddata.Productname, Adddata.Categorylist, Adddata.Productrate2000gm, Adddata.Productrate1750gm, Adddata.Productrate1500gm, Adddata.Productrate1250gm, Adddata.Productrate1kg, Adddata.Productrate500kg, Adddata.Productrate250gm, "show", Adddata.Pictureurl)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}

}
func Updateproduct_DB(Editdata Model.Updateproduct) {
	Queryupdate := "UPDATE easyvegetables SET Easyproduct='" + Editdata.Productname + "' , Rate1KG= '" + fmt.Sprintf("%v", Editdata.Productrate1kg) + "' ,Rate500gm= '" + fmt.Sprintf("%v", Editdata.Productrate500gm) + "' ,Rate250gm= '" + fmt.Sprintf("%v", Editdata.Productrate250gm) + "'  where easyid= " + fmt.Sprintf("%v", Editdata.Productid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}

func Updateimage_DB(getid int, getimage string) {
	Queryupdate := "UPDATE easyvegetables SET pictureurl='" + getimage + "'  where easyid= " + fmt.Sprintf("%v", getid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}
func Updatenonimage_DB(getid int, getimage string) {
	Queryupdate := "UPDATE easynonveg SET pictureurl='" + getimage + "'  where easynonid= " + fmt.Sprintf("%v", getid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}

func Updatefoodproduct_DB(Editdata Model.UpdateFoodproduct) {
	Queryupdate := "UPDATE owndeliver SET dishname='" + Editdata.Dishname + "' , Rate= '" + fmt.Sprintf("%v", Editdata.Dishrate) + "' ,platecount= '" + fmt.Sprintf("%v", Editdata.Platecount) + "'  where id= " + fmt.Sprintf("%v", Editdata.Foodid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}
func Updatenonproduct_DB(Editdata Model.Updatenonproduct) {
	Queryupdate := "UPDATE easynonveg SET Easyproduct='" + Editdata.Productname + "' , Rate2KG= '" + fmt.Sprintf("%v", Editdata.Productrate2kg) + "' ,Rate1750gm= '" + fmt.Sprintf("%v", Editdata.Productrate1750gm) + "' ,Rate1500gm= '" + fmt.Sprintf("%v", Editdata.Productrate1500gm) + "' ,Rate1250gm= '" + fmt.Sprintf("%v", Editdata.Productrate1250gm) + "' ,Rate1KG= '" + fmt.Sprintf("%v", Editdata.Productrate1kg) + "' ,Rate500gm= '" + fmt.Sprintf("%v", Editdata.Productrate500gm) + "' ,Rate250gm= '" + fmt.Sprintf("%v", Editdata.Productrate250gm) + "'  where easynonid= " + fmt.Sprintf("%v", Editdata.Productid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}

func Updatestatus_DB(Data Model.Updatestatus) {

	Queryupdate := "UPDATE easyvegetables SET display ='" + Data.Productstatus + "' where easyid= " + fmt.Sprintf("%v", Data.Productid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}

func Updatestatusnon_DB(Data Model.Updatestatus) {

	Queryupdate := "UPDATE easynonveg SET display ='" + Data.Productstatus + "' where easynonid= " + fmt.Sprintf("%v", Data.Productid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}

func Updatefoodstatus_DB(Data Model.Updatestatus) {

	Queryupdate := "UPDATE owndeliver SET status ='" + Data.Productstatus + "' where id= " + fmt.Sprintf("%v", Data.Productid)

	row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	if err != nil {
		log.Println("Error -DB: update Profile", err, row)
	}
}

func Getordertracking_DB() (Trackresult Model.Trackrespopnse) {
	var Orderdaata Model.Gettrack
	var Orderplaced Model.Ordertrackdata

	rows, err := OpenConnection["Rentmatics"].Query("select orderid,Orderplacedid,loginid,status from  ordertracking where status =?", "Order")
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Orderdaata.Orderid,
			&Orderdaata.Orderplacedid,
			&Orderdaata.Loginid,
			&Orderdaata.Status,
		)
	}

	row, err := OpenConnection["Rentmatics"].Query("select Billid,Phonenumber,Totalproducts,TotalAmount,Date from  easyorderplaced where Orderid =?", Orderdaata.Orderid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}

	for row.Next() {

		row.Scan(
			&Orderplaced.Billid,
			&Orderplaced.Phonenumber,
			&Orderplaced.Totalproducts,
			&Orderplaced.TotalAmount,
			&Orderplaced.Date,
		)

	}
	Trackresult.Trackdata = Orderdaata
	Trackresult.Orderplaced = Orderplaced

	return

}
