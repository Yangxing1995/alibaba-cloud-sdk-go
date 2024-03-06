package iot

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

// QueryOTAFirmware invokes the iot.QueryOTAFirmware API synchronously
func (client *Client) QueryOTAFirmware(request *QueryOTAFirmwareRequest) (response *QueryOTAFirmwareResponse, err error) {
	response = CreateQueryOTAFirmwareResponse()
	err = client.DoAction(request, response)
	return
}

// QueryOTAFirmwareWithChan invokes the iot.QueryOTAFirmware API asynchronously
func (client *Client) QueryOTAFirmwareWithChan(request *QueryOTAFirmwareRequest) (<-chan *QueryOTAFirmwareResponse, <-chan error) {
	responseChan := make(chan *QueryOTAFirmwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryOTAFirmware(request)
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

// QueryOTAFirmwareWithCallback invokes the iot.QueryOTAFirmware API asynchronously
func (client *Client) QueryOTAFirmwareWithCallback(request *QueryOTAFirmwareRequest, callback func(response *QueryOTAFirmwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryOTAFirmwareResponse
		var err error
		defer close(result)
		response, err = client.QueryOTAFirmware(request)
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

// QueryOTAFirmwareRequest is the request struct for api QueryOTAFirmware
type QueryOTAFirmwareRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	FirmwareId        string `position:"Query" name:"FirmwareId"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// QueryOTAFirmwareResponse is the response struct for api QueryOTAFirmware
type QueryOTAFirmwareResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      bool         `json:"Success" xml:"Success"`
	Code         string       `json:"Code" xml:"Code"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	FirmwareInfo FirmwareInfo `json:"FirmwareInfo" xml:"FirmwareInfo"`
}

// CreateQueryOTAFirmwareRequest creates a request to invoke QueryOTAFirmware API
func CreateQueryOTAFirmwareRequest() (request *QueryOTAFirmwareRequest) {
	request = &QueryOTAFirmwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryOTAFirmware", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryOTAFirmwareResponse creates a response to parse from QueryOTAFirmware response
func CreateQueryOTAFirmwareResponse() (response *QueryOTAFirmwareResponse) {
	response = &QueryOTAFirmwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
