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

// CreateWirelessCloudConnector invokes the cc5g.CreateWirelessCloudConnector API synchronously
func (client *Client) CreateWirelessCloudConnector(request *CreateWirelessCloudConnectorRequest) (response *CreateWirelessCloudConnectorResponse, err error) {
	response = CreateCreateWirelessCloudConnectorResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWirelessCloudConnectorWithChan invokes the cc5g.CreateWirelessCloudConnector API asynchronously
func (client *Client) CreateWirelessCloudConnectorWithChan(request *CreateWirelessCloudConnectorRequest) (<-chan *CreateWirelessCloudConnectorResponse, <-chan error) {
	responseChan := make(chan *CreateWirelessCloudConnectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWirelessCloudConnector(request)
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

// CreateWirelessCloudConnectorWithCallback invokes the cc5g.CreateWirelessCloudConnector API asynchronously
func (client *Client) CreateWirelessCloudConnectorWithCallback(request *CreateWirelessCloudConnectorRequest, callback func(response *CreateWirelessCloudConnectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWirelessCloudConnectorResponse
		var err error
		defer close(result)
		response, err = client.CreateWirelessCloudConnector(request)
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

// CreateWirelessCloudConnectorRequest is the request struct for api CreateWirelessCloudConnector
type CreateWirelessCloudConnectorRequest struct {
	*requests.RpcRequest
	UseCase      string                                  `position:"Query" name:"UseCase"`
	ClientToken  string                                  `position:"Query" name:"ClientToken"`
	ISP          string                                  `position:"Query" name:"ISP"`
	Description  string                                  `position:"Query" name:"Description"`
	BusinessType string                                  `position:"Query" name:"BusinessType"`
	NetLinks     *[]CreateWirelessCloudConnectorNetLinks `position:"Query" name:"NetLinks"  type:"Repeated"`
	DryRun       requests.Boolean                        `position:"Query" name:"DryRun"`
	Name         string                                  `position:"Query" name:"Name"`
}

// CreateWirelessCloudConnectorNetLinks is a repeated param struct in CreateWirelessCloudConnectorRequest
type CreateWirelessCloudConnectorNetLinks struct {
	RegionId string    `name:"RegionId"`
	VpcId    string    `name:"VpcId"`
	VSwitchs *[]string `name:"VSwitchs" type:"Repeated"`
	APN      string    `name:"APN"`
}

// CreateWirelessCloudConnectorResponse is the response struct for api CreateWirelessCloudConnector
type CreateWirelessCloudConnectorResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	WirelessCloudConnectorId string `json:"WirelessCloudConnectorId" xml:"WirelessCloudConnectorId"`
}

// CreateCreateWirelessCloudConnectorRequest creates a request to invoke CreateWirelessCloudConnector API
func CreateCreateWirelessCloudConnectorRequest() (request *CreateWirelessCloudConnectorRequest) {
	request = &CreateWirelessCloudConnectorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "CreateWirelessCloudConnector", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateWirelessCloudConnectorResponse creates a response to parse from CreateWirelessCloudConnector response
func CreateCreateWirelessCloudConnectorResponse() (response *CreateWirelessCloudConnectorResponse) {
	response = &CreateWirelessCloudConnectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
