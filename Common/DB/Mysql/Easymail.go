package Mysql

import (
	Model "Easyshopping/Model"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

type SmtpServer struct {
	Host      string
	Port      string
	TlsConfig *tls.Config
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	return header
}

func MailOrder(t interface{}, addr string, number string) {
	var htmlContent string
	switch t1 := t.(type) {
	case Model.Orderplaced:
		Order := t1
		var stringprod []string
		for _, v := range Order.Products {
			var getstring string
			getstring = "<tr style='font-style:sans-serif'><td>" + v.Product + "</td><td>" + fmt.Sprintf("%v", v.Rate) + "</td><td>" + v.Weight + "</td></tr>"
			stringprod = append(stringprod, getstring)
		}
		result := strings.Join(stringprod, "")
		fmt.Println(stringprod)

		htmlContent = "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
			"<div style='font-style:sans-serif'>LOGINID - " + Order.Loginid +
			"<div style='font-style:sans-serif'>ADDRESS - " + addr +
			"</div><div style='font-style:sans-serif'>DATE -" + Order.Date +
			"</div><div style='font-style:sans-serif'>TOTAL AMOUNT -" + fmt.Sprintf("%v", Order.TotalAmount) +
			"</div><div style='font-style:sans-serif'>NO OF PRODUCTS -" + fmt.Sprintf("%v", Order.Noofproducts) +
			"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >PRODUCT</th><th>RATE</th><th>WEIGHT</th></tr>" +
			result + "</tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	case Model.Easyinstance:
		Instancedata := t1
		htmlContent = "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
			"<div style='font-style:sans-serif'>NAME  - " + Instancedata.Deliveryname +
			"<div style='font-style:sans-serif'>ADDRESS - " + Instancedata.Deliveryaddress +
			"</div><div style='font-style:sans-serif'>NUMBER -" + Instancedata.Deliverynumber +
			"</div><div style='font-style:sans-serif'>PRODUCT  -" + Instancedata.Productname +
			"</div><div style='font-style:sans-serif'>MESSAGE" + Instancedata.Message +
			"</div></tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	case Model.Ownorderplaced:
		getorder := t1
		var stringprod []string

		for _, Food := range getorder.Food {
			var getstring string
			getstring = "<tr style='font-style:sans-serif'><td>" + Food.Dishname + "</td><td>" + fmt.Sprintf("%v", Food.Rate) + "</td><td>" + fmt.Sprintf("%v", Food.Platecount) + "</td></tr>"
			stringprod = append(stringprod, getstring)
		}
		result := strings.Join(stringprod, "")

		htmlContent = "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
			"<div style='font-style:sans-serif'>LOGINID - " + getorder.Loginid +
			"<div style='font-style:sans-serif'>ADDRESS - " + addr +
			"</div><div style='font-style:sans-serif'>NUMBER -" + number +
			"</div><div style='font-style:sans-serif'>NO Of ITEMS -" + fmt.Sprintf("%v", getorder.Noofitems) +
			"</div><div style='font-style:sans-serif'>TOTAL AMOUNT -" + fmt.Sprintf("%v", getorder.TotalAmount) +

			"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >DISH </th><th>RATE</th><th>PLATECOUNT</th></tr>" +
			result + "</tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	case Model.Fooddelivery:
		Instancedata := t1
		var stringprod []string
		for _, Food := range Instancedata.Food {
			var getstring string
			getstring = "<tr style='font-style:sans-serif'><td>" + Food.Foodname + "</td><td>" + fmt.Sprintf("%v", Food.Quantity) + "</td></tr>"
			stringprod = append(stringprod, getstring)
		}
		result := strings.Join(stringprod, "")

		htmlContent = "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
			"<div style='font-style:sans-serif'>LOGINID - " + Instancedata.Loginid +
			"<div style='font-style:sans-serif'>ADDRESS - " + Instancedata.Address +
			"</div><div style='font-style:sans-serif'>NUMBER -" + Instancedata.Mobilenumber +
			"</div><div style='font-style:sans-serif'>HOTEL NAME -" + Instancedata.Hotelname +

			"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >FOOD </th><th>QUANTITY</th></tr>" +
			result + "</tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	default:
		_ = t

	}
	mail := Mail{}
	mail.Sender = "Rentmatics@gmail.com"
	mail.To = []string{"sandhiyabalakrishnan6@gmail.com"}
	//mail.Cc = []string{"mnp@gmail.com"}
	//mail.Bcc = []string{"a69@outlook.com"}
	mail.Subject = "E3 NOTIFICATION - New Order Recieved!"
	mail.Body = htmlContent
	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{Host: "smtp.gmail.com", Port: "465"}
	smtpServer.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	auth := smtp.PlainAuth("", mail.Sender, "Motoski@1991", smtpServer.Host)

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TlsConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		log.Panic(err)
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
