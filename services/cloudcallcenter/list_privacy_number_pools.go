package cloudcallcenter

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

// ListPrivacyNumberPools invokes the cloudcallcenter.ListPrivacyNumberPools API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumberpools.html
func (client *Client) ListPrivacyNumberPools(request *ListPrivacyNumberPoolsRequest) (response *ListPrivacyNumberPoolsResponse, err error) {
	response = CreateListPrivacyNumberPoolsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrivacyNumberPoolsWithChan invokes the cloudcallcenter.ListPrivacyNumberPools API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumberpools.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrivacyNumberPoolsWithChan(request *ListPrivacyNumberPoolsRequest) (<-chan *ListPrivacyNumberPoolsResponse, <-chan error) {
	responseChan := make(chan *ListPrivacyNumberPoolsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrivacyNumberPools(request)
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

// ListPrivacyNumberPoolsWithCallback invokes the cloudcallcenter.ListPrivacyNumberPools API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumberpools.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrivacyNumberPoolsWithCallback(request *ListPrivacyNumberPoolsRequest, callback func(response *ListPrivacyNumberPoolsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrivacyNumberPoolsResponse
		var err error
		defer close(result)
		response, err = client.ListPrivacyNumberPools(request)
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

// ListPrivacyNumberPoolsRequest is the request struct for api ListPrivacyNumberPools
type ListPrivacyNumberPoolsRequest struct {
	*requests.RpcRequest
}

// ListPrivacyNumberPoolsResponse is the response struct for api ListPrivacyNumberPools
type ListPrivacyNumberPoolsResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Success            bool               `json:"Success" xml:"Success"`
	Code               string             `json:"Code" xml:"Code"`
	Message            string             `json:"Message" xml:"Message"`
	HttpStatusCode     int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PrivacyNumberPools PrivacyNumberPools `json:"PrivacyNumberPools" xml:"PrivacyNumberPools"`
}

// CreateListPrivacyNumberPoolsRequest creates a request to invoke ListPrivacyNumberPools API
func CreateListPrivacyNumberPoolsRequest() (request *ListPrivacyNumberPoolsRequest) {
	request = &ListPrivacyNumberPoolsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListPrivacyNumberPools", "", "")
	request.Method = requests.POST
	return
}

// CreateListPrivacyNumberPoolsResponse creates a response to parse from ListPrivacyNumberPools response
func CreateListPrivacyNumberPoolsResponse() (response *ListPrivacyNumberPoolsResponse) {
	response = &ListPrivacyNumberPoolsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}