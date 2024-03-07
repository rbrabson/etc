package ftc

import (
	"encoding/json"

	"github.com/rbrabson/ftc/ftchttp"
)

type General struct {
	Name                    string `json:"name"`
	APIVersion              string `json:"apiVersion"`
	ServiceMainifestName    any    `json:"serviceMainifestName"`
	ServiceMainifestVersion any    `json:"serviceMainifestVersion"`
	CodePackageName         string `json:"codePackageName"`
	CodePackageVersion      string `json:"codePackageVersion"`
	Status                  string `json:"status"`
	CurrentSeason           int    `json:"currentSeason"`
	MaxSeason               int    `json:"maxSeason"`
}

func GetGeneral() (*General, error) {
	url := server
	body, err := ftchttp.Get(url)

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
