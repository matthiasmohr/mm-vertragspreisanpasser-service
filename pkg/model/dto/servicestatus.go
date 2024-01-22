package dto

type AppStatus struct {
	Version     string `json:"version"`
	BuildDate   string `json:"buildDate"`
	Description string `json:"description"`
	CommitHash  string `json:"commitId"`
	CommitDate  string `json:"commitDate"`
	BuildBranch string `json:"buildBranch"`
}
