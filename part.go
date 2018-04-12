package asbr

import (
	"encoding/base64"
	"errors"
	"io"
)

// Part is the MMS message part.
type Part struct {
	ContentType  string `xml:"ct,attr"`
	Name         string `xml:"cl,attr"`
	Data         string `xml:"data,attr"`
	Text         string `xml:"text,attr"`
	DecodedMedia []byte
}

// DecodePart will decode the raw base64 string from the Part.
func (p *Part) DecodePart() error {
	if p.Data == "" {
		return errors.New("no data")
	}
	dbuf, err := base64.StdEncoding.DecodeString(p.Data)
	if err != nil {
		return err
	}
	p.DecodedMedia = dbuf
	return nil
}

// WritePart will persist the decoded data to disk.
func (p *Part) WritePart(w io.Writer) (int, error) {
	var err error
	if len(p.DecodedMedia) == 0 {
		err = p.DecodePart()
	}
	if err != nil {
		return 0, err
	}
	return w.Write(p.DecodedMedia)
}
