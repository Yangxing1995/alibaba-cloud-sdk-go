package eais

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

// StartEaisEi invokes the eais.StartEaisEi API synchronously
func (client *Client) StartEaisEi(request *StartEaisEiRequest) (response *StartEaisEiResponse, err error) {
	response = CreateStartEaisEiResponse()
	err = client.DoAction(request, response)
	return
}

// StartEaisEiWithChan invokes the eais.StartEaisEi API asynchronously
func (client *Client) StartEaisEiWithChan(request *StartEaisEiRequest) (<-chan *StartEaisEiResponse, <-chan error) {
	responseChan := make(chan *StartEaisEiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartEaisEi(request)
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

// StartEaisEiWithCallback invokes the eais.StartEaisEi API asynchronously
func (client *Client) StartEaisEiWithCallback(request *StartEaisEiRequest, callback func(response *StartEaisEiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartEaisEiResponse
		var err error
		defer close(result)
		response, err = client.StartEaisEi(request)
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

// StartEaisEiRequest is the request struct for api StartEaisEi
type StartEaisEiRequest struct {
	*requests.RpcRequest
	EiInstanceId string `position:"Query" name:"EiInstanceId"`
}

// StartEaisEiResponse is the response struct for api StartEaisEi
type StartEaisEiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartEaisEiRequest creates a request to invoke StartEaisEi API
func CreateStartEaisEiRequest() (request *StartEaisEiRequest) {
	request = &StartEaisEiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "StartEaisEi", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartEaisEiResponse creates a response to parse from StartEaisEi response
func CreateStartEaisEiResponse() (response *StartEaisEiResponse) {
	response = &StartEaisEiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
