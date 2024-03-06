package adb

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

// DescribeSqlPattern invokes the adb.DescribeSqlPattern API synchronously
func (client *Client) DescribeSqlPattern(request *DescribeSqlPatternRequest) (response *DescribeSqlPatternResponse, err error) {
	response = CreateDescribeSqlPatternResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSqlPatternWithChan invokes the adb.DescribeSqlPattern API asynchronously
func (client *Client) DescribeSqlPatternWithChan(request *DescribeSqlPatternRequest) (<-chan *DescribeSqlPatternResponse, <-chan error) {
	responseChan := make(chan *DescribeSqlPatternResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSqlPattern(request)
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

// DescribeSqlPatternWithCallback invokes the adb.DescribeSqlPattern API asynchronously
func (client *Client) DescribeSqlPatternWithCallback(request *DescribeSqlPatternRequest, callback func(response *DescribeSqlPatternResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSqlPatternResponse
		var err error
		defer close(result)
		response, err = client.DescribeSqlPattern(request)
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

// DescribeSqlPatternRequest is the request struct for api DescribeSqlPattern
type DescribeSqlPatternRequest struct {
	*requests.RpcRequest
	DBClusterId string           `position:"Query" name:"DBClusterId"`
	SqlPattern  string           `position:"Query" name:"SqlPattern"`
	StartTime   string           `position:"Query" name:"StartTime"`
	Type        string           `position:"Query" name:"Type"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Order       string           `position:"Query" name:"Order"`
}

// DescribeSqlPatternResponse is the response struct for api DescribeSqlPattern
type DescribeSqlPatternResponse struct {
	*responses.BaseResponse
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	TotalCount int         `json:"TotalCount" xml:"TotalCount"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Items      []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeSqlPatternRequest creates a request to invoke DescribeSqlPattern API
func CreateDescribeSqlPatternRequest() (request *DescribeSqlPatternRequest) {
	request = &DescribeSqlPatternRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeSqlPattern", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSqlPatternResponse creates a response to parse from DescribeSqlPattern response
func CreateDescribeSqlPatternResponse() (response *DescribeSqlPatternResponse) {
	response = &DescribeSqlPatternResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
