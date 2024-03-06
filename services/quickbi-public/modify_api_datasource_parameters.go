package quickbi_public

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

// ModifyApiDatasourceParameters invokes the quickbi_public.ModifyApiDatasourceParameters API synchronously
func (client *Client) ModifyApiDatasourceParameters(request *ModifyApiDatasourceParametersRequest) (response *ModifyApiDatasourceParametersResponse, err error) {
	response = CreateModifyApiDatasourceParametersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyApiDatasourceParametersWithChan invokes the quickbi_public.ModifyApiDatasourceParameters API asynchronously
func (client *Client) ModifyApiDatasourceParametersWithChan(request *ModifyApiDatasourceParametersRequest) (<-chan *ModifyApiDatasourceParametersResponse, <-chan error) {
	responseChan := make(chan *ModifyApiDatasourceParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApiDatasourceParameters(request)
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

// ModifyApiDatasourceParametersWithCallback invokes the quickbi_public.ModifyApiDatasourceParameters API asynchronously
func (client *Client) ModifyApiDatasourceParametersWithCallback(request *ModifyApiDatasourceParametersRequest, callback func(response *ModifyApiDatasourceParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiDatasourceParametersResponse
		var err error
		defer close(result)
		response, err = client.ModifyApiDatasourceParameters(request)
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

// ModifyApiDatasourceParametersRequest is the request struct for api ModifyApiDatasourceParameters
type ModifyApiDatasourceParametersRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	Parameters  string `position:"Query" name:"Parameters"`
	ApiId       string `position:"Query" name:"ApiId"`
	WorkspaceId string `position:"Query" name:"WorkspaceId"`
}

// ModifyApiDatasourceParametersResponse is the response struct for api ModifyApiDatasourceParameters
type ModifyApiDatasourceParametersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyApiDatasourceParametersRequest creates a request to invoke ModifyApiDatasourceParameters API
func CreateModifyApiDatasourceParametersRequest() (request *ModifyApiDatasourceParametersRequest) {
	request = &ModifyApiDatasourceParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ModifyApiDatasourceParameters", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyApiDatasourceParametersResponse creates a response to parse from ModifyApiDatasourceParameters response
func CreateModifyApiDatasourceParametersResponse() (response *ModifyApiDatasourceParametersResponse) {
	response = &ModifyApiDatasourceParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
