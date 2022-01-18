package waf_openapi

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

// CreateDomain invokes the waf_openapi.CreateDomain API synchronously
func (client *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
	response = CreateCreateDomainResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDomainWithChan invokes the waf_openapi.CreateDomain API asynchronously
func (client *Client) CreateDomainWithChan(request *CreateDomainRequest) (<-chan *CreateDomainResponse, <-chan error) {
	responseChan := make(chan *CreateDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDomain(request)
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

// CreateDomainWithCallback invokes the waf_openapi.CreateDomain API asynchronously
func (client *Client) CreateDomainWithCallback(request *CreateDomainRequest, callback func(response *CreateDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDomainResponse
		var err error
		defer close(result)
		response, err = client.CreateDomain(request)
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

// CreateDomainRequest is the request struct for api CreateDomain
type CreateDomainRequest struct {
	*requests.RpcRequest
	IpFollowStatus       requests.Integer `position:"Query" name:"IpFollowStatus"`
	SniHost              string           `position:"Query" name:"SniHost"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SourceIp             string           `position:"Query" name:"SourceIp"`
	HttpPort             string           `position:"Query" name:"HttpPort"`
	Http2Port            string           `position:"Query" name:"Http2Port"`
	WriteTime            requests.Integer `position:"Query" name:"WriteTime"`
	SniStatus            requests.Integer `position:"Query" name:"SniStatus"`
	Lang                 string           `position:"Query" name:"Lang"`
	AccessHeaderMode     requests.Integer `position:"Query" name:"AccessHeaderMode"`
	AccessType           string           `position:"Query" name:"AccessType"`
	LogHeaders           string           `position:"Query" name:"LogHeaders"`
	AccessHeaders        string           `position:"Query" name:"AccessHeaders"`
	ConnectionTime       requests.Integer `position:"Query" name:"ConnectionTime"`
	ClusterType          requests.Integer `position:"Query" name:"ClusterType"`
	CloudNativeInstances string           `position:"Query" name:"CloudNativeInstances"`
	HttpsRedirect        requests.Integer `position:"Query" name:"HttpsRedirect"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SourceIps            string           `position:"Query" name:"SourceIps"`
	Domain               string           `position:"Query" name:"Domain"`
	IsAccessProduct      requests.Integer `position:"Query" name:"IsAccessProduct"`
	ReadTime             requests.Integer `position:"Query" name:"ReadTime"`
	HttpsPort            string           `position:"Query" name:"HttpsPort"`
	LoadBalancing        requests.Integer `position:"Query" name:"LoadBalancing"`
	HttpToUserIp         requests.Integer `position:"Query" name:"HttpToUserIp"`
}

// CreateDomainResponse is the response struct for api CreateDomain
type CreateDomainResponse struct {
	*responses.BaseResponse
	Cname     string `json:"Cname" xml:"Cname"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDomainRequest creates a request to invoke CreateDomain API
func CreateCreateDomainRequest() (request *CreateDomainRequest) {
	request = &CreateDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "CreateDomain", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDomainResponse creates a response to parse from CreateDomain response
func CreateCreateDomainResponse() (response *CreateDomainResponse) {
	response = &CreateDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
