package ftc

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

// Teams returns a list of FTC teams. The information is returned in `pages`, so multiple requests
// may be required to get all FTC teams.
type Teams struct {
	Teams          []Team `json:"teams,omitempty"`
	TeamCountTotal int    `json:"teamCountTotal,omitempty"`
	TeamCountPage  int    `json:"teamCountPage,omitempty"`
	PageCurrent    int    `json:"pageCurrent,omitempty"`
	PageTotal      int    `json:"pageTotal,omitempty"`
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

// Events is information about FTC events.
type Events struct {
	Event      []Events `json:"events,omitempty"`
	EventCount int      `json:"eventCount,omitempty"`
}

// Event is information about a given FTC event.
type Event struct {
	EventID       string  `json:"eventId,omitempty"`
	Code          string  `json:"code,omitempty"`
	DivisionCode  *string `json:"divisionCode,omitempty"`
	Name          string  `json:"name,omitempty"`
	Remote        bool    `json:"remote,omitempty"`
	Hybrid        bool    `json:"hybrid,omitempty"`
	FieldCount    int     `json:"fieldCount,omitempty"`
	Published     bool    `json:"published,omitempty"`
	Type          string  `json:"type,omitempty"`
	TypeName      string  `json:"typeName,omitempty"`
	RegionCode    string  `json:"regionCode,omitempty"`
	LeagueCode    *string `json:"leagueCode,omitempty"`
	DistrictCode  string  `json:"districtCode,omitempty"`
	Venue         string  `json:"venue,omitempty"`
	Address       string  `json:"address,omitempty"`
	City          string  `json:"city,omitempty"`
	Stateprov     string  `json:"stateprov,omitempty"`
	Country       string  `json:"country,omitempty"`
	Website       string  `json:"website,omitempty"`
	LiveStreamURL string  `json:"liveStreamUrl,omitempty"`
	Coordinates   *string `json:"coordinates,omitempty"`
	Webcasts      *string `json:"webcasts,omitempty"`
	Timezone      string  `json:"timezone,omitempty"`
	DateStart     string  `json:"dateStart,omitempty"`
	DateEnd       string  `json:"dateEnd,omitempty"`
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
