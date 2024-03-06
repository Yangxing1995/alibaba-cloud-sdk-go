package paielasticdatasetaccelerator

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

// DeleteEndpoint invokes the paielasticdatasetaccelerator.DeleteEndpoint API synchronously
func (client *Client) DeleteEndpoint(request *DeleteEndpointRequest) (response *DeleteEndpointResponse, err error) {
	response = CreateDeleteEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEndpointWithChan invokes the paielasticdatasetaccelerator.DeleteEndpoint API asynchronously
func (client *Client) DeleteEndpointWithChan(request *DeleteEndpointRequest) (<-chan *DeleteEndpointResponse, <-chan error) {
	responseChan := make(chan *DeleteEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEndpoint(request)
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

// DeleteEndpointWithCallback invokes the paielasticdatasetaccelerator.DeleteEndpoint API asynchronously
func (client *Client) DeleteEndpointWithCallback(request *DeleteEndpointRequest, callback func(response *DeleteEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEndpointResponse
		var err error
		defer close(result)
		response, err = client.DeleteEndpoint(request)
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

// DeleteEndpointRequest is the request struct for api DeleteEndpoint
type DeleteEndpointRequest struct {
	*requests.RoaRequest
	EndpointId string `position:"Path" name:"EndpointId"`
}

// DeleteEndpointResponse is the response struct for api DeleteEndpoint
type DeleteEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEndpointRequest creates a request to invoke DeleteEndpoint API
func CreateDeleteEndpointRequest() (request *DeleteEndpointRequest) {
	request = &DeleteEndpointRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "DeleteEndpoint", "/api/v1/endpoints/[EndpointId]", "datasetacc", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteEndpointResponse creates a response to parse from DeleteEndpoint response
func CreateDeleteEndpointResponse() (response *DeleteEndpointResponse) {
	response = &DeleteEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
