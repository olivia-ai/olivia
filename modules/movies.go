package modules

import (
	"github.com/olivia-ai/olivia/user"
	"github.com/olivia-ai/olivia/util"

	"github.com/olivia-ai/olivia/language"
)

var genresTag = "movies genres"

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
		information.MovieGenres = append(information.MovieGenres, genres...)
		return information
	})

	return genresTag, response
}
