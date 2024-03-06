package mse

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

// ListInstanceCount invokes the mse.ListInstanceCount API synchronously
func (client *Client) ListInstanceCount(request *ListInstanceCountRequest) (response *ListInstanceCountResponse, err error) {
	response = CreateListInstanceCountResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstanceCountWithChan invokes the mse.ListInstanceCount API asynchronously
func (client *Client) ListInstanceCountWithChan(request *ListInstanceCountRequest) (<-chan *ListInstanceCountResponse, <-chan error) {
	responseChan := make(chan *ListInstanceCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstanceCount(request)
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

// ListInstanceCountWithCallback invokes the mse.ListInstanceCount API asynchronously
func (client *Client) ListInstanceCountWithCallback(request *ListInstanceCountRequest, callback func(response *ListInstanceCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstanceCountResponse
		var err error
		defer close(result)
		response, err = client.ListInstanceCount(request)
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

// ListInstanceCountRequest is the request struct for api ListInstanceCount
type ListInstanceCountRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	RequestPars    string `position:"Query" name:"RequestPars"`
	ClusterType    string `position:"Query" name:"ClusterType"`
	MseVersion     string `position:"Query" name:"MseVersion"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// ListInstanceCountResponse is the response struct for api ListInstanceCount
type ListInstanceCountResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           []int  `json:"Data" xml:"Data"`
}

// CreateListInstanceCountRequest creates a request to invoke ListInstanceCount API
func CreateListInstanceCountRequest() (request *ListInstanceCountRequest) {
	request = &ListInstanceCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListInstanceCount", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstanceCountResponse creates a response to parse from ListInstanceCount response
func CreateListInstanceCountResponse() (response *ListInstanceCountResponse) {
	response = &ListInstanceCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
