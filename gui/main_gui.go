package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Golf Logger Login")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	passwordEntry := widget.NewEntry()
	passwordEntry.SetPlaceHolder("Password")

	errorLabel := widget.NewLabel("")

	loginBtn := widget.NewButton("Login", func() {
		email := emailEntry.Text
		password := passwordEntry.Text

		payload := map[string]string{
			"email":    email,
			"password": password,
		}

		data, err := json.Marshal(payload)
		if err != nil {
			errorLabel.SetText("Error encoding credentials")
			return
		}

		go func() {
			resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(data))
			if err != nil {
				myApp.SendNotification(&fyne.Notification{Title: "Error", Content: "Unable to reach server"})
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			if resp.StatusCode != 200 {
				errorLabel.SetText(fmt.Sprintf("Login failed: %s", body))
				return
			}

			var loginResp struct {
				UserID    string `json:"user_id"`
				AuthToken string `json:"auth_token"`
			}
			if err := json.Unmarshal(body, &loginResp); err != nil {
				errorLabel.SetText("Bad response from server")
				return
			}

			errorLabel.SetText("Login Succesful!")
		}()
	})

	form := container.NewVBox(
		widget.NewLabel("Welcome to Golf Logger"),
		emailEntry,
		passwordEntry,
		loginBtn,
		errorLabel,
	)

	myWindow.SetContent(form)
	myWindow.Resize(fyne.NewSize(500, 800))
	myWindow.ShowAndRun()
}
