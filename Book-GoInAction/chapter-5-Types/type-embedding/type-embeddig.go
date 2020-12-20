package embedding

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user  // Included type (embedded)
	level string
}

// Inner type implementation of notify
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// this changes the inner type implementation
func (u *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		u.name,
		u.email)
}

// Call - Calls the type-embedding content
func Call() {
	fmt.Println("Type Embedding")

	/*
		In go we can modify and alter types, through a thing call type-embedding,
		with that concept we have a inner type and a outer type, the outer type
		implements and modify the inner type.

		The inner type identifiers are promoted to the outer type, in a way
		similar as we have explicit declarated the identifiers in outer type,
		so the outer type have all the inner type have.
	*/

	ad := admin{
		user: user{
			name:  "John Smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// The inner type is promoted
	ad.notify()

	// We can acess the inner type method
	ad.user.notify()

	// Notifier is passed do admin type through type-embedding
	sendNotification(&ad)

	return
}

func sendNotification(n notifier) {
	n.notify()
}
