package ftc

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

// Teams returns a list of FTC teams. The information is returned in `pages`, so multiple requests
// may be required to get all FTC teams.
type Teams struct {
	Teams          []Team `json:"teams"`
	TeamCountTotal int    `json:"teamCountTotal"`
	TeamCountPage  int    `json:"teamCountPage"`
	PageCurrent    int    `json:"pageCurrent"`
	PageTotal      int    `json:"pageTotal"`
}

// Team is information for a given FTC team.
type Team struct {
	TeamNumber        int     `json:"teamNumber,omitempty"`
	DisplayTeamNumber string  `json:"displayTeamNumber,omitempty"`
	NameFull          string  `json:"nameFull,omitempty"`
	NameShort         string  `json:"nameShort,omitempty"`
	SchoolName        *string `json:"schoolName,omitempty"`
	City              string  `json:"city,omitempty"`
	StateProv         string  `json:"stateProv,omitempty"`
	Country           string  `json:"country,omitempty"`
	Website           *string `json:"website,omitempty"`
	RookieYear        int     `json:"rookieYear,omitempty"`
	RobotName         *string `json:"robotName,omitempty"`
	DistrictCode      *string `json:"districtCode,omitempty"`
	HomeCMP           *string `json:"homeCMP,omitempty"`
	HomeRegion        *string `json:"homeRegion,omitempty"`
}

// GetTeams returns a `page` of FTC teams.
func GetTeams(season string) (*Teams, error) {
	url := fmt.Sprintf("%s/%s/teams?teamNumber=7083", server, season)
	body, err := ftchttp.Get(url)

	if err != nil {
		return nil, err
	}

	var output Teams
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// String returns a string representation of Teams. In this case, it is a json string.
func (teams *Teams) String() string {
	body, _ := json.Marshal(teams)
	return string(body)
}
