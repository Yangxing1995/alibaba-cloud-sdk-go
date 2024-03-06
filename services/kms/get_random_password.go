package kms

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

// GetRandomPassword invokes the kms.GetRandomPassword API synchronously
func (client *Client) GetRandomPassword(request *GetRandomPasswordRequest) (response *GetRandomPasswordResponse, err error) {
	response = CreateGetRandomPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// GetRandomPasswordWithChan invokes the kms.GetRandomPassword API asynchronously
func (client *Client) GetRandomPasswordWithChan(request *GetRandomPasswordRequest) (<-chan *GetRandomPasswordResponse, <-chan error) {
	responseChan := make(chan *GetRandomPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRandomPassword(request)
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

// GetRandomPasswordWithCallback invokes the kms.GetRandomPassword API asynchronously
func (client *Client) GetRandomPasswordWithCallback(request *GetRandomPasswordRequest, callback func(response *GetRandomPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRandomPasswordResponse
		var err error
		defer close(result)
		response, err = client.GetRandomPassword(request)
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

// GetRandomPasswordRequest is the request struct for api GetRandomPassword
type GetRandomPasswordRequest struct {
	*requests.RpcRequest
	ExcludeCharacters       string `position:"Query" name:"ExcludeCharacters"`
	PasswordLength          string `position:"Query" name:"PasswordLength"`
	ExcludePunctuation      string `position:"Query" name:"ExcludePunctuation"`
	RequireEachIncludedType string `position:"Query" name:"RequireEachIncludedType"`
	ExcludeNumbers          string `position:"Query" name:"ExcludeNumbers"`
	ExcludeLowercase        string `position:"Query" name:"ExcludeLowercase"`
	ExcludeUppercase        string `position:"Query" name:"ExcludeUppercase"`
}

// GetRandomPasswordResponse is the response struct for api GetRandomPassword
type GetRandomPasswordResponse struct {
	*responses.BaseResponse
	RandomPassword string `json:"RandomPassword" xml:"RandomPassword"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateGetRandomPasswordRequest creates a request to invoke GetRandomPassword API
func CreateGetRandomPasswordRequest() (request *GetRandomPasswordRequest) {
	request = &GetRandomPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "GetRandomPassword", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRandomPasswordResponse creates a response to parse from GetRandomPassword response
func CreateGetRandomPasswordResponse() (response *GetRandomPasswordResponse) {
	response = &GetRandomPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
