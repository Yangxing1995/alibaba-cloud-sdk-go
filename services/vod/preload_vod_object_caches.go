package vod

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

// PreloadVodObjectCaches invokes the vod.PreloadVodObjectCaches API synchronously
func (client *Client) PreloadVodObjectCaches(request *PreloadVodObjectCachesRequest) (response *PreloadVodObjectCachesResponse, err error) {
	response = CreatePreloadVodObjectCachesResponse()
	err = client.DoAction(request, response)
	return
}

// PreloadVodObjectCachesWithChan invokes the vod.PreloadVodObjectCaches API asynchronously
func (client *Client) PreloadVodObjectCachesWithChan(request *PreloadVodObjectCachesRequest) (<-chan *PreloadVodObjectCachesResponse, <-chan error) {
	responseChan := make(chan *PreloadVodObjectCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreloadVodObjectCaches(request)
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

// PreloadVodObjectCachesWithCallback invokes the vod.PreloadVodObjectCaches API asynchronously
func (client *Client) PreloadVodObjectCachesWithCallback(request *PreloadVodObjectCachesRequest, callback func(response *PreloadVodObjectCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreloadVodObjectCachesResponse
		var err error
		defer close(result)
		response, err = client.PreloadVodObjectCaches(request)
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

// PreloadVodObjectCachesRequest is the request struct for api PreloadVodObjectCaches
type PreloadVodObjectCachesRequest struct {
	*requests.RpcRequest
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// PreloadVodObjectCachesResponse is the response struct for api PreloadVodObjectCaches
type PreloadVodObjectCachesResponse struct {
	*responses.BaseResponse
	PreloadTaskId string `json:"PreloadTaskId" xml:"PreloadTaskId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePreloadVodObjectCachesRequest creates a request to invoke PreloadVodObjectCaches API
func CreatePreloadVodObjectCachesRequest() (request *PreloadVodObjectCachesRequest) {
	request = &PreloadVodObjectCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "PreloadVodObjectCaches", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePreloadVodObjectCachesResponse creates a response to parse from PreloadVodObjectCaches response
func CreatePreloadVodObjectCachesResponse() (response *PreloadVodObjectCachesResponse) {
	response = &PreloadVodObjectCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
