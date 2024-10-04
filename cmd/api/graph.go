package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Group struct {
	Name        string `json:"name"`
	ID          int64  `json:"id"`
	ResourceURL string `json:"resource_url"`
}

type ArtistData struct {
	Name   string  `json:"name"`
	ID     int64   `json:"id"`
	Groups []Group `json:"groups"`
}

func (app *application) getGraphHandler(w http.ResponseWriter, r *http.Request) {
	artistId, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	endpoint := app.buildDiscogsEndpoint(pathForArtist(artistId))
	req, err := app.buildRequest("GET", endpoint)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		app.discogsApiErrorResponse(w, r)
		return
	}

	artistData := ArtistData{}
	err = app.readJSON(resp, &artistData)

	fmt.Printf("%#v", artistData)
}

func (app *application) buildDiscogsEndpoint(path string) string {
	var sb strings.Builder

	sb.WriteString(app.config.discogsApiUrl)
	sb.WriteString(path)

	return sb.String()
}

func pathForArtist(id int64) string {
	return fmt.Sprintf("artists/%v", id)
}
