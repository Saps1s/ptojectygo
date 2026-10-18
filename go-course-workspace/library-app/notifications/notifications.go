package notifications

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{ EmailAddress string }

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю email на %s: %q\n", e.EmailAddress, message)
}

type SMSNotifier struct{ PhoneNumber string }

func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю SMS на номер %s: %q\n", s.PhoneNumber, message)
}