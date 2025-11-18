package dcdrsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"time"
)

// Client is the SDK client
type Client struct {
	ServerURL   string
	Token       string
	HTTPClient  *http.Client
	NoSSLVerify bool
}

// NewClient creates a new SDK client
func NewClient(serverURL, token string, noSSLVerify bool) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: noSSLVerify},
	}
	return &Client{
		ServerURL:   serverURL,
		Token:       token,
		HTTPClient:  &http.Client{Transport: tr},
		NoSSLVerify: noSSLVerify,
	}
}

// IdentResponse is the response from the dcdrIdent endpoint
type IdentResponse struct {
	InstanceID string `json:"instance_id"`
}

// Ident corresponds to the dcdrIdent API call
func (c *Client) Ident() (*IdentResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/dcdrIdent", c.ServerURL), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var identResp IdentResponse
	if err := json.NewDecoder(resp.Body).Decode(&identResp); err != nil {
		return nil, err
	}

	return &identResp, nil
}

// Auth corresponds to the dcdrAuth API call
func (c *Client) Auth() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrAuth", c.ServerURL), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// RegisterAppResponse is the response from the dcdrRegister endpoint
type RegisterAppResponse struct {
	AppID string `json:"app_id"`
}

// RegisterApp corresponds to the dcdrRegister API call
func (c *Client) RegisterApp(appName string) (*RegisterAppResponse, error) {
	requestBody, err := json.Marshal(map[string]string{"app_name": appName})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrRegister", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var regAppResp RegisterAppResponse
	if err := json.NewDecoder(resp.Body).Decode(&regAppResp); err != nil {
		return nil, err
	}

	return &regAppResp, nil
}

// SecretCreationRequest is the request to create a secret
type SecretCreationRequest struct {
	AppID      string                 `json:"app_id"`
	SecretName string                 `json:"secret_name"`
	Backend    string                 `json:"backend"`
	MountPath  string                 `json:"mount_path"`
	Data       map[string]interface{} `json:"data"`
}

