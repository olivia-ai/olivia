package ca

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("ca", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Quina és l'àrea de ",
			},
			Responses: []string{
				"La superfície %s és de %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Quina és la capital de ",
				"Doneu-me la capital de ",
			},
			Responses: []string{
				"El capital %s és %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Quina moneda s'utilitza en ",
				"Doni'm la moneda utilitzada en ",
			},
			Responses: []string{
				"La moneda %s és %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Dóna'm el resultat de ",
				"Calcula ",
			},
			Responses: []string{
				"El resultat és %s",
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
				"M'agraden les pel·lícules d'aventura, l'animació",
				"Veig pel·lícules de ciència ficció",
			},
			Responses: []string{
				"Molt bona elecció! Deso aquesta informació al vostre client.",
				"Entès, envio aquesta informació al vostre client",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Cerqueu una pel·lícula de ",
				"Dóna'm una pel·lícula de ",
				"Voldria veure una pel·lícula ",
				"M'agradaria veure una pel·lícula ",
			},
			Responses: []string{
				"He trobat això per a vostès “%s”, que es nota %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Ja he vist aquesta pel·lícula",
			},
			Responses: []string{
				"Veig, aquí hi ha un altre “%s” notat %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
			Context:  modules.MoviesTag,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Estic avorrit",
				"No sé què fer",
			},
			Responses: []string{
				"Us ofereixo una pel·lícula %s “%s” classificada amb %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Ja coneixeu el meu nom?",
			},
			Responses: []string{
				"El vostre nom és %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Em dic ",
				"Podeu trucar-me ",
			},
			Responses: []string{
				"Genial! Hola %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Doneu-me un número aleatori",
				"Generar un número aleatori",
			},
			Responses: []string{
				"El nombre és %s",
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
				"Recordeu-me a cuinar a les 8",
				"Recorda'm que truqui a la mare demà",
				"Tingueu en compte que tinc un examen",
				"Recordeu-me que demà tinc una trucada laboral a les 8 del matí.",
			},
			Responses: []string{
				"Notat! Us recordaré: “%s” per a %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"Què et vaig demanar que recordés?",
				"Doneu-me els meus recordatoris",
			},
			Responses: []string{
				"Em vas demanar que recordi aquestes coses:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Aquí està el meu Spotify IDs",
			},
			Responses: []string{
				"Connexió en curs",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Juga x de x a Spotify",
			},
			Responses: []string{
				"Puc jugar %s de %s a Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["ca"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	vowels := regexp.MustCompile("[aeiou]")

	if vowels.FindStringIndex(strings.ToLower(name))[0] == 0 {
		name = "d'" + name
	} else {
		lastLetter := regexp.MustCompile(".+e")
		article := "del "

		if lastLetter.MatchString(strings.ToLower(name)) {
			article = "de "
		}

		name = fmt.Sprintf("%s%s", article, name)
	}

	return name
}
