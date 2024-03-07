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
	TeamNumber        int     `json:"teamNumber"`
	DisplayTeamNumber string  `json:"displayTeamNumber"`
	NameFull          string  `json:"nameFull"`
	NameShort         string  `json:"nameShort"`
	SchoolName        *string `json:"schoolName"`
	City              string  `json:"city"`
	StateProv         string  `json:"stateProv"`
	Country           string  `json:"country"`
	Website           *string `json:"website"`
	RookieYear        int     `json:"rookieYear"`
	RobotName         *string `json:"robotName"`
	DistrictCode      *string `json:"districtCode"`
	HomeCMP           *string `json:"homeCMP"`
	HomeRegion        *string `json:"homeRegion"`
}

// Events is information about FTC events.
type Events struct {
	Event      []Events `json:"events"`
	EventCount int      `json:"eventCount"`
}

// Event is information about a given FTC event.
type Event struct {
	EventID       string  `json:"eventId"`
	Code          string  `json:"code"`
	DivisionCode  *string `json:"divisionCode"`
	Name          string  `json:"name"`
	Remote        bool    `json:"remote"`
	Hybrid        bool    `json:"hybrid"`
	FieldCount    int     `json:"fieldCount"`
	Published     bool    `json:"published"`
	Type          string  `json:"type"`
	TypeName      string  `json:"typeName"`
	RegionCode    string  `json:"regionCode"`
	LeagueCode    *string `json:"leagueCode"`
	DistrictCode  string  `json:"districtCode"`
	Venue         string  `json:"venue"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	Stateprov     string  `json:"stateprov"`
	Country       string  `json:"country"`
	Website       string  `json:"website"`
	LiveStreamURL string  `json:"liveStreamUrl"`
	Coordinates   *string `json:"coordinates"`
	Webcasts      *string `json:"webcasts"`
	Timezone      string  `json:"timezone"`
	DateStart     string  `json:"dateStart"`
	DateEnd       string  `json:"dateEnd"`
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
