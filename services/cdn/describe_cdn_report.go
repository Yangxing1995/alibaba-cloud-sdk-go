package cdn

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

// DescribeCdnReport invokes the cdn.DescribeCdnReport API synchronously
func (client *Client) DescribeCdnReport(request *DescribeCdnReportRequest) (response *DescribeCdnReportResponse, err error) {
	response = CreateDescribeCdnReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnReportWithChan invokes the cdn.DescribeCdnReport API asynchronously
func (client *Client) DescribeCdnReportWithChan(request *DescribeCdnReportRequest) (<-chan *DescribeCdnReportResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnReport(request)
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

// DescribeCdnReportWithCallback invokes the cdn.DescribeCdnReport API asynchronously
func (client *Client) DescribeCdnReportWithCallback(request *DescribeCdnReportRequest, callback func(response *DescribeCdnReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnReport(request)
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

// DescribeCdnReportRequest is the request struct for api DescribeCdnReport
type DescribeCdnReportRequest struct {
	*requests.RpcRequest
	ReportId   requests.Integer `position:"Query" name:"ReportId"`
	StartTime  string           `position:"Query" name:"StartTime"`
	Area       string           `position:"Query" name:"Area"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	HttpCode   string           `position:"Query" name:"HttpCode"`
	IsOverseas string           `position:"Query" name:"IsOverseas"`
}

// DescribeCdnReportResponse is the response struct for api DescribeCdnReport
type DescribeCdnReportResponse struct {
	*responses.BaseResponse
	Content   string `json:"Content" xml:"Content"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeCdnReportRequest creates a request to invoke DescribeCdnReport API
func CreateDescribeCdnReportRequest() (request *DescribeCdnReportRequest) {
	request = &DescribeCdnReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnReport", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnReportResponse creates a response to parse from DescribeCdnReport response
func CreateDescribeCdnReportResponse() (response *DescribeCdnReportResponse) {
	response = &DescribeCdnReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
