// main.go
package main

import (
	"log"
	"net/mail"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	app.OnRecordAfterCreateRequest("Contact_us").Add(func(e *core.RecordCreateEvent) error {
		message := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: "rezafiansyah.putra@gmail.com"}, {Address: "samueladitia95@gmail.com"}},
			Subject: "POCKETBASE MAILER TEST",
			HTML:    "SEXY BODY",
			// bcc, cc, attachments and custom headers are also supported...
		}

		return app.NewMailClient().Send(message)
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
