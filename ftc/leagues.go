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
	Region           string  `json:"region,omitempty"`
	Code             string  `json:"code,omitempty"`
	Name             string  `json:"name,omitempty"`
	Remote           bool    `json:"remote,omitempty"`
	ParentLeagueCode *string `json:"parentLeagueCode,omitempty"`
	ParentLeagueName *string `json:"parentLeagueName,omitempty"`
	Location         string  `json:"location,omitempty"`
}

// LeagueRankings is the list of rangings for a given league.
type LeagueRankings struct {
	Rankings []LeagueRankings `json:"rankings"`
}

// LeagueRanking is the ranking for a given league.
type LeagueRanking struct {
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
