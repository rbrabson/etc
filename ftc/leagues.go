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
