package clickhouse

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

// DescribeRDSTables invokes the clickhouse.DescribeRDSTables API synchronously
func (client *Client) DescribeRDSTables(request *DescribeRDSTablesRequest) (response *DescribeRDSTablesResponse, err error) {
	response = CreateDescribeRDSTablesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRDSTablesWithChan invokes the clickhouse.DescribeRDSTables API asynchronously
func (client *Client) DescribeRDSTablesWithChan(request *DescribeRDSTablesRequest) (<-chan *DescribeRDSTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeRDSTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRDSTables(request)
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

// DescribeRDSTablesWithCallback invokes the clickhouse.DescribeRDSTables API asynchronously
func (client *Client) DescribeRDSTablesWithCallback(request *DescribeRDSTablesRequest, callback func(response *DescribeRDSTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRDSTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRDSTables(request)
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

// DescribeRDSTablesRequest is the request struct for api DescribeRDSTables
type DescribeRDSTablesRequest struct {
	*requests.RpcRequest
	Schema               string           `position:"Query" name:"Schema"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RdsPassword          string           `position:"Query" name:"RdsPassword"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DbClusterId          string           `position:"Query" name:"DbClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RdsId                string           `position:"Query" name:"RdsId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RdsPort              requests.Integer `position:"Query" name:"RdsPort"`
	RdsVpcUrl            string           `position:"Query" name:"RdsVpcUrl"`
	RdsUserName          string           `position:"Query" name:"RdsUserName"`
}

// DescribeRDSTablesResponse is the response struct for api DescribeRDSTables
type DescribeRDSTablesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Tables    []string `json:"Tables" xml:"Tables"`
}

// CreateDescribeRDSTablesRequest creates a request to invoke DescribeRDSTables API
func CreateDescribeRDSTablesRequest() (request *DescribeRDSTablesRequest) {
	request = &DescribeRDSTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeRDSTables", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRDSTablesResponse creates a response to parse from DescribeRDSTables response
func CreateDescribeRDSTablesResponse() (response *DescribeRDSTablesResponse) {
	response = &DescribeRDSTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
