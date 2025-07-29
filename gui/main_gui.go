package main

import (
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
		errorLabel.SetText("Not implemented yet")
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
