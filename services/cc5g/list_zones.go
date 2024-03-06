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

// ListZones invokes the cc5g.ListZones API synchronously
func (client *Client) ListZones(request *ListZonesRequest) (response *ListZonesResponse, err error) {
	response = CreateListZonesResponse()
	err = client.DoAction(request, response)
	return
}

// ListZonesWithChan invokes the cc5g.ListZones API asynchronously
func (client *Client) ListZonesWithChan(request *ListZonesRequest) (<-chan *ListZonesResponse, <-chan error) {
	responseChan := make(chan *ListZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListZones(request)
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

// ListZonesWithCallback invokes the cc5g.ListZones API asynchronously
func (client *Client) ListZonesWithCallback(request *ListZonesRequest, callback func(response *ListZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListZonesResponse
		var err error
		defer close(result)
		response, err = client.ListZones(request)
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

// ListZonesRequest is the request struct for api ListZones
type ListZonesRequest struct {
	*requests.RpcRequest
}

// ListZonesResponse is the response struct for api ListZones
type ListZonesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Zones     []Zone `json:"Zones" xml:"Zones"`
}

// CreateListZonesRequest creates a request to invoke ListZones API
func CreateListZonesRequest() (request *ListZonesRequest) {
	request = &ListZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "ListZones", "fivegcc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListZonesResponse creates a response to parse from ListZones response
func CreateListZonesResponse() (response *ListZonesResponse) {
	response = &ListZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
