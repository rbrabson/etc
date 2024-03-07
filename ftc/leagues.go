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

type LeagueMembers struct {
	Members []int `json:"members"`
}

// LeagueRankings is the list of rangings for a given league.
type LeagueRankings struct {
	Rankings []LeagueRanking `json:"rankings"`
}

// LeagueRanking is the ranking for a given league.
type LeagueRanking struct {
	Rank              int     `json:"rank"`
	TeamNumber        int     `json:"teamNumber"`
	DisplayTeamNumber string  `json:"displayTeamNumber"`
	TeamName          *string `json:"teamName,omitempty"`
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
	DQ                int     `json:"dq"`
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

// GetLeagueMembers returns the list of members in the league
func GetLeagueMembers(season string, regionCode string, leagueCode string) ([]int, error) {
	url := fmt.Sprintf("%s/%s/leagues/members/%s/%s", server, season, regionCode, leagueCode)
	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output LeagueMembers
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output.Members, nil
}

// GetLeagueRankings returns the team rankings in a given league
func GetLeagueRankings(season string, regionCode string, leagueCode string) ([]LeagueRanking, error) {
	url := fmt.Sprintf("%s/%s/leagues/rankings/%s/%s", server, season, regionCode, leagueCode)
	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output LeagueRankings
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output.Rankings, nil
}

// String returns a string representation of League. In this case, it is a json string.
func (l League) String() string {
	body, _ := json.Marshal(l)
	return string(body)
}

// String returns a string representation of LeagueRanking. In this case, it is a json string.
func (l LeagueRanking) String() string {
	body, _ := json.Marshal(l)
	return string(body)
}
