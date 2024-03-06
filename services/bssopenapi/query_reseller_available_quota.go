package bssopenapi

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

// QueryResellerAvailableQuota invokes the bssopenapi.QueryResellerAvailableQuota API synchronously
func (client *Client) QueryResellerAvailableQuota(request *QueryResellerAvailableQuotaRequest) (response *QueryResellerAvailableQuotaResponse, err error) {
	response = CreateQueryResellerAvailableQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// QueryResellerAvailableQuotaWithChan invokes the bssopenapi.QueryResellerAvailableQuota API asynchronously
func (client *Client) QueryResellerAvailableQuotaWithChan(request *QueryResellerAvailableQuotaRequest) (<-chan *QueryResellerAvailableQuotaResponse, <-chan error) {
	responseChan := make(chan *QueryResellerAvailableQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryResellerAvailableQuota(request)
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

// QueryResellerAvailableQuotaWithCallback invokes the bssopenapi.QueryResellerAvailableQuota API asynchronously
func (client *Client) QueryResellerAvailableQuotaWithCallback(request *QueryResellerAvailableQuotaRequest, callback func(response *QueryResellerAvailableQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryResellerAvailableQuotaResponse
		var err error
		defer close(result)
		response, err = client.QueryResellerAvailableQuota(request)
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

// QueryResellerAvailableQuotaRequest is the request struct for api QueryResellerAvailableQuota
type QueryResellerAvailableQuotaRequest struct {
	*requests.RpcRequest
	ItemCodes string           `position:"Query" name:"ItemCodes"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryResellerAvailableQuotaResponse is the response struct for api QueryResellerAvailableQuota
type QueryResellerAvailableQuotaResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryResellerAvailableQuotaRequest creates a request to invoke QueryResellerAvailableQuota API
func CreateQueryResellerAvailableQuotaRequest() (request *QueryResellerAvailableQuotaRequest) {
	request = &QueryResellerAvailableQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryResellerAvailableQuota", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryResellerAvailableQuotaResponse creates a response to parse from QueryResellerAvailableQuota response
func CreateQueryResellerAvailableQuotaResponse() (response *QueryResellerAvailableQuotaResponse) {
	response = &QueryResellerAvailableQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
