package detect

import (
	"os"
	"path/filepath"
)

type rule struct {
	markers []string
	profile string
}

var rules = []rule{
	{markers: []string{"Cargo.toml"}, profile: "rust"},
	{markers: []string{"go.mod"}, profile: "go"},
	{markers: []string{"pyproject.toml", "setup.py", "requirements.txt", "Pipfile"}, profile: "python"},
	{markers: []string{"package.json"}, profile: "web"},
	{markers: []string{"helmfile.yaml", "kustomization.yaml", "Chart.yaml"}, profile: "k8s"},
	{markers: []string{"Dockerfile"}, profile: "devops"},
}

// Profile scans dir for marker files and returns the matching profile name.
// Returns empty string if no profile is detected.
func Profile(dir string) string {
	for _, r := range rules {
		for _, m := range r.markers {
			if fileExists(filepath.Join(dir, m)) {
				return r.profile
			}
		}
	}
	return ""
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
