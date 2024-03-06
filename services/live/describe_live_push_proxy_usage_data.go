package live

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

// DescribeLivePushProxyUsageData invokes the live.DescribeLivePushProxyUsageData API synchronously
func (client *Client) DescribeLivePushProxyUsageData(request *DescribeLivePushProxyUsageDataRequest) (response *DescribeLivePushProxyUsageDataResponse, err error) {
	response = CreateDescribeLivePushProxyUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLivePushProxyUsageDataWithChan invokes the live.DescribeLivePushProxyUsageData API asynchronously
func (client *Client) DescribeLivePushProxyUsageDataWithChan(request *DescribeLivePushProxyUsageDataRequest) (<-chan *DescribeLivePushProxyUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLivePushProxyUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLivePushProxyUsageData(request)
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

// DescribeLivePushProxyUsageDataWithCallback invokes the live.DescribeLivePushProxyUsageData API asynchronously
func (client *Client) DescribeLivePushProxyUsageDataWithCallback(request *DescribeLivePushProxyUsageDataRequest, callback func(response *DescribeLivePushProxyUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLivePushProxyUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLivePushProxyUsageData(request)
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

// DescribeLivePushProxyUsageDataRequest is the request struct for api DescribeLivePushProxyUsageData
type DescribeLivePushProxyUsageDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	SplitBy    string           `position:"Query" name:"SplitBy"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Region     string           `position:"Query" name:"Region"`
}

// DescribeLivePushProxyUsageDataResponse is the response struct for api DescribeLivePushProxyUsageData
type DescribeLivePushProxyUsageDataResponse struct {
	*responses.BaseResponse
	EndTime       string        `json:"EndTime" xml:"EndTime"`
	StartTime     string        `json:"StartTime" xml:"StartTime"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	PushProxyData PushProxyData `json:"PushProxyData" xml:"PushProxyData"`
}

// CreateDescribeLivePushProxyUsageDataRequest creates a request to invoke DescribeLivePushProxyUsageData API
func CreateDescribeLivePushProxyUsageDataRequest() (request *DescribeLivePushProxyUsageDataRequest) {
	request = &DescribeLivePushProxyUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLivePushProxyUsageData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLivePushProxyUsageDataResponse creates a response to parse from DescribeLivePushProxyUsageData response
func CreateDescribeLivePushProxyUsageDataResponse() (response *DescribeLivePushProxyUsageDataResponse) {
	response = &DescribeLivePushProxyUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
