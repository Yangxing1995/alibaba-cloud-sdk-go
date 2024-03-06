package ens

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

// JoinSecurityGroup invokes the ens.JoinSecurityGroup API synchronously
func (client *Client) JoinSecurityGroup(request *JoinSecurityGroupRequest) (response *JoinSecurityGroupResponse, err error) {
	response = CreateJoinSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// JoinSecurityGroupWithChan invokes the ens.JoinSecurityGroup API asynchronously
func (client *Client) JoinSecurityGroupWithChan(request *JoinSecurityGroupRequest) (<-chan *JoinSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *JoinSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinSecurityGroup(request)
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

// JoinSecurityGroupWithCallback invokes the ens.JoinSecurityGroup API asynchronously
func (client *Client) JoinSecurityGroupWithCallback(request *JoinSecurityGroupRequest, callback func(response *JoinSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.JoinSecurityGroup(request)
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

// JoinSecurityGroupRequest is the request struct for api JoinSecurityGroup
type JoinSecurityGroupRequest struct {
	*requests.RpcRequest
	SecurityGroupId    string `position:"Query" name:"SecurityGroupId"`
	InstanceId         string `position:"Query" name:"InstanceId"`
	NetworkInterfaceId string `position:"Query" name:"NetworkInterfaceId"`
}

// JoinSecurityGroupResponse is the response struct for api JoinSecurityGroup
type JoinSecurityGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateJoinSecurityGroupRequest creates a request to invoke JoinSecurityGroup API
func CreateJoinSecurityGroupRequest() (request *JoinSecurityGroupRequest) {
	request = &JoinSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "JoinSecurityGroup", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateJoinSecurityGroupResponse creates a response to parse from JoinSecurityGroup response
func CreateJoinSecurityGroupResponse() (response *JoinSecurityGroupResponse) {
	response = &JoinSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
