package parser

var CONFIG_LIST = []Config{
	{
		Name: "pfr_team_season",
		Urls: []string{"https://www.pro-football-reference.com/teams/{tag}/{year}.htm"},
		Keys: []string{"tag", "year"},
		Parser: []Parser{
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<div data-template=.Partials\\/Teams\\/Summary.>.*?<h1 itemprop=.name.>.*?<span class=.header_end.>"},
					{Name: "(?ms)<span.*?>.*?<\\/span>.*?<span>(?P<team>(.*?))<\\/span>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>Coach:<\\/strong>.*?<\\/p>"},
					{Name: "(?ms)<a.*?>(?P<hc>(.*?))<\\/a>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>Offensive Coordinator:<\\/strong>.*?<\\/p>"},
					{Name: "(?ms)<a.*?>(?P<oc>(.*?))<\\/a>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>Defensive Coordinator:<\\/strong>.*?<\\/p>"},
					{Name: "(?ms)<a.*?>(?P<dc>(.*?))<\\/a>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>(Chairman.*?|Owner.*?|Principal Owner.*?)<\\/strong>.*?<a.*?>(?P<owner>(.*?))<\\/a>.*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>(Stadium.*?)<\\/strong>.*?<a.*?>(?P<stadium>(.*?))<\\/a>.*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>(Executive VP.*?|General Manager.*?)<\\/strong>.*?<a.*?>(?P<gm>(.*?))<\\/a>.*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>Offensive Scheme:<\\/strong>.*?(?P<offense>(.*?)).*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []Regex{
					{Name: "(?ms)<p><strong>Defensive Alignment:<\\/strong>.*?(?P<defense>(.*?)).*?<\\/p>"},
				},
			},
			{
				Label: "games",
				Regex: []Regex{
					{Name: "(?ms)<caption>Schedule.*?Game.*?<\\/caption>.*?<\\/table>"},
					{Name: "<tr.*?><th.*?>(?P<week>(.*?))<\\/th><td.*?>(?P<day>(.*?))<\\/td><td.*?>(?P<date>(.*?))<\\/td><td.*?>(?P<time>(.*?))<\\/td><td.*?><a href=.(?P<url>(.*?)).>boxscore<\\/a><\\/td><td.*?>(?P<wlt>(.*?))<\\/td><td.*?>(?P<overtime>(.*?))<\\/td><td.*?>(?P<record>(.*?))<\\/td><td.*?>(?P<homeAway>(.*?))<\\/td><td.*?><a.*?>(?P<opponent>(.*?))<\\/a><\\/td><td.*?>(?P<score>(.*?))<\\/td><td.*?>(?P<opponentScore>(.*?))<\\/td><td.*?>(?P<firstDown>(.*?))<\\/td><td.*?>(?P<totalYds>(.*?))<\\/td><td.*?>(?P<passYds>(.*?))<\\/td><td.*?>(?P<rushYds>(.*?))<\\/td><td.*?>(?P<turnovers>(.*?))<\\/td><td.*?>(?P<opponentFirstDowns>(.*?))<\\/td><td.*?>(?P<opponentTotYds>(.*?))<\\/td><td.*?>(?P<opponentPassYds>(.*?))<\\/td><td.*?>(?P<opponentRushYds>(.*?))<\\/td><td.*?>(?P<opponentTurnovers>(.*?))<\\/td><td.*?>(?P<expOffPts>(.*?))<\\/td><td.*?>(?P<expDefPts>(.*?))<\\/td><td.*?>(?P<expSpTeamsPts>(.*?))<\\/td>.*?<\\/tr>"},
				},
			},
		},
	},
	{
		Name: "pfr_season_draft",
		Urls: []string{"https://www.pro-football-reference.com/years/{year}/draft.htm"},
		Keys: []string{"year"},
		Parser: []Parser{
			{
				Label: "picks",
				Regex: []Regex{
					{Name: "(?ms)<caption>Drafted.*?Players.*?<\\/caption>.*?<\\/table>"},
					{Name: "<tr ><th.*?>(?P<round>(.*?))<\\/th><td.*?>(?P<pick>(.*?))<\\/td><td.*?><a href=.\\/teams\\/(?P<tag>(...))\\/...._draft.htm. title=.(?P<team>(.*?)).>.*?<\\/a><\\/td><td.*?><strong><a.*?>(?P<player>(.*?))<\\/a><\\/strong><\\/td><td.*?>(?P<position>(.*?))<\\/td><td.*?>(?P<age>(.*?))<\\/td><td.*?>(?P<yearLast>(.*?))<\\/td><td.*?>(?P<yearFirstTeamAllPro>(.*?))<\\/td><td.*?>(?P<yearProBowl>(.*?))<\\/td><td.*?>(?P<yearPrimaryStarter>(.*?))<\\/td><td.*?>(?P<avCareer>(.*?))<\\/td><td.*?>(?P<avTeam>(.*?))<\\/td><td.*?>(?P<gamesPlayed>(.*?))<\\/td><td.*?>(?P<passCompletions>(.*?))<\\/td><td.*?>(?P<passAttempts>(.*?))<\\/td><td.*?>(?P<passYards>(.*?))<\\/td><td.*?>(?P<passTds>(.*?))<\\/td><td.*?>(?P<passInts>(.*?))<\\/td><td.*?>(?P<rushAttempts>(.*?))<\\/td><td.*?>(?P<rushYards>(.*?))<\\/td><td.*?>(?P<rushTds>(.*?))<\\/td><td.*?>(?P<receptions>(.*?))<\\/td><td.*?>(?P<recYards>(.*?))<\\/td><td.*?>(?P<recTds>(.*?))<\\/td><td.*?>(?P<soloTackles>(.*?))<\\/td><td.*?>(?P<interceptions>(.*?))<\\/td><td.*?>(?P<sacks>(.*?))<\\/td><td.*?><a href=.(?P<collegeUrl>(.*?)).>(?P<college>(.*?))<\\/a><\\/td><td.*?><a href=.(?P<collegeStatsUrl>(.*?)).>College Stats<\\/a><\\/td>.*?<\\/tr>"},
				},
			},
		},
	},
	{
		Name: "wiki_movie_info",
		Urls: []string{},
		Keys: []string{"title", "year"},
		Parser: []Parser{
			{
				Label: "producer",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>Produced by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<producer>(.*?))<\\/a>"},
				},
			}, {
				Label: "director",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>Directed by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<director>(.*?))<\\/a>"},
				},
			}, {
				Label: "screenplay",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>Screenplay by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<writer>(.*?))<\\/a>"},
				},
			}, {
				Label: "cast",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>Starring<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<actor>(.*?))<\\/a>"},
				},
			}, {
				Label: "score",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>Music by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<artist>(.*?))<\\/a>"},
				},
			}, {
				Label: "releaseDate",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>.*?Release date.*?<\\/th>.*?<\\/tr>"},
					{Name: "<span[^>]+>(?P<releaseDate>(\\d{4}-\\d{2}-\\d{2}))<\\/span>"},
				},
			},
			{
				Label: "runtime",
				Regex: []Regex{
					{Name: "(?ms)<tr><th[^>]+>.*?Running time.*?<\\/th>.*?<\\/tr>"},
					{Name: "<td[^>]+>(?P<length>(\\d+)).minutes.*?<\\/td>"},
				},
			},
		},
	},
}
