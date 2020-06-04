package en

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("nl", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Wat is het oppervlakte van ",
				"Geef mij het oppervlakte van ",
			},
			Responses: []string{
				"Het oppervalkte van %s is %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Wat is de hoofdstad van ",
				"Geef mij de hoofdstad van ",
			},
			Responses: []string{
				"De hoofdstad van %s is %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Welke valuta wordt gebruikt in ",
				"Geef mij de valuta die gebruikt wordt in ",
				"Geef mij de valuta van ",
				"Wat is de valuta van ",
				"Welke munteenheid wordt gebruikt in ",
				"Geef mij de munteenheid die gebruikt wordt in ",
				"Geef mij de munteenheid van ",
				"Wat is de munteenheid van ",
			},
			Responses: []string{
				"De valuta van %s is %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Geef mij het resultaat van ",
				"Geef mij de uitkomst van ",
				"Bereken ",
			},
			Responses: []string{
				"Het resultaat is %s",
				"De uitkomst is %s",
				"Dat maakt %s",
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
				"Ik hou van films in het genre avontuur en animatie",
				"Ik kijk science fiction films",
			},
			Responses: []string{
				"Goede keuzes! Ik sla ze in je applicatie op.",
				"Begrepen, ik stuur dit naar je applicatie.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Kan je een film vinden",
				"Geef mij een film van",
				"Zoek een film voor mij van",
				"Ik wil een film kijken van",
			},
			Responses: []string{
				"Ik vond dit voor je “%s” met de beoordeling %.02f/5",
				"Natuurlijk, ik vond de film “%s” beoordeling %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Ik heb deze film al gezien",
				"Oh ik heb deze film al gekeken",
			},
			Responses: []string{
				"Ik begrijp je, hier is er nog een. “%s” welke %.02f/5 is beoordeelt",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Ik verveel mij",
				"Ik weet niet wat ik moet doen",
			},
			Responses: []string{
				"Ik stel je de volgende film voor %s “%s” welke %.02f/5 is beoordeelt",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Weet je mijn naam?",
			},
			Responses: []string{
				"Jouw naam is %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Mijn naam is ",
				"Noem mij ",
			},
			Responses: []string{
				"Geweldig! Hallo %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Geef mij een willekeurig nummer",
				"Genereer een willekeurig nummer",
			},
			Responses: []string{
				"Het nummer is %s",
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
				"Herinner mij er aan om ontbijt te maken om 8 uur",
				"Herinner mij er aan om mijn moeder te bellen",
				"Merk op dat ik een examen heb",
				"Herinner mij aan mijn conference call morgen om 9 uur",
			},
			Responses: []string{
				"Genoteerd! Ik zal je herinneren om: “%s” voor %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"What heb ik je gevraagd om te onthouden",
				"Vertel mij waar je mij aan zult herinneren",
			},
			Responses: []string{
				"Je hebt mij gevraagd het volgende te onthouden:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Hier zijn mijn spotify tokens",
				"Mijn spotify tokens",
			},
			Responses: []string{
				"Login is bezig",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Speel van Spotify",
			},
			Responses: []string{
				"Speelt %s door %s op Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Vertel mij een mop",
				"Laat mij lachen",
			},
			Responses: []string{
				"Hiero, %s",
				"Hier is er een, %s",
			},
			Replacer: modules.JokesReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["nl"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "Verenigde Staten" {
		return "de " + name
	}

	return name
}
