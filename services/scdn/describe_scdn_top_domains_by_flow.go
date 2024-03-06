package scdn

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

// DescribeScdnTopDomainsByFlow invokes the scdn.DescribeScdnTopDomainsByFlow API synchronously
func (client *Client) DescribeScdnTopDomainsByFlow(request *DescribeScdnTopDomainsByFlowRequest) (response *DescribeScdnTopDomainsByFlowResponse, err error) {
	response = CreateDescribeScdnTopDomainsByFlowResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnTopDomainsByFlowWithChan invokes the scdn.DescribeScdnTopDomainsByFlow API asynchronously
func (client *Client) DescribeScdnTopDomainsByFlowWithChan(request *DescribeScdnTopDomainsByFlowRequest) (<-chan *DescribeScdnTopDomainsByFlowResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnTopDomainsByFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnTopDomainsByFlow(request)
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

// DescribeScdnTopDomainsByFlowWithCallback invokes the scdn.DescribeScdnTopDomainsByFlow API asynchronously
func (client *Client) DescribeScdnTopDomainsByFlowWithCallback(request *DescribeScdnTopDomainsByFlowRequest, callback func(response *DescribeScdnTopDomainsByFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnTopDomainsByFlowResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnTopDomainsByFlow(request)
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

// DescribeScdnTopDomainsByFlowRequest is the request struct for api DescribeScdnTopDomainsByFlow
type DescribeScdnTopDomainsByFlowRequest struct {
	*requests.RpcRequest
	Product   string           `position:"Query" name:"Product"`
	Limit     requests.Integer `position:"Query" name:"Limit"`
	EndTime   string           `position:"Query" name:"EndTime"`
	StartTime string           `position:"Query" name:"StartTime"`
}

// DescribeScdnTopDomainsByFlowResponse is the response struct for api DescribeScdnTopDomainsByFlow
type DescribeScdnTopDomainsByFlowResponse struct {
	*responses.BaseResponse
	DomainOnlineCount int64      `json:"DomainOnlineCount" xml:"DomainOnlineCount"`
	EndTime           string     `json:"EndTime" xml:"EndTime"`
	StartTime         string     `json:"StartTime" xml:"StartTime"`
	RequestId         string     `json:"RequestId" xml:"RequestId"`
	DomainCount       int64      `json:"DomainCount" xml:"DomainCount"`
	TopDomains        TopDomains `json:"TopDomains" xml:"TopDomains"`
}

// CreateDescribeScdnTopDomainsByFlowRequest creates a request to invoke DescribeScdnTopDomainsByFlow API
func CreateDescribeScdnTopDomainsByFlowRequest() (request *DescribeScdnTopDomainsByFlowRequest) {
	request = &DescribeScdnTopDomainsByFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnTopDomainsByFlow", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnTopDomainsByFlowResponse creates a response to parse from DescribeScdnTopDomainsByFlow response
func CreateDescribeScdnTopDomainsByFlowResponse() (response *DescribeScdnTopDomainsByFlowResponse) {
	response = &DescribeScdnTopDomainsByFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
