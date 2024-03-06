package fnf

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

// ReportTaskSucceeded invokes the fnf.ReportTaskSucceeded API synchronously
func (client *Client) ReportTaskSucceeded(request *ReportTaskSucceededRequest) (response *ReportTaskSucceededResponse, err error) {
	response = CreateReportTaskSucceededResponse()
	err = client.DoAction(request, response)
	return
}

// ReportTaskSucceededWithChan invokes the fnf.ReportTaskSucceeded API asynchronously
func (client *Client) ReportTaskSucceededWithChan(request *ReportTaskSucceededRequest) (<-chan *ReportTaskSucceededResponse, <-chan error) {
	responseChan := make(chan *ReportTaskSucceededResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportTaskSucceeded(request)
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

// ReportTaskSucceededWithCallback invokes the fnf.ReportTaskSucceeded API asynchronously
func (client *Client) ReportTaskSucceededWithCallback(request *ReportTaskSucceededRequest, callback func(response *ReportTaskSucceededResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportTaskSucceededResponse
		var err error
		defer close(result)
		response, err = client.ReportTaskSucceeded(request)
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

// ReportTaskSucceededRequest is the request struct for api ReportTaskSucceeded
type ReportTaskSucceededRequest struct {
	*requests.RpcRequest
	Output    string `position:"Body" name:"Output"`
	RequestId string `position:"Query" name:"RequestId"`
	TaskToken string `position:"Query" name:"TaskToken"`
}

// ReportTaskSucceededResponse is the response struct for api ReportTaskSucceeded
type ReportTaskSucceededResponse struct {
	*responses.BaseResponse
	EventId   int64  `json:"EventId" xml:"EventId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReportTaskSucceededRequest creates a request to invoke ReportTaskSucceeded API
func CreateReportTaskSucceededRequest() (request *ReportTaskSucceededRequest) {
	request = &ReportTaskSucceededRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "ReportTaskSucceeded", "fnf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportTaskSucceededResponse creates a response to parse from ReportTaskSucceeded response
func CreateReportTaskSucceededResponse() (response *ReportTaskSucceededResponse) {
	response = &ReportTaskSucceededResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
