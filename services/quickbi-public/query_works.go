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

// QueryWorks invokes the quickbi_public.QueryWorks API synchronously
func (client *Client) QueryWorks(request *QueryWorksRequest) (response *QueryWorksResponse, err error) {
	response = CreateQueryWorksResponse()
	err = client.DoAction(request, response)
	return
}

// QueryWorksWithChan invokes the quickbi_public.QueryWorks API asynchronously
func (client *Client) QueryWorksWithChan(request *QueryWorksRequest) (<-chan *QueryWorksResponse, <-chan error) {
	responseChan := make(chan *QueryWorksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWorks(request)
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

// QueryWorksWithCallback invokes the quickbi_public.QueryWorks API asynchronously
func (client *Client) QueryWorksWithCallback(request *QueryWorksRequest, callback func(response *QueryWorksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWorksResponse
		var err error
		defer close(result)
		response, err = client.QueryWorks(request)
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

// QueryWorksRequest is the request struct for api QueryWorks
type QueryWorksRequest struct {
	*requests.RpcRequest
	ApiLevel    string `position:"Query" name:"ApiLevel"`
	WorksId     string `position:"Query" name:"WorksId"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryWorksResponse is the response struct for api QueryWorks
type QueryWorksResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryWorksRequest creates a request to invoke QueryWorks API
func CreateQueryWorksRequest() (request *QueryWorksRequest) {
	request = &QueryWorksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryWorks", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryWorksResponse creates a response to parse from QueryWorks response
func CreateQueryWorksResponse() (response *QueryWorksResponse) {
	response = &QueryWorksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
