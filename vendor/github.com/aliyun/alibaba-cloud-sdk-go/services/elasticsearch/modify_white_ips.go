package elasticsearch

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

// ModifyWhiteIps invokes the elasticsearch.ModifyWhiteIps API synchronously
func (client *Client) ModifyWhiteIps(request *ModifyWhiteIpsRequest) (response *ModifyWhiteIpsResponse, err error) {
	response = CreateModifyWhiteIpsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWhiteIpsWithChan invokes the elasticsearch.ModifyWhiteIps API asynchronously
func (client *Client) ModifyWhiteIpsWithChan(request *ModifyWhiteIpsRequest) (<-chan *ModifyWhiteIpsResponse, <-chan error) {
	responseChan := make(chan *ModifyWhiteIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWhiteIps(request)
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

// ModifyWhiteIpsWithCallback invokes the elasticsearch.ModifyWhiteIps API asynchronously
func (client *Client) ModifyWhiteIpsWithCallback(request *ModifyWhiteIpsRequest, callback func(response *ModifyWhiteIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWhiteIpsResponse
		var err error
		defer close(result)
		response, err = client.ModifyWhiteIps(request)
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

// ModifyWhiteIpsRequest is the request struct for api ModifyWhiteIps
type ModifyWhiteIpsRequest struct {
	*requests.RoaRequest
	ModifyMode  string `position:"Body" name:"modifyMode"`
	InstanceId  string `position:"Path" name:"InstanceId"`
	NodeType    string `position:"Body" name:"nodeType"`
	ClientToken string `position:"Query" name:"clientToken"`
	NetworkType string `position:"Body" name:"networkType"`
}

// ModifyWhiteIpsResponse is the response struct for api ModifyWhiteIps
type ModifyWhiteIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateModifyWhiteIpsRequest creates a request to invoke ModifyWhiteIps API
func CreateModifyWhiteIpsRequest() (request *ModifyWhiteIpsRequest) {
	request = &ModifyWhiteIpsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ModifyWhiteIps", "/openapi/instances/[InstanceId]/actions/modify-white-ips", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyWhiteIpsResponse creates a response to parse from ModifyWhiteIps response
func CreateModifyWhiteIpsResponse() (response *ModifyWhiteIpsResponse) {
	response = &ModifyWhiteIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
