package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("Golf Logger Login")

	win.SetContent(makeLoginScreen(win, ""))

	win.Resize(fyne.NewSize(500, 800))
	win.ShowAndRun()
}

func login(email, password string, errorLabel *widget.Label, win fyne.Window) {
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
			errorLabel.SetText("Unable to reach server")
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

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

		// Success! Update UI
		errorLabel.SetText("Login Successful!")
		win.SetContent(makeMenuScreen(win, loginResp.AuthToken))
	}()
}

func register(email, password string, msgLabel *widget.Label, win fyne.Window) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		msgLabel.SetText("Error encoding credentials")
		return
	}

	go func() {
		resp, err := http.Post("http://localhost:8080/api/register", "application/json", bytes.NewBuffer(data))
		if err != nil {
			msgLabel.SetText("Unable to reach server")
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != 201 {
			msgLabel.SetText(fmt.Sprintf("Register failed: %s", body))
			return
		}
		msg := "Registration successful! Ready to Login."
		win.SetContent(makeLoginScreen(win, msg))
	}()
}

func makeRegisterScreen(win fyne.Window) fyne.CanvasObject {
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	passwordEntry := widget.NewEntry()
	passwordEntry.SetPlaceHolder("Password")

	msgLabel := widget.NewLabel("")

	registerBtn := widget.NewButton("Register Now", func() {
		email := emailEntry.Text
		password := passwordEntry.Text

		register(email, password, msgLabel, win)

	})

	backBtn := widget.NewButton("Back to Login", func() {
		win.SetContent(makeLoginScreen(win, ""))
	})

	form := container.NewVBox(
		widget.NewLabel("Register for Golf Logger"),
		emailEntry,
		passwordEntry,
		registerBtn,
		backBtn,
		msgLabel,
	)

	return container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(form),
		layout.NewSpacer(),
	)
}

func makeLoginScreen(win fyne.Window, msg string) fyne.CanvasObject {
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	passwordEntry := widget.NewEntry()
	passwordEntry.SetPlaceHolder("Password")

	msgLabel := widget.NewLabel(msg)

	loginBtn := widget.NewButton("Login", func() {
		email := emailEntry.Text
		password := passwordEntry.Text

		login(email, password, msgLabel, win)
	})

	registerBtn := widget.NewButton("Register", func() {
		win.SetContent(makeRegisterScreen(win))
	})

	form := container.NewVBox(
		widget.NewLabel("Welcome to Golf Logger"),
		emailEntry,
		passwordEntry,
		loginBtn,
		registerBtn,
		msgLabel,
	)

	return container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(form),
		layout.NewSpacer(),
	)
}

func makeMenuScreen(win fyne.Window, authToken string) fyne.CanvasObject {
	bagBtn := widget.NewButton("My Bag", func() {
		win.SetContent(makeBagScreen(win, authToken))
	})
	rangeMapBtn := widget.NewButton("Range Map", func() {
		//win.SetContent(makeRangeMapScreen(win, authToken))
	})
	logoutBtn := widget.NewButton("Log Out", func() {
		win.SetContent(makeLoginScreen(win, ""))
	})

	form := container.NewVBox(
		bagBtn,
		rangeMapBtn,
		logoutBtn,
	)

	return container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(form),
		layout.NewSpacer(),
	)
}
