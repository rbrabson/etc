package ftc

type Advancements struct {
	AdvancesTo  string        `json:"advancesTo"`
	Slots       int           `json:"slots"`
	Advancement []Advancement `json:"advancement"`
}
type Advancement struct {
	Team        int    `json:"team"`
	DisplayTeam string `json:"displayTeam"`
	Slot        int    `json:"slot"`
	Criteria    string `json:"criteria"`
	Status      string `json:"status"`
}

// GetAdvancement
// GetAdvancedFrom
