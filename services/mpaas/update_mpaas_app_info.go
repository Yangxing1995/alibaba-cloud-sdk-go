package mpaas

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

// UpdateMpaasAppInfo invokes the mpaas.UpdateMpaasAppInfo API synchronously
func (client *Client) UpdateMpaasAppInfo(request *UpdateMpaasAppInfoRequest) (response *UpdateMpaasAppInfoResponse, err error) {
	response = CreateUpdateMpaasAppInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMpaasAppInfoWithChan invokes the mpaas.UpdateMpaasAppInfo API asynchronously
func (client *Client) UpdateMpaasAppInfoWithChan(request *UpdateMpaasAppInfoRequest) (<-chan *UpdateMpaasAppInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateMpaasAppInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMpaasAppInfo(request)
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

// UpdateMpaasAppInfoWithCallback invokes the mpaas.UpdateMpaasAppInfo API asynchronously
func (client *Client) UpdateMpaasAppInfoWithCallback(request *UpdateMpaasAppInfoRequest, callback func(response *UpdateMpaasAppInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMpaasAppInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateMpaasAppInfo(request)
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

// UpdateMpaasAppInfoRequest is the request struct for api UpdateMpaasAppInfo
type UpdateMpaasAppInfoRequest struct {
	*requests.RpcRequest
	SystemType  string           `position:"Body" name:"SystemType"`
	OnexFlag    requests.Boolean `position:"Body" name:"OnexFlag"`
	AppName     string           `position:"Body" name:"AppName"`
	TenantId    string           `position:"Body" name:"TenantId"`
	Identifier  string           `position:"Body" name:"Identifier"`
	IconFileUrl string           `position:"Body" name:"IconFileUrl"`
	AppId       string           `position:"Body" name:"AppId"`
}

// UpdateMpaasAppInfoResponse is the response struct for api UpdateMpaasAppInfo
type UpdateMpaasAppInfoResponse struct {
	*responses.BaseResponse
	ResultMessage string        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContent `json:"ResultContent" xml:"ResultContent"`
}

// CreateUpdateMpaasAppInfoRequest creates a request to invoke UpdateMpaasAppInfo API
func CreateUpdateMpaasAppInfoRequest() (request *UpdateMpaasAppInfoRequest) {
	request = &UpdateMpaasAppInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "UpdateMpaasAppInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateMpaasAppInfoResponse creates a response to parse from UpdateMpaasAppInfo response
func CreateUpdateMpaasAppInfoResponse() (response *UpdateMpaasAppInfoResponse) {
	response = &UpdateMpaasAppInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
