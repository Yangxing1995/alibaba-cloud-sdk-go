package sas

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

// DescribeAffectedMaliciousFileImages invokes the sas.DescribeAffectedMaliciousFileImages API synchronously
func (client *Client) DescribeAffectedMaliciousFileImages(request *DescribeAffectedMaliciousFileImagesRequest) (response *DescribeAffectedMaliciousFileImagesResponse, err error) {
	response = CreateDescribeAffectedMaliciousFileImagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAffectedMaliciousFileImagesWithChan invokes the sas.DescribeAffectedMaliciousFileImages API asynchronously
func (client *Client) DescribeAffectedMaliciousFileImagesWithChan(request *DescribeAffectedMaliciousFileImagesRequest) (<-chan *DescribeAffectedMaliciousFileImagesResponse, <-chan error) {
	responseChan := make(chan *DescribeAffectedMaliciousFileImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAffectedMaliciousFileImages(request)
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

// DescribeAffectedMaliciousFileImagesWithCallback invokes the sas.DescribeAffectedMaliciousFileImages API asynchronously
func (client *Client) DescribeAffectedMaliciousFileImagesWithCallback(request *DescribeAffectedMaliciousFileImagesRequest, callback func(response *DescribeAffectedMaliciousFileImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAffectedMaliciousFileImagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAffectedMaliciousFileImages(request)
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

// DescribeAffectedMaliciousFileImagesRequest is the request struct for api DescribeAffectedMaliciousFileImages
type DescribeAffectedMaliciousFileImagesRequest struct {
	*requests.RpcRequest
	RepoId         string           `position:"Query" name:"RepoId"`
	RepoNamespace  string           `position:"Query" name:"RepoNamespace"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	ImageDigest    string           `position:"Query" name:"ImageDigest"`
	PageSize       string           `position:"Query" name:"PageSize"`
	Lang           string           `position:"Query" name:"Lang"`
	ImageTag       string           `position:"Query" name:"ImageTag"`
	MaliciousMd5   string           `position:"Query" name:"MaliciousMd5"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId      string           `position:"Query" name:"ClusterId"`
	RepoName       string           `position:"Query" name:"RepoName"`
	RepoInstanceId string           `position:"Query" name:"RepoInstanceId"`
	ImageLayer     string           `position:"Query" name:"ImageLayer"`
	RepoRegionId   string           `position:"Query" name:"RepoRegionId"`
	Uuids          *[]string        `position:"Query" name:"Uuids"  type:"Repeated"`
}

// DescribeAffectedMaliciousFileImagesResponse is the response struct for api DescribeAffectedMaliciousFileImages
type DescribeAffectedMaliciousFileImagesResponse struct {
	*responses.BaseResponse
	RequestId                           string                       `json:"RequestId" xml:"RequestId"`
	PageInfo                            PageInfo                     `json:"PageInfo" xml:"PageInfo"`
	AffectedMaliciousFileImagesResponse []AffectedMaliciousFileImage `json:"AffectedMaliciousFileImagesResponse" xml:"AffectedMaliciousFileImagesResponse"`
}

// CreateDescribeAffectedMaliciousFileImagesRequest creates a request to invoke DescribeAffectedMaliciousFileImages API
func CreateDescribeAffectedMaliciousFileImagesRequest() (request *DescribeAffectedMaliciousFileImagesRequest) {
	request = &DescribeAffectedMaliciousFileImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeAffectedMaliciousFileImages", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAffectedMaliciousFileImagesResponse creates a response to parse from DescribeAffectedMaliciousFileImages response
func CreateDescribeAffectedMaliciousFileImagesResponse() (response *DescribeAffectedMaliciousFileImagesResponse) {
	response = &DescribeAffectedMaliciousFileImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
