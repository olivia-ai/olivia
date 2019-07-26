package modules

import (
	"fmt"

	"github.com/olivia-ai/olivia/user"
	"github.com/olivia-ai/olivia/util"

	"github.com/olivia-ai/olivia/language"
)

var (
	genresTag        = "movies genres"
	moviesTag        = "movies search"
	moviesAlreadyTag = "already seen movie"
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
			"Understood, I send these informations to your client.",
		},
		Replacer: GenresReplacer,
	})

	RegisterModule(Module{
		Tag: moviesTag,
		Patterns: []string{
			"Can you find me a movie of",
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
			"Oh I have already watched this movie",
		},
		Responses: []string{
			"Oh I see, here's another one “%s” which is rated %.02f/5",
		},
		Replacer: MovieSearchReplacer,
	})
}

func GenresReplacer(entry, response, token string) (string, string) {
	genres := language.FindMoviesGenres(entry)

	// If there is no genres then reply with a message from res/messages.json
	if len(genres) == 0 {
		responseTag := "no genres"
		return responseTag, util.GetMessage(responseTag)
	}

	// Change the user information to add the new genres
	user.ChangeUserInformations(token, func(information user.Information) user.Information {
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
