package vi

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("vi", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Khu vực của bạn ở ",
				"Gửi tôi khu vực của bạn ",
			},
			Responses: []string{
				"Khu vực %s là %gkm²",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Thủ đô của bạn là gì ",
				"Tên thủ đô của bạn là ",
				"gửi tôi tên thủ đô của bạn ",
			},
			Responses: []string{
				"Thủ đô %s là của %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Loại tiền tệ nào đang được sử dụng ",
				"Gửi tôi loại tiền bạn đang dùng ",
				"Loại tiền của bạn là ",
			},
			Responses: []string{
				"Loại tiền %s là của %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Gửi tôi kết quả ",
				"Tính toán",
			},
			Responses: []string{
				"Kết quả là %s",
				"Tính là %s",
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
				"Thể loại phim bạn hay xem là kinh dị và hài hước",
				"Tôi thích loại phim kinh dị và hài hước",
				"Tôi thích phim chiến tranh",
				"Tôi thích phim hành động",
			},
			Responses: []string{
				"Tuyệt! Tôi đã lưu thông tin các loại phim yêu thích của bạn.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Tìm giúp tôi phim về",
				"Gửi tôi bộ phim về",
			},
			Responses: []string{
				"Tôi tìm thấy phim “%s” cho bạn, với điểm đánh giá %.02f/5",
				"Chắc chắn rồi, Tôi tìm thấy phim “%s”, điểm đánh giá là %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Tôi đã xem bộ phim này",
				"Tôi vừa mới xem bộ phim này",
				"Ồ, tôi đã xem bộ phim này rồi",
			},
			Responses: []string{
				"Ồ tôi biết, tôi tim một bộ phim khác là “%s” có điểm đánh giá %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Tôi chán",
				"Tôi không biết cái tôi làm là gì",
			},
			Responses: []string{
				"Tôi đề xuất cho bạn bộ phim %s “%s”, với điểm đánh giá %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Bạn có biết tên tôi không?",
			},
			Responses: []string{
				"Tên của bạn là %s!",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Tên của tôi là",
				"Bạn có thể gọi tôi là",
			},
			Responses: []string{
				"Xin chào, %s!",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Gửi tôi một số bất kỳ",
			},
			Responses: []string{
				"Số đó là %s",
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
				"Nhắc tôi nấu ăn sáng lúc 8:00 nhé",
				"Nhắc tôi gọi cho mẹ vào thứ ba",
				"Ghi chú lại giúp tôi có bài kiểm tra",
				"Nhắc tôi là tôi có hội thảo vào 9:00 sáng mai",
			},
			Responses: []string{
				"Đã xong! Tôi sẽ nhắn bạn: “%s” lúc %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"Bạn có nhớ nhắn tôi cái gì không",
				"Gửi cho tôi các lời nhắc nhở",
			},
			Responses: []string{
				"Bạn hỏi tôi những lời nhắn là:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Đây là token Spotify của tôi",
				"Khoá Spotify bí mật của tôi",
			},
			Responses: []string{
				"Truy cập theo quy trình",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Chạy trên Spotify",
			},
			Responses: []string{
				"Chạy %s từ %s trên Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Kể chuyện cười",
				"Làm tôi cười",
			},
			Responses: []string{
				"Đây bạn, %s",
				"Đây này, %s",
			},
			Replacer: modules.JokesReplacer,
		},
		{
			Tag: modules.AdvicesTag,
			Patterns: []string{
				"Cho tôi một lời khuyên",
				"Khuyên bảo tôi",
			},
			Responses: []string{
				"Của bạn đây, %s",
				"Hãy lắng nghe, %s",
			},
			Replacer: modules.AdvicesReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["vi"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "United States" {
		return "the " + name
	}

	return name
}
