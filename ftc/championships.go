package ftc

type Championships struct {
	EventCount       int            `json:"eventCount"`
	GameName         string         `json:"gameName"`
	Kickoff          string         `json:"kickoff"`
	RookieStart      int            `json:"rookieStart"`
	TeamCount        int            `json:"teamCount"`
	FRCChampionships []Championship `json:"fRCChampionships"`
}
type Championship struct {
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	Location  string `json:"location"`
}
