package config

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

// ListRemediationTemplates invokes the config.ListRemediationTemplates API synchronously
func (client *Client) ListRemediationTemplates(request *ListRemediationTemplatesRequest) (response *ListRemediationTemplatesResponse, err error) {
	response = CreateListRemediationTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRemediationTemplatesWithChan invokes the config.ListRemediationTemplates API asynchronously
func (client *Client) ListRemediationTemplatesWithChan(request *ListRemediationTemplatesRequest) (<-chan *ListRemediationTemplatesResponse, <-chan error) {
	responseChan := make(chan *ListRemediationTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRemediationTemplates(request)
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

// ListRemediationTemplatesWithCallback invokes the config.ListRemediationTemplates API asynchronously
func (client *Client) ListRemediationTemplatesWithCallback(request *ListRemediationTemplatesRequest, callback func(response *ListRemediationTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRemediationTemplatesResponse
		var err error
		defer close(result)
		response, err = client.ListRemediationTemplates(request)
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

// ListRemediationTemplatesRequest is the request struct for api ListRemediationTemplates
type ListRemediationTemplatesRequest struct {
	*requests.RpcRequest
	ManagedRuleIdentifier string `position:"Query" name:"ManagedRuleIdentifier"`
	RemediationType       string `position:"Query" name:"RemediationType"`
}

// ListRemediationTemplatesResponse is the response struct for api ListRemediationTemplates
type ListRemediationTemplatesResponse struct {
	*responses.BaseResponse
	RequestId            string                `json:"RequestId" xml:"RequestId"`
	RemediationTemplates []RemediationTemplate `json:"RemediationTemplates" xml:"RemediationTemplates"`
}

// CreateListRemediationTemplatesRequest creates a request to invoke ListRemediationTemplates API
func CreateListRemediationTemplatesRequest() (request *ListRemediationTemplatesRequest) {
	request = &ListRemediationTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "ListRemediationTemplates", "", "")
	request.Method = requests.POST
	return
}

// CreateListRemediationTemplatesResponse creates a response to parse from ListRemediationTemplates response
func CreateListRemediationTemplatesResponse() (response *ListRemediationTemplatesResponse) {
	response = &ListRemediationTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
