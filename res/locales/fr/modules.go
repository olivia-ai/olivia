package fr

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("fr", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Quelle est la superficie de ",
				"Donne moi la superficie de ",
			},
			Responses: []string{
				"La superficie %s est de %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Quelle est la capitale de ",
				"Donne moi la capitale de ",
			},
			Responses: []string{
				"La capitale %s est %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Quelle monnaie est utilisée en ",
				"Donne moi la monnaie utilisée en",
				"Donne moi la monnaie de ",
				"Quelle est la monnaie de ",
			},
			Responses: []string{
				"La monnaie %s est %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Donne moi le résultat de ",
				"Calcule ",
			},
			Responses: []string{
				"Le résultat est %s",
				"Cela fait %s",
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
				"J'aime les films d'aventure, animation",
				"Je regarde des films de science-fiction",
			},
			Responses: []string{
				"Très bon choix! Je sauvegarde ces informations dans votre client.",
				"C'est compris, j'envoie ces informations dans votre client",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Trouves un film de ",
				"Donne moi un film de ",
				"J'aimerais regarder un film de ",
				"Je souhaiterais regarder un film de ",
			},
			Responses: []string{
				"J'ai trouvé ceci pour vous “%s” qui est noté %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"J'ai déjà vu ce film",
				"J'ai déjà regardé ce film",
			},
			Responses: []string{
				"Je vois, en voici un autre “%s” noté %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
			Context:  modules.MoviesTag,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Je m'ennuie",
				"Je ne sais pas quoi faire",
			},
			Responses: []string{
				"Je vous propose un film %s “%s” noté %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Connais-tu mon nom",
			},
			Responses: []string{
				"Votre nom est %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Mon nom est ",
				"Tu peux m'appeler ",
			},
			Responses: []string{
				"Super! Bonjour %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Donne moi un nombre aléatoire",
				"Génère un nombre aléatoire",
			},
			Responses: []string{
				"Le nombre est %s",
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
				"Rappelle moi de cuisiner à 8 heures",
				"Rappelle moi d'appeler maman demain",
				"Note que j'ai un examen",
				"Rappelle moi que j'ai un appel de travail demain à 8h",
			},
			Responses: []string{
				"Noté! Je vous rappelerai: “%s” pour le %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"Qu'est-ce que je t'ai demandé de te rappeler",
				"Donnes moi mes rappels",
			},
			Responses: []string{
				"Vous m'avez demandé de me rappeler de ces choses:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Voici mes identifiants spotify",
			},
			Responses: []string{
				"Connexion en cours",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Joue de sur Spotify",
			},
			Responses: []string{
				"Je joue %s de %s sur Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["fr"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
// See https://www.lepointdufle.net/ressources_fle/pays_regle.htm
func ArticleCountries(name string) string {
	exceptions := map[string]string{
		"Belize":              "du ",
		"Cambodge":            "du ",
		"Mexique":             "du ",
		"Mozambique":          "du ",
		"Suriname":            "du ",
		"Zimbabwe":            "du ",
		"Bahreïn":             "du ",
		"Chypre":              "de ",
		"Cuba":                "de ",
		"Djibouti":            "du ",
		"Haïti":               "de ",
		"Israël":              "d'",
		"Madagascar":          "de ",
		"Malte":               "de ",
		"Maurice":             "de ",
		"Monaco":              "de ",
		"Oman":                "d'",
		"Singapour":           "de ",
		"Trinité-et-Tobago":   "de ",
		"Vanuatu":             "du ",
		"Bahamas":             "des ",
		"Bermudes":            "des ",
		"Comores":             "des ",
		"Émirats arabes unis": "des ",
		"États-Unis":          "des ",
		"Fidgi":               "des ",
		"Îles Féroé":          "des ",
		"Pays Bas":            "des ",
		"Philippines":         "des ",
		"Seychelles":          "des ",
	}

	vowels := regexp.MustCompile("[aeiou]")

	if exceptions[name] != "" {
		name = exceptions[name] + name
	} else if vowels.FindStringIndex(strings.ToLower(name))[0] == 0 {
		name = "de l'" + name
	} else {
		lastLetter := regexp.MustCompile(".+e")
		article := "du "

		if lastLetter.MatchString(strings.ToLower(name)) {
			article = "de la "
		}

		name = fmt.Sprintf("%s%s", article, name)
	}

	return name
}
