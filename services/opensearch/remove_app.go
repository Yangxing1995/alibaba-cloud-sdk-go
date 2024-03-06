package opensearch

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

// RemoveApp invokes the opensearch.RemoveApp API synchronously
func (client *Client) RemoveApp(request *RemoveAppRequest) (response *RemoveAppResponse, err error) {
	response = CreateRemoveAppResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveAppWithChan invokes the opensearch.RemoveApp API asynchronously
func (client *Client) RemoveAppWithChan(request *RemoveAppRequest) (<-chan *RemoveAppResponse, <-chan error) {
	responseChan := make(chan *RemoveAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveApp(request)
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

// RemoveAppWithCallback invokes the opensearch.RemoveApp API asynchronously
func (client *Client) RemoveAppWithCallback(request *RemoveAppRequest, callback func(response *RemoveAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveAppResponse
		var err error
		defer close(result)
		response, err = client.RemoveApp(request)
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

// RemoveAppRequest is the request struct for api RemoveApp
type RemoveAppRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// RemoveAppResponse is the response struct for api RemoveApp
type RemoveAppResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    []int  `json:"result" xml:"result"`
}

// CreateRemoveAppRequest creates a request to invoke RemoveApp API
func CreateRemoveAppRequest() (request *RemoveAppRequest) {
	request = &RemoveAppRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "RemoveApp", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateRemoveAppResponse creates a response to parse from RemoveApp response
func CreateRemoveAppResponse() (response *RemoveAppResponse) {
	response = &RemoveAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
