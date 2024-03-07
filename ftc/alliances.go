package ftc

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
