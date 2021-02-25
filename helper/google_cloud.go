package helper

import "os"

func GetProjectID() string {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		projectID = os.Getenv("GCLOUD_PROJECT_ID")
	}

	if projectID == "" {
		// Magic value to attempt to auto detect the project id using application default credentials
		projectID = "*detect-project-id*"
	}

	return projectID
}
