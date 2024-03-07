package ftc

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Leagues struct {
	Leagues     []League `json:"leagues"`
	LeagueCount int      `json:"leagueCount"`
}
type League struct {
	Region           string `json:"region"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	Remote           bool   `json:"remote"`
	ParentLeagueCode any    `json:"parentLeagueCode"`
	ParentLeagueName any    `json:"parentLeagueName"`
	Location         string `json:"location"`
}

type LeagueRankings struct {
	Rankings []LeagueRankings `json:"rankings"`
}
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

func GetLeagues(season string) ([]League, error) {
	url := fmt.Sprintf("%s/%s/leagues", server, season)

	// Setup the HTTP client for the request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Set up HTTPS request with basic authorization.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, authKey)

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status Code: %v", resp.StatusCode)
	}

	// Read the output from the server
	body, err := io.ReadAll(resp.Body)
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
