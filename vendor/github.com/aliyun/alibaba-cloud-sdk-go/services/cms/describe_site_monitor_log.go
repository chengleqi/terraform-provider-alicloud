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

// DescribeSiteMonitorLog invokes the cms.DescribeSiteMonitorLog API synchronously
func (client *Client) DescribeSiteMonitorLog(request *DescribeSiteMonitorLogRequest) (response *DescribeSiteMonitorLogResponse, err error) {
	response = CreateDescribeSiteMonitorLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSiteMonitorLogWithChan invokes the cms.DescribeSiteMonitorLog API asynchronously
func (client *Client) DescribeSiteMonitorLogWithChan(request *DescribeSiteMonitorLogRequest) (<-chan *DescribeSiteMonitorLogResponse, <-chan error) {
	responseChan := make(chan *DescribeSiteMonitorLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSiteMonitorLog(request)
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

// DescribeSiteMonitorLogWithCallback invokes the cms.DescribeSiteMonitorLog API asynchronously
func (client *Client) DescribeSiteMonitorLogWithCallback(request *DescribeSiteMonitorLogRequest, callback func(response *DescribeSiteMonitorLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSiteMonitorLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeSiteMonitorLog(request)
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

// DescribeSiteMonitorLogRequest is the request struct for api DescribeSiteMonitorLog
type DescribeSiteMonitorLogRequest struct {
	*requests.RpcRequest
	City       string           `position:"Query" name:"City"`
	Isp        string           `position:"Query" name:"Isp"`
	StartTime  string           `position:"Query" name:"StartTime"`
	TaskIds    string           `position:"Query" name:"TaskIds"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MetricName string           `position:"Query" name:"MetricName"`
	Length     requests.Integer `position:"Query" name:"Length"`
	EndTime    string           `position:"Query" name:"EndTime"`
	Filter     string           `position:"Query" name:"Filter"`
	Dimensions string           `position:"Query" name:"Dimensions"`
}

// DescribeSiteMonitorLogResponse is the response struct for api DescribeSiteMonitorLog
type DescribeSiteMonitorLogResponse struct {
	*responses.BaseResponse
	NextToken string `json:"NextToken" xml:"NextToken"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDescribeSiteMonitorLogRequest creates a request to invoke DescribeSiteMonitorLog API
func CreateDescribeSiteMonitorLogRequest() (request *DescribeSiteMonitorLogRequest) {
	request = &DescribeSiteMonitorLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSiteMonitorLog", "Cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSiteMonitorLogResponse creates a response to parse from DescribeSiteMonitorLog response
func CreateDescribeSiteMonitorLogResponse() (response *DescribeSiteMonitorLogResponse) {
	response = &DescribeSiteMonitorLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
