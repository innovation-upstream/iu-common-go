package helper

import "os"

// DEPRECATED! Use cloud.google.com/go/compute/metadata.GetProjectID instead
func GetProjectID() string {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		projectID = os.Getenv("GCLOUD_PROJECT_ID")
	}

	if projectID == "" {
		// Magic value to attempt to auto detect the project id using application
		// default credentials (firestore only)
		projectID = "*detect-project-id*"
	}

	return projectID
}
