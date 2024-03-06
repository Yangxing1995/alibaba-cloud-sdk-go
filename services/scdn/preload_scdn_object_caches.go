package scdn

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

// PreloadScdnObjectCaches invokes the scdn.PreloadScdnObjectCaches API synchronously
func (client *Client) PreloadScdnObjectCaches(request *PreloadScdnObjectCachesRequest) (response *PreloadScdnObjectCachesResponse, err error) {
	response = CreatePreloadScdnObjectCachesResponse()
	err = client.DoAction(request, response)
	return
}

// PreloadScdnObjectCachesWithChan invokes the scdn.PreloadScdnObjectCaches API asynchronously
func (client *Client) PreloadScdnObjectCachesWithChan(request *PreloadScdnObjectCachesRequest) (<-chan *PreloadScdnObjectCachesResponse, <-chan error) {
	responseChan := make(chan *PreloadScdnObjectCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreloadScdnObjectCaches(request)
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

// PreloadScdnObjectCachesWithCallback invokes the scdn.PreloadScdnObjectCaches API asynchronously
func (client *Client) PreloadScdnObjectCachesWithCallback(request *PreloadScdnObjectCachesRequest, callback func(response *PreloadScdnObjectCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreloadScdnObjectCachesResponse
		var err error
		defer close(result)
		response, err = client.PreloadScdnObjectCaches(request)
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

// PreloadScdnObjectCachesRequest is the request struct for api PreloadScdnObjectCaches
type PreloadScdnObjectCachesRequest struct {
	*requests.RpcRequest
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	L2Preload     requests.Boolean `position:"Query" name:"L2Preload"`
	Area          string           `position:"Query" name:"Area"`
	WithHeader    string           `position:"Query" name:"WithHeader"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// PreloadScdnObjectCachesResponse is the response struct for api PreloadScdnObjectCaches
type PreloadScdnObjectCachesResponse struct {
	*responses.BaseResponse
	PreloadTaskId string `json:"PreloadTaskId" xml:"PreloadTaskId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePreloadScdnObjectCachesRequest creates a request to invoke PreloadScdnObjectCaches API
func CreatePreloadScdnObjectCachesRequest() (request *PreloadScdnObjectCachesRequest) {
	request = &PreloadScdnObjectCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "PreloadScdnObjectCaches", "", "")
	request.Method = requests.POST
	return
}

// CreatePreloadScdnObjectCachesResponse creates a response to parse from PreloadScdnObjectCaches response
func CreatePreloadScdnObjectCachesResponse() (response *PreloadScdnObjectCachesResponse) {
	response = &PreloadScdnObjectCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
