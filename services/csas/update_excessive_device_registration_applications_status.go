package csas

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

// UpdateExcessiveDeviceRegistrationApplicationsStatus invokes the csas.UpdateExcessiveDeviceRegistrationApplicationsStatus API synchronously
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatus(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) (response *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, err error) {
	response = CreateUpdateExcessiveDeviceRegistrationApplicationsStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateExcessiveDeviceRegistrationApplicationsStatusWithChan invokes the csas.UpdateExcessiveDeviceRegistrationApplicationsStatus API asynchronously
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatusWithChan(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) (<-chan *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateExcessiveDeviceRegistrationApplicationsStatus(request)
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

// UpdateExcessiveDeviceRegistrationApplicationsStatusWithCallback invokes the csas.UpdateExcessiveDeviceRegistrationApplicationsStatus API asynchronously
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatusWithCallback(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest, callback func(response *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateExcessiveDeviceRegistrationApplicationsStatus(request)
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

// UpdateExcessiveDeviceRegistrationApplicationsStatusRequest is the request struct for api UpdateExcessiveDeviceRegistrationApplicationsStatus
type UpdateExcessiveDeviceRegistrationApplicationsStatusRequest struct {
	*requests.RpcRequest
	ApplicationIds *[]string `position:"Body" name:"ApplicationIds"  type:"Repeated"`
	SourceIp       string    `position:"Query" name:"SourceIp"`
	Status         string    `position:"Body" name:"Status"`
}

// UpdateExcessiveDeviceRegistrationApplicationsStatusResponse is the response struct for api UpdateExcessiveDeviceRegistrationApplicationsStatus
type UpdateExcessiveDeviceRegistrationApplicationsStatusResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Applications []Data `json:"Applications" xml:"Applications"`
}

// CreateUpdateExcessiveDeviceRegistrationApplicationsStatusRequest creates a request to invoke UpdateExcessiveDeviceRegistrationApplicationsStatus API
func CreateUpdateExcessiveDeviceRegistrationApplicationsStatusRequest() (request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) {
	request = &UpdateExcessiveDeviceRegistrationApplicationsStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "UpdateExcessiveDeviceRegistrationApplicationsStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateExcessiveDeviceRegistrationApplicationsStatusResponse creates a response to parse from UpdateExcessiveDeviceRegistrationApplicationsStatus response
func CreateUpdateExcessiveDeviceRegistrationApplicationsStatusResponse() (response *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) {
	response = &UpdateExcessiveDeviceRegistrationApplicationsStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
