package letter

import (
	"github.com/ruvido/go-pocketbase-admin/pkg"

    "github.com/spf13/viper"
	gomail "gopkg.in/gomail.v2"
	"log"
	"time"
	// "os"
)

type Server struct {
	User string
	Pass string
	Addr string
	Port int
}

// ======================================
var totalAddresses = 0
var countPad       = 0
// ======================================

func Send ( lett Letter, people []admin.User ) {

	batchSize 		:= viper.GetInt("sending.batch")
	sleepInterval 	:= viper.GetInt("sending.wait")
	if batchSize == 0 {batchSize = 500}
	countPad = 0
	totalAddresses = len(people)

	log.Println(batchSize,sleepInterval)

	for len(people) > 0 {
		if batchSize > len(people) {
			batchSize = len(people)
			sleepInterval = 0
		}       
		activeUsers := people[:batchSize]

		smtpSend (
			lett, 
			activeUsers)

		people = people[batchSize:]
		log.Println("Waiting sleep interval ...")
		// time.Sleep(sleepInterval * time.Second)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
		countPad += batchSize
	}
}

func smtpSend ( lett Letter, people []admin.User ) {
	smtp := new(Server)
	smtp.User = viper.GetString("smtp.user")
	smtp.Pass = viper.GetString("smtp.password")
	smtp.Addr = viper.GetString("smtp.address")
	smtp.Port = viper.GetInt("smtp.port")

	if countPad == 0 {
		log.Println("smtpSend",smtp.User)
		log.Println("smtpSend",smtp.Pass)
		log.Println("smtpSend",smtp.Addr)
		log.Println("smtpSend",smtp.Port)
	}
// ---------------------------------------------------

	d := gomail.NewDialer(
		smtp.Addr, smtp.Port, smtp.User, smtp.Pass)
	s, err := d.Dial()
	if err != nil {panic(err)}

// ---------------------------------------------------

	m := gomail.NewMessage()
	for ii,pp := range people {
		log.Printf("SEND> %5d/%-5d %q",countPad+ii+1,totalAddresses,pp.Email)
		m.SetHeader("From", viper.GetString("smtp.sender"))
		m.SetHeader("To", pp.Email )
		m.SetHeader("Subject", lett.Title )
		m.SetBody("text/html", lett.Content )
		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", pp.Email, err)
		}
		m.Reset()
	}	

// ---------------------------------------------------
}

// func SendNewsletter( letter Letter, people []User ){
//
// 	user     	:= "PM-B-broadcast-A0iIuI2wvc75VAMmIkhn1"
// 	password	:= "n7BFb_NWH9KOGH4LhU80WjCURAsGGEzAJS7a"
// 	server		:= "smtp-broadcasts.postmarkapp.com"
// 	d := gomail.NewDialer(server, 587, user, password )
// 	s, err := d.Dial()
// 	// _, err := d.Dial()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	countPad := 0 //TESTTTT
// 	totalAddresses := 199 //TESTTTT
//
// 	m := gomail.NewMessage()
// 	for ii,pp := range people {
// 		log.Printf("SEND> %5d/%-5d - %q",countPad+ii+1,totalAddresses,pp.Addr)
// 		m.SetHeader("From", "5pani2pesci <newsletter@5p2p.it>")
// 		m.SetHeader("To", pp.Addr )
// 		m.SetHeader("Subject", letter.Title )
// 		m.SetBody("text/html", letter.Content )
// 		if err := gomail.Send(s, m); err != nil {
// 			log.Printf("Could not send email to %q: %v", pp.Addr, err)
// 		}
// 		m.Reset()
// 	}	
// }

// func SendMessage( letter Letter, address string ){
//
// 	user     	:= "PM-B-broadcast-A0iIuI2wvc75VAMmIkhn1"
// 	password	:= "n7BFb_NWH9KOGH4LhU80WjCURAsGGEzAJS7a"
// 	server		:= "smtp-broadcasts.postmarkapp.com"
// 	d := gomail.NewDialer(server, 587, user, password )
// 	s, err := d.Dial()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "5pani2pesci <newsletter@5p2p.it>")
// 	m.SetHeader("To", address )
// 	m.SetHeader("Subject", letter.Title )
// 	m.SetBody("text/html", letter.Content )
// 	if err := gomail.Send(s, m); err != nil {
// 		log.Printf("Could not send email to %q: %v", address, err)
// 	}
// 	m.Reset()
// }
//
