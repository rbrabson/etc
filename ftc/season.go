package ftc

import (
	"fmt"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

type Teams struct {
	Teams          []Team `json:"teams"`
	TeamCountTotal int    `json:"teamCountTotal"`
	TeamCountPage  int    `json:"teamCountPage"`
	PageCurrent    int    `json:"pageCurrent"`
	PageTotal      int    `json:"pageTotal"`
}
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

type Events struct {
	Event      []Events `json:"events"`
	EventCount int      `json:"eventCount"`
}
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

func GetTeams(season string) ([]string, error) {
	url := fmt.Sprintf("%s/%s/teams?teamNumber=7083", server, season)
	body, err := ftchttp.Get(url)

	if err != nil {
		return nil, err
	}
	fmt.Println("Body", string(body))

	return nil, nil
}
