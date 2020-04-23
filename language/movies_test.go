package language

import (
	"reflect"
	"testing"
)

func TestSerializeMovies(t *testing.T) {
	movies := SerializeMovies()
	excepted := "Toy Story (1995)"

	if movies[0].Name != excepted {
		t.Errorf("SerializeMovies() failed, excepted %s got %s.", excepted, movies[0].Name)
	}
}

func TestSearchMovie(t *testing.T) {
	movie := SearchMovie("Adventure", "0")
	excepted := "2001: A Space Odyssey (1968)"

	if movie.Name != excepted {
		t.Errorf("SearchMovie() failed, excepted %s got %s.", excepted, movie.Name)
	}
}

func TestFindMoviesGenres(t *testing.T) {
	sentence := "I like movies of adventure, sci-fi"
	excepted := []string{"Adventure", "Sci-Fi"}
	genres := FindMoviesGenres("en", sentence)

	if !reflect.DeepEqual(excepted, genres) {
		t.Errorf("FindMoviesGenres() failed, excepted %s got %s.", excepted, genres)
	}
}
