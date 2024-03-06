package ververica

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

// TableExists invokes the ververica.TableExists API synchronously
func (client *Client) TableExists(request *TableExistsRequest) (response *TableExistsResponse, err error) {
	response = CreateTableExistsResponse()
	err = client.DoAction(request, response)
	return
}

// TableExistsWithChan invokes the ververica.TableExists API asynchronously
func (client *Client) TableExistsWithChan(request *TableExistsRequest) (<-chan *TableExistsResponse, <-chan error) {
	responseChan := make(chan *TableExistsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TableExists(request)
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

// TableExistsWithCallback invokes the ververica.TableExists API asynchronously
func (client *Client) TableExistsWithCallback(request *TableExistsRequest, callback func(response *TableExistsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TableExistsResponse
		var err error
		defer close(result)
		response, err = client.TableExists(request)
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

// TableExistsRequest is the request struct for api TableExists
type TableExistsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Database  string `position:"Query" name:"database"`
	Cat       string `position:"Path" name:"cat"`
	Namespace string `position:"Path" name:"namespace"`
	Table     string `position:"Query" name:"table"`
}

// TableExistsResponse is the response struct for api TableExists
type TableExistsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateTableExistsRequest creates a request to invoke TableExists API
func CreateTableExistsRequest() (request *TableExistsRequest) {
	request = &TableExistsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "TableExists", "/pop/workspaces/[workspace]/catalog/v1beta2/namespaces/[namespace]/catalogs/[cat]:tableExists", "", "")
	request.Method = requests.GET
	return
}

// CreateTableExistsResponse creates a response to parse from TableExists response
func CreateTableExistsResponse() (response *TableExistsResponse) {
	response = &TableExistsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
