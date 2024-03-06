package oceanbasepro

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

// ModifyInstanceTags invokes the oceanbasepro.ModifyInstanceTags API synchronously
func (client *Client) ModifyInstanceTags(request *ModifyInstanceTagsRequest) (response *ModifyInstanceTagsResponse, err error) {
	response = CreateModifyInstanceTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceTagsWithChan invokes the oceanbasepro.ModifyInstanceTags API asynchronously
func (client *Client) ModifyInstanceTagsWithChan(request *ModifyInstanceTagsRequest) (<-chan *ModifyInstanceTagsResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceTags(request)
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

// ModifyInstanceTagsWithCallback invokes the oceanbasepro.ModifyInstanceTags API asynchronously
func (client *Client) ModifyInstanceTagsWithCallback(request *ModifyInstanceTagsRequest, callback func(response *ModifyInstanceTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceTagsResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceTags(request)
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

// ModifyInstanceTagsRequest is the request struct for api ModifyInstanceTags
type ModifyInstanceTagsRequest struct {
	*requests.RpcRequest
	Tags       string `position:"Body" name:"Tags"`
	InstanceId string `position:"Body" name:"InstanceId"`
}

// ModifyInstanceTagsResponse is the response struct for api ModifyInstanceTags
type ModifyInstanceTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateModifyInstanceTagsRequest creates a request to invoke ModifyInstanceTags API
func CreateModifyInstanceTagsRequest() (request *ModifyInstanceTagsRequest) {
	request = &ModifyInstanceTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ModifyInstanceTags", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceTagsResponse creates a response to parse from ModifyInstanceTags response
func CreateModifyInstanceTagsResponse() (response *ModifyInstanceTagsResponse) {
	response = &ModifyInstanceTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
