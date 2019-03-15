package config

type GatewaySpec struct {
	// PathPrefixes list the set of path prefixes that should be handled by
	// this gateway.
	PathPrefixes []string

	// UseSubdomains indicates whether or not this gateway uses subdomains
	// for IPFS resources instead of paths. That is: http://CID.ipfs.GATEWAY/...
	//
	// If this flag is set, any /ipns/$id and/or /ipfs/$id paths in PathPrefixes
	// will be permanently redirected to http://$id.[ipns|ipfs].$gateway/.
	//
	// We do not support using both paths and subdomains for a single domain
	// for security reasons.
	UseSubdomains bool

	// RootRedirect is the path to which requests to `/` on this gateway
	// should be redirected.
	RootRedirect string
}

// Gateway contains options for the HTTP gateway server.
type Gateway struct {

	// HTTPHeaders configures the headers that should be returned by this
	// gateway.
	HTTPHeaders map[string][]string // HTTP headers to return with the gateway

	// Writable enables PUT/POST request handling by this gateway. Usually,
	// writing is done through the API, not the gateway.
	Writable bool

	// FIXME: Not yet implemented
	APICommands []string

	// NoFetch configures the gateway to _not_ fetch blocks in response to
	// requests.
	NoFetch bool

	// GatewaySpec configures the default behavior for this gateway.
	GatewaySpec

	// PublicGateways configures behavior of known public gateways.
	PublicGateways map[string]*GatewaySpec
}
