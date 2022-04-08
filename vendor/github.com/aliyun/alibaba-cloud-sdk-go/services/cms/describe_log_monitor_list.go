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

// DescribeLogMonitorList invokes the cms.DescribeLogMonitorList API synchronously
func (client *Client) DescribeLogMonitorList(request *DescribeLogMonitorListRequest) (response *DescribeLogMonitorListResponse, err error) {
	response = CreateDescribeLogMonitorListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogMonitorListWithChan invokes the cms.DescribeLogMonitorList API asynchronously
func (client *Client) DescribeLogMonitorListWithChan(request *DescribeLogMonitorListRequest) (<-chan *DescribeLogMonitorListResponse, <-chan error) {
	responseChan := make(chan *DescribeLogMonitorListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogMonitorList(request)
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

// DescribeLogMonitorListWithCallback invokes the cms.DescribeLogMonitorList API asynchronously
func (client *Client) DescribeLogMonitorListWithCallback(request *DescribeLogMonitorListRequest, callback func(response *DescribeLogMonitorListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogMonitorListResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogMonitorList(request)
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

// DescribeLogMonitorListRequest is the request struct for api DescribeLogMonitorList
type DescribeLogMonitorListRequest struct {
	*requests.RpcRequest
	GroupId     requests.Integer `position:"Query" name:"GroupId"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	SearchValue string           `position:"Query" name:"SearchValue"`
}

// DescribeLogMonitorListResponse is the response struct for api DescribeLogMonitorList
type DescribeLogMonitorListResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	PageSize       int          `json:"PageSize" xml:"PageSize"`
	PageNumber     int          `json:"PageNumber" xml:"PageNumber"`
	Total          int64        `json:"Total" xml:"Total"`
	LogMonitorList []LogMonitor `json:"LogMonitorList" xml:"LogMonitorList"`
}

// CreateDescribeLogMonitorListRequest creates a request to invoke DescribeLogMonitorList API
func CreateDescribeLogMonitorListRequest() (request *DescribeLogMonitorListRequest) {
	request = &DescribeLogMonitorListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeLogMonitorList", "Cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLogMonitorListResponse creates a response to parse from DescribeLogMonitorList response
func CreateDescribeLogMonitorListResponse() (response *DescribeLogMonitorListResponse) {
	response = &DescribeLogMonitorListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
