package modules

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/olivia-ai/olivia/user"
	"github.com/olivia-ai/olivia/util"

	"github.com/olivia-ai/olivia/language"
)

var (
	genresTag        = "movies genres"
	moviesTag        = "movies search"
	moviesAlreadyTag = "already seen movie"
	moviesDataTag    = "movies search from data"
)

func init() {
	RegisterModule(Module{
		Tag: genresTag,
		Patterns: []string{
			"I like movies of",
			"I watch movies of",
		},
		Responses: []string{
			"Great choices! I save them into your client.",
			"Understood, I send this information to your client.",
		},
		Replacer: GenresReplacer,
	})

	RegisterModule(Module{
		Tag: moviesTag,
		Patterns: []string{
			"Can you find me a movie of",
			"Give me a movie of",
			"Find me a film of",
			"I would like to watch a movie of",
		},
		Responses: []string{
			"I found this for you “%s” which is rated %.02f/5",
			"Sure, I found this movie “%s” rated %.02f/5",
		},
		Replacer: MovieSearchReplacer,
	})

	RegisterModule(Module{
		Tag: moviesAlreadyTag,
		Patterns: []string{
			"I already saw this movie",
			"I have already watched this film",
			"Oh I have already watched this movie",
		},
		Responses: []string{
			"Oh I see, here's another one “%s” which is rated %.02f/5",
		},
		Replacer: MovieSearchReplacer,
	})

	RegisterModule(Module{
		Tag: moviesDataTag,
		Patterns: []string{
			"I'm bored",
			"I don't know what to do",
		},
		Responses: []string{
			"I propose you a movie of %s “%s” which is rated %.02f/5",
		},
		Replacer: MovieSearchFromInformationReplacer,
	})
}

// GenresReplacer gets the genre specified in the message and adds it to the user information.
// See modules/modules.go#Module.Replacer() for more details.
func GenresReplacer(entry, response, token string) (string, string) {
	genres := language.FindMoviesGenres(entry)

	// If there is no genres then reply with a message from res/messages.json
	if len(genres) == 0 {
		responseTag := "no genres"
		return responseTag, util.GetMessage(responseTag)
	}

	// Change the user information to add the new genres
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		for _, genre := range genres {
			// Append the genre only is it isn't already in the information
			if util.Contains(information.MovieGenres, genre) {
				continue
			}

			information.MovieGenres = append(information.MovieGenres, genre)
		}
		return information
	})

	return genresTag, response
}

// MovieSearchReplacer replaces the patterns contained inside the response by the movie's name
// and rating from the genre specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func MovieSearchReplacer(entry, response, token string) (string, string) {
	genres := language.FindMoviesGenres(entry)

	// If there is no genres then reply with a message from res/messages.json
	if len(genres) == 0 {
		responseTag := "no genres"
		return responseTag, util.GetMessage(responseTag)
	}

	movie := language.SearchMovie(genres[0], token)

	return moviesTag, fmt.Sprintf(response, movie.Name, movie.Rating)
}

// MovieSearchFromInformationReplacer replaces the patterns contained inside the response by the movie's name
// and rating from the genre in the user's information.
// See modules/modules.go#Module.Replacer() for more details.
func MovieSearchFromInformationReplacer(_, response, token string) (string, string) {
	// If there is no genres then reply with a message from res/messages.json
	genres := user.GetUserInformation(token).MovieGenres
	if len(genres) == 0 {
		responseTag := "no genres saved"
		return responseTag, util.GetMessage(responseTag)
	}

	movie := language.SearchMovie(genres[rand.Intn(len(genres))], token)
	genresJoined := strings.Join(genres, ", ")
	return moviesDataTag, fmt.Sprintf(response, genresJoined, movie.Name, movie.Rating)
}
