package ddoscoo

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

// ModifyBlackholeStatus invokes the ddoscoo.ModifyBlackholeStatus API synchronously
func (client *Client) ModifyBlackholeStatus(request *ModifyBlackholeStatusRequest) (response *ModifyBlackholeStatusResponse, err error) {
	response = CreateModifyBlackholeStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBlackholeStatusWithChan invokes the ddoscoo.ModifyBlackholeStatus API asynchronously
func (client *Client) ModifyBlackholeStatusWithChan(request *ModifyBlackholeStatusRequest) (<-chan *ModifyBlackholeStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyBlackholeStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBlackholeStatus(request)
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

// ModifyBlackholeStatusWithCallback invokes the ddoscoo.ModifyBlackholeStatus API asynchronously
func (client *Client) ModifyBlackholeStatusWithCallback(request *ModifyBlackholeStatusRequest, callback func(response *ModifyBlackholeStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBlackholeStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyBlackholeStatus(request)
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

// ModifyBlackholeStatusRequest is the request struct for api ModifyBlackholeStatus
type ModifyBlackholeStatusRequest struct {
	*requests.RpcRequest
	BlackholeStatus string `position:"Query" name:"BlackholeStatus"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

// ModifyBlackholeStatusResponse is the response struct for api ModifyBlackholeStatus
type ModifyBlackholeStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBlackholeStatusRequest creates a request to invoke ModifyBlackholeStatus API
func CreateModifyBlackholeStatusRequest() (request *ModifyBlackholeStatusRequest) {
	request = &ModifyBlackholeStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyBlackholeStatus", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBlackholeStatusResponse creates a response to parse from ModifyBlackholeStatus response
func CreateModifyBlackholeStatusResponse() (response *ModifyBlackholeStatusResponse) {
	response = &ModifyBlackholeStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
