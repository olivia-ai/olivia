package en

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("de", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Wie groß ist das Land ",
			},
			Responses: []string{
				"Die Größe von %s ist %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Was ist die Hauptstadt von ",
				"Wie ist die Hauptstadt von ",
				"Gib mir die Hauptstadt von ",
			},
			Responses: []string{
				"Die Hauptstadt von %s ist %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Welche Währung wird benutzt in ",
				"Gib mir die Währung von ",
				"Wie ist die Währung von ",
			},
			Responses: []string{
				"Die Währung von %s ist %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Gib mir das Ergebnis von ",
				"Rechne ",
			},
			Responses: []string{
				"Das Ergebnis ist %s",
				"Das macht %s",
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
				"Ich mag Abenteuer- und Animationfilme",
				"Ich schaue gerne Sci-Fi Filme",
			},
			Responses: []string{
				"Tolle Auswahl! Ich speichere sie in Ihrem Client.",
				"Verstanden, Ich sende diese Informationen an Ihren Client.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Kannst du mir den Film finden",
				"Gib mir den Film von",
				"Finde mir den Film",
				"Ich würde gerne den Film schauen von ",
			},
			Responses: []string{
				"Ich habe Ihn für dich gefunden “%s” die Bewertung ist %.02f/5",
				"Klar, Ich habe diesen Film gefunden “%s” welcher mit %.02f/5 bewertet ist.",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Ich habe diesen Film schon gesehen",
				"Ich habe diesen Film bereits gesehen",
				"Oh, ich habe diesen Film schon gesehen",
			},
			Responses: []string{
				"Ok, hier ist noch einer “%s” welcher mit %.02f/5 bewertet ist.",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Ich langweile mich",
				"Ich weiß nicht was ich tun soll",
			},
			Responses: []string{
				"Ich schlage Dir einen Film vor %s “%s” welcher mit %.02f/5 bewertet ist.",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Weißt du meinen Namen",
				"Weißt du wie ich Heiße",
			},
			Responses: []string{
				"Dein Name ist %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Mein Name ist ",
				"Nenn mich ",
			},
			Responses: []string{
				"Super! Hallo %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Gib mir eine Zufallszahl",
				"Generiere mir eine Zufallszahl",
			},
			Responses: []string{
				"Die Zufallszahl ist %s",
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
				"Erinnern Sie mich daran, um 20 Uhr ein Frühstück vorzubereiten",
				"Erinnere mich daran, Mama am Dienstag anzurufen",
				"Notiere dass ich eine Prüfung habe",
				"Erinnern Sie mich daran, dass ich morgen um 21 Uhr eine Telefonkonferenz habe",
			},
			Responses: []string{
				"Gemerkt! Ich erinnere dich um “%s” für %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"Woran habe ich dich gebeten, mich zu erinnern",
				"Gib mir alle Erinnerungen",
			},
			Responses: []string{
				"Sie haben mich gebeten, Sie an folgende Sachen zu erinnern:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Hier sind meine Spotify Token",
				"Meine Spotify Zugangsdaten",
			},
			Responses: []string{
				"Sie werden eingeloggt",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Spiele auf Spotify",
				"Spiele von Spotify",
			},
			Responses: []string{
				"Spiele %s von %s auf Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["de"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "vereinigten Staaten" {
		return "die " + name
	}

	return name
}
