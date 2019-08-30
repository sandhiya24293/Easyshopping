
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
			getstring = "\nPRODUCT - " + v.Product + "\nRATE :" + fmt.Sprintf("%v", v.Rate) + "\nWEIGHT :" + v.Weight
			stringprod = append(stringprod, getstring)
		}
		result := strings.Join(stringprod, "")
		fmt.Println(stringprod)

		htmlContent = "E3 NEW ORDER RECIEVED : " +
			"\n\nLOGINID - " + Order.Loginid +
			"\nADDRESS - " + addr +
			"\nDATE - " + Order.Date +
			"\nTOTAL AMOUNT - " + fmt.Sprintf("%v", Order.TotalAmount) + result +

			"\n\n Please Check E3 Admin Panel for more detail ...!"

	case Model.Easyinstance:
		Instancedata := t1
		htmlContent = "\n\nE3 NEW ORDER : " +
			"\n\nNAME  - " + Instancedata.Deliveryname +
			"\nADDRESS - " + Instancedata.Deliveryaddress +
			"\nNUMBER -" + Instancedata.Deliverynumber +
			"\nPRODUCT  -" + Instancedata.Productname +
			"\nMESSAGE" + Instancedata.Message +
			"\n\nPlease Check E3 Admin Panel for more detail ...!"

	case Model.Ownorderplaced:
		getorder := t1
		var stringprod []string

		for _, Food := range getorder.Food {
			var getstring string
			getstring = "\nDISH : " + Food.Dishname + "\nRATE : " + fmt.Sprintf("%v", Food.Rate) + "\nPLATECOUNT" + fmt.Sprintf("%v", Food.Platecount)
			stringprod = append(stringprod, getstring)
		}
		result := strings.Join(stringprod, "")

		htmlContent = "E3 NEW ORDER : " +
			"\n\nLOGINID - " + getorder.Loginid +
			"\nADDRESS - " + addr +
			"\nNUMBER -" + number +
			"\nNO Of ITEMS -" + fmt.Sprintf("%v", getorder.Noofitems) +
			"\nTOTAL AMOUNT -" + fmt.Sprintf("%v", getorder.TotalAmount) +
			result + "\nPlease Check E3 Admin Panel for more detail ...!"

	case Model.Fooddelivery:
		Instancedata := t1
		var stringprod []string
		for _, Food := range Instancedata.Food {
			var getstring string
			getstring = "\nFOOD - " + Food.Foodname + "\nQUANTITY - " + fmt.Sprintf("%v", Food.Quantity)
			stringprod = append(stringprod, getstring)
		}
		result := strings.Join(stringprod, "")

		htmlContent = "\n\nE3 NEW ORDER :  " +
			"\nLOGINID - " + Instancedata.Loginid +
			"\nADDRESS - " + Instancedata.Address +
			"\nNUMBER -" + Instancedata.Mobilenumber +
			"\nHOTEL NAME -" + Instancedata.Hotelname +

			result + "\nPlease Check E3 Admin Panel for more detail ...!"

	default:
		_ = t

	}
	mail := Mail{}
	mail.Sender = "Rentmatics@gmail.com"
	mail.To = []string{"e3easyshopping@gmail.com"}
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
