/* 
 * IronFunctions
 *
 * null
 *
 * OpenAPI spec version: 0.0.3
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package functions

import (
	"net/url"
	"encoding/json"
)

type TasksApi struct {
	Configuration Configuration
}

func NewTasksApi() *TasksApi {
	configuration := NewConfiguration()
	return &TasksApi{
		Configuration: *configuration,
	}
}

func NewTasksApiWithBasePath(basePath string) *TasksApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &TasksApi{
		Configuration: *configuration,
	}
}

/**
 * Get next task.
 * Gets the next task in the queue, ready for processing. Titan may return &lt;&#x3D;n tasks. Consumers should start processing tasks in order. Each returned task is set to &#x60;status&#x60; \&quot;running\&quot; and &#x60;started_at&#x60; is set to the current time. No other consumer can retrieve this task.
 *
 * @param n Number of tasks to return.
 * @return *TasksWrapper
 */
func (a TasksApi) TasksGet(n int32) (*TasksWrapper, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/tasks"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("n", a.Configuration.APIClient.ParameterToString(n, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TasksWrapper)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

