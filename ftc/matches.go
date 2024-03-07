package ftc

// Type of FTC match
type MatchType string

const (
	QUALIFIER MatchType = "qual"
	PLAYOFF   MatchType = "playoff"
)

type Matches struct {
	Matches []Match `json:"matches"`
}
type Match struct {
	ActualStartTime string      `json:"actualStartTime"`
	Description     string      `json:"description"`
	TournamentLevel string      `json:"tournamentLevel"`
	Series          int         `json:"series"`
	MatchNumber     int         `json:"matchNumber"`
	ScoreRedFinal   int         `json:"scoreRedFinal"`
	ScoreRedFoul    int         `json:"scoreRedFoul"`
	ScoreRedAuto    int         `json:"scoreRedAuto"`
	ScoreBlueFinal  int         `json:"scoreBlueFinal"`
	ScoreBlueFoul   int         `json:"scoreBlueFoul"`
	ScoreBlueAuto   int         `json:"scoreBlueAuto"`
	PostResultTime  string      `json:"postResultTime"`
	Teams           []MatchTeam `json:"teams"`
	ModifiedOn      string      `json:"modifiedOn"`
}
type MatchTeam struct {
	TeamNumber int    `json:"teamNumber"`
	Station    string `json:"station"`
	DQ         bool   `json:"dq"`
	OnField    bool   `json:"onField"`
}