// CreateSecret corresponds to the dcdrCreateSecret API call
func (c *Client) CreateSecret(createReq *SecretCreationRequest) error {
	requestBody, err := json.Marshal(createReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrCreateSecret", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// SecretRequest is the request for secret operations
type SecretRequest struct {
	AppID      string `json:"app_id"`
	SecretName string `json:"secret_name"`
}

// GetSecretResponse is the response from the dcdrGet endpoint
type GetSecretResponse struct {
	Data map[string]interface{} `json:"data"`
}

// GetSecret corresponds to the dcdrGet API call
func (c *Client) GetSecret(getReq *SecretRequest) (map[string]interface{}, error) {
	requestBody, err := json.Marshal(getReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrGet", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var secretData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&secretData); err != nil {
		return nil, err
	}

	return secretData, nil
}

// TaintSecret corresponds to the dcdrTaint API call
func (c *Client) TaintSecret(taintReq *SecretRequest) error {
	requestBody, err := json.Marshal(taintReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrTaint", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// UntaintSecret corresponds to the dcdrUntaint API call
func (c *Client) UntaintSecret(untaintReq *SecretRequest) error {
	requestBody, err := json.Marshal(untaintReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrUntaint", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// DestroySecret corresponds to the dcdrDestroy API call
func (c *Client) DestroySecret(destroyReq *SecretRequest) error {
	requestBody, err := json.Marshal(destroyReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrDestroy", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// IsTaintedResponse is the response from the dcdrIsTainted endpoint
type IsTaintedResponse struct {
	Tainted bool `json:"tainted"`
}

// IsTainted corresponds to the dcdrIsTainted API call
func (c *Client) IsTainted(isTaintedReq *SecretRequest) (*IsTaintedResponse, error) {
	requestBody, err := json.Marshal(isTaintedReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrIsTainted", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var isTaintedResp IsTaintedResponse
	if err := json.NewDecoder(resp.Body).Decode(&isTaintedResp); err != nil {
		return nil, err
	}

	return &isTaintedResp, nil
}

// RotateSecret corresponds to the dcdrRotate API call
func (c *Client) RotateSecret(rotateReq *SecretRequest) error {
	requestBody, err := json.Marshal(rotateReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrRotate", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotImplemented {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return fmt.Errorf("not implemented")
}

// App is a struct for application details
type App struct {
	AppID   string `json:"app_id"`
	AppName string `json:"app_name"`
}

// ListAppsResponse is the response from the dcdrListApps endpoint
type ListAppsResponse []App

// ListApps corresponds to the dcdrListApps API call
func (c *Client) ListApps() (ListAppsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/dcdrListApps", c.ServerURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var listAppsResp ListAppsResponse
	if err := json.NewDecoder(resp.Body).Decode(&listAppsResp); err != nil {
		return nil, err
	}

	return listAppsResp, nil
}

// Secret is a struct for secret details
type Secret struct {
	SecretName string `json:"secret_name"`
	Backend    string `json:"backend"`
	MountPath  string `json:"mount_path"`
	Tainted    bool   `json:"tainted"`
}

// ListSecretsResponse is the response from the dcdrListSecrets endpoint
type ListSecretsResponse []Secret

// ListSecrets corresponds to the dcdrListSecrets API call
func (c *Client) ListSecrets(appID string) (ListSecretsResponse, error) {
	requestBody, err := json.Marshal(map[string]string{"app_id": appID})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrListSecrets", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var listSecretsResp ListSecretsResponse
	if err := json.NewDecoder(resp.Body).Decode(&listSecretsResp); err != nil {
		return nil, err
	}

	return listSecretsResp, nil
}

// Backend is a struct for backend details
type Backend struct {
	Backend         string `json:"backend"`
	NumApplications int    `json:"num_applications"`
	NumSecrets      int    `json:"num_secrets"`
}

// ListBackendsResponse is the response from the dcdrListBackends endpoint
type ListBackendsResponse []Backend

// ListBackends corresponds to the dcdrListBackends API call
func (c *Client) ListBackends() (ListBackendsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/dcdrListBackends", c.ServerURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var listBackendsResp ListBackendsResponse
	if err := json.NewDecoder(resp.Body).Decode(&listBackendsResp); err != nil {
		return nil, err
	}

	return listBackendsResp, nil
}

// DeleteApp corresponds to the dcdrDeleteApp API call
func (c *Client) DeleteApp(appID string) error {
	requestBody, err := json.Marshal(map[string]string{"app_id": appID})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrDeleteApp", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// WhoamiResponse is the response from the dcdrWhoami endpoint
type WhoamiResponse struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	IsRoot   bool   `json:"is_root"`
	AppID    string `json:"app_id,omitempty"`
	AppName  string `json:"app_name,omitempty"`
}

// Whoami corresponds to the dcdrWhoami API call
func (c *Client) Whoami() (*WhoamiResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/dcdrWhoami", c.ServerURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var whoamiResp WhoamiResponse
	if err := json.NewDecoder(resp.Body).Decode(&whoamiResp); err != nil {
		return nil, err
	}

	return &whoamiResp, nil
}

// CreateAppUserRequest is the request to create an application user
type CreateAppUserRequest struct {
	AppID    string `json:"app_id"`
	UserName string `json:"user_name"`
}

// CreateAppUserResponse is the response from the dcdrAppUser/create endpoint
type CreateAppUserResponse struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

// CreateAppUser corresponds to the dcdrAppUser/create API call
func (c *Client) CreateAppUser(createReq *CreateAppUserRequest) (*CreateAppUserResponse, error) {
	requestBody, err := json.Marshal(createReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrAppUser/create", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var createAppUserResp CreateAppUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&createAppUserResp); err != nil {
		return nil, err
	}

	return &createAppUserResp, nil
}

// AppUser is a struct for application user details
type AppUser struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Status   string `json:"status"`
}

// ListAppUsersResponse is the response from the dcdrAppUser/list endpoint
type ListAppUsersResponse []AppUser

// ListAppUsers corresponds to the dcdrAppUser/list API call
func (c *Client) ListAppUsers(appID string) (ListAppUsersResponse, error) {
	requestBody, err := json.Marshal(map[string]string{"app_id": appID})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/dcdrAppUser/list", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var listAppUsersResp ListAppUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&listAppUsersResp); err != nil {
		return nil, err
	}

	return listAppUsersResp, nil
}

// AppUserRequest is the request for application user operations
type AppUserRequest struct {
	AppID  string `json:"app_id"`
	UserID string `json:"user_id"`
}

// SuspendAppUser corresponds to the dcdrAppUser/suspend API call
func (c *Client) SuspendAppUser(suspendReq *AppUserRequest) error {
	requestBody, err := json.Marshal(suspendReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrAppUser/suspend", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// UnsuspendAppUser corresponds to the dcdrAppUser/unsuspend API call
func (c *Client) UnsuspendAppUser(unsuspendReq *AppUserRequest) error {
	requestBody, err := json.Marshal(unsuspendReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrAppUser/unsuspend", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// DeleteAppUser corresponds to the dcdrAppUser/delete API call
func (c *Client) DeleteAppUser(deleteReq *AppUserRequest) error {
	requestBody, err := json.Marshal(deleteReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrAppUser/delete", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// GetAppUserTokenResponse is the response from the dcdrAppUser/getToken endpoint
type GetAppUserTokenResponse struct {
	Token string `json:"token"`
}

// GetAppUserToken corresponds to the dcdrAppUser/getToken API call
func (c *Client) GetAppUserToken(getTokenReq *AppUserRequest) (*GetAppUserTokenResponse, error) {
	requestBody, err := json.Marshal(getTokenReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcdrAppUser/getToken", c.ServerURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var getTokenResp GetAppUserTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&getTokenResp); err != nil {
		return nil, err
	}

	return &getTokenResp, nil
}

// Logout corresponds to the dcdrLogout API call
func (c *Client) Logout() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/logout", c.ServerURL), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

// DownloadAuditLogs downloads the audit log bundle.
func (c *Client) DownloadAuditLogs(format string, outFile string) (string, error) {
	if format != "csv" && format != "json" {
		return "", fmt.Errorf("format must be either 'csv' or 'json'")
	}

	url := fmt.Sprintf("%s/api/dcdrAudit/download?format=%s", c.ServerURL, format)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	if outFile == "" {
		disposition := resp.Header.Get("Content-Disposition")
		if disposition != "" {
			_, params, err := mime.ParseMediaType(disposition)
			if err == nil {
				outFile = params["filename"]
			}
		}
	}

	if outFile == "" {
		timestamp := time.Now().Format("2006-01-02-15-04-05")
		outFile = fmt.Sprintf("dcdr-audit-logs-%s.zip", timestamp)
	}

	f, err := os.Create(outFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}

	return outFile, nil
}
