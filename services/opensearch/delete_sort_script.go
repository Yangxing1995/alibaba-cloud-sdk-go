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

// DeleteSortScript invokes the opensearch.DeleteSortScript API synchronously
func (client *Client) DeleteSortScript(request *DeleteSortScriptRequest) (response *DeleteSortScriptResponse, err error) {
	response = CreateDeleteSortScriptResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSortScriptWithChan invokes the opensearch.DeleteSortScript API asynchronously
func (client *Client) DeleteSortScriptWithChan(request *DeleteSortScriptRequest) (<-chan *DeleteSortScriptResponse, <-chan error) {
	responseChan := make(chan *DeleteSortScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSortScript(request)
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

// DeleteSortScriptWithCallback invokes the opensearch.DeleteSortScript API asynchronously
func (client *Client) DeleteSortScriptWithCallback(request *DeleteSortScriptRequest, callback func(response *DeleteSortScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSortScriptResponse
		var err error
		defer close(result)
		response, err = client.DeleteSortScript(request)
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

// DeleteSortScriptRequest is the request struct for api DeleteSortScript
type DeleteSortScriptRequest struct {
	*requests.RoaRequest
	AppVersionId     string `position:"Path" name:"appVersionId"`
	ScriptName       string `position:"Path" name:"scriptName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// DeleteSortScriptResponse is the response struct for api DeleteSortScript
type DeleteSortScriptResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateDeleteSortScriptRequest creates a request to invoke DeleteSortScript API
func CreateDeleteSortScriptRequest() (request *DeleteSortScriptRequest) {
	request = &DeleteSortScriptRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DeleteSortScript", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appVersionId]/sort-scripts/[scriptName]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteSortScriptResponse creates a response to parse from DeleteSortScript response
func CreateDeleteSortScriptResponse() (response *DeleteSortScriptResponse) {
	response = &DeleteSortScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
