package vs

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

// CreateTemplate invokes the vs.CreateTemplate API synchronously
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	response = CreateCreateTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTemplateWithChan invokes the vs.CreateTemplate API asynchronously
func (client *Client) CreateTemplateWithChan(request *CreateTemplateRequest) (<-chan *CreateTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTemplate(request)
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

// CreateTemplateWithCallback invokes the vs.CreateTemplate API asynchronously
func (client *Client) CreateTemplateWithCallback(request *CreateTemplateRequest, callback func(response *CreateTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateTemplate(request)
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

// CreateTemplateRequest is the request struct for api CreateTemplate
type CreateTemplateRequest struct {
	*requests.RpcRequest
	OssEndpoint      string           `position:"Query" name:"OssEndpoint"`
	JpgOverwrite     string           `position:"Query" name:"JpgOverwrite"`
	StartTime        string           `position:"Query" name:"StartTime"`
	Type             string           `position:"Query" name:"Type"`
	JpgOnDemand      string           `position:"Query" name:"JpgOnDemand"`
	ShowLog          string           `position:"Query" name:"ShowLog"`
	OssBucket        string           `position:"Query" name:"OssBucket"`
	TransConfigsJSON string           `position:"Query" name:"TransConfigsJSON"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	Name             string           `position:"Query" name:"Name"`
	Interval         requests.Integer `position:"Query" name:"Interval"`
	Region           string           `position:"Query" name:"Region"`
	HlsTs            string           `position:"Query" name:"HlsTs"`
	Description      string           `position:"Query" name:"Description"`
	OssFilePrefix    string           `position:"Query" name:"OssFilePrefix"`
	Retention        requests.Integer `position:"Query" name:"Retention"`
	HlsM3u8          string           `position:"Query" name:"HlsM3u8"`
	EndTime          string           `position:"Query" name:"EndTime"`
	Trigger          string           `position:"Query" name:"Trigger"`
	JpgSequence      string           `position:"Query" name:"JpgSequence"`
	Mp4              string           `position:"Query" name:"Mp4"`
	Flv              string           `position:"Query" name:"Flv"`
	Callback         string           `position:"Query" name:"Callback"`
	FileFormat       string           `position:"Query" name:"FileFormat"`
}

// CreateTemplateResponse is the response struct for api CreateTemplate
type CreateTemplateResponse struct {
	*responses.BaseResponse
	Id        string `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTemplateRequest creates a request to invoke CreateTemplate API
func CreateCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "CreateTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTemplateResponse creates a response to parse from CreateTemplate response
func CreateCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
