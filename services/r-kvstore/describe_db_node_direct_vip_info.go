package r_kvstore

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

// DescribeDBNodeDirectVipInfo invokes the r_kvstore.DescribeDBNodeDirectVipInfo API synchronously
func (client *Client) DescribeDBNodeDirectVipInfo(request *DescribeDBNodeDirectVipInfoRequest) (response *DescribeDBNodeDirectVipInfoResponse, err error) {
	response = CreateDescribeDBNodeDirectVipInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBNodeDirectVipInfoWithChan invokes the r_kvstore.DescribeDBNodeDirectVipInfo API asynchronously
func (client *Client) DescribeDBNodeDirectVipInfoWithChan(request *DescribeDBNodeDirectVipInfoRequest) (<-chan *DescribeDBNodeDirectVipInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDBNodeDirectVipInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBNodeDirectVipInfo(request)
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

// DescribeDBNodeDirectVipInfoWithCallback invokes the r_kvstore.DescribeDBNodeDirectVipInfo API asynchronously
func (client *Client) DescribeDBNodeDirectVipInfoWithCallback(request *DescribeDBNodeDirectVipInfoRequest, callback func(response *DescribeDBNodeDirectVipInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBNodeDirectVipInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBNodeDirectVipInfo(request)
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

// DescribeDBNodeDirectVipInfoRequest is the request struct for api DescribeDBNodeDirectVipInfo
type DescribeDBNodeDirectVipInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeDBNodeDirectVipInfoResponse is the response struct for api DescribeDBNodeDirectVipInfo
type DescribeDBNodeDirectVipInfoResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	InstanceId    string        `json:"InstanceId" xml:"InstanceId"`
	DirectVipInfo DirectVipInfo `json:"DirectVipInfo" xml:"DirectVipInfo"`
}

// CreateDescribeDBNodeDirectVipInfoRequest creates a request to invoke DescribeDBNodeDirectVipInfo API
func CreateDescribeDBNodeDirectVipInfoRequest() (request *DescribeDBNodeDirectVipInfoRequest) {
	request = &DescribeDBNodeDirectVipInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeDBNodeDirectVipInfo", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBNodeDirectVipInfoResponse creates a response to parse from DescribeDBNodeDirectVipInfo response
func CreateDescribeDBNodeDirectVipInfoResponse() (response *DescribeDBNodeDirectVipInfoResponse) {
	response = &DescribeDBNodeDirectVipInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
