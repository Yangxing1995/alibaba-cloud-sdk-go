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

// DescribeLiveDomainRecordUsageData invokes the live.DescribeLiveDomainRecordUsageData API synchronously
func (client *Client) DescribeLiveDomainRecordUsageData(request *DescribeLiveDomainRecordUsageDataRequest) (response *DescribeLiveDomainRecordUsageDataResponse, err error) {
	response = CreateDescribeLiveDomainRecordUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainRecordUsageDataWithChan invokes the live.DescribeLiveDomainRecordUsageData API asynchronously
func (client *Client) DescribeLiveDomainRecordUsageDataWithChan(request *DescribeLiveDomainRecordUsageDataRequest) (<-chan *DescribeLiveDomainRecordUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainRecordUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainRecordUsageData(request)
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

// DescribeLiveDomainRecordUsageDataWithCallback invokes the live.DescribeLiveDomainRecordUsageData API asynchronously
func (client *Client) DescribeLiveDomainRecordUsageDataWithCallback(request *DescribeLiveDomainRecordUsageDataRequest, callback func(response *DescribeLiveDomainRecordUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainRecordUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainRecordUsageData(request)
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

// DescribeLiveDomainRecordUsageDataRequest is the request struct for api DescribeLiveDomainRecordUsageData
type DescribeLiveDomainRecordUsageDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	SplitBy    string           `position:"Query" name:"SplitBy"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainRecordUsageDataResponse is the response struct for api DescribeLiveDomainRecordUsageData
type DescribeLiveDomainRecordUsageDataResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	RecordUsageData RecordUsageData `json:"RecordUsageData" xml:"RecordUsageData"`
}

// CreateDescribeLiveDomainRecordUsageDataRequest creates a request to invoke DescribeLiveDomainRecordUsageData API
func CreateDescribeLiveDomainRecordUsageDataRequest() (request *DescribeLiveDomainRecordUsageDataRequest) {
	request = &DescribeLiveDomainRecordUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainRecordUsageData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainRecordUsageDataResponse creates a response to parse from DescribeLiveDomainRecordUsageData response
func CreateDescribeLiveDomainRecordUsageDataResponse() (response *DescribeLiveDomainRecordUsageDataResponse) {
	response = &DescribeLiveDomainRecordUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
