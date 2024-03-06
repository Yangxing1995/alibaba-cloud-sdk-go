package dbs

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

// StopBackupPlan invokes the dbs.StopBackupPlan API synchronously
func (client *Client) StopBackupPlan(request *StopBackupPlanRequest) (response *StopBackupPlanResponse, err error) {
	response = CreateStopBackupPlanResponse()
	err = client.DoAction(request, response)
	return
}

// StopBackupPlanWithChan invokes the dbs.StopBackupPlan API asynchronously
func (client *Client) StopBackupPlanWithChan(request *StopBackupPlanRequest) (<-chan *StopBackupPlanResponse, <-chan error) {
	responseChan := make(chan *StopBackupPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopBackupPlan(request)
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

// StopBackupPlanWithCallback invokes the dbs.StopBackupPlan API asynchronously
func (client *Client) StopBackupPlanWithCallback(request *StopBackupPlanRequest, callback func(response *StopBackupPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopBackupPlanResponse
		var err error
		defer close(result)
		response, err = client.StopBackupPlan(request)
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

// StopBackupPlanRequest is the request struct for api StopBackupPlan
type StopBackupPlanRequest struct {
	*requests.RpcRequest
	StopMethod   string `position:"Query" name:"StopMethod"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	BackupPlanId string `position:"Query" name:"BackupPlanId"`
	OwnerId      string `position:"Query" name:"OwnerId"`
}

// StopBackupPlanResponse is the response struct for api StopBackupPlan
type StopBackupPlanResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
}

// CreateStopBackupPlanRequest creates a request to invoke StopBackupPlan API
func CreateStopBackupPlanRequest() (request *StopBackupPlanRequest) {
	request = &StopBackupPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "StopBackupPlan", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopBackupPlanResponse creates a response to parse from StopBackupPlan response
func CreateStopBackupPlanResponse() (response *StopBackupPlanResponse) {
	response = &StopBackupPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
