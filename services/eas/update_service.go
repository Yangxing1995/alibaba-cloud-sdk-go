package eas

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

// UpdateService invokes the eas.UpdateService API synchronously
func (client *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
	response = CreateUpdateServiceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServiceWithChan invokes the eas.UpdateService API asynchronously
func (client *Client) UpdateServiceWithChan(request *UpdateServiceRequest) (<-chan *UpdateServiceResponse, <-chan error) {
	responseChan := make(chan *UpdateServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateService(request)
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

// UpdateServiceWithCallback invokes the eas.UpdateService API asynchronously
func (client *Client) UpdateServiceWithCallback(request *UpdateServiceRequest, callback func(response *UpdateServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServiceResponse
		var err error
		defer close(result)
		response, err = client.UpdateService(request)
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

// UpdateServiceRequest is the request struct for api UpdateService
type UpdateServiceRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"ServiceName"`
	ClusterId   string `position:"Path" name:"ClusterId"`
	Body        string `position:"Body" name:"body"`
}

// UpdateServiceResponse is the response struct for api UpdateService
type UpdateServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateServiceRequest creates a request to invoke UpdateService API
func CreateUpdateServiceRequest() (request *UpdateServiceRequest) {
	request = &UpdateServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "UpdateService", "/api/v2/services/[ClusterId]/[ServiceName]", "eas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateServiceResponse creates a response to parse from UpdateService response
func CreateUpdateServiceResponse() (response *UpdateServiceResponse) {
	response = &UpdateServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
