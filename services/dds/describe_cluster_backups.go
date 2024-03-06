package dds

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

// DescribeClusterBackups invokes the dds.DescribeClusterBackups API synchronously
func (client *Client) DescribeClusterBackups(request *DescribeClusterBackupsRequest) (response *DescribeClusterBackupsResponse, err error) {
	response = CreateDescribeClusterBackupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterBackupsWithChan invokes the dds.DescribeClusterBackups API asynchronously
func (client *Client) DescribeClusterBackupsWithChan(request *DescribeClusterBackupsRequest) (<-chan *DescribeClusterBackupsResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterBackupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterBackups(request)
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

// DescribeClusterBackupsWithCallback invokes the dds.DescribeClusterBackups API asynchronously
func (client *Client) DescribeClusterBackupsWithCallback(request *DescribeClusterBackupsRequest, callback func(response *DescribeClusterBackupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterBackupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterBackups(request)
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

// DescribeClusterBackupsRequest is the request struct for api DescribeClusterBackups
type DescribeClusterBackupsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime              string           `position:"Query" name:"StartTime"`
	SecurityToken          string           `position:"Query" name:"SecurityToken"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	BackupId               string           `position:"Query" name:"BackupId"`
	EndTime                string           `position:"Query" name:"EndTime"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	IsOnlyGetClusterBackUp requests.Boolean `position:"Query" name:"IsOnlyGetClusterBackUp"`
	PageNo                 requests.Integer `position:"Query" name:"PageNo"`
}

// DescribeClusterBackupsResponse is the response struct for api DescribeClusterBackups
type DescribeClusterBackupsResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	MaxResults     int             `json:"MaxResults" xml:"MaxResults"`
	PageNumber     int             `json:"PageNumber" xml:"PageNumber"`
	PageSize       int             `json:"PageSize" xml:"PageSize"`
	ClusterBackups []ClusterBackup `json:"ClusterBackups" xml:"ClusterBackups"`
}

// CreateDescribeClusterBackupsRequest creates a request to invoke DescribeClusterBackups API
func CreateDescribeClusterBackupsRequest() (request *DescribeClusterBackupsRequest) {
	request = &DescribeClusterBackupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeClusterBackups", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterBackupsResponse creates a response to parse from DescribeClusterBackups response
func CreateDescribeClusterBackupsResponse() (response *DescribeClusterBackupsResponse) {
	response = &DescribeClusterBackupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
