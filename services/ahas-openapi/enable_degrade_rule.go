package ahas_openapi

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

// EnableDegradeRule invokes the ahas_openapi.EnableDegradeRule API synchronously
func (client *Client) EnableDegradeRule(request *EnableDegradeRuleRequest) (response *EnableDegradeRuleResponse, err error) {
	response = CreateEnableDegradeRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableDegradeRuleWithChan invokes the ahas_openapi.EnableDegradeRule API asynchronously
func (client *Client) EnableDegradeRuleWithChan(request *EnableDegradeRuleRequest) (<-chan *EnableDegradeRuleResponse, <-chan error) {
	responseChan := make(chan *EnableDegradeRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableDegradeRule(request)
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

// EnableDegradeRuleWithCallback invokes the ahas_openapi.EnableDegradeRule API asynchronously
func (client *Client) EnableDegradeRuleWithCallback(request *EnableDegradeRuleRequest, callback func(response *EnableDegradeRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableDegradeRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableDegradeRule(request)
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

// EnableDegradeRuleRequest is the request struct for api EnableDegradeRule
type EnableDegradeRuleRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	RuleId       requests.Integer `position:"Query" name:"RuleId"`
}

// EnableDegradeRuleResponse is the response struct for api EnableDegradeRule
type EnableDegradeRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnableDegradeRuleRequest creates a request to invoke EnableDegradeRule API
func CreateEnableDegradeRuleRequest() (request *EnableDegradeRuleRequest) {
	request = &EnableDegradeRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "EnableDegradeRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableDegradeRuleResponse creates a response to parse from EnableDegradeRule response
func CreateEnableDegradeRuleResponse() (response *EnableDegradeRuleResponse) {
	response = &EnableDegradeRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
