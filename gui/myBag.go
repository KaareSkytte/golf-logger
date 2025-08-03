package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

func makeBagScreen(win fyne.Window, authToken string) fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("My Bag", makeMyBagTab(win, authToken)),
		container.NewTabItem("All Clubs", makeAllClubsTab(win, authToken)),
	)

	return tabs
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

func makeAllClubsTab(win fyne.Window, authToken string) fyne.CanvasObject {
	clubsLabel := widget.NewLabel("loading...")
	content := container.NewVBox(clubsLabel)

	go func() {
		clubs, err := getAllClubs(authToken)
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
