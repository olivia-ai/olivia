package en

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("el", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Ποιά είναι η περιοχή της",
				"Ποιά είναι η περιοχή του",
				"Πες μου την περιοχή της",
				"Πες μου την περιοχή του",

			},
			Responses: []string{
				"Η περιοχή της/του %s είναι %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Ποια είναι η πρωτεύουσα της",
				"Ποια είναι η πρωτεύουσα του",
				"Πες μου την πρωτεύουσα της",
				"Πες μου την πρωτεύουσα του",
			},
			Responses: []string{
				"Η πρωτεύουσα της/του %s είναι η/το %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Τι νόμισμα χρησιμοποιείται στην",
				"Τι νόμισμα χρησιμοποιείται στον",
				"Τι νόμισμα χρησιμοποιείται στο",
				"Πες μου τι νόμισμα χρησιμοποείται στην",
				"Πες μου τι νόμισμα χρησιμοποείται στον",
				"Πες μου τι νόμισμα χρησιμοποείται στη",
			},
			Responses: []string{
				"Το νόμισμα της/του %s είναι %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Πες μου το αποτελέσμα του",
				"Υπολόγισε",
			},
			Responses: []string{
				"Το αποτέλεσμα είναι %s",
				"Αυτό μας κάνει %s",
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
				"Μου αρέσουν οι ταινίες δράσης και κινουμένων σχεδίων",
				"Βλέπω ταινίες επιστημονικής φαντασίας",
			},
			Responses: []string{
				"Πολύ καλές επιλογές. Τις αποθήκευσα στα στοιχεία σου.",
				"Κατάλαβα, στέλνω τις προτιμήσεις αυτές στα στοιχεία σου.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Μπορείς να μου βρεις την ταινία",
				"Βρες μου μια ταινία",
				"Ψάξε μια ταινία",
				"Θα ήθελα να δω μια ταινία",
			},
			Responses: []string{
				"Βρήκα αυτήν για σένα “%s” που έχει αξιολογηθεί με %.02f/5",
				"Αμέ, βρήκα αυτή την ταινία “%s” με αξιολόγηση %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Έχω ήδη δει αυτή την ταινία",
				"Την έχω δεί",
				"Αυτή την ταινία την έχω δει",
			},
			Responses: []string{
				"Α δεν ήξερα, να μια άλλη “%s” με βαθμολογία %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Βαριέμαι",
				"Δεν ξέρω τι να κάνω",
			},
			Responses: []string{
				"Σου προτείνω να δεις αυτή την ταινία %s “%s” που είναι βαθμολογημένη με %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Ξέρεις το όνομα μου?",
			},
			Responses: []string{
				"Αμέ, το όνομα σου είναι %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Το όνομα μου είναι",
				"Μπορείς να με φωνάζεις",
			},
			Responses: []string{
				"Τέλεια! Γεια σου %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Δώσε μου ένα τυχαίο αριθμό",
				"Πάραξε ένα τυχαίο αριθμό",
			},
			Responses: []string{
				"Ο τυχαίος αριθμός είναι:  %s",
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
				"Υπενθύμισε μου να μαγειρέψω πρωϊνό στις 8πμ",
				"Θύμισε μου να καλέσω τη μαμά την Τρίτη",
				"Σημείωσε ότι έχω εξετάσεις",
				"Θύμισε μου ότι έχω ένα συνέδριο αύριο το πρωί στις 9πμ",
			},
			Responses: []string{
				"Σημειώθηκε! Θα σε ενημερώσω για : “%s” για τις %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"Τι σου ζήτησα να μου θυμίσεις",
				"Πες μου τις υπενθυμίσεις μου",
			},
			Responses: []string{
				"Μου ζήτησες να θυμάμαι τα παρακάτω:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Εδώ είναι τα στοιχεία του Spotify μου",
				"Οι κωδικοί του Spotify μου",
			},
			Responses: []string{
				"Είσοδος σε εξέλιξη",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Παίξε από το Spotify",
			},
			Responses: []string{
				"Παίζει %s από %s στο Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Πες μου ένα αστείο",
				"Κάνε με να γελάσω",
			},
			Responses: []string{
				"Εδώ είσαι, %s",
				"Να ένα καλό, %s",
			},
			Replacer: modules.JokesReplacer,
		},
		{
			Tag: modules.AdvicesTag,
			Patterns: []string{
				"Δώσε μου μια συμβουλή",
				"Συμβούλεψε με",
			},
			Responses: []string{
				"Εδώ είσαι, %s",
				"Να ένα καλό, %s",
				"Άκου προσεκτικά, %s",
			},
			Replacer: modules.AdvicesReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["el"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "Greece" {
		return "Η " + name
	}

	return name
}
