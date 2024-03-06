package cc5g

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

// CreateWirelessCloudConnectorGroup invokes the cc5g.CreateWirelessCloudConnectorGroup API synchronously
func (client *Client) CreateWirelessCloudConnectorGroup(request *CreateWirelessCloudConnectorGroupRequest) (response *CreateWirelessCloudConnectorGroupResponse, err error) {
	response = CreateCreateWirelessCloudConnectorGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWirelessCloudConnectorGroupWithChan invokes the cc5g.CreateWirelessCloudConnectorGroup API asynchronously
func (client *Client) CreateWirelessCloudConnectorGroupWithChan(request *CreateWirelessCloudConnectorGroupRequest) (<-chan *CreateWirelessCloudConnectorGroupResponse, <-chan error) {
	responseChan := make(chan *CreateWirelessCloudConnectorGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWirelessCloudConnectorGroup(request)
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

// CreateWirelessCloudConnectorGroupWithCallback invokes the cc5g.CreateWirelessCloudConnectorGroup API asynchronously
func (client *Client) CreateWirelessCloudConnectorGroupWithCallback(request *CreateWirelessCloudConnectorGroupRequest, callback func(response *CreateWirelessCloudConnectorGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWirelessCloudConnectorGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateWirelessCloudConnectorGroup(request)
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

// CreateWirelessCloudConnectorGroupRequest is the request struct for api CreateWirelessCloudConnectorGroup
type CreateWirelessCloudConnectorGroupRequest struct {
	*requests.RpcRequest
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	Description string           `position:"Query" name:"Description"`
	Name        string           `position:"Query" name:"Name"`
}

// CreateWirelessCloudConnectorGroupResponse is the response struct for api CreateWirelessCloudConnectorGroup
type CreateWirelessCloudConnectorGroupResponse struct {
	*responses.BaseResponse
	RequestId                     string `json:"RequestId" xml:"RequestId"`
	WirelessCloudConnectorGroupId string `json:"WirelessCloudConnectorGroupId" xml:"WirelessCloudConnectorGroupId"`
}

// CreateCreateWirelessCloudConnectorGroupRequest creates a request to invoke CreateWirelessCloudConnectorGroup API
func CreateCreateWirelessCloudConnectorGroupRequest() (request *CreateWirelessCloudConnectorGroupRequest) {
	request = &CreateWirelessCloudConnectorGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "CreateWirelessCloudConnectorGroup", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateWirelessCloudConnectorGroupResponse creates a response to parse from CreateWirelessCloudConnectorGroup response
func CreateCreateWirelessCloudConnectorGroupResponse() (response *CreateWirelessCloudConnectorGroupResponse) {
	response = &CreateWirelessCloudConnectorGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
