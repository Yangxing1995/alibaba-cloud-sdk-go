package ccc

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

// AddPhoneTags invokes the ccc.AddPhoneTags API synchronously
func (client *Client) AddPhoneTags(request *AddPhoneTagsRequest) (response *AddPhoneTagsResponse, err error) {
	response = CreateAddPhoneTagsResponse()
	err = client.DoAction(request, response)
	return
}

// AddPhoneTagsWithChan invokes the ccc.AddPhoneTags API asynchronously
func (client *Client) AddPhoneTagsWithChan(request *AddPhoneTagsRequest) (<-chan *AddPhoneTagsResponse, <-chan error) {
	responseChan := make(chan *AddPhoneTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPhoneTags(request)
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

// AddPhoneTagsWithCallback invokes the ccc.AddPhoneTags API asynchronously
func (client *Client) AddPhoneTagsWithCallback(request *AddPhoneTagsRequest, callback func(response *AddPhoneTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPhoneTagsResponse
		var err error
		defer close(result)
		response, err = client.AddPhoneTags(request)
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

// AddPhoneTagsRequest is the request struct for api AddPhoneTags
type AddPhoneTagsRequest struct {
	*requests.RpcRequest
	RegionNameProvince string           `position:"Query" name:"RegionNameProvince"`
	Type               requests.Integer `position:"Query" name:"Type"`
	Concurrency        requests.Integer `position:"Query" name:"Concurrency"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
	Provider           string           `position:"Query" name:"Provider"`
	PhoneNumberList    *[]string        `position:"Query" name:"PhoneNumberList"  type:"Repeated"`
	ServiceTag         string           `position:"Query" name:"ServiceTag"`
	SipTag             string           `position:"Query" name:"SipTag"`
	RegionNameCity     string           `position:"Query" name:"RegionNameCity"`
}

// AddPhoneTagsResponse is the response struct for api AddPhoneTags
type AddPhoneTagsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateAddPhoneTagsRequest creates a request to invoke AddPhoneTags API
func CreateAddPhoneTagsRequest() (request *AddPhoneTagsRequest) {
	request = &AddPhoneTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "AddPhoneTags", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddPhoneTagsResponse creates a response to parse from AddPhoneTags response
func CreateAddPhoneTagsResponse() (response *AddPhoneTagsResponse) {
	response = &AddPhoneTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
