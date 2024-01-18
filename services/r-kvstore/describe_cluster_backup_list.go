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

// DescribeClusterBackupList invokes the r_kvstore.DescribeClusterBackupList API synchronously
func (client *Client) DescribeClusterBackupList(request *DescribeClusterBackupListRequest) (response *DescribeClusterBackupListResponse, err error) {
	response = CreateDescribeClusterBackupListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterBackupListWithChan invokes the r_kvstore.DescribeClusterBackupList API asynchronously
func (client *Client) DescribeClusterBackupListWithChan(request *DescribeClusterBackupListRequest) (<-chan *DescribeClusterBackupListResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterBackupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterBackupList(request)
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

// DescribeClusterBackupListWithCallback invokes the r_kvstore.DescribeClusterBackupList API asynchronously
func (client *Client) DescribeClusterBackupListWithCallback(request *DescribeClusterBackupListRequest, callback func(response *DescribeClusterBackupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterBackupListResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterBackupList(request)
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

// DescribeClusterBackupListRequest is the request struct for api DescribeClusterBackupList
type DescribeClusterBackupListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ClusterBackupId      string           `position:"Query" name:"ClusterBackupId"`
}

// DescribeClusterBackupListResponse is the response struct for api DescribeClusterBackupList
type DescribeClusterBackupListResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	MaxResults     int             `json:"MaxResults" xml:"MaxResults"`
	PageNumber     int             `json:"PageNumber" xml:"PageNumber"`
	PageSize       int             `json:"PageSize" xml:"PageSize"`
	ClusterBackups []ClusterBackup `json:"ClusterBackups" xml:"ClusterBackups"`
}

// CreateDescribeClusterBackupListRequest creates a request to invoke DescribeClusterBackupList API
func CreateDescribeClusterBackupListRequest() (request *DescribeClusterBackupListRequest) {
	request = &DescribeClusterBackupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeClusterBackupList", "redisa", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeClusterBackupListResponse creates a response to parse from DescribeClusterBackupList response
func CreateDescribeClusterBackupListResponse() (response *DescribeClusterBackupListResponse) {
	response = &DescribeClusterBackupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
