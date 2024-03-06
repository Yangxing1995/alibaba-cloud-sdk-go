package paifeaturestore

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

// UpdateModelFeatureFGFeature invokes the paifeaturestore.UpdateModelFeatureFGFeature API synchronously
func (client *Client) UpdateModelFeatureFGFeature(request *UpdateModelFeatureFGFeatureRequest) (response *UpdateModelFeatureFGFeatureResponse, err error) {
	response = CreateUpdateModelFeatureFGFeatureResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateModelFeatureFGFeatureWithChan invokes the paifeaturestore.UpdateModelFeatureFGFeature API asynchronously
func (client *Client) UpdateModelFeatureFGFeatureWithChan(request *UpdateModelFeatureFGFeatureRequest) (<-chan *UpdateModelFeatureFGFeatureResponse, <-chan error) {
	responseChan := make(chan *UpdateModelFeatureFGFeatureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateModelFeatureFGFeature(request)
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

// UpdateModelFeatureFGFeatureWithCallback invokes the paifeaturestore.UpdateModelFeatureFGFeature API asynchronously
func (client *Client) UpdateModelFeatureFGFeatureWithCallback(request *UpdateModelFeatureFGFeatureRequest, callback func(response *UpdateModelFeatureFGFeatureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateModelFeatureFGFeatureResponse
		var err error
		defer close(result)
		response, err = client.UpdateModelFeatureFGFeature(request)
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

// UpdateModelFeatureFGFeatureRequest is the request struct for api UpdateModelFeatureFGFeature
type UpdateModelFeatureFGFeatureRequest struct {
	*requests.RoaRequest
	ModelFeatureId string `position:"Path" name:"ModelFeatureId"`
	Body           string `position:"Body" name:"body"`
	InstanceId     string `position:"Path" name:"InstanceId"`
}

// UpdateModelFeatureFGFeatureResponse is the response struct for api UpdateModelFeatureFGFeature
type UpdateModelFeatureFGFeatureResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateUpdateModelFeatureFGFeatureRequest creates a request to invoke UpdateModelFeatureFGFeature API
func CreateUpdateModelFeatureFGFeatureRequest() (request *UpdateModelFeatureFGFeatureRequest) {
	request = &UpdateModelFeatureFGFeatureRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "UpdateModelFeatureFGFeature", "/api/v1/instances/[InstanceId]/modelfeatures/[ModelFeatureId]/fgfeature", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateModelFeatureFGFeatureResponse creates a response to parse from UpdateModelFeatureFGFeature response
func CreateUpdateModelFeatureFGFeatureResponse() (response *UpdateModelFeatureFGFeatureResponse) {
	response = &UpdateModelFeatureFGFeatureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
