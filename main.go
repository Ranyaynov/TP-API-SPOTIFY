package main

import (
    "encoding/json"
    "log"
    "net/http"
    "spotify-tp/templates"
)

// STRUCTS SPOTIFY SIMPLIFIÉS

type SpotifySearch struct {
    Artists struct {
        Items []struct {
            ID string `json:"id"`
        } `json:"items"`
    } `json:"artists"`
    Tracks struct {
        Items []Track `json:"items"`
    } `json:"tracks"`
}

type Album struct {
    Name   string `json:"name"`
    Images []struct {
        URL string `json:"url"`
    } `json:"images"`
    ReleaseDate string `json:"release_date"`
    TotalTracks int    `json:"total_tracks"`
}

type Track struct {
    Name   string `json:"name"`
    Album  Album  `json:"album"`
    Artists []struct {
        Name string `json:"name"`
    } `json:"artists"`
    ExternalURLs struct {
        Spotify string `json:"spotify"`
    } `json:"external_urls"`
}


// ROUTE 1 : ALBUMS DE DAMSO
func DamsoHandler(w http.ResponseWriter, r *http.Request) {

    searchURL := "https://api.spotify.com/v1/search?q=damso&type=artist"
    data, _ := SpotifyRequest(searchURL)

    var result SpotifySearch
    json.Unmarshal(data, &result)

    if len(result.Artists.Items) == 0 {
        http.Error(w, "Artiste non trouvé", 500)
        return
    }

    artistID := result.Artists.Items[0].ID
    albumsURL := "https://api.spotify.com/v1/artists/" + artistID + "/albums"

    albumsData, _ := SpotifyRequest(albumsURL)

    var albumResponse struct {
        Items []Album `json:"items"`
    }

    json.Unmarshal(albumsData, &albumResponse)

    templates.DamsoPage(albumResponse.Items).Render(r.Context(), w)
}


// ROUTE 2 : MUSIQUE MALADRESSE
func MaladresseHandler(w http.ResponseWriter, r *http.Request) {

    url := "https://api.spotify.com/v1/search?q=maladresse laylow&type=track"
    data, _ := SpotifyRequest(url)

    var result SpotifySearch
    json.Unmarshal(data, &result)

    if len(result.Tracks.Items) == 0 {
        http.Error(w, "Musique non trouvée", 500)
        return
    }

    track := result.Tracks.Items[0]

    templates.MaladressePage(track).Render(r.Context(), w)
}


func main() {
    http.HandleFunc("/album/damso", DamsoHandler)
    http.HandleFunc("/track/laylow", MaladresseHandler)

    log.Println("Serveur lancé sur : http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}