package builder

import "fmt"

type email struct {
	from, to, subject, body string
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder) // execute func(b *EmailBuilder) in an argument, on the caller side.
	sendEmailImpl(&builder.email)
}

func sendEmailImpl(email *email) {
	fmt.Println("this mail has been sent;")
	fmt.Printf("from:    %s\n", email.to)
	fmt.Printf("to:      %s\n", email.from)
	fmt.Printf("subject: %s\n", email.subject)
	fmt.Printf("body:    %s\n", email.body)
}
