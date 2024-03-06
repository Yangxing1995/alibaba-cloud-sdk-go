package linkvisual

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

// UpdateEventRecordPlan invokes the linkvisual.UpdateEventRecordPlan API synchronously
func (client *Client) UpdateEventRecordPlan(request *UpdateEventRecordPlanRequest) (response *UpdateEventRecordPlanResponse, err error) {
	response = CreateUpdateEventRecordPlanResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEventRecordPlanWithChan invokes the linkvisual.UpdateEventRecordPlan API asynchronously
func (client *Client) UpdateEventRecordPlanWithChan(request *UpdateEventRecordPlanRequest) (<-chan *UpdateEventRecordPlanResponse, <-chan error) {
	responseChan := make(chan *UpdateEventRecordPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEventRecordPlan(request)
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

// UpdateEventRecordPlanWithCallback invokes the linkvisual.UpdateEventRecordPlan API asynchronously
func (client *Client) UpdateEventRecordPlanWithCallback(request *UpdateEventRecordPlanRequest, callback func(response *UpdateEventRecordPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEventRecordPlanResponse
		var err error
		defer close(result)
		response, err = client.UpdateEventRecordPlan(request)
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

// UpdateEventRecordPlanRequest is the request struct for api UpdateEventRecordPlan
type UpdateEventRecordPlanRequest struct {
	*requests.RpcRequest
	EventTypes        string           `position:"Query" name:"EventTypes"`
	PreRecordDuration requests.Integer `position:"Query" name:"PreRecordDuration"`
	RecordDuration    requests.Integer `position:"Query" name:"RecordDuration"`
	TemplateId        string           `position:"Query" name:"TemplateId"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	Name              string           `position:"Query" name:"Name"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	PlanId            string           `position:"Query" name:"PlanId"`
}

// UpdateEventRecordPlanResponse is the response struct for api UpdateEventRecordPlan
type UpdateEventRecordPlanResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUpdateEventRecordPlanRequest creates a request to invoke UpdateEventRecordPlan API
func CreateUpdateEventRecordPlanRequest() (request *UpdateEventRecordPlanRequest) {
	request = &UpdateEventRecordPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "UpdateEventRecordPlan", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateEventRecordPlanResponse creates a response to parse from UpdateEventRecordPlan response
func CreateUpdateEventRecordPlanResponse() (response *UpdateEventRecordPlanResponse) {
	response = &UpdateEventRecordPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
