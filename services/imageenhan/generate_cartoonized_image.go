package imageenhan

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

// GenerateCartoonizedImage invokes the imageenhan.GenerateCartoonizedImage API synchronously
func (client *Client) GenerateCartoonizedImage(request *GenerateCartoonizedImageRequest) (response *GenerateCartoonizedImageResponse, err error) {
	response = CreateGenerateCartoonizedImageResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateCartoonizedImageWithChan invokes the imageenhan.GenerateCartoonizedImage API asynchronously
func (client *Client) GenerateCartoonizedImageWithChan(request *GenerateCartoonizedImageRequest) (<-chan *GenerateCartoonizedImageResponse, <-chan error) {
	responseChan := make(chan *GenerateCartoonizedImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateCartoonizedImage(request)
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

// GenerateCartoonizedImageWithCallback invokes the imageenhan.GenerateCartoonizedImage API asynchronously
func (client *Client) GenerateCartoonizedImageWithCallback(request *GenerateCartoonizedImageRequest, callback func(response *GenerateCartoonizedImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateCartoonizedImageResponse
		var err error
		defer close(result)
		response, err = client.GenerateCartoonizedImage(request)
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

// GenerateCartoonizedImageRequest is the request struct for api GenerateCartoonizedImage
type GenerateCartoonizedImageRequest struct {
	*requests.RpcRequest
	ImageType string           `position:"Body" name:"ImageType"`
	Index     string           `position:"Body" name:"Index"`
	Async     requests.Boolean `position:"Body" name:"Async"`
	ImageUrl  string           `position:"Body" name:"ImageUrl"`
}

// GenerateCartoonizedImageResponse is the response struct for api GenerateCartoonizedImage
type GenerateCartoonizedImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGenerateCartoonizedImageRequest creates a request to invoke GenerateCartoonizedImage API
func CreateGenerateCartoonizedImageRequest() (request *GenerateCartoonizedImageRequest) {
	request = &GenerateCartoonizedImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageenhan", "2019-09-30", "GenerateCartoonizedImage", "imageenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGenerateCartoonizedImageResponse creates a response to parse from GenerateCartoonizedImage response
func CreateGenerateCartoonizedImageResponse() (response *GenerateCartoonizedImageResponse) {
	response = &GenerateCartoonizedImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
