package oceanbasepro

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

// SearchOmsOpenAPIMonitorMetric invokes the oceanbasepro.SearchOmsOpenAPIMonitorMetric API synchronously
func (client *Client) SearchOmsOpenAPIMonitorMetric(request *SearchOmsOpenAPIMonitorMetricRequest) (response *SearchOmsOpenAPIMonitorMetricResponse, err error) {
	response = CreateSearchOmsOpenAPIMonitorMetricResponse()
	err = client.DoAction(request, response)
	return
}

// SearchOmsOpenAPIMonitorMetricWithChan invokes the oceanbasepro.SearchOmsOpenAPIMonitorMetric API asynchronously
func (client *Client) SearchOmsOpenAPIMonitorMetricWithChan(request *SearchOmsOpenAPIMonitorMetricRequest) (<-chan *SearchOmsOpenAPIMonitorMetricResponse, <-chan error) {
	responseChan := make(chan *SearchOmsOpenAPIMonitorMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchOmsOpenAPIMonitorMetric(request)
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

// SearchOmsOpenAPIMonitorMetricWithCallback invokes the oceanbasepro.SearchOmsOpenAPIMonitorMetric API asynchronously
func (client *Client) SearchOmsOpenAPIMonitorMetricWithCallback(request *SearchOmsOpenAPIMonitorMetricRequest, callback func(response *SearchOmsOpenAPIMonitorMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchOmsOpenAPIMonitorMetricResponse
		var err error
		defer close(result)
		response, err = client.SearchOmsOpenAPIMonitorMetric(request)
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

// SearchOmsOpenAPIMonitorMetricRequest is the request struct for api SearchOmsOpenAPIMonitorMetric
type SearchOmsOpenAPIMonitorMetricRequest struct {
	*requests.RpcRequest
	EndTime       requests.Integer `position:"Body" name:"EndTime"`
	BeginTime     requests.Integer `position:"Body" name:"BeginTime"`
	MaxPointNum   requests.Integer `position:"Body" name:"MaxPointNum"`
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	Metric        string           `position:"Body" name:"Metric"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	WorkerGradeId string           `position:"Body" name:"WorkerGradeId"`
	ProjectId     string           `position:"Body" name:"ProjectId"`
}

// SearchOmsOpenAPIMonitorMetricResponse is the response struct for api SearchOmsOpenAPIMonitorMetric
type SearchOmsOpenAPIMonitorMetricResponse struct {
	*responses.BaseResponse
	Success     bool        `json:"Success" xml:"Success"`
	Code        string      `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	Advice      string      `json:"Advice" xml:"Advice"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalCount  int64       `json:"TotalCount" xml:"TotalCount"`
	Cost        string      `json:"Cost" xml:"Cost"`
	ErrorDetail ErrorDetail `json:"ErrorDetail" xml:"ErrorDetail"`
	Data        []DataItem  `json:"Data" xml:"Data"`
}

// CreateSearchOmsOpenAPIMonitorMetricRequest creates a request to invoke SearchOmsOpenAPIMonitorMetric API
func CreateSearchOmsOpenAPIMonitorMetricRequest() (request *SearchOmsOpenAPIMonitorMetricRequest) {
	request = &SearchOmsOpenAPIMonitorMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "SearchOmsOpenAPIMonitorMetric", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchOmsOpenAPIMonitorMetricResponse creates a response to parse from SearchOmsOpenAPIMonitorMetric response
func CreateSearchOmsOpenAPIMonitorMetricResponse() (response *SearchOmsOpenAPIMonitorMetricResponse) {
	response = &SearchOmsOpenAPIMonitorMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
