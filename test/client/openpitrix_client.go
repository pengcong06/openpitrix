// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/client/access_manager"
	"openpitrix.io/openpitrix/test/client/account_manager"
	"openpitrix.io/openpitrix/test/client/app_manager"
	"openpitrix.io/openpitrix/test/client/attachment_service"
	"openpitrix.io/openpitrix/test/client/category_manager"
	"openpitrix.io/openpitrix/test/client/isv_manager"
	"openpitrix.io/openpitrix/test/client/release_manager"
	"openpitrix.io/openpitrix/test/client/repo_indexer"
	"openpitrix.io/openpitrix/test/client/repo_manager"
	"openpitrix.io/openpitrix/test/client/runtime_manager"
	"openpitrix.io/openpitrix/test/client/service_config"
	"openpitrix.io/openpitrix/test/client/token_manager"
)

// Default openpitrix HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new openpitrix HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Openpitrix {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new openpitrix HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Openpitrix {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new openpitrix client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Openpitrix {
	cli := new(Openpitrix)
	cli.Transport = transport

	cli.AccessManager = access_manager.New(transport, formats)

	cli.AccountManager = account_manager.New(transport, formats)

	cli.AppManager = app_manager.New(transport, formats)

	cli.AttachmentService = attachment_service.New(transport, formats)

	cli.CategoryManager = category_manager.New(transport, formats)

	cli.IsvManager = isv_manager.New(transport, formats)

	cli.ReleaseManager = release_manager.New(transport, formats)

	cli.RepoIndexer = repo_indexer.New(transport, formats)

	cli.RepoManager = repo_manager.New(transport, formats)

	cli.RuntimeManager = runtime_manager.New(transport, formats)

	cli.ServiceConfig = service_config.New(transport, formats)

	cli.TokenManager = token_manager.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Openpitrix is a client for openpitrix
type Openpitrix struct {
	AccessManager *access_manager.Client

	AccountManager *account_manager.Client

	AppManager *app_manager.Client

	AttachmentService *attachment_service.Client

	CategoryManager *category_manager.Client

	IsvManager *isv_manager.Client

	ReleaseManager *release_manager.Client

	RepoIndexer *repo_indexer.Client

	RepoManager *repo_manager.Client

	RuntimeManager *runtime_manager.Client

	ServiceConfig *service_config.Client

	TokenManager *token_manager.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Openpitrix) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AccessManager.SetTransport(transport)

	c.AccountManager.SetTransport(transport)

	c.AppManager.SetTransport(transport)

	c.AttachmentService.SetTransport(transport)

	c.CategoryManager.SetTransport(transport)

	c.IsvManager.SetTransport(transport)

	c.ReleaseManager.SetTransport(transport)

	c.RepoIndexer.SetTransport(transport)

	c.RepoManager.SetTransport(transport)

	c.RuntimeManager.SetTransport(transport)

	c.ServiceConfig.SetTransport(transport)

	c.TokenManager.SetTransport(transport)

}
