package ftc

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

type AllianceSelections struct {
	Selections []AllianceSelection `json:"selections"`
	Count      int                 `json:"count"`
}
type AllianceSelection struct {
	Index  int    `json:"index"`
	Team   int    `json:"team"`
	Result string `json:"result"`
}

type Alliances struct {
	Alliances []Alliance `json:"alliances"`
	Count     int        `json:"count"`
}
type Alliance struct {
	Number         int     `json:"number"`
	Name           string  `json:"name"`
	Captain        int     `json:"captain"`
	CaptainDisplay string  `json:"captainDisplay"`
	Round1         int     `json:"round1,omitempty"`
	Round1Display  string  `json:"round1Display,omitempty"`
	Round2         int     `json:"round2,omitempty"`
	Round2Display  string  `json:"round2Display,omitempty"`
	Round3         *int    `json:"round3,omitempty"`
	Backup         *string `json:"backup,omitempty"`
	BackupReplaced *string `json:"backupReplaced,omitempty"`
}

// GetEventAlliances returns the alliance selectsions for the playoffs for the given event.
func GetEventAlliances(season string, eventCode string) ([]Alliance, error) {
	url := fmt.Sprintf("%s/%s/alliances/%s", server, season, eventCode)

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output []Alliance
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output, nil
}

// GetAllianceSelections returns the teams that were selected into alliances for the given event.
func GetAllianceSelections(season string, eventCode string) ([]AllianceSelection, error) {
	url := fmt.Sprintf("%s/%s/alliances/%s/selection", server, season, eventCode)

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output []AllianceSelection
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output, nil
}

func (a AllianceSelection) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}

func (a Alliances) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}

func (a Alliance) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}
