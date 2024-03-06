package r_kvstore

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

// SwitchInstanceHA invokes the r_kvstore.SwitchInstanceHA API synchronously
func (client *Client) SwitchInstanceHA(request *SwitchInstanceHARequest) (response *SwitchInstanceHAResponse, err error) {
	response = CreateSwitchInstanceHAResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchInstanceHAWithChan invokes the r_kvstore.SwitchInstanceHA API asynchronously
func (client *Client) SwitchInstanceHAWithChan(request *SwitchInstanceHARequest) (<-chan *SwitchInstanceHAResponse, <-chan error) {
	responseChan := make(chan *SwitchInstanceHAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchInstanceHA(request)
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

// SwitchInstanceHAWithCallback invokes the r_kvstore.SwitchInstanceHA API asynchronously
func (client *Client) SwitchInstanceHAWithCallback(request *SwitchInstanceHARequest, callback func(response *SwitchInstanceHAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchInstanceHAResponse
		var err error
		defer close(result)
		response, err = client.SwitchInstanceHA(request)
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

// SwitchInstanceHARequest is the request struct for api SwitchInstanceHA
type SwitchInstanceHARequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SwitchMode           requests.Integer `position:"Query" name:"SwitchMode"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	NodeId               string           `position:"Query" name:"NodeId"`
	Product              string           `position:"Query" name:"Product"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SwitchType           string           `position:"Query" name:"SwitchType"`
	Category             string           `position:"Query" name:"Category"`
	SwitchZone           requests.Integer `position:"Query" name:"SwitchZone"`
}

// SwitchInstanceHAResponse is the response struct for api SwitchInstanceHA
type SwitchInstanceHAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchInstanceHARequest creates a request to invoke SwitchInstanceHA API
func CreateSwitchInstanceHARequest() (request *SwitchInstanceHARequest) {
	request = &SwitchInstanceHARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchInstanceHA", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchInstanceHAResponse creates a response to parse from SwitchInstanceHA response
func CreateSwitchInstanceHAResponse() (response *SwitchInstanceHAResponse) {
	response = &SwitchInstanceHAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
