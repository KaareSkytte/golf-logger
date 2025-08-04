package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Club struct {
	ClubName string `json:"ClubName"`
	ClubType string `json:"ClubType"`
	Distance int    `json:"Distance"`
	InBag    bool   `json:"InBag"`
}

var typeOrder = []string{"Driver", "Wood", "Hybrid", "Iron", "Wedge", "Putter"}

func getMyBag(authToken string) ([]Club, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/bag", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}

	var clubs []Club
	if err := json.NewDecoder(resp.Body).Decode(&clubs); err != nil {
		return nil, fmt.Errorf("bad JSON: %w", err)
	}

	return clubs, nil
}

func getAllClubs(authToken string) ([]Club, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/bag/full", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}

	var clubs []Club
	if err := json.NewDecoder(resp.Body).Decode(&clubs); err != nil {
		return nil, fmt.Errorf("bad JSON: %w", err)
	}

	return clubs, nil
}

func addRemoveFromBag(authToken, clubName string, inBag bool) error {
	reqBody, err := json.Marshal(map[string]interface{}{
		"clubName": clubName,
		"inBag":    inBag,
	})
	if err != nil {
		return fmt.Errorf("failed to encode body: %w", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/bag/club", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}

	return nil
}

func makeBagScreen(win fyne.Window, authToken string) fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("My Bag", makeMyBagTab(win, authToken)),
		container.NewTabItem("All Clubs", makeAllClubsTab(win, authToken)),
	)
	backBtn := widget.NewButton("Back", func() {
		win.SetContent(makeMenuScreen(win, authToken))
	})

	form := container.NewVBox(
		tabs,
		backBtn,
	)
	return form
}

func makeMyBagTab(win fyne.Window, authToken string) fyne.CanvasObject {
	clubsLabel := widget.NewLabel("loading...")
	content := container.NewVBox(clubsLabel)

	go func() {
		clubs, err := getMyBag(authToken)
		if err != nil {
			clubsLabel.SetText(fmt.Sprintf("Error: %v", err))
			return
		}

		clubList := make([]fyne.CanvasObject, 0, len(clubs))
		for _, c := range clubs {
			clubList = append(clubList,
				widget.NewLabel(fmt.Sprintf("%s (%s): %d m", c.ClubName, c.ClubType, c.Distance)),
			)
		}
		content.Objects = clubList
		content.Refresh()
	}()

	return content
}

func categorizeClubs(clubs []Club) map[string][]Club {
	categories := make(map[string][]Club)
	for _, c := range clubs {
		categories[c.ClubType] = append(categories[c.ClubType], c)
	}
	return categories
}

func buildClubRow(c Club, authToken string, refresh func()) fyne.CanvasObject {
	status := "out of bag"
	btnLabel := "Add"
	newStatus := true

	if c.InBag {
		status = "In bag"
		btnLabel = "Remove"
		newStatus = false
	}

	return container.NewHBox(
		widget.NewLabel(fmt.Sprintf("%s: %d m", c.ClubName, c.Distance)),
		widget.NewLabel(status),
		widget.NewButton(btnLabel, func() {
			go func() {
				addRemoveFromBag(authToken, c.ClubName, newStatus)
				refresh()
			}()
		}),
	)
}

func makeAllClubsTab(win fyne.Window, authToken string) fyne.CanvasObject {
	content := container.NewVBox(widget.NewLabel("loading..."))

	var refresh func()
	refresh = func() {
		go func() {
			clubs, err := getAllClubs(authToken)
			ui := []fyne.CanvasObject{}
			if err != nil {
				ui = append(ui, widget.NewLabel(fmt.Sprintf("Error: %v", err)))
			} else {
				clubsByType := categorizeClubs(clubs)
				for _, t := range typeOrder {
					list := clubsByType[t]
					if len(list) == 0 {
						continue
					}

					ui = append(ui, widget.NewLabelWithStyle(t, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
					for _, c := range list {
						row := buildClubRow(c, authToken, refresh)
						ui = append(ui, row)
					}
				}
			}
			time.AfterFunc(0, func() {
				content.Objects = ui
				content.Refresh()
			})
		}()
	}
	refresh()
	return content
}
