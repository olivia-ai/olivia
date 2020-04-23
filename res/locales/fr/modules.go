package fr

import (
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("fr", []modules.Module{
		// AREA
		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Quelle est la superficie de ",
				"Donne moi la superficie de ",
			},
			Responses: []string{
				"La superficie de %s est de %gkm²",
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
				"La capitale de %s est %s",
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
				"La monnaie utilisée en %s est %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
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

	// MOVIES
	language.MoviesGenres["fr"] = []string{
		"Action", "Aventure", "Animation", "Enfant", "Comédie", "Crime", "Documentaire", "Drama", "Fantaisie",
		"Film-Noir", "Horreur", "Musical", "Mystère", "Romance", "Science-fiction", "Thriller", "Guerre", "Western",
	}
}
