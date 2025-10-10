package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

type SMSNotifier struct {
	PhoneNumber string
}

func (en EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю email на:%s, %s ",en.EmailAddress, message)
}

func (sn SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю sms на номер:%s,%s ",sn.PhoneNumber ,message)
}
