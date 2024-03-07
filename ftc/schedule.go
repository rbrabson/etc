package ftc

type Schedules struct {
	Schedule []Schedule `json:"schedule"`
}
type Schedule struct {
	Description     string          `json:"description,omitempty"`
	Field           string          `json:"field,omitempty"`
	TournamentLevel string          `json:"tournamentLevel,omitempty"`
	StartTime       string          `json:"startTime,omitempty"`
	Series          int             `json:"series,omitempty"`
	MatchNumber     int             `json:"matchNumber,omitempty"`
	Teams           []ScheduledTeam `json:"teams"`
	ModifiedOn      string          `json:"modifiedOn,omitempty"`
}
type ScheduledTeam struct {
	TeamNumber        int    `json:"teamNumber,omitempty"`
	DisplayTeamNumber string `json:"displayTeamNumber,omitempty"`
	Station           string `json:"station,omitempty"`
	Team              string `json:"team,omitempty"`
	TeamName          string `json:"teamName,omitempty"`
	Surrogate         bool   `json:"surrogate,omitempty"`
	NoShow            bool   `json:"noShow,omitempty"`
}
