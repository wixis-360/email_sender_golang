package main

import (
	"Email_Sender_Using_GoLang/config"
	"Email_Sender_Using_GoLang/models"
)

func main() {
	// define user
	user:=models.User{Name: "0 Example",Address: "Wixis 360"}
	//calling send email method
	config.SendMail(user)
}
