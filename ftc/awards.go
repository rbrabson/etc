package ftc

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
