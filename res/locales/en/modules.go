package en

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("en", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"What is the area of ",
				"Give me the area of ",
			},
			Responses: []string{
				"The area of %s is %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"What is the capital of ",
				"What's the capital of ",
				"Give me the capital of ",
			},
			Responses: []string{
				"The capital of %s is %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Which currency is used in ",
				"Give me the used currency of ",
				"Give me the currency of ",
				"What is the currency of ",
			},
			Responses: []string{
				"The currency of %s is %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Give me the result of ",
				"Calculate ",
			},
			Responses: []string{
				"The result is %s",
				"That makes %s",
			},
			Replacer: modules.MathReplacer,
		},

		// MOVIES
		// A translation of movies genres is also required in `language/movies.go`, please don't forget
		// to translate it.
		// Otherwise, remove the registration of the Movies modules in this file.

		{
			Tag: modules.GenresTag,
			Patterns: []string{
				"My favorite movie genres are Comedy, Horror",
				"I like the Comedy, Horror genres",
                "I like movies about War",
                "I like Action movies",
			},
			Responses: []string{
				"Great choices! I saved this movie genre information to your client.",
				"Understood, I saved this movie genre information to your client.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Find me a movie about",
				"Give me a movie about",
				"Find me a film about",
			},
			Responses: []string{
				"I found the movie “%s” for you, which is rated %.02f/5",
				"Sure, I found this movie “%s”, which is rated %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"I already saw this movie",
				"I have already watched this film",
				"Oh I have already watched this movie",
				"I have already seen this movie",
			},
			Responses: []string{
				"Oh I see, here's another one “%s” which is rated %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"I'm bored",
				"I don't know what to do",
			},
			Responses: []string{
				"I propose you watch the %s movie “%s”, which is rated %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Do you know my name?",
			},
			Responses: []string{
				"Your name is %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"My name is ",
				"You can call me ",
			},
			Responses: []string{
				"Great! Hi %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Give me a random number",
				"Generate a random number",
			},
			Responses: []string{
				"The number is %s",
			},
			Replacer: modules.RandomNumberReplacer,
		},

		// REMINDERS
		// Translations are required in `language/date/date`, `language/date/rules` and in `language/reason`,
		// please don't forget to translate it.
		// Otherwise, remove the registration of the Reminders modules in this file.

		{
			Tag: modules.ReminderSetterTag,
			Patterns: []string{
				"Remind me to cook a breakfast at 8pm",
				"Remind me to call mom tuesday",
				"Note that I have an exam",
				"Remind me that I have a conference call tomorrow at 9pm",
			},
			Responses: []string{
				"Noted! I will remind you: “%s” for the %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"What did I ask for you to remember",
				"Give me my reminders",
			},
			Responses: []string{
				"You asked me to remember those things:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Here are my spotify tokens",
				"My spotify secrets",
			},
			Responses: []string{
				"Login in progress",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Play from on Spotify",
			},
			Responses: []string{
				"Playing %s from %s on Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Tell me a joke",
				"Make me laugh",
			},
			Responses: []string{
				"Here you go, %s",
				"Here's one, %s",
			},
			Replacer: modules.JokesReplacer,
		},
		{
			Tag: modules.AdvicesTag,
			Patterns: []string{
				"Give me an advice",
				"Advise me",
			},
			Responses: []string{
				"Here you go, %s",
				"Here's one, %s",
				"Listen closely, %s",
			},
			Replacer: modules.AdvicesReplacer,
},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["en"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "United States" {
		return "the " + name
	}

	return name
}
