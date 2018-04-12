package asbr

import "strconv"

// An Addr type can only be one of the following protocol data units (PDU).
// You are likely to only encounter From and To values from the Android messaging apps.
// BCC and CC are included just in case.
const (
	BCC  = 0x81
	CC   = 0x82
	From = 0x89
	To   = 0x97
)

// Addr is the address information for a MMS message
type Addr struct {
	Address string `xml:"address,attr"`
	Type    string `xml:"type,attr"`
}

// IsSender verifies the Addr is the sender of the MMS message.
func (a Addr) IsSender() (bool, error) {
	return a.compareType(From)
}

// IsRecipient verifies the Addr is the recipient of the MMS message.
func (a Addr) IsRecipient() (bool, error) {
	return a.compareType(To)
}

func (a Addr) compareType(pdu int) (bool, error) {
	i, err := strconv.Atoi(a.Type)
	if err != nil {
		return false, err
	}
	return i == pdu, nil
}
