package arms

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

// OpenArmsServiceSecondVersion invokes the arms.OpenArmsServiceSecondVersion API synchronously
func (client *Client) OpenArmsServiceSecondVersion(request *OpenArmsServiceSecondVersionRequest) (response *OpenArmsServiceSecondVersionResponse, err error) {
	response = CreateOpenArmsServiceSecondVersionResponse()
	err = client.DoAction(request, response)
	return
}

// OpenArmsServiceSecondVersionWithChan invokes the arms.OpenArmsServiceSecondVersion API asynchronously
func (client *Client) OpenArmsServiceSecondVersionWithChan(request *OpenArmsServiceSecondVersionRequest) (<-chan *OpenArmsServiceSecondVersionResponse, <-chan error) {
	responseChan := make(chan *OpenArmsServiceSecondVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenArmsServiceSecondVersion(request)
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

// OpenArmsServiceSecondVersionWithCallback invokes the arms.OpenArmsServiceSecondVersion API asynchronously
func (client *Client) OpenArmsServiceSecondVersionWithCallback(request *OpenArmsServiceSecondVersionRequest, callback func(response *OpenArmsServiceSecondVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenArmsServiceSecondVersionResponse
		var err error
		defer close(result)
		response, err = client.OpenArmsServiceSecondVersion(request)
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

// OpenArmsServiceSecondVersionRequest is the request struct for api OpenArmsServiceSecondVersion
type OpenArmsServiceSecondVersionRequest struct {
	*requests.RpcRequest
	Type string `position:"Query" name:"Type"`
}

// OpenArmsServiceSecondVersionResponse is the response struct for api OpenArmsServiceSecondVersion
type OpenArmsServiceSecondVersionResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOpenArmsServiceSecondVersionRequest creates a request to invoke OpenArmsServiceSecondVersion API
func CreateOpenArmsServiceSecondVersionRequest() (request *OpenArmsServiceSecondVersionRequest) {
	request = &OpenArmsServiceSecondVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "OpenArmsServiceSecondVersion", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenArmsServiceSecondVersionResponse creates a response to parse from OpenArmsServiceSecondVersion response
func CreateOpenArmsServiceSecondVersionResponse() (response *OpenArmsServiceSecondVersionResponse) {
	response = &OpenArmsServiceSecondVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
