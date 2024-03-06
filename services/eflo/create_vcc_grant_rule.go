package eflo

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

// CreateVccGrantRule invokes the eflo.CreateVccGrantRule API synchronously
func (client *Client) CreateVccGrantRule(request *CreateVccGrantRuleRequest) (response *CreateVccGrantRuleResponse, err error) {
	response = CreateCreateVccGrantRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVccGrantRuleWithChan invokes the eflo.CreateVccGrantRule API asynchronously
func (client *Client) CreateVccGrantRuleWithChan(request *CreateVccGrantRuleRequest) (<-chan *CreateVccGrantRuleResponse, <-chan error) {
	responseChan := make(chan *CreateVccGrantRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVccGrantRule(request)
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

// CreateVccGrantRuleWithCallback invokes the eflo.CreateVccGrantRule API asynchronously
func (client *Client) CreateVccGrantRuleWithCallback(request *CreateVccGrantRuleRequest, callback func(response *CreateVccGrantRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVccGrantRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateVccGrantRule(request)
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

// CreateVccGrantRuleRequest is the request struct for api CreateVccGrantRule
type CreateVccGrantRuleRequest struct {
	*requests.RpcRequest
	ErId          string `position:"Body" name:"ErId"`
	GrantTenantId string `position:"Body" name:"GrantTenantId"`
	InstanceId    string `position:"Body" name:"InstanceId"`
}

// CreateVccGrantRuleResponse is the response struct for api CreateVccGrantRule
type CreateVccGrantRuleResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateCreateVccGrantRuleRequest creates a request to invoke CreateVccGrantRule API
func CreateCreateVccGrantRuleRequest() (request *CreateVccGrantRuleRequest) {
	request = &CreateVccGrantRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "CreateVccGrantRule", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVccGrantRuleResponse creates a response to parse from CreateVccGrantRule response
func CreateCreateVccGrantRuleResponse() (response *CreateVccGrantRuleResponse) {
	response = &CreateVccGrantRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
