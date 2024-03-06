package ons

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

// OnsTopicCreate invokes the ons.OnsTopicCreate API synchronously
func (client *Client) OnsTopicCreate(request *OnsTopicCreateRequest) (response *OnsTopicCreateResponse, err error) {
	response = CreateOnsTopicCreateResponse()
	err = client.DoAction(request, response)
	return
}

// OnsTopicCreateWithChan invokes the ons.OnsTopicCreate API asynchronously
func (client *Client) OnsTopicCreateWithChan(request *OnsTopicCreateRequest) (<-chan *OnsTopicCreateResponse, <-chan error) {
	responseChan := make(chan *OnsTopicCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTopicCreate(request)
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

// OnsTopicCreateWithCallback invokes the ons.OnsTopicCreate API asynchronously
func (client *Client) OnsTopicCreateWithCallback(request *OnsTopicCreateRequest, callback func(response *OnsTopicCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTopicCreateResponse
		var err error
		defer close(result)
		response, err = client.OnsTopicCreate(request)
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

// OnsTopicCreateRequest is the request struct for api OnsTopicCreate
type OnsTopicCreateRequest struct {
	*requests.RpcRequest
	MessageType requests.Integer `position:"Query" name:"MessageType"`
	Remark      string           `position:"Query" name:"Remark"`
	UserId      string           `position:"Query" name:"UserId"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	Topic       string           `position:"Query" name:"Topic"`
}

// OnsTopicCreateResponse is the response struct for api OnsTopicCreate
type OnsTopicCreateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsTopicCreateRequest creates a request to invoke OnsTopicCreate API
func CreateOnsTopicCreateRequest() (request *OnsTopicCreateRequest) {
	request = &OnsTopicCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsTopicCreate", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsTopicCreateResponse creates a response to parse from OnsTopicCreate response
func CreateOnsTopicCreateResponse() (response *OnsTopicCreateResponse) {
	response = &OnsTopicCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
