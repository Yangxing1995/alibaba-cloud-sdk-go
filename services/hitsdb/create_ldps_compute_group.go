package hitsdb

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

// CreateLdpsComputeGroup invokes the hitsdb.CreateLdpsComputeGroup API synchronously
func (client *Client) CreateLdpsComputeGroup(request *CreateLdpsComputeGroupRequest) (response *CreateLdpsComputeGroupResponse, err error) {
	response = CreateCreateLdpsComputeGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLdpsComputeGroupWithChan invokes the hitsdb.CreateLdpsComputeGroup API asynchronously
func (client *Client) CreateLdpsComputeGroupWithChan(request *CreateLdpsComputeGroupRequest) (<-chan *CreateLdpsComputeGroupResponse, <-chan error) {
	responseChan := make(chan *CreateLdpsComputeGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLdpsComputeGroup(request)
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

// CreateLdpsComputeGroupWithCallback invokes the hitsdb.CreateLdpsComputeGroup API asynchronously
func (client *Client) CreateLdpsComputeGroupWithCallback(request *CreateLdpsComputeGroupRequest, callback func(response *CreateLdpsComputeGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLdpsComputeGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateLdpsComputeGroup(request)
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

// CreateLdpsComputeGroupRequest is the request struct for api CreateLdpsComputeGroup
type CreateLdpsComputeGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Properties           string           `position:"Query" name:"Properties"`
}

// CreateLdpsComputeGroupResponse is the response struct for api CreateLdpsComputeGroup
type CreateLdpsComputeGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLdpsComputeGroupRequest creates a request to invoke CreateLdpsComputeGroup API
func CreateCreateLdpsComputeGroupRequest() (request *CreateLdpsComputeGroupRequest) {
	request = &CreateLdpsComputeGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "CreateLdpsComputeGroup", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLdpsComputeGroupResponse creates a response to parse from CreateLdpsComputeGroup response
func CreateCreateLdpsComputeGroupResponse() (response *CreateLdpsComputeGroupResponse) {
	response = &CreateLdpsComputeGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
