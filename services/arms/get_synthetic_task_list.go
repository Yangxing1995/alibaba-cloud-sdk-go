package arms

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

// GetSyntheticTaskList invokes the arms.GetSyntheticTaskList API synchronously
func (client *Client) GetSyntheticTaskList(request *GetSyntheticTaskListRequest) (response *GetSyntheticTaskListResponse, err error) {
	response = CreateGetSyntheticTaskListResponse()
	err = client.DoAction(request, response)
	return
}

// GetSyntheticTaskListWithChan invokes the arms.GetSyntheticTaskList API asynchronously
func (client *Client) GetSyntheticTaskListWithChan(request *GetSyntheticTaskListRequest) (<-chan *GetSyntheticTaskListResponse, <-chan error) {
	responseChan := make(chan *GetSyntheticTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSyntheticTaskList(request)
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

// GetSyntheticTaskListWithCallback invokes the arms.GetSyntheticTaskList API asynchronously
func (client *Client) GetSyntheticTaskListWithCallback(request *GetSyntheticTaskListRequest, callback func(response *GetSyntheticTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSyntheticTaskListResponse
		var err error
		defer close(result)
		response, err = client.GetSyntheticTaskList(request)
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

// GetSyntheticTaskListRequest is the request struct for api GetSyntheticTaskList
type GetSyntheticTaskListRequest struct {
	*requests.RpcRequest
	TaskType   string           `position:"Query" name:"TaskType"`
	TaskName   string           `position:"Query" name:"TaskName"`
	PageNum    requests.Integer `position:"Query" name:"PageNum"`
	Url        string           `position:"Query" name:"Url"`
	TaskStatus string           `position:"Query" name:"TaskStatus"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Order      string           `position:"Query" name:"Order"`
	Direction  string           `position:"Query" name:"Direction"`
}

// GetSyntheticTaskListResponse is the response struct for api GetSyntheticTaskList
type GetSyntheticTaskListResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo `json:"PageInfo" xml:"PageInfo"`
}

// CreateGetSyntheticTaskListRequest creates a request to invoke GetSyntheticTaskList API
func CreateGetSyntheticTaskListRequest() (request *GetSyntheticTaskListRequest) {
	request = &GetSyntheticTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetSyntheticTaskList", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSyntheticTaskListResponse creates a response to parse from GetSyntheticTaskList response
func CreateGetSyntheticTaskListResponse() (response *GetSyntheticTaskListResponse) {
	response = &GetSyntheticTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
