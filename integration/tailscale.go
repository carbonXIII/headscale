package integration

import (
	"net/netip"
	"net/url"

	"github.com/juanfont/headscale/integration/dockertestutil"
	"github.com/juanfont/headscale/integration/tsic"
	"tailscale.com/ipn/ipnstate"
)

// nolint
type TailscaleClient interface {
	Hostname() string
	Shutdown() error
	Version() string
	Execute(
		command []string,
		options ...dockertestutil.ExecuteCommandOption,
	) (string, string, error)
	Login(loginServer, authKey string) error
	LoginWithURL(loginServer string) (*url.URL, error)
	Logout() error
	IPs() ([]netip.Addr, error)
	FQDN() (string, error)
	Status() (*ipnstate.Status, error)
	WaitForNeedsLogin() error
	WaitForRunning() error
	WaitForPeers(expected int) error
	Ping(hostnameOrIP string, opts ...tsic.PingOption) error
	Curl(url string, opts ...tsic.CurlOption) (string, error)
	ID() string
}
