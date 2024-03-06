package apds

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

// SyncMigrationJob invokes the apds.SyncMigrationJob API synchronously
func (client *Client) SyncMigrationJob(request *SyncMigrationJobRequest) (response *SyncMigrationJobResponse, err error) {
	response = CreateSyncMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// SyncMigrationJobWithChan invokes the apds.SyncMigrationJob API asynchronously
func (client *Client) SyncMigrationJobWithChan(request *SyncMigrationJobRequest) (<-chan *SyncMigrationJobResponse, <-chan error) {
	responseChan := make(chan *SyncMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncMigrationJob(request)
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

// SyncMigrationJobWithCallback invokes the apds.SyncMigrationJob API asynchronously
func (client *Client) SyncMigrationJobWithCallback(request *SyncMigrationJobRequest, callback func(response *SyncMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.SyncMigrationJob(request)
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

// SyncMigrationJobRequest is the request struct for api SyncMigrationJob
type SyncMigrationJobRequest struct {
	*requests.RoaRequest
	Regions string `position:"Query" name:"regions"`
	JobType string `position:"Query" name:"jobType"`
}

// SyncMigrationJobResponse is the response struct for api SyncMigrationJob
type SyncMigrationJobResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateSyncMigrationJobRequest creates a request to invoke SyncMigrationJob API
func CreateSyncMigrationJobRequest() (request *SyncMigrationJobRequest) {
	request = &SyncMigrationJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "SyncMigrationJob", "/okss-services/migration-job/sync-migration-job", "", "")
	request.Method = requests.POST
	return
}

// CreateSyncMigrationJobResponse creates a response to parse from SyncMigrationJob response
func CreateSyncMigrationJobResponse() (response *SyncMigrationJobResponse) {
	response = &SyncMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
