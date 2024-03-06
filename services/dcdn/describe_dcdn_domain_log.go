package dcdn

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

// DescribeDcdnDomainLog invokes the dcdn.DescribeDcdnDomainLog API synchronously
func (client *Client) DescribeDcdnDomainLog(request *DescribeDcdnDomainLogRequest) (response *DescribeDcdnDomainLogResponse, err error) {
	response = CreateDescribeDcdnDomainLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainLogWithChan invokes the dcdn.DescribeDcdnDomainLog API asynchronously
func (client *Client) DescribeDcdnDomainLogWithChan(request *DescribeDcdnDomainLogRequest) (<-chan *DescribeDcdnDomainLogResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainLog(request)
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

// DescribeDcdnDomainLogWithCallback invokes the dcdn.DescribeDcdnDomainLog API asynchronously
func (client *Client) DescribeDcdnDomainLogWithCallback(request *DescribeDcdnDomainLogRequest, callback func(response *DescribeDcdnDomainLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainLog(request)
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

// DescribeDcdnDomainLogRequest is the request struct for api DescribeDcdnDomainLog
type DescribeDcdnDomainLogRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EndTime    string           `position:"Query" name:"EndTime"`
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeDcdnDomainLogResponse is the response struct for api DescribeDcdnDomainLog
type DescribeDcdnDomainLogResponse struct {
	*responses.BaseResponse
	DomainName       string                                  `json:"DomainName" xml:"DomainName"`
	RequestId        string                                  `json:"RequestId" xml:"RequestId"`
	DomainLogDetails DomainLogDetailsInDescribeDcdnDomainLog `json:"DomainLogDetails" xml:"DomainLogDetails"`
}

// CreateDescribeDcdnDomainLogRequest creates a request to invoke DescribeDcdnDomainLog API
func CreateDescribeDcdnDomainLogRequest() (request *DescribeDcdnDomainLogRequest) {
	request = &DescribeDcdnDomainLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainLog", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainLogResponse creates a response to parse from DescribeDcdnDomainLog response
func CreateDescribeDcdnDomainLogResponse() (response *DescribeDcdnDomainLogResponse) {
	response = &DescribeDcdnDomainLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
