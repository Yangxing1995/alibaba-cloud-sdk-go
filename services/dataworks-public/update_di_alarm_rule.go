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

// UpdateDIAlarmRule invokes the dataworks_public.UpdateDIAlarmRule API synchronously
func (client *Client) UpdateDIAlarmRule(request *UpdateDIAlarmRuleRequest) (response *UpdateDIAlarmRuleResponse, err error) {
	response = CreateUpdateDIAlarmRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDIAlarmRuleWithChan invokes the dataworks_public.UpdateDIAlarmRule API asynchronously
func (client *Client) UpdateDIAlarmRuleWithChan(request *UpdateDIAlarmRuleRequest) (<-chan *UpdateDIAlarmRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateDIAlarmRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDIAlarmRule(request)
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

// UpdateDIAlarmRuleWithCallback invokes the dataworks_public.UpdateDIAlarmRule API asynchronously
func (client *Client) UpdateDIAlarmRuleWithCallback(request *UpdateDIAlarmRuleRequest, callback func(response *UpdateDIAlarmRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDIAlarmRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateDIAlarmRule(request)
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

// UpdateDIAlarmRuleRequest is the request struct for api UpdateDIAlarmRule
type UpdateDIAlarmRuleRequest struct {
	*requests.RpcRequest
	MetricType           string           `position:"Body" name:"MetricType"`
	TriggerConditions    string           `position:"Body" name:"TriggerConditions"`
	Description          string           `position:"Body" name:"Description"`
	NotificationSettings string           `position:"Body" name:"NotificationSettings"`
	Enabled              requests.Boolean `position:"Body" name:"Enabled"`
	DIAlarmRuleId        requests.Integer `position:"Body" name:"DIAlarmRuleId"`
}

// UpdateDIAlarmRuleResponse is the response struct for api UpdateDIAlarmRule
type UpdateDIAlarmRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDIAlarmRuleRequest creates a request to invoke UpdateDIAlarmRule API
func CreateUpdateDIAlarmRuleRequest() (request *UpdateDIAlarmRuleRequest) {
	request = &UpdateDIAlarmRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateDIAlarmRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDIAlarmRuleResponse creates a response to parse from UpdateDIAlarmRule response
func CreateUpdateDIAlarmRuleResponse() (response *UpdateDIAlarmRuleResponse) {
	response = &UpdateDIAlarmRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
