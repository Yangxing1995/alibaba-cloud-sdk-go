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

// QueryZnodeDetail invokes the mse.QueryZnodeDetail API synchronously
func (client *Client) QueryZnodeDetail(request *QueryZnodeDetailRequest) (response *QueryZnodeDetailResponse, err error) {
	response = CreateQueryZnodeDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryZnodeDetailWithChan invokes the mse.QueryZnodeDetail API asynchronously
func (client *Client) QueryZnodeDetailWithChan(request *QueryZnodeDetailRequest) (<-chan *QueryZnodeDetailResponse, <-chan error) {
	responseChan := make(chan *QueryZnodeDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryZnodeDetail(request)
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

// QueryZnodeDetailWithCallback invokes the mse.QueryZnodeDetail API asynchronously
func (client *Client) QueryZnodeDetailWithCallback(request *QueryZnodeDetailRequest, callback func(response *QueryZnodeDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryZnodeDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryZnodeDetail(request)
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

// QueryZnodeDetailRequest is the request struct for api QueryZnodeDetail
type QueryZnodeDetailRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	Path           string `position:"Query" name:"Path"`
	RequestPars    string `position:"Query" name:"RequestPars"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// QueryZnodeDetailResponse is the response struct for api QueryZnodeDetail
type QueryZnodeDetailResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   string `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryZnodeDetailRequest creates a request to invoke QueryZnodeDetail API
func CreateQueryZnodeDetailRequest() (request *QueryZnodeDetailRequest) {
	request = &QueryZnodeDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryZnodeDetail", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryZnodeDetailResponse creates a response to parse from QueryZnodeDetail response
func CreateQueryZnodeDetailResponse() (response *QueryZnodeDetailResponse) {
	response = &QueryZnodeDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
