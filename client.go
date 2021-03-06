package bimface

import (
	"bimface/config"
	"bimface/consts"
	"bimface/http"
	"bimface/service"
)

// Client for binface SDK
type Client struct {
	credential *config.Credential
	endpoint   *config.Endpoint

	serviceClient      *http.ServiceClient
	AccessTokenService *service.AccessTokenService
	SupportFileService *service.SupportFileService

	AppendFileService     *service.AppendFileService
	CategoryTreeService   *service.CategoryTreeService
	CompareService        *service.CompareService
	DownloadService       *service.DownloadService
	ElementService        *service.ElementService
	ElevService           *service.ElevService
	IntegrateService      *service.IntegrateService
	OfflineDatabagService *service.OfflineDatabagService
	PropertyService       *service.PropertyService
	ShareLinkService      *service.ShareLinkService
	TranslateService      *service.TranslateService
	UploadService         *service.UploadService
	ViewTokenService      *service.ViewTokenService
}

// NewClient create an bimface client.
func NewClient(appKey string, appSecret string, endpoint *config.Endpoint) *Client {
	if endpoint == nil {
		endpoint = config.NewEndpoint(consts.APIHost, consts.FileHost)
	}

	o := &Client{
		credential:    config.NewCredential(appKey, appSecret),
		endpoint:      endpoint,
		serviceClient: http.NewServiceClient(),
	}
	o.AccessTokenService = service.NewAccessTokenService(o.serviceClient, o.endpoint, o.credential)
	o.AccessTokenService = service.NewAccessTokenService(o.serviceClient, o.endpoint, o.credential)
	o.SupportFileService = service.NewSupportFileService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)

	o.AppendFileService = service.NewAppendFileService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService, o.SupportFileService)
	o.CategoryTreeService = service.NewCategoryTreeService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.CompareService = service.NewCompareService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.DownloadService = service.NewDownloadService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ElementService = service.NewElementService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ElevService = service.NewElevService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.IntegrateService = service.NewIntegrateService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.OfflineDatabagService = service.NewOfflineDatabagService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.PropertyService = service.NewPropertyService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ShareLinkService = service.NewShareLinkService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.TranslateService = service.NewTranslateService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.UploadService = service.NewUploadService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ViewTokenService = service.NewViewTokenService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)

	return o
}
