package dms_enterprise

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

// RevokeTemplateAuthority invokes the dms_enterprise.RevokeTemplateAuthority API synchronously
func (client *Client) RevokeTemplateAuthority(request *RevokeTemplateAuthorityRequest) (response *RevokeTemplateAuthorityResponse, err error) {
	response = CreateRevokeTemplateAuthorityResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeTemplateAuthorityWithChan invokes the dms_enterprise.RevokeTemplateAuthority API asynchronously
func (client *Client) RevokeTemplateAuthorityWithChan(request *RevokeTemplateAuthorityRequest) (<-chan *RevokeTemplateAuthorityResponse, <-chan error) {
	responseChan := make(chan *RevokeTemplateAuthorityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeTemplateAuthority(request)
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

// RevokeTemplateAuthorityWithCallback invokes the dms_enterprise.RevokeTemplateAuthority API asynchronously
func (client *Client) RevokeTemplateAuthorityWithCallback(request *RevokeTemplateAuthorityRequest, callback func(response *RevokeTemplateAuthorityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeTemplateAuthorityResponse
		var err error
		defer close(result)
		response, err = client.RevokeTemplateAuthority(request)
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

// RevokeTemplateAuthorityRequest is the request struct for api RevokeTemplateAuthority
type RevokeTemplateAuthorityRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	TemplateId requests.Integer `position:"Query" name:"TemplateId"`
	UserIds    string           `position:"Query" name:"UserIds"`
}

// RevokeTemplateAuthorityResponse is the response struct for api RevokeTemplateAuthority
type RevokeTemplateAuthorityResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       bool   `json:"Result" xml:"Result"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateRevokeTemplateAuthorityRequest creates a request to invoke RevokeTemplateAuthority API
func CreateRevokeTemplateAuthorityRequest() (request *RevokeTemplateAuthorityRequest) {
	request = &RevokeTemplateAuthorityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "RevokeTemplateAuthority", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRevokeTemplateAuthorityResponse creates a response to parse from RevokeTemplateAuthority response
func CreateRevokeTemplateAuthorityResponse() (response *RevokeTemplateAuthorityResponse) {
	response = &RevokeTemplateAuthorityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
