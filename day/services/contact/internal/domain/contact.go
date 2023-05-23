package domain

import (
	"strconv"
)

type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return c.FirstName + " " + c.LastName
}

func (c *Contact) SetFullName(firstName, lastName string) {
	c.FirstName = firstName
	c.LastName = lastName
}

func (c *Contact) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

func (c *Contact) IsValidPhoneNumber() bool {
	_, err := strconv.Atoi(c.PhoneNumber)
	return err == nil
}
