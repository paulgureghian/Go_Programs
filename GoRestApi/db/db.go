package db

import "github.com/paulgureghian/Go_Programs/GoRestApi/models"

var MovieList = []models.Movie{

	{
		Id:         "1",
		Title:      "The Godfather",
		Year:       "1972",
		Category:   "Crime/Drama",
		ImdbRating: "9.2/10",
	},

	{

		Id:         "2",
		Title:      "The English Patient",
		Year:       "1996",
		Category:   "Romance/Drama",
		ImdbRating: "7.4/10",
	},

	{
		Id:         "3",
		Title:      "The Great Gatsby",
		Year:       "2013",
		Category:   "Romance/Drama",
		ImdbRating: "7.2/10",
	},
}
