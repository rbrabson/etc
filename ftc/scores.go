package ftc

import (
	"encoding/json"
	"strings"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

type Scores struct {
	MatchScores []MatchScores `json:"matchScores"`
}
type MatchScores struct {
	MatchLevel    string          `json:"matchLevel"`
	MatchSeries   int             `json:"matchSeries"`
	MatchNumber   int             `json:"matchNumber"`
	Randomization int             `json:"randomization"`
	Alliances     []MatchAlliance `json:"alliances"`
}

type MatchAlliance struct {
	Alliance                string `json:"alliance"`
	Team                    int    `json:"team"`
	InitTeamProp1           bool   `json:"initTeamProp1"`
	InitTeamProp2           bool   `json:"initTeamProp2"`
	Robot1Auto              bool   `json:"robot1Auto"`
	Robot2Auto              bool   `json:"robot2Auto"`
	SpikeMarkPixel1         bool   `json:"spikeMarkPixel1"`
	SpikeMarkPixel2         bool   `json:"spikeMarkPixel2"`
	TargetBackdropPixel1    bool   `json:"targetBackdropPixel1"`
	TargetBackdropPixel2    bool   `json:"targetBackdropPixel2"`
	AutoBackdrop            int    `json:"autoBackdrop"`
	AutoBackstage           int    `json:"autoBackstage"`
	DcBackdrop              int    `json:"dcBackdrop"`
	DcBackstage             int    `json:"dcBackstage"`
	Mosaics                 int    `json:"mosaics"`
	MaxSetLine              int    `json:"maxSetLine"`
	EgRobot1                string `json:"egRobot1"`
	EgRobot2                string `json:"egRobot2"`
	Drone1                  int    `json:"drone1"`
	Drone2                  int    `json:"drone2"`
	MinorPenalties          int    `json:"minorPenalties"`
	MajorPenalties          int    `json:"majorPenalties"`
	AutoNavigatingPoints    int    `json:"autoNavigatingPoints"`
	AutoRandomizationPoints int    `json:"autoRandomizationPoints"`
	AutoBackstagePoints     int    `json:"autoBackstagePoints"`
	AutoBackdropPoints      int    `json:"autoBackdropPoints"`
	DcBackdropPoints        int    `json:"dcBackdropPoints"`
	DcBackstagePoints       int    `json:"dcBackstagePoints"`
	MosaicPoints            int    `json:"mosaicPoints"`
	SetBonusPoints          int    `json:"setBonusPoints"`
	EgLocationPoints        int    `json:"egLocationPoints"`
	EgDronePoints           int    `json:"egDronePoints"`
	AutoPoints              int    `json:"autoPoints"`
	DcPoints                int    `json:"dcPoints"`
	EndgamePoints           int    `json:"endgamePoints"`
	PenaltyPointsCommitted  int    `json:"penaltyPointsCommitted"`
	PrePenaltyTotal         int    `json:"prePenaltyTotal"`
	TotalPoints             int    `json:"totalPoints"`
}

func GetEventMatchResults(season string, eventCode string, tournamentLevel MatchType, teamNumber ...string) ([]MatchScores, error) {
	sb := strings.Builder{}
	sb.WriteString(server)
	sb.WriteString("/")
	sb.WriteString(season)
	sb.WriteString("/matches/")
	sb.WriteString(eventCode)
	sb.WriteString("?")
	sb.WriteString(string(tournamentLevel))
	if len(teamNumber) > 0 {
		sb.WriteString("&")
		sb.WriteString(teamNumber[0])
	}
	url := sb.String()

	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output Scores
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return output.MatchScores, nil
}

func (ms MatchScores) String() string {
	body, _ := json.Marshal(ms)
	return string(body)
}

func (ma MatchAlliance) String() string {
	body, _ := json.Marshal(ma)
	return string(body)
}
