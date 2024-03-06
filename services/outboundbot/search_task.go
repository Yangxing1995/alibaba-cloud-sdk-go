package outboundbot

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

// SearchTask invokes the outboundbot.SearchTask API synchronously
func (client *Client) SearchTask(request *SearchTaskRequest) (response *SearchTaskResponse, err error) {
	response = CreateSearchTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SearchTaskWithChan invokes the outboundbot.SearchTask API asynchronously
func (client *Client) SearchTaskWithChan(request *SearchTaskRequest) (<-chan *SearchTaskResponse, <-chan error) {
	responseChan := make(chan *SearchTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchTask(request)
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

// SearchTaskWithCallback invokes the outboundbot.SearchTask API asynchronously
func (client *Client) SearchTaskWithCallback(request *SearchTaskRequest, callback func(response *SearchTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchTaskResponse
		var err error
		defer close(result)
		response, err = client.SearchTask(request)
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

// SearchTaskRequest is the request struct for api SearchTask
type SearchTaskRequest struct {
	*requests.RpcRequest
	ActualTimeLte        requests.Integer `position:"Query" name:"ActualTimeLte"`
	OtherId              string           `position:"Query" name:"OtherId"`
	TaskCreateTimeLte    requests.Integer `position:"Query" name:"TaskCreateTimeLte"`
	JobId                string           `position:"Query" name:"JobId"`
	TaskCreateTimeGte    requests.Integer `position:"Query" name:"TaskCreateTimeGte"`
	CalledNumber         string           `position:"Query" name:"CalledNumber"`
	UserIdMatch          string           `position:"Query" name:"UserIdMatch"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ScriptNameQuery      string           `position:"Query" name:"ScriptNameQuery"`
	PageIndex            requests.Integer `position:"Query" name:"PageIndex"`
	SortOrder            string           `position:"Query" name:"SortOrder"`
	TaskStatusStringList string           `position:"Query" name:"TaskStatusStringList"`
	JobGroupNameQuery    string           `position:"Query" name:"JobGroupNameQuery"`
	TaskId               string           `position:"Query" name:"TaskId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	RecordingDurationGte requests.Integer `position:"Query" name:"RecordingDurationGte"`
	CallDurationLte      requests.Integer `position:"Query" name:"CallDurationLte"`
	JobGroupId           string           `position:"Query" name:"JobGroupId"`
	SortBy               string           `position:"Query" name:"SortBy"`
	JobStatusStringList  string           `position:"Query" name:"JobStatusStringList"`
	ActualTimeGte        requests.Integer `position:"Query" name:"ActualTimeGte"`
	CallDurationGte      requests.Integer `position:"Query" name:"CallDurationGte"`
	RecordingDurationLte requests.Integer `position:"Query" name:"RecordingDurationLte"`
}

// SearchTaskResponse is the response struct for api SearchTask
type SearchTaskResponse struct {
	*responses.BaseResponse
	HttpStatusCode     int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageIndex          int              `json:"PageIndex" xml:"PageIndex"`
	RequestId          string           `json:"RequestId" xml:"RequestId"`
	Success            bool             `json:"Success" xml:"Success"`
	Code               string           `json:"Code" xml:"Code"`
	Message            string           `json:"Message" xml:"Message"`
	PageSize           int              `json:"PageSize" xml:"PageSize"`
	Total              int64            `json:"Total" xml:"Total"`
	SearchTaskInfoList []SearchTaskInfo `json:"SearchTaskInfoList" xml:"SearchTaskInfoList"`
}

// CreateSearchTaskRequest creates a request to invoke SearchTask API
func CreateSearchTaskRequest() (request *SearchTaskRequest) {
	request = &SearchTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SearchTask", "", "")
	request.Method = requests.GET
	return
}

// CreateSearchTaskResponse creates a response to parse from SearchTask response
func CreateSearchTaskResponse() (response *SearchTaskResponse) {
	response = &SearchTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
