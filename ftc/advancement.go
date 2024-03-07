package ftc

// Advancements is the list of teams that advanced from or to a given tournament
type Advancements struct {
	AdvancesTo  string        `json:"advancesTo"`
	Slots       int           `json:"slots"`
	Advancement []Advancement `json:"advancement"`
}

// Advancement is the advancement information for a given team
type Advancement struct {
	Team        int    `json:"team"`
	DisplayTeam string `json:"displayTeam"`
	Slot        int    `json:"slot"`
	Criteria    string `json:"criteria"`
	Status      string `json:"status"`
}

// GetAdvancement
// GetAdvancedFrom
