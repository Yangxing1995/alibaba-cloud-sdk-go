package cloud_siem

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

// ListDataSourceLogs invokes the cloud_siem.ListDataSourceLogs API synchronously
func (client *Client) ListDataSourceLogs(request *ListDataSourceLogsRequest) (response *ListDataSourceLogsResponse, err error) {
	response = CreateListDataSourceLogsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataSourceLogsWithChan invokes the cloud_siem.ListDataSourceLogs API asynchronously
func (client *Client) ListDataSourceLogsWithChan(request *ListDataSourceLogsRequest) (<-chan *ListDataSourceLogsResponse, <-chan error) {
	responseChan := make(chan *ListDataSourceLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataSourceLogs(request)
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

// ListDataSourceLogsWithCallback invokes the cloud_siem.ListDataSourceLogs API asynchronously
func (client *Client) ListDataSourceLogsWithCallback(request *ListDataSourceLogsRequest, callback func(response *ListDataSourceLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataSourceLogsResponse
		var err error
		defer close(result)
		response, err = client.ListDataSourceLogs(request)
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

// ListDataSourceLogsRequest is the request struct for api ListDataSourceLogs
type ListDataSourceLogsRequest struct {
	*requests.RpcRequest
	CloudCode            string `position:"Body" name:"CloudCode"`
	AccountId            string `position:"Body" name:"AccountId"`
	DataSourceInstanceId string `position:"Body" name:"DataSourceInstanceId"`
}

// ListDataSourceLogsResponse is the response struct for api ListDataSourceLogs
type ListDataSourceLogsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListDataSourceLogsRequest creates a request to invoke ListDataSourceLogs API
func CreateListDataSourceLogsRequest() (request *ListDataSourceLogsRequest) {
	request = &ListDataSourceLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListDataSourceLogs", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataSourceLogsResponse creates a response to parse from ListDataSourceLogs response
func CreateListDataSourceLogsResponse() (response *ListDataSourceLogsResponse) {
	response = &ListDataSourceLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
