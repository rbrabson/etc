package ftc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

// Advancements is the list of teams that advanced from or to a given tournament
type AdvancementsTo struct {
	AdvancesTo  string        `json:"advancesTo"`
	Slots       int           `json:"slots"`
	Advancement []Advancement `json:"advancement"`
}

type AdvancementsFrom []struct {
	AdvancedFrom       string        `json:"advancedFrom"`
	AdvancedFromRegion any           `json:"advancedFromRegion"`
	Slots              int           `json:"slots"`
	Advancement        []Advancement `json:"advancement"`
}

// Advancement is the advancement information for a given team
type Advancement struct {
	Team        int    `json:"team"`
	DisplayTeam string `json:"displayTeam"`
	Slot        int    `json:"slot"`
	Criteria    string `json:"criteria"`
	Status      string `json:"status"`
}

// GetAdvancementFrom returns the source events from which teams advanced from to reach
// the specified event.
func GetAdvancementFrom(season string, eventCode string) (*AdvancementsFrom, error) {
	url := fmt.Sprintf("%s/%s/advancement/%s", server, season, eventCode)

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output AdvancementsFrom
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return &output, nil
}

// GetAdvancementTo returns the list of teams advancing from the event and to which event
// they are advancing.
func GetAdvancementTo(season string, eventCode string, excludeSkipped ...bool) (*AdvancementsTo, error) {
	sb := strings.Builder{}
	sb.WriteString(server)
	sb.WriteString("/")
	sb.WriteString(season)
	sb.WriteString("/advancement/")
	sb.WriteString(eventCode)
	if len(excludeSkipped) > 0 {
		sb.WriteString("?excludedSkipped=")
		if excludeSkipped[0] {
			sb.WriteString("true")
		} else {
			sb.WriteString("false")
		}
	}
	url := sb.String()

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output AdvancementsTo
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return &output, nil
}
