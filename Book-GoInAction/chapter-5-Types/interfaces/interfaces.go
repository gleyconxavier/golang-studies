package interfaces

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func (u *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		u.name,
		u.email)
}

func Call() {

	fmt.Println("Interfaces")
	/*
		Interfaces are types that simple declare behaviors,
		those behaviors are not implemented directly by interfaces,
		but for types declared by the user, through methods.
	*/

	u := user{"Bill", "bill@email.com"}
	a := admin{"The Badass", "adminthebadass@email.com"}

	sendNotification(&u) // The *user receiver only accepts pointers
	sendNotification(&a) // The *admin receiver only accepts pointers

	/*
		The motive behind that is cause is not every time we can get
		the value pointer, passing by value.
	*/

	return
}

func sendNotification(n notifier) {
	n.notify()
}
