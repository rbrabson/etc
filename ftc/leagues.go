package ftc

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

// Leagues is the data for the FTC leagues
type Leagues struct {
	Leagues     []League `json:"leagues"`
	LeagueCount int      `json:"leagueCount"`
}

// League is the data for a given FTC league
type League struct {
	Region           string `json:"region"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	Remote           bool   `json:"remote"`
	ParentLeagueCode any    `json:"parentLeagueCode"`
	ParentLeagueName any    `json:"parentLeagueName"`
	Location         string `json:"location"`
}

// LeagueRankings is the list of rangings for a given league.
type LeagueRankings struct {
	Rankings []LeagueRankings `json:"rankings"`
}

// LeagueRanking is the ranking for a given league.
type LeagueRanking struct {
	Rank              int     `json:"rank"`
	TeamNumber        int     `json:"teamNumber"`
	DisplayTeamNumber string  `json:"displayTeamNumber"`
	TeamName          any     `json:"teamName"`
	SortOrder1        float64 `json:"sortOrder1"`
	SortOrder2        float64 `json:"sortOrder2"`
	SortOrder3        float64 `json:"sortOrder3"`
	SortOrder4        float64 `json:"sortOrder4"`
	SortOrder5        float64 `json:"sortOrder5"`
	SortOrder6        float64 `json:"sortOrder6"`
	Wins              int     `json:"wins"`
	Losses            int     `json:"losses"`
	Ties              int     `json:"ties"`
	QualAverage       int     `json:"qualAverage"`
	Dq                int     `json:"dq"`
	MatchesPlayed     int     `json:"matchesPlayed"`
	MatchesCounted    int     `json:"matchesCounted"`
}

// GetLeagues returns the list of rankings for FTC leagues
func GetLeagues(season string) ([]League, error) {
	url := fmt.Sprintf("%s/%s/leagues", server, season)
	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output Leagues
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output.Leagues, nil
}

// String returns a string representation of Leagues. In this case, it is a json string.
func (general *League) String() string {
	body, _ := json.Marshal(general)
	return string(body)
}
