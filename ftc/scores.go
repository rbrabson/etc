package ftc

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
