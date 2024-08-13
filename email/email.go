package email

import "fmt"

func Send(content string, to string) {
	fmt.Printf("Sending to \"%s\"\nContent:\n%s\n", to, content)
}

func Content(title string, name string, nights uint) string {
	text := "Dear %s. %s,\nYour room reservation for %d night(s) is confirmed. Have a nice day"
	return fmt.Sprintf(text, title, name, nights)
}
