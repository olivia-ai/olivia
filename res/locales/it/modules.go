package it

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("it", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Qual è la superficie ",
				"Dammi la superficie ",
			},
			Responses: []string{
				"La superficie %s è %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Qual è la capitale ",
				"Dimmi la capitale ",
				"Dammi la capitale ",
				"Come si chiama la capitale",
			},
			Responses: []string{
				"La capitale %s è %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Che moneta è usata in ",
				"Dimmi la moneta usata in ",
				"Qual è la moneta in ",
				"Che moneta si usa in ",
			},
			Responses: []string{
				"La moneta %s è %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Dammi il risultato di ",
				"Calcola ",
			},
			Responses: []string{
				"Il risultato è %s",
				"Il risultato ottenuto è %s",
			},
			Replacer: modules.MathReplacer,
		},

		// MOVIES
		// A translation of movies genres is also required in `language/movies.go`, please don't forget
		// to translate it.
		// Otherwise, remove the registration of the Movies modules in this file.

		/*{
			Tag: modules.GenresTag,
			Patterns: []string{
				"I like movies of adventure, animation",
				"I watch movies of sci-fi",
			},
			Responses: []string{
				"Great choices! I save them into your client.",
				"Understood, I send this information to your client.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
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
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"I already saw this movie",
				"I have already watched this film",
				"Oh I have already watched this movie",
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
				"I propose you a movie of %s “%s” which is rated %.02f/5",
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
				"Mi hai chiesto di ricordarti queste cose:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.
		*/
		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"I miei token di spotify",
				"I miei segreti spotify",
			},
			Responses: []string{
				"Login in corso",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Riproduci su Spotify",
			},
			Responses: []string{
				"In riproduzione %s da %s su Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["it"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	firstCharacters := name[0]
	lastCharacters := name[len(name)-1]

	switch {
	case name == "Stati Uniti":
		return "degli " + name
	case name == "Filippine":
		return "delle " + name
	case firstCharacters == 'A' || firstCharacters == 'E' || firstCharacters == 'I' ||
		firstCharacters == 'O' || firstCharacters == 'U':
		return "dell'" + name
	case lastCharacters == 'a':
		return "della " + name
	default:
		return "del " + name
	}
}
