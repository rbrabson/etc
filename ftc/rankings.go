package ftc

// Rankings is the list of rangings for a given league.
type Rankings struct {
	Rankings []Ranking `json:"rankings"`
}

// LeagueRanking is the ranking for a given league.
type Ranking struct {
	Rank              int     `json:"rank,omitempty"`
	TeamNumber        int     `json:"teamNumber,omitempty"`
	DisplayTeamNumber string  `json:"displayTeamNumber,omitempty"`
	TeamName          *string `json:"teamName,omitempty"`
	SortOrder1        float64 `json:"sortOrder1,omitempty"`
	SortOrder2        float64 `json:"sortOrder2,omitempty"`
	SortOrder3        float64 `json:"sortOrder3,omitempty"`
	SortOrder4        float64 `json:"sortOrder4,omitempty"`
	SortOrder5        float64 `json:"sortOrder5,omitempty"`
	SortOrder6        float64 `json:"sortOrder6,omitempty"`
	Wins              int     `json:"wins,omitempty"`
	Losses            int     `json:"losses,omitempty"`
	Ties              int     `json:"ties,omitempty"`
	QualAverage       int     `json:"qualAverage,omitempty"`
	Dq                int     `json:"dq,omitempty"`
	MatchesPlayed     int     `json:"matchesPlayed,omitempty"`
	MatchesCounted    int     `json:"matchesCounted,omitempty"`
}
