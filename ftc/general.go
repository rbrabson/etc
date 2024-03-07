package ftc

import (
	"encoding/json"

	"github.com/rbrabson/ftc/internal/ftchttp"
)

// General provives information for the FTC server API
type General struct {
	Name                    string  `json:"name,omitempty"`
	APIVersion              string  `json:"apiVersion,omitempty"`
	ServiceMainifestName    *string `json:"serviceMainifestName,omitempty"`
	ServiceMainifestVersion *string `json:"serviceMainifestVersion,omitempty"`
	CodePackageName         string  `json:"codePackageName,omitempty"`
	CodePackageVersion      string  `json:"codePackageVersion,omitempty"`
	Status                  string  `json:"status,omitempty"`
	CurrentSeason           int     `json:"currentSeason,omitempty"`
	MaxSeason               int     `json:"maxSeason,omitempty"`
}

// GetGeneral returns the information for the FTC server API
func GetGeneral() (*General, error) {
	url := server
	body, err := ftchttp.Get(url)
	if err != nil {
		return nil, err
	}

	var output General
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	// Return the output
	return &output, nil
}

// String returns a string representation of General. In this case, it is a json string.
func (general *General) String() string {
	body, _ := json.Marshal(general)
	return string(body)
}
