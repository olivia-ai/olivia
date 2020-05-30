package tr

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("tr", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Alanı nedir ",
				"Bana alanını ver ",
			},
			Responses: []string{
				"%s bölgesinin alanı %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Başkenti ",
				"Başkenti neresidir",
				"ülkenin başkenti ",
			},
			Responses: []string{
				"%s ülkesinin başkenti %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Hangi para birimini kullanır ",
				"para birimini söyle ",
				"bana para birimini ver ",
				"para birimi nedir ",
			},
			Responses: []string{
				"%s ülkesinin para birimi %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Sonucunu ver ",
				"Hesapla ",
			},
			Responses: []string{
				"Sonuç %s",
				"Hesaplanan sonuç %s",
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
				"Macera, animasyon filmlerini severim",
				"Bilim kurgu filmleri izlerim",
			},
			Responses: []string{
				"Harika seçimler! Bunları müşterinize saklıyorum.",
				"Anlaşıldı, bu bilgileri müşterinize gönderiyorum.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Bana bir film bulabilir misin",
				"Bana bir film verir misin",
				"bu türlerde bir film bakar mısın",
				"Bu tarz filmler izlemekten hoşlanırım: ",
			},
			Responses: []string{
				"Bunları senin için buldum “%s” ve puanı %.02f/5 olanları listeledim.",
				"Tabii ki, Bu filmi buldum “%s” puanı da: %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Bu filmi çoktan gördüm",
				"Bu filmi çoktan izledim",
				"Ohooo, bu filmi çoktan izledim",
			},
			Responses: []string{
				"Ah anlıyorum, İşte burada tam da aynı tadda “%s” hem de puanı %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"sıkıldım",
				"Ne yapacağımı bilmiyorum",
			},
			Responses: []string{
				"Sana bir film öneriim mi %s “%s” hem de puanı %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Adımı biliyor musun acaba?",
			},
			Responses: []string{
				"Senin adın %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Benim adım ",
				"Bana şey diye de seslenebilirsin, ımmm ",
			},
			Responses: []string{
				"Harika! Selam %s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Rastgele bir sayı söyle",
				"Rastgele bir sayı oluştur",
			},
			Responses: []string{
				"İşte sayın %s",
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
				"Bana saat 8'de kahvaltı yapmamı hatırlat",
				"Salı günü annemi aramamı hatırlat",
				"Bir sınavım olduğunu not al",
				"Yarın saat 9'da bir konferans görüşmem olduğunu hatırlat",
			},
			Responses: []string{
				"Kaydettim bile! Size hatırlatacağım: “%s” için %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"Ne hatırlamanı istedim",
				"Bana hatırlatıcılarımı ver",
			},
			Responses: []string{
				"Benden şu şeyleri hatırlamamı istedin:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"İşte benim spotify tokenlarım",
				"Spotify bilgilerim",
			},
			Responses: []string{
				"Giriş devam ediyor",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Spotify'dan birşeyler çal",
			},
			Responses: []string{
				"%s sanatçısının %s eseri Spotify ayrıcalığı ile sizlerle.(Bu sahnede alkış tutabilrsin)",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Bana komik bir şeyler söyle",
				"Beni güldür",
			},
			Responses: []string{
				"Hadi bakalım, %s",
				"Işte bir tane, %s",
			},
			Replacer: modules.JokesReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["tr"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "Türkiye" {
		return "bir " + name
	}

	return name
}
