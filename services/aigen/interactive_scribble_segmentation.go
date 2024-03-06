package aigen

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

// InteractiveScribbleSegmentation invokes the aigen.InteractiveScribbleSegmentation API synchronously
func (client *Client) InteractiveScribbleSegmentation(request *InteractiveScribbleSegmentationRequest) (response *InteractiveScribbleSegmentationResponse, err error) {
	response = CreateInteractiveScribbleSegmentationResponse()
	err = client.DoAction(request, response)
	return
}

// InteractiveScribbleSegmentationWithChan invokes the aigen.InteractiveScribbleSegmentation API asynchronously
func (client *Client) InteractiveScribbleSegmentationWithChan(request *InteractiveScribbleSegmentationRequest) (<-chan *InteractiveScribbleSegmentationResponse, <-chan error) {
	responseChan := make(chan *InteractiveScribbleSegmentationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InteractiveScribbleSegmentation(request)
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

// InteractiveScribbleSegmentationWithCallback invokes the aigen.InteractiveScribbleSegmentation API asynchronously
func (client *Client) InteractiveScribbleSegmentationWithCallback(request *InteractiveScribbleSegmentationRequest, callback func(response *InteractiveScribbleSegmentationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InteractiveScribbleSegmentationResponse
		var err error
		defer close(result)
		response, err = client.InteractiveScribbleSegmentation(request)
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

// InteractiveScribbleSegmentationRequest is the request struct for api InteractiveScribbleSegmentation
type InteractiveScribbleSegmentationRequest struct {
	*requests.RpcRequest
	PosScribbleImageUrl string `position:"Body" name:"PosScribbleImageUrl"`
	IntegratedMaskUrl   string `position:"Body" name:"IntegratedMaskUrl"`
	MaskImageUrl        string `position:"Body" name:"MaskImageUrl"`
	ReturnForm          string `position:"Body" name:"ReturnForm"`
	NegScribbleImageUrl string `position:"Body" name:"NegScribbleImageUrl"`
	ReturnFormat        string `position:"Body" name:"ReturnFormat"`
	EdgeFeathering      string `position:"Body" name:"EdgeFeathering"`
	ImageUrl            string `position:"Body" name:"ImageUrl"`
	PostprocessOption   string `position:"Body" name:"PostprocessOption"`
}

// InteractiveScribbleSegmentationResponse is the response struct for api InteractiveScribbleSegmentation
type InteractiveScribbleSegmentationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateInteractiveScribbleSegmentationRequest creates a request to invoke InteractiveScribbleSegmentation API
func CreateInteractiveScribbleSegmentationRequest() (request *InteractiveScribbleSegmentationRequest) {
	request = &InteractiveScribbleSegmentationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aigen", "2024-01-11", "InteractiveScribbleSegmentation", "", "")
	request.Method = requests.POST
	return
}

// CreateInteractiveScribbleSegmentationResponse creates a response to parse from InteractiveScribbleSegmentation response
func CreateInteractiveScribbleSegmentationResponse() (response *InteractiveScribbleSegmentationResponse) {
	response = &InteractiveScribbleSegmentationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
