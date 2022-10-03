package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radenrishwan/otakudesu-api/scrape"
	"github.com/radenrishwan/otakudesu-api/utils"
)

func main() {
	r := mux.NewRouter()

	r.Use(utils.ErrorHandler)

	r.HandleFunc("/api/home", scrape.HomePage)
	r.HandleFunc("/api/anime-list", scrape.AnimeList)
	r.HandleFunc("/api/genres", scrape.AnimeGenreList)
	r.HandleFunc("/api/anime/ongoing", scrape.AnimeOnGoing)
	r.HandleFunc("/api/anime/complete", scrape.AnimeComplete)
	r.HandleFunc("/api/anime/genre/{genre}", scrape.AnimeFindByGenre)
	r.HandleFunc("/api/anime/{id}", scrape.AnimeDetail)
	r.HandleFunc("/api/episode/{id}", scrape.EpisodeDetail)
	r.HandleFunc("/api/search", scrape.FindAnime)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := map[string]any{
			"message": "Hi, Welcome to Unofficial Otakudesu API 🐔",
			"author":  "seior",
			"github":  "https://github.com/radenrishwan/otakudesu-api",
		}

		bytes, err := json.Marshal(utils.DefaultResponse[any]{
			Code: 200,
			Data: result,
		})
		utils.PanicIfError(err)

		utils.NewSuccessResponse(string(bytes), w, r)
	})

	log.Fatalln(http.ListenAndServe(":8080", r))
}
