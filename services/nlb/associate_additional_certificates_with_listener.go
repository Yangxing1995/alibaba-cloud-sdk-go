package nlb

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

// AssociateAdditionalCertificatesWithListener invokes the nlb.AssociateAdditionalCertificatesWithListener API synchronously
func (client *Client) AssociateAdditionalCertificatesWithListener(request *AssociateAdditionalCertificatesWithListenerRequest) (response *AssociateAdditionalCertificatesWithListenerResponse, err error) {
	response = CreateAssociateAdditionalCertificatesWithListenerResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateAdditionalCertificatesWithListenerWithChan invokes the nlb.AssociateAdditionalCertificatesWithListener API asynchronously
func (client *Client) AssociateAdditionalCertificatesWithListenerWithChan(request *AssociateAdditionalCertificatesWithListenerRequest) (<-chan *AssociateAdditionalCertificatesWithListenerResponse, <-chan error) {
	responseChan := make(chan *AssociateAdditionalCertificatesWithListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateAdditionalCertificatesWithListener(request)
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

// AssociateAdditionalCertificatesWithListenerWithCallback invokes the nlb.AssociateAdditionalCertificatesWithListener API asynchronously
func (client *Client) AssociateAdditionalCertificatesWithListenerWithCallback(request *AssociateAdditionalCertificatesWithListenerRequest, callback func(response *AssociateAdditionalCertificatesWithListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateAdditionalCertificatesWithListenerResponse
		var err error
		defer close(result)
		response, err = client.AssociateAdditionalCertificatesWithListener(request)
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

// AssociateAdditionalCertificatesWithListenerRequest is the request struct for api AssociateAdditionalCertificatesWithListener
type AssociateAdditionalCertificatesWithListenerRequest struct {
	*requests.RpcRequest
	ClientToken              string           `position:"Body" name:"ClientToken"`
	ListenerId               string           `position:"Body" name:"ListenerId"`
	DryRun                   requests.Boolean `position:"Body" name:"DryRun"`
	AdditionalCertificateIds *[]string        `position:"Body" name:"AdditionalCertificateIds"  type:"Repeated"`
}

// AssociateAdditionalCertificatesWithListenerResponse is the response struct for api AssociateAdditionalCertificatesWithListener
type AssociateAdditionalCertificatesWithListenerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	JobId          string `json:"JobId" xml:"JobId"`
}

// CreateAssociateAdditionalCertificatesWithListenerRequest creates a request to invoke AssociateAdditionalCertificatesWithListener API
func CreateAssociateAdditionalCertificatesWithListenerRequest() (request *AssociateAdditionalCertificatesWithListenerRequest) {
	request = &AssociateAdditionalCertificatesWithListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "AssociateAdditionalCertificatesWithListener", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateAdditionalCertificatesWithListenerResponse creates a response to parse from AssociateAdditionalCertificatesWithListener response
func CreateAssociateAdditionalCertificatesWithListenerResponse() (response *AssociateAdditionalCertificatesWithListenerResponse) {
	response = &AssociateAdditionalCertificatesWithListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
