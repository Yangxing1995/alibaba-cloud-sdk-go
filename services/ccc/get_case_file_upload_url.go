package ccc

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

// GetCaseFileUploadUrl invokes the ccc.GetCaseFileUploadUrl API synchronously
func (client *Client) GetCaseFileUploadUrl(request *GetCaseFileUploadUrlRequest) (response *GetCaseFileUploadUrlResponse, err error) {
	response = CreateGetCaseFileUploadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetCaseFileUploadUrlWithChan invokes the ccc.GetCaseFileUploadUrl API asynchronously
func (client *Client) GetCaseFileUploadUrlWithChan(request *GetCaseFileUploadUrlRequest) (<-chan *GetCaseFileUploadUrlResponse, <-chan error) {
	responseChan := make(chan *GetCaseFileUploadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCaseFileUploadUrl(request)
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

// GetCaseFileUploadUrlWithCallback invokes the ccc.GetCaseFileUploadUrl API asynchronously
func (client *Client) GetCaseFileUploadUrlWithCallback(request *GetCaseFileUploadUrlRequest, callback func(response *GetCaseFileUploadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCaseFileUploadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetCaseFileUploadUrl(request)
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

// GetCaseFileUploadUrlRequest is the request struct for api GetCaseFileUploadUrl
type GetCaseFileUploadUrlRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
}

// GetCaseFileUploadUrlResponse is the response struct for api GetCaseFileUploadUrl
type GetCaseFileUploadUrlResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetCaseFileUploadUrlRequest creates a request to invoke GetCaseFileUploadUrl API
func CreateGetCaseFileUploadUrlRequest() (request *GetCaseFileUploadUrlRequest) {
	request = &GetCaseFileUploadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetCaseFileUploadUrl", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCaseFileUploadUrlResponse creates a response to parse from GetCaseFileUploadUrl response
func CreateGetCaseFileUploadUrlResponse() (response *GetCaseFileUploadUrlResponse) {
	response = &GetCaseFileUploadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
