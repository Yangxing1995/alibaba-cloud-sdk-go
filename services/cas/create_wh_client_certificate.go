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

// CreateWHClientCertificate invokes the cas.CreateWHClientCertificate API synchronously
func (client *Client) CreateWHClientCertificate(request *CreateWHClientCertificateRequest) (response *CreateWHClientCertificateResponse, err error) {
	response = CreateCreateWHClientCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWHClientCertificateWithChan invokes the cas.CreateWHClientCertificate API asynchronously
func (client *Client) CreateWHClientCertificateWithChan(request *CreateWHClientCertificateRequest) (<-chan *CreateWHClientCertificateResponse, <-chan error) {
	responseChan := make(chan *CreateWHClientCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWHClientCertificate(request)
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

// CreateWHClientCertificateWithCallback invokes the cas.CreateWHClientCertificate API asynchronously
func (client *Client) CreateWHClientCertificateWithCallback(request *CreateWHClientCertificateRequest, callback func(response *CreateWHClientCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWHClientCertificateResponse
		var err error
		defer close(result)
		response, err = client.CreateWHClientCertificate(request)
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

// CreateWHClientCertificateRequest is the request struct for api CreateWHClientCertificate
type CreateWHClientCertificateRequest struct {
	*requests.RpcRequest
	Country          string           `position:"Query" name:"Country"`
	Csr              string           `position:"Query" name:"Csr"`
	Immediately      requests.Integer `position:"Query" name:"Immediately"`
	Years            requests.Integer `position:"Query" name:"Years"`
	CommonName       string           `position:"Query" name:"CommonName"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	SanValue         string           `position:"Query" name:"SanValue"`
	State            string           `position:"Query" name:"State"`
	Algorithm        string           `position:"Query" name:"Algorithm"`
	Months           requests.Integer `position:"Query" name:"Months"`
	AfterTime        requests.Integer `position:"Query" name:"AfterTime"`
	Locality         string           `position:"Query" name:"Locality"`
	SanType          requests.Integer `position:"Query" name:"SanType"`
	Organization     string           `position:"Query" name:"Organization"`
	Days             requests.Integer `position:"Query" name:"Days"`
	BeforeTime       requests.Integer `position:"Query" name:"BeforeTime"`
	ParentIdentifier string           `position:"Query" name:"ParentIdentifier"`
	OrganizationUnit string           `position:"Query" name:"OrganizationUnit"`
}

// CreateWHClientCertificateResponse is the response struct for api CreateWHClientCertificate
type CreateWHClientCertificateResponse struct {
	*responses.BaseResponse
	Identifier            string `json:"Identifier" xml:"Identifier"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	RootX509Certificate   string `json:"RootX509Certificate" xml:"RootX509Certificate"`
	ParentX509Certificate string `json:"ParentX509Certificate" xml:"ParentX509Certificate"`
	X509Certificate       string `json:"X509Certificate" xml:"X509Certificate"`
	CertificateChain      string `json:"CertificateChain" xml:"CertificateChain"`
}

// CreateCreateWHClientCertificateRequest creates a request to invoke CreateWHClientCertificate API
func CreateCreateWHClientCertificateRequest() (request *CreateWHClientCertificateRequest) {
	request = &CreateWHClientCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "CreateWHClientCertificate", "cas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateWHClientCertificateResponse creates a response to parse from CreateWHClientCertificate response
func CreateCreateWHClientCertificateResponse() (response *CreateWHClientCertificateResponse) {
	response = &CreateWHClientCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
