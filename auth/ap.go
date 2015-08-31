package auth

import (
	"time"

	"github.com/gravitational/teleport/backend"
	"github.com/gravitational/teleport/services"
)

// AccessPoint is a interface needed by nodes to control the access
// to the node, and provide heartbeats
type AccessPoint interface {
	// GetServers returns a list of registered servers
	GetServers() ([]services.Server, error)

	// UpsertServer registers server presence, permanently if ttl is 0 or
	// for the specified duration with second resolution if it's >= 1 second
	UpsertServer(s services.Server, ttl time.Duration) error

	// GetUserCAPub returns the user certificate authority public key
	GetUserCAPub() ([]byte, error)

	// GetUserKeys returns a list of authorized keys for a given user
	// in a OpenSSH key authorized_keys format
	GetUserKeys(user string) ([]services.AuthorizedKey, error)

	// GetWebSessionsKeys returns a list of generated public keys
	// associated with user web session
	GetWebSessionsKeys(user string) ([]services.AuthorizedKey, error)

	// GetRemoteCerts returns a list of trusted remote certificates
	GetRemoteCerts(ctype, fqdn string) ([]services.RemoteCert, error)
}

type BackendAccessPoint struct {
	*services.CAService
	*services.PresenceService
	*services.ProvisioningService
	*services.UserService
	*services.WebService
}

func NewBackendAccessPoint(bk backend.Backend) *BackendAccessPoint {
	ap := BackendAccessPoint{}
	ap.CAService = services.NewCAService(bk)
	ap.PresenceService = services.NewPresenceService(bk)
	ap.ProvisioningService = services.NewProvisioningService(bk)
	ap.UserService = services.NewUserService(bk)
	ap.WebService = services.NewWebService(bk)

	return &ap
}