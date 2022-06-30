package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ankorGH/nalo/v1"
)

var to, message, username, password, from string

func main() {
	flag.StringVar(&from, "from", "Nyn", "Message title header")
	flag.StringVar(&message, "message", "Hello, User", "Message")
	flag.StringVar(&username, "username", "", "username - nalo credentials for sending sms")
	flag.StringVar(&password, "password", "", "password - nalo credentials for sending sms")
	flag.StringVar(&to, "to", "to", "to - contact of destination")
	flag.Parse()

	n := nalo.New(context.Background(), nalo.Credentials{
		Username: username,
		Password: password,
	}, nil)

	balance, err := n.GetBalance()
	if err != nil {
		fmt.Printf("Error getting balance: %v", err)
		return
	}

	fmt.Println("Credit Balance: ", balance.Balance)
	fmt.Println("Total sms balance: ", balance.TotalSMS)

	if balance.TotalSMS <= 0 {
		fmt.Println("Insufficient balance to send sms")
		return
	}

	sms, err := n.SendSMS(message, to, from, nalo.Delivery_Active, nalo.MessageType_PlainTextISO)
	if err != nil {
		fmt.Printf("Error sending sms: %v", err)
		return
	}

	fmt.Println(sms)
}
