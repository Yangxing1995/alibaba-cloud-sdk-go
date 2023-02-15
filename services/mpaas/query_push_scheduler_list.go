package mpaas

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

// QueryPushSchedulerList invokes the mpaas.QueryPushSchedulerList API synchronously
func (client *Client) QueryPushSchedulerList(request *QueryPushSchedulerListRequest) (response *QueryPushSchedulerListResponse, err error) {
	response = CreateQueryPushSchedulerListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPushSchedulerListWithChan invokes the mpaas.QueryPushSchedulerList API asynchronously
func (client *Client) QueryPushSchedulerListWithChan(request *QueryPushSchedulerListRequest) (<-chan *QueryPushSchedulerListResponse, <-chan error) {
	responseChan := make(chan *QueryPushSchedulerListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPushSchedulerList(request)
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

// QueryPushSchedulerListWithCallback invokes the mpaas.QueryPushSchedulerList API asynchronously
func (client *Client) QueryPushSchedulerListWithCallback(request *QueryPushSchedulerListRequest, callback func(response *QueryPushSchedulerListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPushSchedulerListResponse
		var err error
		defer close(result)
		response, err = client.QueryPushSchedulerList(request)
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

// QueryPushSchedulerListRequest is the request struct for api QueryPushSchedulerList
type QueryPushSchedulerListRequest struct {
	*requests.RpcRequest
	StartTime   requests.Integer `position:"Body" name:"StartTime"`
	Type        requests.Integer `position:"Body" name:"Type"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	UniqueId    string           `position:"Body" name:"UniqueId"`
	EndTime     requests.Integer `position:"Body" name:"EndTime"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// QueryPushSchedulerListResponse is the response struct for api QueryPushSchedulerList
type QueryPushSchedulerListResponse struct {
	*responses.BaseResponse
	ResultMessage string        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContent `json:"ResultContent" xml:"ResultContent"`
}

// CreateQueryPushSchedulerListRequest creates a request to invoke QueryPushSchedulerList API
func CreateQueryPushSchedulerListRequest() (request *QueryPushSchedulerListRequest) {
	request = &QueryPushSchedulerListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "QueryPushSchedulerList", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryPushSchedulerListResponse creates a response to parse from QueryPushSchedulerList response
func CreateQueryPushSchedulerListResponse() (response *QueryPushSchedulerListResponse) {
	response = &QueryPushSchedulerListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
