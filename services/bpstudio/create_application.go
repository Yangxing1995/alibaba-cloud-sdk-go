package bpstudio

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

// CreateApplication invokes the bpstudio.CreateApplication API synchronously
func (client *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
	response = CreateCreateApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApplicationWithChan invokes the bpstudio.CreateApplication API asynchronously
func (client *Client) CreateApplicationWithChan(request *CreateApplicationRequest) (<-chan *CreateApplicationResponse, <-chan error) {
	responseChan := make(chan *CreateApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApplication(request)
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

// CreateApplicationWithCallback invokes the bpstudio.CreateApplication API asynchronously
func (client *Client) CreateApplicationWithCallback(request *CreateApplicationRequest, callback func(response *CreateApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApplicationResponse
		var err error
		defer close(result)
		response, err = client.CreateApplication(request)
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

// CreateApplicationRequest is the request struct for api CreateApplication
type CreateApplicationRequest struct {
	*requests.RpcRequest
	Variables       map[string]string             `position:"Body" name:"Variables"  type:"Map"`
	ClientToken     string                        `position:"Body" name:"ClientToken"`
	Instances       *[]CreateApplicationInstances `position:"Body" name:"Instances"  type:"Json"`
	Configuration   map[string]string             `position:"Body" name:"Configuration"  type:"Map"`
	TemplateId      string                        `position:"Body" name:"TemplateId"`
	ResourceGroupId string                        `position:"Body" name:"ResourceGroupId"`
	AreaId          string                        `position:"Body" name:"AreaId"`
	Name            string                        `position:"Body" name:"Name"`
}

// CreateApplicationInstances is a repeated param struct in CreateApplicationRequest
type CreateApplicationInstances struct {
	NodeName string `name:"NodeName"`
	NodeType string `name:"NodeType"`
	Id       string `name:"Id"`
}

// CreateApplicationResponse is the response struct for api CreateApplication
type CreateApplicationResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateCreateApplicationRequest creates a request to invoke CreateApplication API
func CreateCreateApplicationRequest() (request *CreateApplicationRequest) {
	request = &CreateApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BPStudio", "2021-09-31", "CreateApplication", "bpstudio", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateApplicationResponse creates a response to parse from CreateApplication response
func CreateCreateApplicationResponse() (response *CreateApplicationResponse) {
	response = &CreateApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
