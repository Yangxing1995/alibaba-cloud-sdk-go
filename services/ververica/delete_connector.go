package ververica

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

// DeleteConnector invokes the ververica.DeleteConnector API synchronously
func (client *Client) DeleteConnector(request *DeleteConnectorRequest) (response *DeleteConnectorResponse, err error) {
	response = CreateDeleteConnectorResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteConnectorWithChan invokes the ververica.DeleteConnector API asynchronously
func (client *Client) DeleteConnectorWithChan(request *DeleteConnectorRequest) (<-chan *DeleteConnectorResponse, <-chan error) {
	responseChan := make(chan *DeleteConnectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteConnector(request)
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

// DeleteConnectorWithCallback invokes the ververica.DeleteConnector API asynchronously
func (client *Client) DeleteConnectorWithCallback(request *DeleteConnectorRequest, callback func(response *DeleteConnectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteConnectorResponse
		var err error
		defer close(result)
		response, err = client.DeleteConnector(request)
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

// DeleteConnectorRequest is the request struct for api DeleteConnector
type DeleteConnectorRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Name      string `position:"Path" name:"name"`
	Namespace string `position:"Path" name:"namespace"`
}

// DeleteConnectorResponse is the response struct for api DeleteConnector
type DeleteConnectorResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateDeleteConnectorRequest creates a request to invoke DeleteConnector API
func CreateDeleteConnectorRequest() (request *DeleteConnectorRequest) {
	request = &DeleteConnectorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "DeleteConnector", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/connectors/[name]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteConnectorResponse creates a response to parse from DeleteConnector response
func CreateDeleteConnectorResponse() (response *DeleteConnectorResponse) {
	response = &DeleteConnectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
