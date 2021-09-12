var WIKI_MOVIE = []Config{
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Produced by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<producer>(.*?))<\\/a>"},
		},
	}, {
		Label: "director",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Directed by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<director>(.*?))<\\/a>"},
		},
	}, {
		Label: "screenplay",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Screenplay by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<writer>(.*?))<\\/a>"},
		},
	}, {
		Label: "cast",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Starring<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<actor>(.*?))<\\/a>"},
		},
	}, {
		Label: "score",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Music by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<artist>(.*?))<\\/a>"},
		},
	}, {
		Label: "releaseDate",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>.*?Release date.*?<\\/th>.*?<\\/tr>"},
			{Name: "", Value: "<span[^>]+>(?P<releaseDate>(\\d{4}-\\d{2}-\\d{2}))<\\/span>"},
		},
	},
	{
		Label: "runtime",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>.*?Running time.*?<\\/th>.*?<\\/tr>"},
			{Name: "", Value: "<td[^>]+>(?P<length>(\\d+)).minutes.*?<\\/td>"},
		},
	},
}