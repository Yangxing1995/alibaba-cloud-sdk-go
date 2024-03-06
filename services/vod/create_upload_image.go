package vod

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

// CreateUploadImage invokes the vod.CreateUploadImage API synchronously
func (client *Client) CreateUploadImage(request *CreateUploadImageRequest) (response *CreateUploadImageResponse, err error) {
	response = CreateCreateUploadImageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUploadImageWithChan invokes the vod.CreateUploadImage API asynchronously
func (client *Client) CreateUploadImageWithChan(request *CreateUploadImageRequest) (<-chan *CreateUploadImageResponse, <-chan error) {
	responseChan := make(chan *CreateUploadImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUploadImage(request)
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

// CreateUploadImageWithCallback invokes the vod.CreateUploadImage API asynchronously
func (client *Client) CreateUploadImageWithCallback(request *CreateUploadImageRequest, callback func(response *CreateUploadImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUploadImageResponse
		var err error
		defer close(result)
		response, err = client.CreateUploadImage(request)
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

// CreateUploadImageRequest is the request struct for api CreateUploadImage
type CreateUploadImageRequest struct {
	*requests.RpcRequest
	Description      string           `position:"Query" name:"Description"`
	Title            string           `position:"Query" name:"Title"`
	StorageLocation  string           `position:"Query" name:"StorageLocation"`
	UserData         string           `position:"Query" name:"UserData"`
	CateId           requests.Integer `position:"Query" name:"CateId"`
	ImageType        string           `position:"Query" name:"ImageType"`
	ImageExt         string           `position:"Query" name:"ImageExt"`
	Tags             string           `position:"Query" name:"Tags"`
	OriginalFileName string           `position:"Query" name:"OriginalFileName"`
	AppId            string           `position:"Query" name:"AppId"`
}

// CreateUploadImageResponse is the response struct for api CreateUploadImage
type CreateUploadImageResponse struct {
	*responses.BaseResponse
	FileURL       string `json:"FileURL" xml:"FileURL"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	UploadAddress string `json:"UploadAddress" xml:"UploadAddress"`
	ImageURL      string `json:"ImageURL" xml:"ImageURL"`
	ImageId       string `json:"ImageId" xml:"ImageId"`
	UploadAuth    string `json:"UploadAuth" xml:"UploadAuth"`
}

// CreateCreateUploadImageRequest creates a request to invoke CreateUploadImage API
func CreateCreateUploadImageRequest() (request *CreateUploadImageRequest) {
	request = &CreateUploadImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "CreateUploadImage", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateUploadImageResponse creates a response to parse from CreateUploadImage response
func CreateCreateUploadImageResponse() (response *CreateUploadImageResponse) {
	response = &CreateUploadImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
