package en

import (
	"github.com/olivia-ai/olivia/modules"
)

func init() {
	modules.RegisterModules("zh", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"的面积是多少",
				"告诉我的面积",
			},
			Responses: []string{
				"%s的面积为%g平方千米",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"的首都在哪",
				"哪是的首都",
			},
			Responses: []string{
				"%s的首都是%s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"使用哪种货币",
				"的货币是什么",
				"什么是的货币",
			},
			Responses: []string{
				"%s的货币是%s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"告诉我的结果",
				"等于几",
				"等于多少",
				"计算",
				"的结果是",
			},
			Responses: []string{
				"结果是%s",
				"答案是%s",
				"等于%s",
				"是%s",
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
				"我最喜欢的电影类型是喜剧和恐怖",
				"我喜欢恐怖电影",
				"我喜欢喜剧、恐怖类型的电影",
				"我喜欢关于战争的电影",
				"我喜欢以战争为题材的电影",
				"我喜欢以战争为主题的电影",
				"我喜欢动作片",
			},
			Responses: []string{
				"我记住了你的电影偏好",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"帮我找部电影",
				"推荐部电影",
				"我想看部电影",
			},
			Responses: []string{
				"我为您找到了电影“%s”，分级为%.02f5",
				"好巧不巧，我找到了电影“%s”，分级为%.02f5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"我看过这部电影了",
				"我不喜欢这部电影",
				"我对这部电影不感兴趣",
			},
			Responses: []string{
				"我又找到了另一个“%s”，分级为%.02f5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				//"I'm bored",
				//"I don't know what to do",
				//The original English sentence is above, but I think "boring" has nothing to do with recommended movies; Shouldn't boring talk to robots?
				//I don't think it is appropriate to use the word "boring" here; I'm not sure if it will take up other uses of the word "boring"
				"随便推荐部电影吧",
				"我想打发时间",
			},
			Responses: []string{
				"我建议你看%s电影“%s”，分级为%.02f5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"你知道我的名字吗？",
				"我是谁？",
				"我叫什么名字？",
				"你认识我吗？",
				"我的名字是什么？",
				"你知道我叫什么名字吗？",
			},
			Responses: []string{
				"我知道，你叫%s",
				"你叫%s，对吧！",
				"你肯定是%s",
				"我认识你，你叫%s",
				"我知道，你的名字是%s",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"我的名字是",
				"我是",
				"我的名字叫",
				"你可以叫我",
			},
			Responses: []string{
				"很高兴认识你，%s",
				"你好，%s",
				"很抱歉以这种方式认识你，%s",
				"我记住你了，%s",
				"认识你是我的荣幸，%s",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"给我一个随机数",
				"生成随机数",
				"我想要一个随机数",
				"来一个随机数",
			},
			Responses: []string{
				"好的，它是：%s",
				"它来了：%s",
				"这个数是：%s",
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
				"提醒我上午8点做早餐",
				"提醒我星期二给妈妈打电话",
				"记得告诉我，我有一个考试",
				"提醒我明天上午9点有电话会议",
			},
			Responses: []string{
				"好的，我会提醒您在%s去“%s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"我有日程安排吗？",
				"我最近需要做什么？",
			},
			Responses: []string{
				"我将提醒您:\n%s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		//I can't understand these. It seems that this is a foreign music platform?
		//I may connect with Netease Cloud or Cool Dog slowly; At present, I have ideas about these two platforms
		//If someone can provide help to connect with Himalayan FM or QQ music, then Guo Degang can speak crosstalk

		//To add another word, the above is not necessarily possible;
		//If the artificial neural network of this project does not meet my expectations,
		//I may not continue to work; I haven't tried it yet. It was in English at that time.
		//I questioned the performance of neural networks
		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Here are my spotify tokens",
				"My spotify secrets",
			},
			Responses: []string{
				"Login in progress",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Play from on Spotify",
			},
			Responses: []string{
				"Playing %s from %s on Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Tell me a joke",
				"Make me laugh",
			},
			Responses: []string{
				"Here you go, %s",
				"Here's one, %s",
			},
			Replacer: modules.JokesReplacer,
		},
		{
			Tag: modules.AdvicesTag,
			Patterns: []string{
				"Give me an advice",
				"Advise me",
			},
			Responses: []string{
				"Here you go, %s",
				"Here's one, %s",
				"Listen closely, %s",
			},
			Replacer: modules.AdvicesReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["zh"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "中国" {
		return "这是 " + name
	}

	return name
}
