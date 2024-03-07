package ftc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

type Awards struct {
	Awards []Award `json:"awards"`
}
type Award struct {
	AwardID     int    `json:"awardId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ForPerson   bool   `json:"forPerson"`
}

type TeamAwards struct {
	Awards []TeamAward `json:"awards"`
}
type TeamAward struct {
	AwardID      int    `json:"awardId"`
	EventCode    string `json:"eventCode"`
	Name         string `json:"name"`
	Series       int    `json:"series"`
	TeamNumber   int    `json:"teamNumber"`
	SchoolName   any    `json:"schoolName"`
	FullTeamName string `json:"fullTeamName"`
	Person       any    `json:"person"`
}

func GetAwardListing(season string) ([]Award, error) {
	url := fmt.Sprintf("{{baseUrl}}/v2.0/%s/awards/list", season)

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output []Award
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output, nil
}

// GetEventAwards gets the list of awards at a given event
func GetEventAwards(season string, eventCode string, teamNumber ...string) ([]TeamAward, error) {
	sb := strings.Builder{}
	sb.WriteString(server)
	sb.WriteString("/")
	sb.WriteString("/awards/")
	sb.WriteString(eventCode)
	if len(teamNumber) > 0 {
		sb.WriteString("?teamNumber")
		sb.WriteString(teamNumber[0])
	}
	url := sb.String()

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output []TeamAward
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output, nil
}

// GetTeamAwards gets the list of awards for a given team
func GetTeamAwards(season string, teamNumber string, eventCode ...string) ([]TeamAward, error) {
	sb := strings.Builder{}
	sb.WriteString(server)
	sb.WriteString("/")
	sb.WriteString("/awards/")
	sb.WriteString(teamNumber)
	if len(teamNumber) > 0 {
		sb.WriteString("?eventCode")
		sb.WriteString(eventCode[0])
	}
	url := sb.String()

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output []TeamAward
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output, nil
}

func (a Award) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}

func (a TeamAward) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}
