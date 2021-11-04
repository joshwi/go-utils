package parser

import (
	"github.com/joshwi/go-utils/utils"
)

var CONFIG_LIST = []utils.Config{
	{
		Name:   "wiki_film",
		Urls:   []string{"https://en.wikipedia.org/wiki/Lists_of_American_films"},
		Params: []string{},
		Parser: []utils.Parser{
			{
				Label: "year",
				Regex: []utils.Regex{
					{Name: "(?ms)<li><a href=\"(?P<url>([\\/[\\w+]+))\" title=\"List of American films of (?P<year>(\\d{4}))\">"},
				},
			},
		},
	},
	{
		Name:   "wiki_film_year",
		Urls:   []string{"https://en.wikipedia.org/wiki/List_of_American_films_of_{year}"},
		Params: []string{"year"},
		Parser: []utils.Parser{
			{
				Label: "movie",
				Regex: []utils.Regex{
					{Name: "(?ms)<tbody>.*<\\/tbody>"},
					{Name: "(?ms)<td><i><a href=\"(?P<url>((\\/[^\\s\\n]+)+))\" title=\".*?\">(?P<title>(.*?))<\\/a>\\s{0,}<\\/i>\\s{0,}<\\/td>"},
				},
			},
		},
	},
	{
		Name:   "wiki_movie",
		Urls:   []string{"https://en.wikipedia.org{url}"},
		Params: []string{"url"},
		Parser: []utils.Parser{
			{
				Label: "producer",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>Produced by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<producer>(.*?))<\\/a>"},
				},
			}, {
				Label: "director",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>Directed by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<director>(.*?))<\\/a>"},
				},
			}, {
				Label: "screenplay",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>Screenplay by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<writer>(.*?))<\\/a>"},
				},
			}, {
				Label: "cast",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>Starring<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<actor>(.*?))<\\/a>"},
				},
			}, {
				Label: "score",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>Music by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
					{Name: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<artist>(.*?))<\\/a>"},
				},
			}, {
				Label: "releaseDate",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>.*?Release date.*?<\\/th>.*?<\\/tr>"},
					{Name: "<span[^>]+>(?P<releaseDate>(\\d{4}-\\d{2}-\\d{2}))<\\/span>"},
				},
			},
			{
				Label: "runtime",
				Regex: []utils.Regex{
					{Name: "(?ms)<tr><th[^>]+>.*?Running time.*?<\\/th>.*?<\\/tr>"},
					{Name: "<td[^>]+>(?P<length>(\\d+)).minutes.*?<\\/td>"},
				},
			},
		},
	},
	{
		Name:   "pfr_map_team",
		Urls:   []string{"https://www.pro-football-reference.com/teams"},
		Params: []string{},
		Parser: []utils.Parser{
			{
				Label: "teams",
				Regex: []utils.Regex{
					{Name: "(?ms)<div class=.table_container. id=.div_teams_active.>.*?<\\/table>"},
					{Name: "<th[^>]+><a href=.\\/teams\\/(?P<tag>(.*?))\\/.>(?P<team>(.*?))<\\/a><\\/th><td[^>]+>(?P<year_min>(.*?))<\\/td><td[^>]+>(?P<year_max>(.*?))<\\/td><td[^>]+>(?P<win>(.*?))<\\/td><td[^>]+>(?P<loss>(.*?))<\\/td><td[^>]+>(?P<tie>(.*?))<\\/td><td[^>]+>(?P<win_perc>(.*?))<\\/td><td[^>]+><a href=.(?P<top_player_url>(.*?)). title[^>]+>(?P<top_player>(.*?))<\\/a><\\/td><td[^>]+><a href=.(?P<top_pass_url>(.*?)). title[^>]+>(?P<top_pass>(.*?))<\\/a><\\/td><td[^>]+><a href=.(?P<top_rush_url>(.*?)). title[^>]+>(?P<top_rush>(.*?))<\\/a><\\/td><td[^>]+><a href=.(?P<top_rec_url>(.*?)). title[^>]+>(?P<top_rec>(.*?))<\\/a><\\/td><td[^>]+><a href=.(?P<top_coach_url>(.*?)). title[^>]+>(?P<top_coach>(.*?))<\\/a><\\/td><td[^>]+>(?P<playoff_yrs>(.*?))<\\/td><td[^>]+>(?P<playoff_win>(.*?))<\\/td><td[^>]+>(?P<playoff_loss>(.*?))<\\/td><td[^>]+>(?P<playoff_perc>(.*?))<\\/td><td[^>]+>(?P<champs>(.*?))<\\/td><td[^>]+>(?P<sb_champs>(.*?))<\\/td><td[^>]+>(?P<conf_champs>(.*?))<\\/td><td[^>]+>(?P<div_champs>(.*?))<\\/td>"},
				},
			},
		},
	},
	{
		Name:   "pfr_map_season",
		Urls:   []string{"https://www.pro-football-reference.com/years"},
		Params: []string{},
		Parser: []utils.Parser{
			{
				Label: "years",
				Regex: []utils.Regex{
					{Name: "(?ms)<div class=\"table_container\" id=\"div_years\">.*?<\\/table>"},
					{Name: "<th.*?><a href=\"(?P<url>(\\/years\\/\\d{4}\\/))\">(?P<year>(\\d{4}))<\\/a><\\/th>"},
				},
			},
		},
	},
	{
		Name:   "pfr_map_player",
		Urls:   []string{"https://www.pro-football-reference.com/players/{letter}"},
		Params: []string{"letter"},
		Parser: []utils.Parser{
			{
				Label: "players",
				Regex: []utils.Regex{
					{Name: "(?ms)<div class=.section_content. id=.div_players.>.*?<\\/div>"},
					{Name: "<p>.*?<a href=.(?P<url>(.*?)).>(?P<name>(.*?))<\\/a>(.*?)\\((?P<pos>(.*?))\\)(.*?)(?P<rookie_year>([0-9]{4}))-(?P<last_year>([0-9]{4})).*?<\\/p>"},
				},
			},
		},
	},
	{
		Name:   "pfr_team_season",
		Urls:   []string{"https://www.pro-football-reference.com/teams/{tag}/{year}.htm"},
		Params: []string{"tag", "year"},
		Parser: []utils.Parser{
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<div data-template=.Partials\\/Teams\\/Summary.>.*?<h1 itemprop=.name.>.*?<span class=.header_end.>"},
					{Name: "(?ms)<span.*?>.*?<\\/span>.*?<span>(?P<team>(.*?))<\\/span>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>Coach:<\\/strong>.*?<\\/p>"},
					{Name: "(?ms)<a.*?>(?P<hc>(.*?))<\\/a>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>Offensive Coordinator:<\\/strong>.*?<\\/p>"},
					{Name: "(?ms)<a.*?>(?P<oc>(.*?))<\\/a>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>Defensive Coordinator:<\\/strong>.*?<\\/p>"},
					{Name: "(?ms)<a.*?>(?P<dc>(.*?))<\\/a>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>(Chairman.*?|Owner.*?|Principal Owner.*?)<\\/strong>.*?<a.*?>(?P<owner>(.*?))<\\/a>.*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>(Stadium.*?)<\\/strong>.*?<a.*?>(?P<stadium>(.*?))<\\/a>.*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>(Executive VP.*?|General Manager.*?)<\\/strong>.*?<a.*?>(?P<gm>(.*?))<\\/a>.*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>Offensive Scheme:<\\/strong>.*?(?P<off>(.*?)).*?<\\/p>"},
				},
			},
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<p><strong>Defensive Alignment:<\\/strong>.*?(?P<def>(.*?)).*?<\\/p>"},
				},
			},
			// {
			// 	Label: "games",
			// 	Regex: []utils.Regex{
			// 		{Name: "(?ms)<caption>Schedule.*?Game.*?<\\/caption>.*?<\\/table>"},
			// 		{Name: "<tr.*?><th.*?>(?P<week>(.*?))<\\/th><td.*?>(?P<day>(.*?))<\\/td><td.*?>(?P<date>(.*?))<\\/td><td.*?>(?P<time>(.*?))<\\/td><td.*?><a href=.(?P<url>(.*?)).>boxscore<\\/a><\\/td><td.*?>(?P<wlt>(.*?))<\\/td><td.*?>(?P<overtime>(.*?))<\\/td><td.*?>(?P<record>(.*?))<\\/td><td.*?>(?P<homeAway>(.*?))<\\/td><td.*?><a.*?>(?P<opponent>(.*?))<\\/a><\\/td><td.*?>(?P<score>(.*?))<\\/td><td.*?>(?P<opponentScore>(.*?))<\\/td><td.*?>(?P<firstDown>(.*?))<\\/td><td.*?>(?P<totalYds>(.*?))<\\/td><td.*?>(?P<passYds>(.*?))<\\/td><td.*?>(?P<rushYds>(.*?))<\\/td><td.*?>(?P<turnovers>(.*?))<\\/td><td.*?>(?P<opponentFirstDowns>(.*?))<\\/td><td.*?>(?P<opponentTotYds>(.*?))<\\/td><td.*?>(?P<opponentPassYds>(.*?))<\\/td><td.*?>(?P<opponentRushYds>(.*?))<\\/td><td.*?>(?P<opponentTurnovers>(.*?))<\\/td><td.*?>(?P<expOffPts>(.*?))<\\/td><td.*?>(?P<expDefPts>(.*?))<\\/td><td.*?>(?P<expSpTeamsPts>(.*?))<\\/td>.*?<\\/tr>"},
			// 	},
			// },
			{
				Label: "games",
				Regex: []utils.Regex{
					{Name: "(?ms)<caption>Schedule.*?Game.*?<\\/caption>.*?<\\/table>"},
					{Name: "<tr.*?><th.*?>(?P<week>(.*?))<\\/th><td.*?>(?P<day>(.*?))<\\/td><td.*?>(?P<date>(.*?))<\\/td><td.*?>(?P<time>(.*?))<\\/td><td.*?><a href=.(?P<url>(.*?)).>boxscore<\\/a><\\/td><td.*?>(?P<wlt>(.*?))<\\/td><td.*?>(?P<ot>(.*?))<\\/td><td.*?>(?P<record>(.*?))<\\/td><td.*?>(?P<home_away>(.*?))<\\/td><td.*?><a.*?>(?P<opp>(.*?))<\\/a><\\/td><td.*?>(?P<score>(.*?))<\\/td><td.*?>(?P<opp_score>(.*?))<\\/td><td.*?>(?P<first_downs>(.*?))<\\/td><td.*?>(?P<tot_yds>(.*?))<\\/td><td.*?>(?P<pass_yds>(.*?))<\\/td><td.*?>(?P<rush_yds>(.*?))<\\/td><td.*?>(?P<to>(.*?))<\\/td><td.*?>(?P<opp_first_downs>(.*?))<\\/td><td.*?>(?P<opp_tot_yds>(.*?))<\\/td><td.*?>(?P<opp_pass_yds>(.*?))<\\/td><td.*?>(?P<opp_rush_yds>(.*?))<\\/td><td.*?>(?P<opp_to>(.*?))<\\/td><td.*?>(?P<exp_off_pts>(.*?))<\\/td><td.*?>(?P<exp_def_pts>(.*?))<\\/td><td.*?>(?P<exp_spt_pts>(.*?))<\\/td>.*?<\\/tr>"},
				},
			},
		},
	},
	{
		Name:   "pfr_season_draft",
		Urls:   []string{"https://www.pro-football-reference.com/years/{year}/draft.htm"},
		Params: []string{"year"},
		Parser: []utils.Parser{
			{
				Label: "picks",
				Regex: []utils.Regex{
					{Name: "(?ms)<caption>Drafted.*?Players.*?<\\/caption>.*?<\\/table>"},
					{Name: "<tr ><th.*?data-stat=\"draft_round\".*?>(?P<round>(.*?))<\\/th><td.*?data-stat=\"draft_pick\".*?>(?P<pick>(.*?))<\\/td><td.*?data-stat=\"team\".*?><a href=\"\\/teams\\/(?P<tag>(...))\\/...._draft.htm\" title=\"(?P<team>(.*?))\">.*?<\\/a><\\/td><td.*?data-stat=\"player\".*?>(<strong>|)<a href=\"(?P<url_player>(\\/players\\/\\w\\/\\w+.htm))\">(?P<player>(.*?))<\\/a>(<\\/strong>|)<\\/td><td.*?data-stat=\"pos\".*?>(?P<pos>(.*?))<\\/td><td.*?data-stat=\"age\".*?>(?P<age>(.*?))<\\/td><td.*?data-stat=\"year_max\".*?>(?P<year_max>(.*?))<\\/td><td.*?data-stat=\"all_pros_first_team\".*?>(?P<first_tap>(.*?))<\\/td><td.*?data-stat=\"pro_bowls\".*?>(?P<pro_bowl>(.*?))<\\/td><td.*?data-stat=\"years_as_primary_starter\".*?>(?P<starter>(.*?))<\\/td><td.*?data-stat=\"career_av\".*?>(?P<av_career>(.*?))<\\/td><td.*?data-stat=\"draft_av\".*?>(?P<av_team>(.*?))<\\/td><td.*?data-stat=\"g\".*?>(?P<games>(.*?))<\\/td><td.*?data-stat=\"pass_cmp\".*?>(?P<pass_cmp>(.*?))<\\/td><td.*?data-stat=\"pass_att\".*?>(?P<pass_att>(.*?))<\\/td><td.*?data-stat=\"pass_yds\".*?>(?P<pass_yd>(.*?))<\\/td><td.*?data-stat=\"pass_td\".*?>(?P<pass_td>(.*?))<\\/td><td.*?data-stat=\"pass_int\".*?>(?P<pass_int>(.*?))<\\/td><td.*?data-stat=\"rush_att\".*?>(?P<rush_att>(.*?))<\\/td><td.*?data-stat=\"rush_yds\".*?>(?P<rush_yd>(.*?))<\\/td><td.*?data-stat=\"rush_td\".*?>(?P<rush_td>(.*?))<\\/td><td.*?data-stat=\"rec\".*?>(?P<rec>(.*?))<\\/td><td.*?data-stat=\"rec_yds\".*?>(?P<rec_yd>(.*?))<\\/td><td.*?data-stat=\"rec_td\".*?>(?P<rec_td>(.*?))<\\/td><td.*?data-stat=\"tackles_solo\".*?>(?P<solo>(.*?))<\\/td><td.*?data-stat=\"def_int\".*?>(?P<int>(.*?))<\\/td><td.*?data-stat=\"sacks\".*?>(?P<sack>(.*?))<\\/td><td.*?data-stat=\"college_id\".*?><a href=\"(?P<url_college>(.*?))\">(?P<college>(.*?))<\\/a><\\/td><td.*?data-stat=\"college_link\".*?>(<a href=\"(?P<url_college_stats>(.*?))\">|)(College Stats|)(<\\/a>|)<\\/td>.*?<\\/tr>"},
				},
			},
		},
	},
	{
		Name:   "pfr_boxscore_games",
		Urls:   []string{"https://www.pro-football-reference.com/boxscores/{id}.htm"},
		Params: []string{"id"},
		Parser: []utils.Parser{
			{
				Label: "officials",
				Regex: []utils.Regex{
					{Name: "(?ms)<div id=.all_officials.*?<table[^>]+>.*?<\\/table>"},
					{Name: "<th[^>]+>(?P<type>(.*?))<\\/th><td[^>]+><a href=.(?P<url>(.*?)).>(?P<name>(.*?))<\\/a><\\/td>"},
				},
			},
			{
				Label: "home_snap_counts",
				Regex: []utils.Regex{
					{Name: "(?ms)<div id=.all_home_snap_counts..*?<\\/table>"},
					{Name: "<th.*?><a href=.(?P<player_url>(.*?)).>(?P<player>(.*?))<a\\/><\\/th><td.*?>(?P<position>(.*?))<\\/td><td.*?>(?P<off>(.*?))<\\/td><td.*?>(?P<off_perc>(.*?))<\\/td><td.*?>(?P<def>(.*?))<\\/td><td.*?>(?P<def_perc>(.*?))<\\/td><td.*?>(?P<st>(.*?))<\\/td><td.*?>(?P<st_perc>(.*?))<\\/td>"},
				},
			},
			{
				Label: "away_snap_counts",
				Regex: []utils.Regex{
					{Name: "(?ms)<div id=.all_vis_snap_counts..*?<\\/table>"},
					{Name: "<th.*?><a href=.(?P<player_url>(.*?)).>(?P<player>(.*?))<a\\/><\\/th><td.*?>(?P<position>(.*?))<\\/td><td.*?>(?P<off>(.*?))<\\/td><td.*?>(?P<off_perc>(.*?))<\\/td><td.*?>(?P<def>(.*?))<\\/td><td.*?>(?P<def_perc>(.*?))<\\/td><td.*?>(?P<st>(.*?))<\\/td><td.*?>(?P<st_perc>(.*?))<\\/td>"},
				},
			},
			{
				Label: "home_drives",
				Regex: []utils.Regex{
					{Name: "(?ms)<div id=.all_home_drives..*?<\\/table>"},
					{Name: "<th.*?>(?P<number>(.*?))<\\/th><td.*?>(?P<quarter>(.*?))<\\/td><td.*?>(?P<time>(.*?))<\\/td><td.*?>(?P<location>(.*?))<\\/td><td.*?><span class=.tooltip. tip=.(?P<pass>(\\d+)) Pass, (?P<rush>(\\d+)) Rush, (?P<penalty>(\\d+)) Penalty.>(?P<plays>(.*?))<\\/span><\\/td><td.*?>(?P<length>(.*?))<\\/td><td.*?>(?P<net_yds>(.*?))<\\/td><td.*?>(?P<result>(.*?))<\\/td>"},
				},
			},
			{
				Label: "away_drives",
				Regex: []utils.Regex{
					{Name: "(?ms)<div id=.all_vis_drives..*?<\\/table>"},
					{Name: "<th.*?>(?P<number>(.*?))<\\/th><td.*?>(?P<quarter>(.*?))<\\/td><td.*?>(?P<time>(.*?))<\\/td><td.*?>(?P<location>(.*?))<\\/td><td.*?><span class=.tooltip. tip=.(?P<pass>(\\d+)) Pass, (?P<rush>(\\d+)) Rush, (?P<penalty>(\\d+)) Penalty.>(?P<plays>(.*?))<\\/span><\\/td><td.*?>(?P<length>(.*?))<\\/td><td.*?>(?P<net_yds>(.*?))<\\/td><td.*?>(?P<result>(.*?))<\\/td>"},
				},
			},
			{
				Label: "plays",
				Regex: []utils.Regex{
					{Name: "(?ms)<caption>Full Play-By-Play.*?<\\/caption>.*?<\\/table>"},
					{Name: "<th.*?>(?P<quarter>(.*?))<\\/th><td.*?>(?P<time>(.*?))<\\/td><td.*?>(?P<down>(.*?))<\\/td><td.*?>(?P<distance>(.*?))<\\/td><td.*?>(?P<location>(.*?))<\\/td><td.*?>(?P<description>(.*?))<\\/td>"},
				},
			},
		},
	},
	{
		Name:   "genius_song_lyrics",
		Urls:   []string{},
		Params: []string{},
		Parser: []utils.Parser{
			{
				Label: "",
				Regex: []utils.Regex{
					{Name: "(?ms)<div class=.lyrics.>.*?<!--sse-->(?<unsynchronisedLyrics>(.*?))<!--\\/sse-->.*?<\\/div>"},
				},
			},
		},
	},
}
