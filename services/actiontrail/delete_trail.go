package actiontrail

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

// DeleteTrail invokes the actiontrail.DeleteTrail API synchronously
func (client *Client) DeleteTrail(request *DeleteTrailRequest) (response *DeleteTrailResponse, err error) {
	response = CreateDeleteTrailResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTrailWithChan invokes the actiontrail.DeleteTrail API asynchronously
func (client *Client) DeleteTrailWithChan(request *DeleteTrailRequest) (<-chan *DeleteTrailResponse, <-chan error) {
	responseChan := make(chan *DeleteTrailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTrail(request)
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

// DeleteTrailWithCallback invokes the actiontrail.DeleteTrail API asynchronously
func (client *Client) DeleteTrailWithCallback(request *DeleteTrailRequest, callback func(response *DeleteTrailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTrailResponse
		var err error
		defer close(result)
		response, err = client.DeleteTrail(request)
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

// DeleteTrailRequest is the request struct for api DeleteTrail
type DeleteTrailRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// DeleteTrailResponse is the response struct for api DeleteTrail
type DeleteTrailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTrailRequest creates a request to invoke DeleteTrail API
func CreateDeleteTrailRequest() (request *DeleteTrailRequest) {
	request = &DeleteTrailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2020-07-06", "DeleteTrail", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTrailResponse creates a response to parse from DeleteTrail response
func CreateDeleteTrailResponse() (response *DeleteTrailResponse) {
	response = &DeleteTrailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
