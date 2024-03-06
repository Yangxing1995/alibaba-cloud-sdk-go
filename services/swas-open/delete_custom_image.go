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

// DeleteCustomImage invokes the swas_open.DeleteCustomImage API synchronously
func (client *Client) DeleteCustomImage(request *DeleteCustomImageRequest) (response *DeleteCustomImageResponse, err error) {
	response = CreateDeleteCustomImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCustomImageWithChan invokes the swas_open.DeleteCustomImage API asynchronously
func (client *Client) DeleteCustomImageWithChan(request *DeleteCustomImageRequest) (<-chan *DeleteCustomImageResponse, <-chan error) {
	responseChan := make(chan *DeleteCustomImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCustomImage(request)
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

// DeleteCustomImageWithCallback invokes the swas_open.DeleteCustomImage API asynchronously
func (client *Client) DeleteCustomImageWithCallback(request *DeleteCustomImageRequest, callback func(response *DeleteCustomImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCustomImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteCustomImage(request)
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

// DeleteCustomImageRequest is the request struct for api DeleteCustomImage
type DeleteCustomImageRequest struct {
	*requests.RpcRequest
	ImageId     string `position:"Query" name:"ImageId"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

// DeleteCustomImageResponse is the response struct for api DeleteCustomImage
type DeleteCustomImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCustomImageRequest creates a request to invoke DeleteCustomImage API
func CreateDeleteCustomImageRequest() (request *DeleteCustomImageRequest) {
	request = &DeleteCustomImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DeleteCustomImage", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCustomImageResponse creates a response to parse from DeleteCustomImage response
func CreateDeleteCustomImageResponse() (response *DeleteCustomImageResponse) {
	response = &DeleteCustomImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
