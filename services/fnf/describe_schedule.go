package fnf

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

// DescribeSchedule invokes the fnf.DescribeSchedule API synchronously
func (client *Client) DescribeSchedule(request *DescribeScheduleRequest) (response *DescribeScheduleResponse, err error) {
	response = CreateDescribeScheduleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScheduleWithChan invokes the fnf.DescribeSchedule API asynchronously
func (client *Client) DescribeScheduleWithChan(request *DescribeScheduleRequest) (<-chan *DescribeScheduleResponse, <-chan error) {
	responseChan := make(chan *DescribeScheduleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSchedule(request)
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

// DescribeScheduleWithCallback invokes the fnf.DescribeSchedule API asynchronously
func (client *Client) DescribeScheduleWithCallback(request *DescribeScheduleRequest, callback func(response *DescribeScheduleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScheduleResponse
		var err error
		defer close(result)
		response, err = client.DescribeSchedule(request)
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

// DescribeScheduleRequest is the request struct for api DescribeSchedule
type DescribeScheduleRequest struct {
	*requests.RpcRequest
	ScheduleName string `position:"Query" name:"ScheduleName"`
	RequestId    string `position:"Query" name:"RequestId"`
	FlowName     string `position:"Query" name:"FlowName"`
}

// DescribeScheduleResponse is the response struct for api DescribeSchedule
type DescribeScheduleResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Description      string `json:"Description" xml:"Description"`
	ScheduleId       string `json:"ScheduleId" xml:"ScheduleId"`
	Payload          string `json:"Payload" xml:"Payload"`
	ScheduleName     string `json:"ScheduleName" xml:"ScheduleName"`
	CreatedTime      string `json:"CreatedTime" xml:"CreatedTime"`
	LastModifiedTime string `json:"LastModifiedTime" xml:"LastModifiedTime"`
	CronExpression   string `json:"CronExpression" xml:"CronExpression"`
	Enable           bool   `json:"Enable" xml:"Enable"`
}

// CreateDescribeScheduleRequest creates a request to invoke DescribeSchedule API
func CreateDescribeScheduleRequest() (request *DescribeScheduleRequest) {
	request = &DescribeScheduleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "DescribeSchedule", "fnf", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeScheduleResponse creates a response to parse from DescribeSchedule response
func CreateDescribeScheduleResponse() (response *DescribeScheduleResponse) {
	response = &DescribeScheduleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
