package main

import (
	"fmt"
	"log"
//	"os"

	"dcdr.local/sdk/dcdr-sdk-go-v1"
)

func main() {
	// Replace with your server URL and root token
	client := dcdrsdk.NewClient("https://localhost:8301", "fJH4EBwZLwraWq4ftsQoSIIvb6", true)

	// 1. Get Server Identity
	fmt.Println("1. Getting server identity...")
	ident, err := client.Ident()
	if err != nil {
		log.Fatalf("Failed to get server identity: %v", err)
	}
	fmt.Printf("   - Server Instance ID: %s\n", ident.InstanceID)

	// 2. Authenticate
	fmt.Println("2. Authenticating...")
	if err := client.Auth(); err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}
	fmt.Println("   - Authentication successful.")

	// 3. Register a new application
	fmt.Println("3. Registering a new application...")
	regAppResp, err := client.RegisterApp("MyTestApp")
	if err != nil {
		log.Fatalf("Failed to register application: %v", err)
	}
	appID := regAppResp.AppID
	fmt.Printf("   - Application registered with ID: %s\n", appID)

	// 4. Create an application user
	fmt.Println("4. Creating an application user...")
	createAppUserResp, err := client.CreateAppUser(&dcdrsdk.CreateAppUserRequest{
		AppID:    appID,
		UserName: "testuser",
	})
	if err != nil {
		log.Fatalf("Failed to create application user: %v", err)
	}
	fmt.Printf("   - Application user created with ID: %s\n", createAppUserResp.UserID)
	appUserToken := createAppUserResp.Token

	// 5. Create a new client for the application user
	fmt.Println("5. Creating a new client for the application user...")
	appClient := dcdrsdk.NewClient("https://localhost:8301", appUserToken, true)

	// 6. Create a secret
	fmt.Println("6. Creating a secret...")
	err = appClient.CreateSecret(&dcdrsdk.SecretCreationRequest{
		AppID:      appID,
		SecretName: "MySecret",
		Backend:    "vault-1", // Make sure this backend is configured in your server
		MountPath:  "test",
		Data: map[string]interface{}{
			"username": "myuser",
			"password": "mypassword",
		},
	})
	if err != nil {
		log.Fatalf("Failed to create secret: %v", err)
	}
	fmt.Println("   - Secret 'MySecret' created successfully.")

	// 7. Get the secret
	fmt.Println("7. Getting the secret...")
	secret, err := appClient.GetSecret(&dcdrsdk.SecretRequest{AppID: appID, SecretName: "MySecret"})
	if err != nil {
		log.Fatalf("Failed to get secret: %v", err)
	}
	fmt.Printf("   - Secret 'MySecret' data: %v\n", secret)

	// 8. Taint the secret
	fmt.Println("8. Tainting the secret...")
	if err := appClient.TaintSecret(&dcdrsdk.SecretRequest{AppID: appID, SecretName: "MySecret"}); err != nil {
		log.Fatalf("Failed to taint secret: %v", err)
	}
	fmt.Println("   - Secret 'MySecret' tainted successfully.")

	// 9. Check if the secret is tainted
	fmt.Println("9. Checking if the secret is tainted...")
	isTaintedResp, err := appClient.IsTainted(&dcdrsdk.SecretRequest{AppID: appID, SecretName: "MySecret"})
	if err != nil {
		log.Fatalf("Failed to check if secret is tainted: %v", err)
	}
	fmt.Printf("   - Is 'MySecret' tainted? %v\n", isTaintedResp.Tainted)

	// 10. Untaint the secret
	fmt.Println("10. Untainting the secret...")
	if err := appClient.UntaintSecret(&dcdrsdk.SecretRequest{AppID: appID, SecretName: "MySecret"}); err != nil {
		log.Fatalf("Failed to untaint secret: %v", err)
	}
	fmt.Println("   - Secret 'MySecret' untainted successfully.")

	// 11. List applications
	fmt.Println("11. Listing applications...")
	apps, err := client.ListApps()
	if err != nil {
		log.Fatalf("Failed to list applications: %v", err)
	}
	fmt.Printf("   - Found %d application(s).\n", len(apps))
	for _, app := range apps {
		fmt.Printf("     - App ID: %s, App Name: %s\n", app.AppID, app.AppName)
	}

	// 12. List secrets for the application
	fmt.Println("12. Listing secrets for the application...")
	secrets, err := client.ListSecrets(appID)
	if err != nil {
		log.Fatalf("Failed to list secrets: %v", err)
	}
	fmt.Printf("   - Found %d secret(s) for app '%s'.\n", len(secrets), appID)
	for _, s := range secrets {
		fmt.Printf("     - Secret Name: %s, Backend: %s, Tainted: %v\n", s.SecretName, s.Backend, s.Tainted)
	}

	// 13. List backends
	fmt.Println("13. Listing backends...")
	backends, err := client.ListBackends()
	if err != nil {
		log.Fatalf("Failed to list backends: %v", err)
	}
	fmt.Printf("   - Found %d backend(s).\n", len(backends))
	for _, b := range backends {
		fmt.Printf("     - Backend: %s, Applications: %d, Secrets: %d\n", b.Backend, b.NumApplications, b.NumSecrets)
	}

	// 14. Destroy the secret
	fmt.Println("14. Destroying the secret...")
	if err := appClient.DestroySecret(&dcdrsdk.SecretRequest{AppID: appID, SecretName: "MySecret"}); err != nil {
		log.Fatalf("Failed to destroy secret: %v", err)
	}
	fmt.Println("   - Secret 'MySecret' destroyed successfully.")

	// 15. Delete the application user
	fmt.Println("15. Deleting the application user...")
	if err := client.DeleteAppUser(&dcdrsdk.AppUserRequest{AppID: appID, UserID: createAppUserResp.UserID}); err != nil {
		log.Fatalf("Failed to delete application user: %v", err)
	}
	fmt.Println("   - Application user deleted successfully.")

	// 16. Delete the application
	fmt.Println("16. Deleting the application...")
	if err := client.DeleteApp(appID); err != nil {
		log.Fatalf("Failed to delete application: %v", err)
	}
	fmt.Println("   - Application deleted successfully.")

	// 17. Download Audit Logs
	fmt.Println("17. Downloading audit logs...")
	downloadedFile, err := client.DownloadAuditLogs("csv", "")
	if err != nil {
		log.Fatalf("Failed to download audit logs: %v", err)
	}
	fmt.Printf("   - Audit logs downloaded to: %s\n", downloadedFile)
	// Clean up the downloaded file
	//if err := os.Remove(downloadedFile); err != nil {
//		log.Printf("Warning: failed to clean up downloaded file: %v", err)
//	}

	// 18. Logout
	fmt.Println("18. Logging out...")
	if err := client.Logout(); err != nil {
		log.Fatalf("Failed to logout: %v", err)
	}
	fmt.Println("   - Logout successful.")
}
