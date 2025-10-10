package main

import (
	"fmt"
)

func main(){
	user1 := Reader{
		ID: 1,
		FirstName: "Ilona",
		LastName: "Valieva",
		IsActive: true,
	}

	fmt.Println(user1)
	user1.DisplayReader()
	user1.Deactivate()
	user1.Activate()
	fmt.Println(user1)


	notifiers := []Notifier{}
	EmailNotifier:= EmailNotifier{EmailAddress: "ilonava09@gmail.com"}
	SMSNotifier:=SMSNotifier{PhoneNumber: "+79163158170"}

	notifiers=append(notifiers, EmailNotifier,SMSNotifier)
	message:="Ваша книга просрочена!"

	for _, notifier:=range notifiers{
		notifier.Notify(message)
	}
}


