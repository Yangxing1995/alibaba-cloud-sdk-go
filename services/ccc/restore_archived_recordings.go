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

// RestoreArchivedRecordings invokes the ccc.RestoreArchivedRecordings API synchronously
func (client *Client) RestoreArchivedRecordings(request *RestoreArchivedRecordingsRequest) (response *RestoreArchivedRecordingsResponse, err error) {
	response = CreateRestoreArchivedRecordingsResponse()
	err = client.DoAction(request, response)
	return
}

// RestoreArchivedRecordingsWithChan invokes the ccc.RestoreArchivedRecordings API asynchronously
func (client *Client) RestoreArchivedRecordingsWithChan(request *RestoreArchivedRecordingsRequest) (<-chan *RestoreArchivedRecordingsResponse, <-chan error) {
	responseChan := make(chan *RestoreArchivedRecordingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestoreArchivedRecordings(request)
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

// RestoreArchivedRecordingsWithCallback invokes the ccc.RestoreArchivedRecordings API asynchronously
func (client *Client) RestoreArchivedRecordingsWithCallback(request *RestoreArchivedRecordingsRequest, callback func(response *RestoreArchivedRecordingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestoreArchivedRecordingsResponse
		var err error
		defer close(result)
		response, err = client.RestoreArchivedRecordings(request)
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

// RestoreArchivedRecordingsRequest is the request struct for api RestoreArchivedRecordings
type RestoreArchivedRecordingsRequest struct {
	*requests.RpcRequest
	ContactIds string `position:"Query" name:"ContactIds"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// RestoreArchivedRecordingsResponse is the response struct for api RestoreArchivedRecordings
type RestoreArchivedRecordingsResponse struct {
	*responses.BaseResponse
	Code           string                   `json:"Code" xml:"Code"`
	HttpStatusCode int                      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                   `json:"Message" xml:"Message"`
	RequestId      string                   `json:"RequestId" xml:"RequestId"`
	Data           []RecordingRestoreDetail `json:"Data" xml:"Data"`
}

// CreateRestoreArchivedRecordingsRequest creates a request to invoke RestoreArchivedRecordings API
func CreateRestoreArchivedRecordingsRequest() (request *RestoreArchivedRecordingsRequest) {
	request = &RestoreArchivedRecordingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "RestoreArchivedRecordings", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestoreArchivedRecordingsResponse creates a response to parse from RestoreArchivedRecordings response
func CreateRestoreArchivedRecordingsResponse() (response *RestoreArchivedRecordingsResponse) {
	response = &RestoreArchivedRecordingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
