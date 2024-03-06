package qualitycheck

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

// VerifyFile invokes the qualitycheck.VerifyFile API synchronously
func (client *Client) VerifyFile(request *VerifyFileRequest) (response *VerifyFileResponse, err error) {
	response = CreateVerifyFileResponse()
	err = client.DoAction(request, response)
	return
}

// VerifyFileWithChan invokes the qualitycheck.VerifyFile API asynchronously
func (client *Client) VerifyFileWithChan(request *VerifyFileRequest) (<-chan *VerifyFileResponse, <-chan error) {
	responseChan := make(chan *VerifyFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyFile(request)
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

// VerifyFileWithCallback invokes the qualitycheck.VerifyFile API asynchronously
func (client *Client) VerifyFileWithCallback(request *VerifyFileRequest, callback func(response *VerifyFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyFileResponse
		var err error
		defer close(result)
		response, err = client.VerifyFile(request)
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

// VerifyFileRequest is the request struct for api VerifyFile
type VerifyFileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// VerifyFileResponse is the response struct for api VerifyFile
type VerifyFileResponse struct {
	*responses.BaseResponse
	Code      string  `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	Data      float64 `json:"Data" xml:"Data"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Success   bool    `json:"Success" xml:"Success"`
}

// CreateVerifyFileRequest creates a request to invoke VerifyFile API
func CreateVerifyFileRequest() (request *VerifyFileRequest) {
	request = &VerifyFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "VerifyFile", "", "")
	request.Method = requests.POST
	return
}

// CreateVerifyFileResponse creates a response to parse from VerifyFile response
func CreateVerifyFileResponse() (response *VerifyFileResponse) {
	response = &VerifyFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
