package domain

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

// CancelDomainVerification invokes the domain.CancelDomainVerification API synchronously
func (client *Client) CancelDomainVerification(request *CancelDomainVerificationRequest) (response *CancelDomainVerificationResponse, err error) {
	response = CreateCancelDomainVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// CancelDomainVerificationWithChan invokes the domain.CancelDomainVerification API asynchronously
func (client *Client) CancelDomainVerificationWithChan(request *CancelDomainVerificationRequest) (<-chan *CancelDomainVerificationResponse, <-chan error) {
	responseChan := make(chan *CancelDomainVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelDomainVerification(request)
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

// CancelDomainVerificationWithCallback invokes the domain.CancelDomainVerification API asynchronously
func (client *Client) CancelDomainVerificationWithCallback(request *CancelDomainVerificationRequest, callback func(response *CancelDomainVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelDomainVerificationResponse
		var err error
		defer close(result)
		response, err = client.CancelDomainVerification(request)
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

// CancelDomainVerificationRequest is the request struct for api CancelDomainVerification
type CancelDomainVerificationRequest struct {
	*requests.RpcRequest
	ActionType   string `position:"Query" name:"ActionType"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// CancelDomainVerificationResponse is the response struct for api CancelDomainVerification
type CancelDomainVerificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelDomainVerificationRequest creates a request to invoke CancelDomainVerification API
func CreateCancelDomainVerificationRequest() (request *CancelDomainVerificationRequest) {
	request = &CancelDomainVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "CancelDomainVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelDomainVerificationResponse creates a response to parse from CancelDomainVerification response
func CreateCancelDomainVerificationResponse() (response *CancelDomainVerificationResponse) {
	response = &CancelDomainVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
