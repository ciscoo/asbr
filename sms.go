package asbr

import (
	"strconv"
)

// SMS message types.
const (
	MessageTypeAll = iota
	MessageTypeInbox
	MessageTypeSent
	MessageTypeDraft
	MessageTypeOutbox
	MessageTypeFailed
	MessageTypeQueued
)

// SMS is a text-based SMS message.
type SMS struct {
	Address      string `xml:"address,attr"`
	Date         string `xml:"date,attr"`
	DateSent     string `xml:"date_sent,attr"`
	ReadableDate string `xml:"readable_date,attr"`
	Type         string `xml:"type,attr"`
	Body         string `xml:"body,attr"`
	ContactName  string `xml:"contact_name,attr"`
}

// IsMessageType verifies the SMS is of the supplied type.
func (s SMS) IsMessageType(mt int) (bool, error) {
	i, err := strconv.Atoi(s.Type)
	if err != nil {
		return false, err
	}
	return i == mt, nil
}
