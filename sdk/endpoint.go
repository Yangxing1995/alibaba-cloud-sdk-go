package sdk

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type Resolver interface {
	TryResolve(client *Client, request requests.AcsRequest) (endpoint string, err error)
}

// GetDefaultResolver get default resolver
func GetDefaultResolver() Resolver {
	return &DefaultResolver{}
}

type DefaultResolver struct {
}

func (resolver *DefaultResolver) TryResolve(client *Client, request requests.AcsRequest) (endpoint string, err error) {
	regionId := client.regionId
	if len(request.GetRegionId()) > 0 {
		regionId = request.GetRegionId()
	}

	// resolve endpoint
	endpoint = request.GetDomain()

	if endpoint == "" && client.Domain != "" {
		endpoint = client.Domain
	}

	if endpoint == "" {
		endpoint = endpoints.GetEndpointFromMap(regionId, request.GetProduct())
	}

	if endpoint == "" && client.EndpointType != "" &&
		(request.GetProduct() != "Sts" || len(request.GetQueryParams()) == 0) {
		if client.EndpointMap != nil && client.Network == "" || client.Network == "public" {
			endpoint = client.EndpointMap[regionId]
		}

		if endpoint == "" {
			endpoint, err = client.GetEndpointRules(regionId, request.GetProduct())
			if err != nil {
				return
			}
		}
	}

	if endpoint == "" {
		resolveParam := &endpoints.ResolveParam{
			Domain:               request.GetDomain(),
			Product:              request.GetProduct(),
			RegionId:             regionId,
			LocationProduct:      request.GetLocationServiceCode(),
			LocationEndpointType: request.GetLocationEndpointType(),
			CommonApi:            client.ProcessCommonRequest,
		}
		endpoint, err = endpoints.Resolve(resolveParam)
		if err != nil {
			return
		}
	}
	return
}
