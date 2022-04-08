package cms

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

// SendDryRunSystemEvent invokes the cms.SendDryRunSystemEvent API synchronously
func (client *Client) SendDryRunSystemEvent(request *SendDryRunSystemEventRequest) (response *SendDryRunSystemEventResponse, err error) {
	response = CreateSendDryRunSystemEventResponse()
	err = client.DoAction(request, response)
	return
}

// SendDryRunSystemEventWithChan invokes the cms.SendDryRunSystemEvent API asynchronously
func (client *Client) SendDryRunSystemEventWithChan(request *SendDryRunSystemEventRequest) (<-chan *SendDryRunSystemEventResponse, <-chan error) {
	responseChan := make(chan *SendDryRunSystemEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendDryRunSystemEvent(request)
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

// SendDryRunSystemEventWithCallback invokes the cms.SendDryRunSystemEvent API asynchronously
func (client *Client) SendDryRunSystemEventWithCallback(request *SendDryRunSystemEventRequest, callback func(response *SendDryRunSystemEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendDryRunSystemEventResponse
		var err error
		defer close(result)
		response, err = client.SendDryRunSystemEvent(request)
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

// SendDryRunSystemEventRequest is the request struct for api SendDryRunSystemEvent
type SendDryRunSystemEventRequest struct {
	*requests.RpcRequest
	Product      string `position:"Query" name:"Product"`
	GroupId      string `position:"Query" name:"GroupId"`
	EventName    string `position:"Query" name:"EventName"`
	EventContent string `position:"Query" name:"EventContent"`
}

// SendDryRunSystemEventResponse is the response struct for api SendDryRunSystemEvent
type SendDryRunSystemEventResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateSendDryRunSystemEventRequest creates a request to invoke SendDryRunSystemEvent API
func CreateSendDryRunSystemEventRequest() (request *SendDryRunSystemEventRequest) {
	request = &SendDryRunSystemEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "SendDryRunSystemEvent", "Cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendDryRunSystemEventResponse creates a response to parse from SendDryRunSystemEvent response
func CreateSendDryRunSystemEventResponse() (response *SendDryRunSystemEventResponse) {
	response = &SendDryRunSystemEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
