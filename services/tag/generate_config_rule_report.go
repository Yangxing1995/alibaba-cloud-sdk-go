package tag

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

// GenerateConfigRuleReport invokes the tag.GenerateConfigRuleReport API synchronously
func (client *Client) GenerateConfigRuleReport(request *GenerateConfigRuleReportRequest) (response *GenerateConfigRuleReportResponse, err error) {
	response = CreateGenerateConfigRuleReportResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateConfigRuleReportWithChan invokes the tag.GenerateConfigRuleReport API asynchronously
func (client *Client) GenerateConfigRuleReportWithChan(request *GenerateConfigRuleReportRequest) (<-chan *GenerateConfigRuleReportResponse, <-chan error) {
	responseChan := make(chan *GenerateConfigRuleReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateConfigRuleReport(request)
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

// GenerateConfigRuleReportWithCallback invokes the tag.GenerateConfigRuleReport API asynchronously
func (client *Client) GenerateConfigRuleReportWithCallback(request *GenerateConfigRuleReportRequest, callback func(response *GenerateConfigRuleReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateConfigRuleReportResponse
		var err error
		defer close(result)
		response, err = client.GenerateConfigRuleReport(request)
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

// GenerateConfigRuleReportRequest is the request struct for api GenerateConfigRuleReport
type GenerateConfigRuleReportRequest struct {
	*requests.RpcRequest
	TargetId             string           `position:"Query" name:"TargetId"`
	TargetType           string           `position:"Query" name:"TargetType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UserType             string           `position:"Query" name:"UserType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GenerateConfigRuleReportResponse is the response struct for api GenerateConfigRuleReport
type GenerateConfigRuleReportResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ReportId  string `json:"ReportId" xml:"ReportId"`
}

// CreateGenerateConfigRuleReportRequest creates a request to invoke GenerateConfigRuleReport API
func CreateGenerateConfigRuleReportRequest() (request *GenerateConfigRuleReportRequest) {
	request = &GenerateConfigRuleReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "GenerateConfigRuleReport", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGenerateConfigRuleReportResponse creates a response to parse from GenerateConfigRuleReport response
func CreateGenerateConfigRuleReportResponse() (response *GenerateConfigRuleReportResponse) {
	response = &GenerateConfigRuleReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
