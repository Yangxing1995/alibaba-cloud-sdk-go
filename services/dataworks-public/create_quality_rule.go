package dataworks_public

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

// CreateQualityRule invokes the dataworks_public.CreateQualityRule API synchronously
func (client *Client) CreateQualityRule(request *CreateQualityRuleRequest) (response *CreateQualityRuleResponse, err error) {
	response = CreateCreateQualityRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQualityRuleWithChan invokes the dataworks_public.CreateQualityRule API asynchronously
func (client *Client) CreateQualityRuleWithChan(request *CreateQualityRuleRequest) (<-chan *CreateQualityRuleResponse, <-chan error) {
	responseChan := make(chan *CreateQualityRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQualityRule(request)
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

// CreateQualityRuleWithCallback invokes the dataworks_public.CreateQualityRule API asynchronously
func (client *Client) CreateQualityRuleWithCallback(request *CreateQualityRuleRequest, callback func(response *CreateQualityRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQualityRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateQualityRule(request)
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

// CreateQualityRuleRequest is the request struct for api CreateQualityRule
type CreateQualityRuleRequest struct {
	*requests.RpcRequest
	TaskSetting       string           `position:"Body" name:"TaskSetting"`
	Trend             string           `position:"Body" name:"Trend"`
	BlockType         requests.Integer `position:"Body" name:"BlockType"`
	PropertyType      string           `position:"Body" name:"PropertyType"`
	EntityId          requests.Integer `position:"Body" name:"EntityId"`
	RuleName          string           `position:"Body" name:"RuleName"`
	Checker           requests.Integer `position:"Body" name:"Checker"`
	Operator          string           `position:"Body" name:"Operator"`
	Property          string           `position:"Body" name:"Property"`
	WarningThreshold  string           `position:"Body" name:"WarningThreshold"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	MethodName        string           `position:"Body" name:"MethodName"`
	ProjectName       string           `position:"Body" name:"ProjectName"`
	RuleType          requests.Integer `position:"Body" name:"RuleType"`
	TemplateId        requests.Integer `position:"Body" name:"TemplateId"`
	ExpectValue       string           `position:"Body" name:"ExpectValue"`
	WhereCondition    string           `position:"Body" name:"WhereCondition"`
	CriticalThreshold string           `position:"Body" name:"CriticalThreshold"`
	Comment           string           `position:"Body" name:"Comment"`
	PredictType       requests.Integer `position:"Body" name:"PredictType"`
}

// CreateQualityRuleResponse is the response struct for api CreateQualityRule
type CreateQualityRuleResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateCreateQualityRuleRequest creates a request to invoke CreateQualityRule API
func CreateCreateQualityRuleRequest() (request *CreateQualityRuleRequest) {
	request = &CreateQualityRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateQualityRule", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateQualityRuleResponse creates a response to parse from CreateQualityRule response
func CreateCreateQualityRuleResponse() (response *CreateQualityRuleResponse) {
	response = &CreateQualityRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
