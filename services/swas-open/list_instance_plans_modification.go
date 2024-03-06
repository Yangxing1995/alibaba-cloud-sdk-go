package swas_open

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

// ListInstancePlansModification invokes the swas_open.ListInstancePlansModification API synchronously
func (client *Client) ListInstancePlansModification(request *ListInstancePlansModificationRequest) (response *ListInstancePlansModificationResponse, err error) {
	response = CreateListInstancePlansModificationResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstancePlansModificationWithChan invokes the swas_open.ListInstancePlansModification API asynchronously
func (client *Client) ListInstancePlansModificationWithChan(request *ListInstancePlansModificationRequest) (<-chan *ListInstancePlansModificationResponse, <-chan error) {
	responseChan := make(chan *ListInstancePlansModificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstancePlansModification(request)
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

// ListInstancePlansModificationWithCallback invokes the swas_open.ListInstancePlansModification API asynchronously
func (client *Client) ListInstancePlansModificationWithCallback(request *ListInstancePlansModificationRequest, callback func(response *ListInstancePlansModificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstancePlansModificationResponse
		var err error
		defer close(result)
		response, err = client.ListInstancePlansModification(request)
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

// ListInstancePlansModificationRequest is the request struct for api ListInstancePlansModification
type ListInstancePlansModificationRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListInstancePlansModificationResponse is the response struct for api ListInstancePlansModification
type ListInstancePlansModificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Plans     []Plan `json:"Plans" xml:"Plans"`
}

// CreateListInstancePlansModificationRequest creates a request to invoke ListInstancePlansModification API
func CreateListInstancePlansModificationRequest() (request *ListInstancePlansModificationRequest) {
	request = &ListInstancePlansModificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ListInstancePlansModification", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstancePlansModificationResponse creates a response to parse from ListInstancePlansModification response
func CreateListInstancePlansModificationResponse() (response *ListInstancePlansModificationResponse) {
	response = &ListInstancePlansModificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
