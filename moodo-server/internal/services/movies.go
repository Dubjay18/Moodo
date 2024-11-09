package services

import (
	"fmt"
	"moodo-server/internal/lib"

	"os"

	"github.com/go-resty/resty/v2"
)

// TMDb API configuration
const TMDB_BASE_URL = "https://api.themoviedb.org/3"

var apiKey = os.Getenv("TMDB_API_KEY")

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	GenreIDs []int  `json:"genre_ids"`
}

type TMDBResponse struct {
	Results      []Movie `json:"results"`
	Page         int     `json:"page"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// Fetches movies by genre from TMDb with pagination support
func fetchMoviesByGenre(genreID int, page int) ([]Movie, int, int, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"api_key":     apiKey,
			"with_genres": fmt.Sprintf("%d", genreID),
			"language":    "en-US",
			"page":        fmt.Sprintf("%d", page),
		}).
		SetResult(&TMDBResponse{}).
		Get(TMDB_BASE_URL + "/discover/movie")

	if err != nil {
		return nil, 0, 0, err
	}

	tmdbResp := resp.Result().(*TMDBResponse)
	return tmdbResp.Results, tmdbResp.Page, tmdbResp.TotalPages, nil
}

// Mood-based recommendation algorithm with pagination
func RecommendMoviesByMood(mood string, page int) ([]Movie, int, int, error) {
	var recommendedMovies []Movie
	var totalPages int

	// Map mood to genre IDs
	if genreIDs, found := lib.MoodToGenreMapping[mood]; found {
		for _, genreID := range genreIDs {
			movies, currentPage, totalPageCount, err := fetchMoviesByGenre(genreID, page)
			if err != nil {
				return nil, 0, 0, err
			}
			recommendedMovies = append(recommendedMovies, movies...)
			totalPages = totalPageCount

			// Break after fetching one page per genre for simplicity
			// You could customize this logic to adjust how many pages are fetched per genre
			if currentPage >= totalPages {
				break
			}
		}
	} else {
		return nil, 0, 0, fmt.Errorf("mood not found")
	}

	// Return the list of movies with the pagination details
	return recommendedMovies, page, totalPages, nil
}
