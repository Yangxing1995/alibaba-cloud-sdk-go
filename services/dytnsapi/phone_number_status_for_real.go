package dytnsapi

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

// PhoneNumberStatusForReal invokes the dytnsapi.PhoneNumberStatusForReal API synchronously
func (client *Client) PhoneNumberStatusForReal(request *PhoneNumberStatusForRealRequest) (response *PhoneNumberStatusForRealResponse, err error) {
	response = CreatePhoneNumberStatusForRealResponse()
	err = client.DoAction(request, response)
	return
}

// PhoneNumberStatusForRealWithChan invokes the dytnsapi.PhoneNumberStatusForReal API asynchronously
func (client *Client) PhoneNumberStatusForRealWithChan(request *PhoneNumberStatusForRealRequest) (<-chan *PhoneNumberStatusForRealResponse, <-chan error) {
	responseChan := make(chan *PhoneNumberStatusForRealResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PhoneNumberStatusForReal(request)
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

// PhoneNumberStatusForRealWithCallback invokes the dytnsapi.PhoneNumberStatusForReal API asynchronously
func (client *Client) PhoneNumberStatusForRealWithCallback(request *PhoneNumberStatusForRealRequest, callback func(response *PhoneNumberStatusForRealResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PhoneNumberStatusForRealResponse
		var err error
		defer close(result)
		response, err = client.PhoneNumberStatusForReal(request)
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

// PhoneNumberStatusForRealRequest is the request struct for api PhoneNumberStatusForReal
type PhoneNumberStatusForRealRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	RouteName            string           `position:"Query" name:"RouteName"`
	Mask                 string           `position:"Query" name:"Mask"`
	ResultCount          string           `position:"Query" name:"ResultCount"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AuthCode             string           `position:"Query" name:"AuthCode"`
	InputNumber          string           `position:"Query" name:"InputNumber"`
	FlowName             string           `position:"Query" name:"FlowName"`
	PhoneStatusSceneCode string           `position:"Query" name:"PhoneStatusSceneCode"`
}

// PhoneNumberStatusForRealResponse is the response struct for api PhoneNumberStatusForReal
type PhoneNumberStatusForRealResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePhoneNumberStatusForRealRequest creates a request to invoke PhoneNumberStatusForReal API
func CreatePhoneNumberStatusForRealRequest() (request *PhoneNumberStatusForRealRequest) {
	request = &PhoneNumberStatusForRealRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "PhoneNumberStatusForReal", "", "")
	request.Method = requests.POST
	return
}

// CreatePhoneNumberStatusForRealResponse creates a response to parse from PhoneNumberStatusForReal response
func CreatePhoneNumberStatusForRealResponse() (response *PhoneNumberStatusForRealResponse) {
	response = &PhoneNumberStatusForRealResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
