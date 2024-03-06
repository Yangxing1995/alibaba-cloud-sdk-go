package paifeaturestore

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

// GetDatasourceTable invokes the paifeaturestore.GetDatasourceTable API synchronously
func (client *Client) GetDatasourceTable(request *GetDatasourceTableRequest) (response *GetDatasourceTableResponse, err error) {
	response = CreateGetDatasourceTableResponse()
	err = client.DoAction(request, response)
	return
}

// GetDatasourceTableWithChan invokes the paifeaturestore.GetDatasourceTable API asynchronously
func (client *Client) GetDatasourceTableWithChan(request *GetDatasourceTableRequest) (<-chan *GetDatasourceTableResponse, <-chan error) {
	responseChan := make(chan *GetDatasourceTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDatasourceTable(request)
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

// GetDatasourceTableWithCallback invokes the paifeaturestore.GetDatasourceTable API asynchronously
func (client *Client) GetDatasourceTableWithCallback(request *GetDatasourceTableRequest, callback func(response *GetDatasourceTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDatasourceTableResponse
		var err error
		defer close(result)
		response, err = client.GetDatasourceTable(request)
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

// GetDatasourceTableRequest is the request struct for api GetDatasourceTable
type GetDatasourceTableRequest struct {
	*requests.RoaRequest
	InstanceId   string `position:"Path" name:"InstanceId"`
	DatasourceId string `position:"Path" name:"DatasourceId"`
	TableName    string `position:"Path" name:"TableName"`
}

// GetDatasourceTableResponse is the response struct for api GetDatasourceTable
type GetDatasourceTableResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	TableName string       `json:"TableName" xml:"TableName"`
	Fields    []FieldsItem `json:"Fields" xml:"Fields"`
}

// CreateGetDatasourceTableRequest creates a request to invoke GetDatasourceTable API
func CreateGetDatasourceTableRequest() (request *GetDatasourceTableRequest) {
	request = &GetDatasourceTableRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetDatasourceTable", "/api/v1/instances/[InstanceId]/datasources/[DatasourceId]/tables/[TableName]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetDatasourceTableResponse creates a response to parse from GetDatasourceTable response
func CreateGetDatasourceTableResponse() (response *GetDatasourceTableResponse) {
	response = &GetDatasourceTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
