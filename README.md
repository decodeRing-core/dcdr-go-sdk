<a name="top"></a>
[![decodeRing Core Server](https://decodering.org/wp-content/uploads/2025/11/Git-Banner-2-scaled.png)](https://decodering.org)
![Go](https://img.shields.io/badge/Go-1.24.4-blue) ![OS](https://img.shields.io/badge/OS-Linux_Windows_MacOS-green) ![CPU](https://img.shields.io/badge/CPU-x64-FF8C00) ![Release](https://img.shields.io/badge/Release-v1.0.0-blue) ![Release Date](https://img.shields.io/badge/Release_Date-November_2025-blue) ![License](https://img.shields.io/badge/License-Apache_2.0-blue)

⭐ Star us on GitHub — your support means a lot to us! 🙏😊

## Table of Contents
- [About](#-about)
- [Installation](#-installation)
- [Usage](#-usage)
- [Example Program](#-example-program)

## 🚀 About

This SDK provides a Go interface for accessing secrets through a decodeRing server.

## 💾 Installation

```bash
go get github.com/decodeRing-core/dcdr-go-sdk
```

## 📘 Usage

First, create a new client:

```go
client := dcdrsdk.NewClient("http://localhost:8301", "your-token", true)
```

### 📂 Functions

Here is a list of the available functions:

*   `NewClient(serverURL, token string, noSSLVerify bool) *Client`: Creates a new SDK client.
*   `Ident() (*IdentResponse, error)`: Corresponds to `dcdrIdent`.
*   `Auth() error`: Corresponds to `dcdrAuth`.
*   `RegisterApp(appName string) (*RegisterAppResponse, error)`: Corresponds to `dcdrRegister`.
*   `CreateSecret(req *SecretCreationRequest) error`: Corresponds to `dcdrCreateSecret`.
*   `GetSecret(req *SecretRequest) (*GetSecretResponse, error)`: Corresponds to `dcdrGet`.
*   `TaintSecret(req *SecretRequest) error`: Corresponds to `dcdrTaint`.
*   `UntaintSecret(req *SecretRequest) error`: Corresponds to `dcdrUntaint`.
*   `DestroySecret(req *SecretRequest) error`: Corresponds to `dcdrDestroy`.
*   `IsTainted(req *SecretRequest) (*IsTaintedResponse, error)`: Corresponds to `dcdrIsTainted`.
*   `RotateSecret(req *SecretRequest) error`: Corresponds to `dcdrRotate`.
*   `ListApps() (*ListAppsResponse, error)`: Corresponds to `dcdrListApps`.
*   `ListSecrets(appID string) (*ListSecretsResponse, error)`: Corresponds to `dcdrListSecrets`.
*   `ListBackends() (*ListBackendsResponse, error)`: Corresponds to `dcdrListBackends`.
*   `DeleteApp(appID string) error`: Corresponds to `dcdrDeleteApp`.
*   `Whoami() (*WhoamiResponse, error)`: Corresponds to `dcdrWhoami`.
*   `CreateAppUser(req *CreateAppUserRequest) (*CreateAppUserResponse, error)`: Corresponds to `dcdrAppUser/create`.
*   `ListAppUsers(appID string) (*ListAppUsersResponse, error)`: Corresponds to `dcdrAppUser/list`.
*   `SuspendAppUser(req *AppUserRequest) error`: Corresponds to `dcdrAppUser/suspend`.
*   `UnsuspendAppUser(req *AppUserRequest) error`: Corresponds to `dcdrAppUser/unsuspend`.
*   `DeleteAppUser(req *AppUserRequest) error`: Corresponds to `dcdrAppUser/delete`.
*   `GetAppUserToken(req *AppUserRequest) (*GetAppUserTokenResponse, error)`: Corresponds to `dcdrAppUser/getToken`.
*   `Logout() error`: Corresponds to `dcdrLogout`.
*   `DownloadAuditLogs(format string, outFile string) (string, error)`: Downloads the audit log bundle.

[Back to top](#top)

## 👀 Example Program

There is an example go program in the `example/` folder.

[Back to top](#top)
