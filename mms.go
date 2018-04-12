package asbr

// MMS is a MMS message
type MMS struct {
	Date         string `xml:"date,attr"`
	DateSent     string `xml:"date_sent,attr"`
	ReadableDate string `xml:"readable_date"`
	Address      string `xml:"address,attr"`
	ContactName  string `xml:"contact_name,attr"`
	Parts        []Part `xml:"parts>part"`
	Addrs        []Addr `xml:"addrs>addr"`
}

// Text retrieves the message body from the MMS, if any.
func (m MMS) Text() string {
	var t string
	for _, part := range m.Parts {
		if part.ContentType == "text/plain" {
			t = part.Text
			break
		}
	}
	return t
}

// Media retrieves the MMS media, if any.
func (m MMS) Media() ([]byte, error) {
	var p Part
	for _, part := range m.Parts {
		if part.ContentType == "application/smil" || part.ContentType == "text/plain" {
			continue
		}
		p = part
	}
	if err := p.DecodePart(); err != nil {
		return nil, err
	}
	return p.DecodedMedia, nil
}

// Recipent retrieves the recipent of the MMS message.
func (m MMS) Recipent() (Addr, error) {
	var recipent Addr
	for _, addr := range m.Addrs {
		ok, err := addr.IsRecipient()
		if err != nil {
			return recipent, err
		}
		if ok {
			recipent = addr
			break
		}
	}
	return recipent, nil
}

// Sender retrives the sender of the MMS message.
func (m MMS) Sender() (Addr, error) {
	var sender Addr
	for _, addr := range m.Addrs {
		ok, err := addr.IsSender()
		if err != nil {
			return sender, err
		}
		if ok {
			sender = addr
			break
		}
	}
	return sender, nil
}
