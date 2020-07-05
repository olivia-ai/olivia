package language

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivia-ai/olivia/user"

	"github.com/olivia-ai/olivia/util"
)

// Movie is the serializer from res/datasets/movies.csv
type Movie struct {
	Name   string
	Genres []string
	Rating float64
}

var (
	// MoviesGenres initializes movies genres in different languages
	MoviesGenres = map[string][]string{
		"en": {
			"Action", "Adventure", "Animation", "Children", "Comedy", "Crime", "Documentary", "Drama", "Fantasy",
			"Film-Noir", "Horror", "Musical", "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western",
		},
		"de": {
			"Action", "Abenteuer", "Animation", "Kinder", "Komödie", "Verbrechen", "Dokumentarfilm", "Drama", "Fantasie",
			"Film-Noir", "Horror", "Musical", "Mystery", "Romance", "Sci-Fi", "Thriller", "Krieg", "Western",
		},
		"fr": {
			"Action", "Aventure", "Animation", "Enfant", "Comédie", "Crime", "Documentaire", "Drama", "Fantaisie",
			"Film-Noir", "Horreur", "Musical", "Mystère", "Romance", "Science-fiction", "Thriller", "Guerre", "Western",
		},
		"es": {
			"Acción", "Aventura", "Animación", "Infantil", "Comedia", "Crimen", "Documental", "Drama", "Fantasía",
			"Cine Negro", "Terror", "Musical", "Misterio", "Romance", "Ciencia Ficción", "Thriller", "Guerra", "Western",
		},
		"ca": {
			"Acció", "Aventura", "Animació", "Nen", "Comèdia", "Crim", "Documental", "Drama", "Fantasia",
			"Film-Noir", "Horror", "Musical", "Misteri", "Romanç", "Ciència-ficció", "Thriller", "War", "Western",
		},
		"it": {
			"Azione", "Avventura", "Animazione", "Bambini", "Commedia", "Poliziesco", "Documentario", "Dramma", "Fantasia",
			"Film-Noir", "Orrore", "Musical", "Mistero", "Romantico", "Fantascienza", "Giallo", "Guerra", "Western",
		},
		"nl": {
			"Actie", "Avontuur", "Animatie", "Kinderen", "Komedie", "Krimi", "Documentaire", "Drama", "Fantasie",
			"Film-Noir", "Horror", "Musical", "Mysterie", "Romantiek", "Sci-Fi", "Thriller", "Oorlog", "Western",
		},
		"el": {
			"Δράση", "Περιπέτεια", "Κινούμενα Σχέδια", "Παιδικά", "Κωμωδία", "Έγκλημα", "Ντοκιμαντέρ", "Δράμα", "Φαντασία",
			"Film-Noir", "Τρόμου", "Μουσική", "Μυστηρίου", "Ρομαντική", "Επιστημονική Φαντασία", "Θρίλλερ", "Πολέμου", "Western",
		},
	}
	movies = SerializeMovies()
)

// SerializeMovies retrieves the content of res/datasets/movies.csv and serialize it
func SerializeMovies() (movies []Movie) {
	path := "res/datasets/movies.csv"
	bytes, err := os.Open(path)
	if err != nil {
		bytes, _ = os.Open("../" + path)
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

// SearchMovie search a movie for a given genre
func SearchMovie(genre, userToken string) (output Movie) {
	for _, movie := range movies {
		userMovieBlacklist := user.GetUserInformation(userToken).MovieBlacklist
		// Continue if the movie is not from the request genre or if this movie has already been suggested
		if !util.Contains(movie.Genres, genre) || util.Contains(userMovieBlacklist, movie.Name) {
			continue
		}

		if reflect.DeepEqual(output, Movie{}) || movie.Rating > output.Rating {
			output = movie
		}
	}

	// Add the found movie to the user blacklist
	user.ChangeUserInformation(userToken, func(information user.Information) user.Information {
		information.MovieBlacklist = append(information.MovieBlacklist, output.Name)
		return information
	})

	return
}

// FindMoviesGenres returns an array of genres found in the entry string
func FindMoviesGenres(locale, content string) (output []string) {
	for i, genre := range MoviesGenres[locale] {
		if LevenshteinContains(strings.ToUpper(content), strings.ToUpper(genre), 2) {
			output = append(output, MoviesGenres["en"][i])
		}
	}

	return
}
