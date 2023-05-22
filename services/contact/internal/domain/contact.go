package domain

import (
	"strings"
	"unicode"
)

type Contact struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
}

type Group struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewContact(id int64, firstName, lastName, patronymic, phoneNumber string) *Contact {
	contact := &Contact{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Patronymic:  patronymic,
		PhoneNumber: sanitizePhoneNumber(phoneNumber),
	}

	contact.FullName = generateFullName(contact)
	return contact
}

func sanitizePhoneNumber(phoneNumber string) string {
	// Remove non-digit characters from the phone number
	var sb strings.Builder
	for _, ch := range phoneNumber {
		if unicode.IsDigit(ch) {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func generateFullName(contact *Contact) string {
	// Combine last name and patronymic to form full name
	return contact.LastName + " " + contact.Patronymic
}

func NewGroup(id int64, name string) *Group {
	group := &Group{
		Id:   id,
		Name: truncateString(name, 250),
	}
	return group
}

func truncateString(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	return str[:maxLength]
}
