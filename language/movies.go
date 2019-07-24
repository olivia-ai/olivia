package language

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	Name   string
	Genres []string
	Rating float64
}

var genres = []string{
	"Action", "Adventure", "Animation", "Children", "Comedy", "Crime", "Documentary", "Drama", "Fantasy",
	"Film-Noir", "Horror", "Musical", "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western",
}

func SerializeMovies() (movies []Movie) {
	path := "res/movies.csv"
	bytes, err := os.Open(path)
	if err != nil {
		bytes, err = os.Open("../" + path)
	}

	reader := csv.NewReader(bufio.NewReader(bytes))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Convert the string to a float
		rating, _ := strconv.ParseFloat(line[3], 64)

		movies = append(movies, Movie{
			Name:   line[1],
			Genres: strings.Split(line[2], "|"),
			Rating: rating,
		})
	}

	return
}

// FindMoviesGenres returns an array of genres found in the entry string
func FindMoviesGenres(content string) (output []string) {
	for _, genre := range genres {
		if strings.Contains(strings.ToUpper(content), strings.ToUpper(genre)) {
			output = append(output, genre)
		}
	}

	return
}
