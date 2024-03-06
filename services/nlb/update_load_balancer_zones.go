package nlb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateLoadBalancerZones invokes the nlb.UpdateLoadBalancerZones API synchronously
func (client *Client) UpdateLoadBalancerZones(request *UpdateLoadBalancerZonesRequest) (response *UpdateLoadBalancerZonesResponse, err error) {
	response = CreateUpdateLoadBalancerZonesResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLoadBalancerZonesWithChan invokes the nlb.UpdateLoadBalancerZones API asynchronously
func (client *Client) UpdateLoadBalancerZonesWithChan(request *UpdateLoadBalancerZonesRequest) (<-chan *UpdateLoadBalancerZonesResponse, <-chan error) {
	responseChan := make(chan *UpdateLoadBalancerZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLoadBalancerZones(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateLoadBalancerZonesWithCallback invokes the nlb.UpdateLoadBalancerZones API asynchronously
func (client *Client) UpdateLoadBalancerZonesWithCallback(request *UpdateLoadBalancerZonesRequest, callback func(response *UpdateLoadBalancerZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLoadBalancerZonesResponse
		var err error
		defer close(result)
		response, err = client.UpdateLoadBalancerZones(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateLoadBalancerZonesRequest is the request struct for api UpdateLoadBalancerZones
type UpdateLoadBalancerZonesRequest struct {
	*requests.RpcRequest
	ClientToken    string                                 `position:"Body" name:"ClientToken"`
	DryRun         requests.Boolean                       `position:"Body" name:"DryRun"`
	ZoneMappings   *[]UpdateLoadBalancerZonesZoneMappings `position:"Body" name:"ZoneMappings"  type:"Repeated"`
	LoadBalancerId string                                 `position:"Body" name:"LoadBalancerId"`
}

// UpdateLoadBalancerZonesZoneMappings is a repeated param struct in UpdateLoadBalancerZonesRequest
type UpdateLoadBalancerZonesZoneMappings struct {
	VSwitchId          string `name:"VSwitchId"`
	ZoneId             string `name:"ZoneId"`
	PrivateIPv4Address string `name:"PrivateIPv4Address"`
	AllocationId       string `name:"AllocationId"`
	EipType            string `name:"EipType"`
}

// UpdateLoadBalancerZonesResponse is the response struct for api UpdateLoadBalancerZones
type UpdateLoadBalancerZonesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	JobId          string `json:"JobId" xml:"JobId"`
}

// CreateUpdateLoadBalancerZonesRequest creates a request to invoke UpdateLoadBalancerZones API
func CreateUpdateLoadBalancerZonesRequest() (request *UpdateLoadBalancerZonesRequest) {
	request = &UpdateLoadBalancerZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "UpdateLoadBalancerZones", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLoadBalancerZonesResponse creates a response to parse from UpdateLoadBalancerZones response
func CreateUpdateLoadBalancerZonesResponse() (response *UpdateLoadBalancerZonesResponse) {
	response = &UpdateLoadBalancerZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
