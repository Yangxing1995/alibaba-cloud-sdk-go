package polardb

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

// ModifyDBClusterDeletion invokes the polardb.ModifyDBClusterDeletion API synchronously
func (client *Client) ModifyDBClusterDeletion(request *ModifyDBClusterDeletionRequest) (response *ModifyDBClusterDeletionResponse, err error) {
	response = CreateModifyDBClusterDeletionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterDeletionWithChan invokes the polardb.ModifyDBClusterDeletion API asynchronously
func (client *Client) ModifyDBClusterDeletionWithChan(request *ModifyDBClusterDeletionRequest) (<-chan *ModifyDBClusterDeletionResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterDeletionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterDeletion(request)
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

// ModifyDBClusterDeletionWithCallback invokes the polardb.ModifyDBClusterDeletion API asynchronously
func (client *Client) ModifyDBClusterDeletionWithCallback(request *ModifyDBClusterDeletionRequest, callback func(response *ModifyDBClusterDeletionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterDeletionResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterDeletion(request)
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

// ModifyDBClusterDeletionRequest is the request struct for api ModifyDBClusterDeletion
type ModifyDBClusterDeletionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Protection           requests.Boolean `position:"Query" name:"Protection"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBClusterDeletionResponse is the response struct for api ModifyDBClusterDeletion
type ModifyDBClusterDeletionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterDeletionRequest creates a request to invoke ModifyDBClusterDeletion API
func CreateModifyDBClusterDeletionRequest() (request *ModifyDBClusterDeletionRequest) {
	request = &ModifyDBClusterDeletionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterDeletion", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterDeletionResponse creates a response to parse from ModifyDBClusterDeletion response
func CreateModifyDBClusterDeletionResponse() (response *ModifyDBClusterDeletionResponse) {
	response = &ModifyDBClusterDeletionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
