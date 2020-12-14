package methods

import "fmt"

type user struct {
	name  string
	email string
}

// When we use the receptor, to receive by value, it'll create a copy of the
// value received
func (u user) notify() {
	fmt.Printf("Sending User email to %s<%s>\n",
		u.name,
		u.email,
	)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func Call() {
	fmt.Println("Methods")

	/*
		Methods are functions that add behaviors to types declared by the user
		they receive extra params between the reserved word "func" and the
		function name.
	*/

	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()

	// Pointers of type user can also call methods with value receivers
	lisa := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify()

	// Values of type user can also call methods declared with pointer receivers
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can also call methods declared with pointer
	// receivers
	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()

	// The receptor links the function to a type when a function have a receptor
	// than it's called method

	return
}
