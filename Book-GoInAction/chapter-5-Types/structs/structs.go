package structs

import "fmt"

// Declaring a struct
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func (u user) notify() {
	fmt.Printf("Sending user e-mail to %s<%s>\n",
		u.name,
		u.email)
}

func Call() {
	fmt.Println("Structs")

	var bill user
	bill.notify()

	// When declaring a variable with zero values is preferable to use "var"
	// when declaring a variable with values is preferable to use ":="

	// Declaring using a literal struct
	lisa := user{
		name:       "lisa",
		email:      "lisa@gmail.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(lisa)

	// Another way to declare, this way we hide the keys, passing only values
	eduart := user{"Eduart", "eduart@gmail.com", 321, false}
	fmt.Println(eduart)

	// Declaring a admin type (struct who have a struct)
	fred := admin{
		person: user{
			name:       "fred",
			email:      "fred@gmail.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(fred)

	return
}
