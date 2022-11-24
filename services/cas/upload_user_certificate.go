package cas

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

// UploadUserCertificate invokes the cas.UploadUserCertificate API synchronously
func (client *Client) UploadUserCertificate(request *UploadUserCertificateRequest) (response *UploadUserCertificateResponse, err error) {
	response = CreateUploadUserCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// UploadUserCertificateWithChan invokes the cas.UploadUserCertificate API asynchronously
func (client *Client) UploadUserCertificateWithChan(request *UploadUserCertificateRequest) (<-chan *UploadUserCertificateResponse, <-chan error) {
	responseChan := make(chan *UploadUserCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadUserCertificate(request)
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

// UploadUserCertificateWithCallback invokes the cas.UploadUserCertificate API asynchronously
func (client *Client) UploadUserCertificateWithCallback(request *UploadUserCertificateRequest, callback func(response *UploadUserCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadUserCertificateResponse
		var err error
		defer close(result)
		response, err = client.UploadUserCertificate(request)
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

// UploadUserCertificateRequest is the request struct for api UploadUserCertificate
type UploadUserCertificateRequest struct {
	*requests.RpcRequest
	EncryptCert       string `position:"Query" name:"EncryptCert"`
	Cert              string `position:"Query" name:"Cert"`
	SourceIp          string `position:"Query" name:"SourceIp"`
	Key               string `position:"Query" name:"Key"`
	EncryptPrivateKey string `position:"Query" name:"EncryptPrivateKey"`
	SignPrivateKey    string `position:"Query" name:"SignPrivateKey"`
	SignCert          string `position:"Query" name:"SignCert"`
	Name              string `position:"Query" name:"Name"`
}

// UploadUserCertificateResponse is the response struct for api UploadUserCertificate
type UploadUserCertificateResponse struct {
	*responses.BaseResponse
	CertId    int64  `json:"CertId" xml:"CertId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUploadUserCertificateRequest creates a request to invoke UploadUserCertificate API
func CreateUploadUserCertificateRequest() (request *UploadUserCertificateRequest) {
	request = &UploadUserCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "UploadUserCertificate", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadUserCertificateResponse creates a response to parse from UploadUserCertificate response
func CreateUploadUserCertificateResponse() (response *UploadUserCertificateResponse) {
	response = &UploadUserCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
