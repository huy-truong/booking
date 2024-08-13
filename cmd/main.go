package main

import (
	"github.com/huy-truong/booking/email"
)

func main() {
	email.Send(email.Content("Mr", "TRUONG", 10), "truonghuy1801@gmail.com")
}
