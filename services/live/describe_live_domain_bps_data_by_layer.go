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

// DescribeLiveDomainBpsDataByLayer invokes the live.DescribeLiveDomainBpsDataByLayer API synchronously
func (client *Client) DescribeLiveDomainBpsDataByLayer(request *DescribeLiveDomainBpsDataByLayerRequest) (response *DescribeLiveDomainBpsDataByLayerResponse, err error) {
	response = CreateDescribeLiveDomainBpsDataByLayerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainBpsDataByLayerWithChan invokes the live.DescribeLiveDomainBpsDataByLayer API asynchronously
func (client *Client) DescribeLiveDomainBpsDataByLayerWithChan(request *DescribeLiveDomainBpsDataByLayerRequest) (<-chan *DescribeLiveDomainBpsDataByLayerResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainBpsDataByLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainBpsDataByLayer(request)
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

// DescribeLiveDomainBpsDataByLayerWithCallback invokes the live.DescribeLiveDomainBpsDataByLayer API asynchronously
func (client *Client) DescribeLiveDomainBpsDataByLayerWithCallback(request *DescribeLiveDomainBpsDataByLayerRequest, callback func(response *DescribeLiveDomainBpsDataByLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainBpsDataByLayerResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainBpsDataByLayer(request)
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

// DescribeLiveDomainBpsDataByLayerRequest is the request struct for api DescribeLiveDomainBpsDataByLayer
type DescribeLiveDomainBpsDataByLayerRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	Layer          string           `position:"Query" name:"Layer"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeLiveDomainBpsDataByLayerResponse is the response struct for api DescribeLiveDomainBpsDataByLayer
type DescribeLiveDomainBpsDataByLayerResponse struct {
	*responses.BaseResponse
	DataInterval    string          `json:"DataInterval" xml:"DataInterval"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	BpsDataInterval BpsDataInterval `json:"BpsDataInterval" xml:"BpsDataInterval"`
}

// CreateDescribeLiveDomainBpsDataByLayerRequest creates a request to invoke DescribeLiveDomainBpsDataByLayer API
func CreateDescribeLiveDomainBpsDataByLayerRequest() (request *DescribeLiveDomainBpsDataByLayerRequest) {
	request = &DescribeLiveDomainBpsDataByLayerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainBpsDataByLayer", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainBpsDataByLayerResponse creates a response to parse from DescribeLiveDomainBpsDataByLayer response
func CreateDescribeLiveDomainBpsDataByLayerResponse() (response *DescribeLiveDomainBpsDataByLayerResponse) {
	response = &DescribeLiveDomainBpsDataByLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
