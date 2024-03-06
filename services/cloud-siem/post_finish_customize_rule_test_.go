package cloud_siem

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

// PostFinishCustomizeRuleTest invokes the cloud_siem.PostFinishCustomizeRuleTest API synchronously
func (client *Client) PostFinishCustomizeRuleTest(request *PostFinishCustomizeRuleTestRequest) (response *PostFinishCustomizeRuleTestResponse, err error) {
	response = CreatePostFinishCustomizeRuleTestResponse()
	err = client.DoAction(request, response)
	return
}

// PostFinishCustomizeRuleTestWithChan invokes the cloud_siem.PostFinishCustomizeRuleTest API asynchronously
func (client *Client) PostFinishCustomizeRuleTestWithChan(request *PostFinishCustomizeRuleTestRequest) (<-chan *PostFinishCustomizeRuleTestResponse, <-chan error) {
	responseChan := make(chan *PostFinishCustomizeRuleTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PostFinishCustomizeRuleTest(request)
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

// PostFinishCustomizeRuleTestWithCallback invokes the cloud_siem.PostFinishCustomizeRuleTest API asynchronously
func (client *Client) PostFinishCustomizeRuleTestWithCallback(request *PostFinishCustomizeRuleTestRequest, callback func(response *PostFinishCustomizeRuleTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PostFinishCustomizeRuleTestResponse
		var err error
		defer close(result)
		response, err = client.PostFinishCustomizeRuleTest(request)
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

// PostFinishCustomizeRuleTestRequest is the request struct for api PostFinishCustomizeRuleTest
type PostFinishCustomizeRuleTestRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// PostFinishCustomizeRuleTestResponse is the response struct for api PostFinishCustomizeRuleTest
type PostFinishCustomizeRuleTestResponse struct {
	*responses.BaseResponse
}

// CreatePostFinishCustomizeRuleTestRequest creates a request to invoke PostFinishCustomizeRuleTest API
func CreatePostFinishCustomizeRuleTestRequest() (request *PostFinishCustomizeRuleTestRequest) {
	request = &PostFinishCustomizeRuleTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "PostFinishCustomizeRuleTest", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePostFinishCustomizeRuleTestResponse creates a response to parse from PostFinishCustomizeRuleTest response
func CreatePostFinishCustomizeRuleTestResponse() (response *PostFinishCustomizeRuleTestResponse) {
	response = &PostFinishCustomizeRuleTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
