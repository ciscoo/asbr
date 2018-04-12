# Android SMS Backup Reader (ASBR)

A small package that provides types to unmarshal XML produced by [SMS Backup & Restore](https://play.google.com/store/apps/details?id=com.riteshsahu.SMSBackupRestore) developed by [SyncTech](http://synctech.com.au/).

## Sample usage

```go
package main

import (
	"encoding/xml"
	"io"
	"log"
  "os"
  
  "github.com/ciscoo/asbr"
)

func main() {
	r, err := os.Open("./sms-20171005162519.xml")
	if err != nil {
		log.Fatal(err)
	}
	d := xml.NewDecoder(r)
	for {
		var err error
		var s SMS
		var m MMS

		token, err := d.Token()
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			log.Fatal(err)
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "sms" {
				err = d.DecodeElement(&s, &t)
				// Do something with the SMS
			}
			if t.Name.Local == "mms" {
				err = d.DecodeElement(&m, &t)
				// Do something with the MMS
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
```
