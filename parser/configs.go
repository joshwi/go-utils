package parser

var PFR_TEAM_SEASON = []Config{
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<div data-template=.Partials\\/Teams\\/Summary.>.*?<h1 itemprop=.name.>.*?<span class=.header_end.>"},
			{Name: "", Value: "(?ms)<span.*?>.*?<\\/span>.*?<span>(?P<team>(.*?))<\\/span>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>Coach:<\\/strong>.*?<\\/p>"},
			{Name: "", Value: "(?ms)<a.*?>(?P<hc>(.*?))<\\/a>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>Offensive Coordinator:<\\/strong>.*?<\\/p>"},
			{Name: "", Value: "(?ms)<a.*?>(?P<oc>(.*?))<\\/a>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>Defensive Coordinator:<\\/strong>.*?<\\/p>"},
			{Name: "", Value: "(?ms)<a.*?>(?P<dc>(.*?))<\\/a>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>(Chairman.*?|Owner.*?|Principal Owner.*?)<\\/strong>.*?<a.*?>(?P<owner>(.*?))<\\/a>.*?<\\/p>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>(Stadium.*?)<\\/strong>.*?<a.*?>(?P<stadium>(.*?))<\\/a>.*?<\\/p>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>(Executive VP.*?|General Manager.*?)<\\/strong>.*?<a.*?>(?P<gm>(.*?))<\\/a>.*?<\\/p>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>Offensive Scheme:<\\/strong>.*?(?P<offense>(.*?)).*?<\\/p>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<p><strong>Defensive Alignment:<\\/strong>.*?(?P<defense>(.*?)).*?<\\/p>"},
		},
	},
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<caption>Schedule.*?Game.*?<\\/caption>.*?<\\/table>"},
			{Name: "", Value: "<tr.*?><th.*?>(?P<week>(.*?))<\\/th><td.*?>(?P<day>(.*?))<\\/td><td.*?>(?P<date>(.*?))<\\/td><td.*?>(?P<time>(.*?))<\\/td><td.*?><a href=.(?P<url>(.*?)).>boxscore<\\/a><\\/td><td.*?>(?P<wlt>(.*?))<\\/td><td.*?>(?P<overtime>(.*?))<\\/td><td.*?>(?P<record>(.*?))<\\/td><td.*?>(?P<homeAway>(.*?))<\\/td><td.*?><a.*?>(?P<opponent>(.*?))<\\/a><\\/td><td.*?>(?P<score>(.*?))<\\/td><td.*?>(?P<opponentScore>(.*?))<\\/td><td.*?>(?P<firstDown>(.*?))<\\/td><td.*?>(?P<totalYds>(.*?))<\\/td><td.*?>(?P<passYds>(.*?))<\\/td><td.*?>(?P<rushYds>(.*?))<\\/td><td.*?>(?P<turnovers>(.*?))<\\/td><td.*?>(?P<opponentFirstDowns>(.*?))<\\/td><td.*?>(?P<opponentTotYds>(.*?))<\\/td><td.*?>(?P<opponentPassYds>(.*?))<\\/td><td.*?>(?P<opponentRushYds>(.*?))<\\/td><td.*?>(?P<opponentTurnovers>(.*?))<\\/td><td.*?>(?P<expOffPts>(.*?))<\\/td><td.*?>(?P<expDefPts>(.*?))<\\/td><td.*?>(?P<expSpTeamsPts>(.*?))<\\/td>.*?<\\/tr>"},
		},
	},
}

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
