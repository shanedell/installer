package cs

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

// DescribeExternalAgent invokes the cs.DescribeExternalAgent API synchronously
func (client *Client) DescribeExternalAgent(request *DescribeExternalAgentRequest) (response *DescribeExternalAgentResponse, err error) {
	response = CreateDescribeExternalAgentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExternalAgentWithChan invokes the cs.DescribeExternalAgent API asynchronously
func (client *Client) DescribeExternalAgentWithChan(request *DescribeExternalAgentRequest) (<-chan *DescribeExternalAgentResponse, <-chan error) {
	responseChan := make(chan *DescribeExternalAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExternalAgent(request)
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

// DescribeExternalAgentWithCallback invokes the cs.DescribeExternalAgent API asynchronously
func (client *Client) DescribeExternalAgentWithCallback(request *DescribeExternalAgentRequest, callback func(response *DescribeExternalAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExternalAgentResponse
		var err error
		defer close(result)
		response, err = client.DescribeExternalAgent(request)
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

// DescribeExternalAgentRequest is the request struct for api DescribeExternalAgent
type DescribeExternalAgentRequest struct {
	*requests.RoaRequest
	PrivateIpAddress string `position:"Query" name:"PrivateIpAddress"`
	ClusterId        string `position:"Path" name:"ClusterId"`
}

// DescribeExternalAgentResponse is the response struct for api DescribeExternalAgent
type DescribeExternalAgentResponse struct {
	*responses.BaseResponse
	Config string `json:"config" xml:"config"`
}

// CreateDescribeExternalAgentRequest creates a request to invoke DescribeExternalAgent API
func CreateDescribeExternalAgentRequest() (request *DescribeExternalAgentRequest) {
	request = &DescribeExternalAgentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeExternalAgent", "/k8s/[ClusterId]/external/agent/deployment", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeExternalAgentResponse creates a response to parse from DescribeExternalAgent response
func CreateDescribeExternalAgentResponse() (response *DescribeExternalAgentResponse) {
	response = &DescribeExternalAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}