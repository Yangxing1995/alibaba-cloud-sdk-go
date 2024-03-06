package dypnsapi_intl

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

// CheckVerification invokes the dypnsapi_intl.CheckVerification API synchronously
func (client *Client) CheckVerification(request *CheckVerificationRequest) (response *CheckVerificationResponse, err error) {
	response = CreateCheckVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// CheckVerificationWithChan invokes the dypnsapi_intl.CheckVerification API asynchronously
func (client *Client) CheckVerificationWithChan(request *CheckVerificationRequest) (<-chan *CheckVerificationResponse, <-chan error) {
	responseChan := make(chan *CheckVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckVerification(request)
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

// CheckVerificationWithCallback invokes the dypnsapi_intl.CheckVerification API asynchronously
func (client *Client) CheckVerificationWithCallback(request *CheckVerificationRequest, callback func(response *CheckVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckVerificationResponse
		var err error
		defer close(result)
		response, err = client.CheckVerification(request)
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

// CheckVerificationRequest is the request struct for api CheckVerification
type CheckVerificationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Code                 string           `position:"Query" name:"Code"`
	ServiceSid           string           `position:"Query" name:"ServiceSid"`
	RouteName            string           `position:"Query" name:"RouteName"`
	VerificationId       string           `position:"Query" name:"VerificationId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	To                   string           `position:"Query" name:"To"`
}

// CheckVerificationResponse is the response struct for api CheckVerification
type CheckVerificationResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Message   string                 `json:"Message" xml:"Message"`
	Model     map[string]interface{} `json:"Model" xml:"Model"`
	Code      string                 `json:"Code" xml:"Code"`
	Success   string                 `json:"Success" xml:"Success"`
}

// CreateCheckVerificationRequest creates a request to invoke CheckVerification API
func CreateCheckVerificationRequest() (request *CheckVerificationRequest) {
	request = &CheckVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi-intl", "2017-07-25", "CheckVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckVerificationResponse creates a response to parse from CheckVerification response
func CreateCheckVerificationResponse() (response *CheckVerificationResponse) {
	response = &CheckVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
