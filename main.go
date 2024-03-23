// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/samueladitia95/bnw-be/utils"
)

func main() {
	app := pocketbase.New()

	app.OnRecordAfterCreateRequest("Contact_us").Add(func(e *core.RecordCreateEvent) error {
		var (
			mailTo  string
			payload map[string]any
		)

		payload = e.Record.PublicExport()

		mailTo = utils.EmailSubjectMapping(payload["subject"].(string))

		message := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: mailTo}},
			Subject: "B&W Get In Touch",
			HTML: fmt.Sprintf(`<p>Hello, %s</p>
			<p>This client want to get in touch with these details want to get in touch with you:</p><br>
			<p>Name: %s</p><br>
			<p>Email: %s</p><br>
			<p>Phone: %s</p><br>
			<p>Message: %s</p><br>
			<p>
			  Thanks<br/>
			</p>`, "B&W Team", payload["name"], payload["sender"], payload["phone_number"], payload["message"]),
			// bcc, cc, attachments and custom headers are also supported...
		}

		return app.NewMailClient().Send(message)
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/verify-retailer", func(c echo.Context) error {
			data := struct {
				Id       string `json:"id" form:"id"`
				Password       string `json:"password" form:"password"`
			}{}

			if err := c.Bind(&data); err != nil {
					return apis.NewBadRequestError("Failed to read request data", err)
			}

			record, _ := app.Dao().FindRecordById("Products", data.Id)
			if record == nil {
				return c.JSON(http.StatusNotFound, map[string]any{"message": "Not Found"})
			}

			password := record.GetString("password")

			if (password != data.Password) {
				return c.JSON(http.StatusUnauthorized, map[string]any{"message": "Failed Auth"})
			}
			
			return c.JSON(http.StatusOK, map[string]any{"message": "Ok"})
		})

    return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
