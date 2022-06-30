# NALO

A `go` client library of sending sms using [nalo](httpswww.nalosolutions.com/).

## Getting started

Install package

```bash
go get -u github.com/ankorgh/nalo
```

### Setup a client

```go
n := nalo.New(context.Background(), nalo.Credentials{
		Username: "",
		Password: "",
	}, nil)
```

### Get balance

```go
balance, err := n.GetBalance()
if err != nil {
  log.Fatalf("Error getting balance: %v", err)
}

fmt.Println("Credit Balance: ", balance.Balance)
fmt.Println("Total sms balance: ", balance.TotalSMS)
```

### Send Sms

```
sms, err := n.SendSMS(message, to, from, nalo.Delivery_Active, nalo.MessageType_PlainTextISO)
if err != nil {
  log.Fatalf("Error sending sms: %v", err)
}

fmt.Println(sms)
```