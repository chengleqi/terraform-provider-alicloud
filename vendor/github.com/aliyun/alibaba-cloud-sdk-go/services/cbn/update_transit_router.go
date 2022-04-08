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

// UpdateTransitRouter invokes the cbn.UpdateTransitRouter API synchronously
func (client *Client) UpdateTransitRouter(request *UpdateTransitRouterRequest) (response *UpdateTransitRouterResponse, err error) {
	response = CreateUpdateTransitRouterResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTransitRouterWithChan invokes the cbn.UpdateTransitRouter API asynchronously
func (client *Client) UpdateTransitRouterWithChan(request *UpdateTransitRouterRequest) (<-chan *UpdateTransitRouterResponse, <-chan error) {
	responseChan := make(chan *UpdateTransitRouterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTransitRouter(request)
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

// UpdateTransitRouterWithCallback invokes the cbn.UpdateTransitRouter API asynchronously
func (client *Client) UpdateTransitRouterWithCallback(request *UpdateTransitRouterRequest, callback func(response *UpdateTransitRouterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTransitRouterResponse
		var err error
		defer close(result)
		response, err = client.UpdateTransitRouter(request)
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

// UpdateTransitRouterRequest is the request struct for api UpdateTransitRouter
type UpdateTransitRouterRequest struct {
	*requests.RpcRequest
	TransitRouterName        string           `position:"Query" name:"TransitRouterName"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken              string           `position:"Query" name:"ClientToken"`
	DryRun                   requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	TransitRouterId          string           `position:"Query" name:"TransitRouterId"`
	TransitRouterDescription string           `position:"Query" name:"TransitRouterDescription"`
}

// UpdateTransitRouterResponse is the response struct for api UpdateTransitRouter
type UpdateTransitRouterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateTransitRouterRequest creates a request to invoke UpdateTransitRouter API
func CreateUpdateTransitRouterRequest() (request *UpdateTransitRouterRequest) {
	request = &UpdateTransitRouterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "UpdateTransitRouter", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTransitRouterResponse creates a response to parse from UpdateTransitRouter response
func CreateUpdateTransitRouterResponse() (response *UpdateTransitRouterResponse) {
	response = &UpdateTransitRouterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
