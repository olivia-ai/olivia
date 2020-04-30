package es

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("es", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"¿Cuál es el área de ",
				"Dame el área de ",
			},
			Responses: []string{
				"El área %s es %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"¿Cuál es la capital de ",
				"¿Cuál es la capital de ",
				"Dame el capital de ",
			},
			Responses: []string{
				"La capital %s es %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"¿Qué moneda se utiliza en ",
				"Dame la moneda usada de ",
				"¿Cuál es la moneda de ",
			},
			Responses: []string{
				"La moneda %s es %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Dame el resultado de ",
				"Calcule ",
			},
			Responses: []string{
				"El resultado es %s",
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
				"Me gustan las películas de aventura, la animación",
				"Veo películas de ciencia-ficción",
			},
			Responses: []string{
				"¡Grandes elecciones! Las guardo en tu cliente.",
				"Entendido, le envío esta información a su cliente.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"¿Puedes encontrarme una película de",
				"Dame una película de",
				"Encuéntrame una película de",
				"Me gustaría ver una película de",
			},
			Responses: []string{
				"Encontré esto para ustedes “%s” que está clasificado %.02f/5",
				"Claro, encontré esta película “%s” clasificada %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Ya he visto esta película",
			},
			Responses: []string{
				"Oh, ya veo, aquí hay otro “%s” que está clasificado como %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Estoy aburrido",
				"No sé qué hacer",
			},
			Responses: []string{
				"Te propongo una película de %s “%s” que está clasificada %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"¿Sabe mi nombre?",
			},
			Responses: []string{
				"¡Tu nombre es %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Mi nombre es ",
				"Puedes llamarme ",
			},
			Responses: []string{
				"¡Grandioso! Hola.",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Dame un número al azar",
				"Generar un número aleatorio",
			},
			Responses: []string{
				"El número es %s",
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
				"Recuérdame que prepare un desayuno a las 8:00",
				"Recuérdame que llame a mamá el martes.",
				"Tengan en cuenta que tengo un examen",
				"Recuérdame que tengo una conferencia telefónica mañana a las 9 p.m.",
			},
			Responses: []string{
				"¡Anotado! Te lo recordaré: “%s” para el %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"¿Qué te pedí que recordaras?",
				"Dame mis recordatorios",
			},
			Responses: []string{
				"Me pediste que recordara esas cosas:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Aquí están mis fichas de Spotify",
				"Mis secretos spotify",
			},
			Responses: []string{
				"Inicio de sesión en curso",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Juega desde en Spotify",
			},
			Responses: []string{
				"Jugando %s de %s en Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["es"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	lastCharacters := name[len(name)-2:]
	var article string

	if lastCharacters == "as" {
		article = "de las "
	} else if lastCharacters == "os" {
		article = "de los "
	} else if string(lastCharacters[1]) == "a" {
		article = "de "
	} else {
		article = "del "
	}

	return article + name
}
