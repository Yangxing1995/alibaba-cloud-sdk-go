package cbn

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

// ListTransitRouterMulticastGroups invokes the cbn.ListTransitRouterMulticastGroups API synchronously
func (client *Client) ListTransitRouterMulticastGroups(request *ListTransitRouterMulticastGroupsRequest) (response *ListTransitRouterMulticastGroupsResponse, err error) {
	response = CreateListTransitRouterMulticastGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterMulticastGroupsWithChan invokes the cbn.ListTransitRouterMulticastGroups API asynchronously
func (client *Client) ListTransitRouterMulticastGroupsWithChan(request *ListTransitRouterMulticastGroupsRequest) (<-chan *ListTransitRouterMulticastGroupsResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterMulticastGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterMulticastGroups(request)
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

// ListTransitRouterMulticastGroupsWithCallback invokes the cbn.ListTransitRouterMulticastGroups API asynchronously
func (client *Client) ListTransitRouterMulticastGroupsWithCallback(request *ListTransitRouterMulticastGroupsRequest, callback func(response *ListTransitRouterMulticastGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterMulticastGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterMulticastGroups(request)
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

// ListTransitRouterMulticastGroupsRequest is the request struct for api ListTransitRouterMulticastGroups
type ListTransitRouterMulticastGroupsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                   requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                       string           `position:"Query" name:"ClientToken"`
	NetworkInterfaceIds               *[]string        `position:"Query" name:"NetworkInterfaceIds"  type:"Repeated"`
	VSwitchIds                        *[]string        `position:"Query" name:"VSwitchIds"  type:"Repeated"`
	TransitRouterMulticastDomainId    string           `position:"Query" name:"TransitRouterMulticastDomainId"`
	IsGroupSource                     requests.Boolean `position:"Query" name:"IsGroupSource"`
	ConnectPeerIds                    *[]string        `position:"Query" name:"ConnectPeerIds"  type:"Repeated"`
	NextToken                         string           `position:"Query" name:"NextToken"`
	GroupIpAddress                    string           `position:"Query" name:"GroupIpAddress"`
	ResourceId                        string           `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount              string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                      string           `position:"Query" name:"OwnerAccount"`
	PeerTransitRouterMulticastDomains *[]string        `position:"Query" name:"PeerTransitRouterMulticastDomains"  type:"Repeated"`
	OwnerId                           requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType                      string           `position:"Query" name:"ResourceType"`
	Version                           string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId         string           `position:"Query" name:"TransitRouterAttachmentId"`
	MaxResults                        requests.Integer `position:"Query" name:"MaxResults"`
	IsGroupMember                     requests.Boolean `position:"Query" name:"IsGroupMember"`
}

// ListTransitRouterMulticastGroupsResponse is the response struct for api ListTransitRouterMulticastGroups
type ListTransitRouterMulticastGroupsResponse struct {
	*responses.BaseResponse
	RequestId                    string                        `json:"RequestId" xml:"RequestId"`
	TotalCount                   int                           `json:"TotalCount" xml:"TotalCount"`
	MaxResults                   int                           `json:"MaxResults" xml:"MaxResults"`
	NextToken                    string                        `json:"NextToken" xml:"NextToken"`
	TransitRouterMulticastGroups []TransitRouterMulticastGroup `json:"TransitRouterMulticastGroups" xml:"TransitRouterMulticastGroups"`
}

// CreateListTransitRouterMulticastGroupsRequest creates a request to invoke ListTransitRouterMulticastGroups API
func CreateListTransitRouterMulticastGroupsRequest() (request *ListTransitRouterMulticastGroupsRequest) {
	request = &ListTransitRouterMulticastGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterMulticastGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterMulticastGroupsResponse creates a response to parse from ListTransitRouterMulticastGroups response
func CreateListTransitRouterMulticastGroupsResponse() (response *ListTransitRouterMulticastGroupsResponse) {
	response = &ListTransitRouterMulticastGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
