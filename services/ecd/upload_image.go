package ecd

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

// UploadImage invokes the ecd.UploadImage API synchronously
func (client *Client) UploadImage(request *UploadImageRequest) (response *UploadImageResponse, err error) {
	response = CreateUploadImageResponse()
	err = client.DoAction(request, response)
	return
}

// UploadImageWithChan invokes the ecd.UploadImage API asynchronously
func (client *Client) UploadImageWithChan(request *UploadImageRequest) (<-chan *UploadImageResponse, <-chan error) {
	responseChan := make(chan *UploadImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadImage(request)
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

// UploadImageWithCallback invokes the ecd.UploadImage API asynchronously
func (client *Client) UploadImageWithCallback(request *UploadImageRequest, callback func(response *UploadImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadImageResponse
		var err error
		defer close(result)
		response, err = client.UploadImage(request)
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

// UploadImageRequest is the request struct for api UploadImage
type UploadImageRequest struct {
	*requests.RpcRequest
	EnableSecurityCheck requests.Boolean `position:"Query" name:"EnableSecurityCheck"`
	GpuCategory         requests.Boolean `position:"Query" name:"GpuCategory"`
	Description         string           `position:"Query" name:"Description"`
	OssObjectPath       string           `position:"Query" name:"OssObjectPath"`
	ImageName           string           `position:"Query" name:"ImageName"`
	LicenseType         string           `position:"Query" name:"LicenseType"`
	OsType              string           `position:"Query" name:"OsType"`
	DataDiskSize        requests.Integer `position:"Query" name:"DataDiskSize"`
	ProtocolType        string           `position:"Query" name:"ProtocolType"`
	GpuDriverType       string           `position:"Query" name:"GpuDriverType"`
}

// UploadImageResponse is the response struct for api UploadImage
type UploadImageResponse struct {
	*responses.BaseResponse
	ImageId   string `json:"ImageId" xml:"ImageId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUploadImageRequest creates a request to invoke UploadImage API
func CreateUploadImageRequest() (request *UploadImageRequest) {
	request = &UploadImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "UploadImage", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadImageResponse creates a response to parse from UploadImage response
func CreateUploadImageResponse() (response *UploadImageResponse) {
	response = &UploadImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
