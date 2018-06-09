package main

// build with `go build gcflags "-m -m"`

func main() {
	user1 := createUser1()
	user2 := createUser2()

	println("u1", &user1, "u2", &user2)
}

type user struct {
	name  string
	email string
}

func createUser1() user {
	User := user{
		email: "a@a.com",
		name:  "aa",
	}
	println("v1", &User)
	return User
}

func createUser2() *user { // ย้ายไป Heap หลังจาก run เสร็จ
	User := user{
		email: "b@b.com",
		name:  "bb",
	}
	println("v2", &User)
	return &User
}
