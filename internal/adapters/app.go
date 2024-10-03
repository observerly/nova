/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nox/adapters/app
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package adapters

/*****************************************************************************************************************/

import (
	"context"
	"os"
	"path/filepath"

	"cloud.google.com/go/compute/metadata"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

/*****************************************************************************************************************/

func SetupFirebaseApp() (*firebase.App, error) {
	// Check if the application is running on Google Cloud Platform:
	onGCP := metadata.OnGCE()

	// If the application is running on Google Cloud Platform, we can use the default credentials:
	if onGCP {
		return firebase.NewApp(context.Background(), nil)
	}

	// Get the $PWD of the workspace root:
	pwd, _ := os.Getwd()

	// Get the path to the service account credentials:
	sa := filepath.Join(pwd, "gcp-service-account.json")

	// Return the Firebase app with the provided credentials:
	return firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(sa))
}

/*****************************************************************************************************************/
