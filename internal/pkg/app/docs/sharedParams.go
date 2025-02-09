// Copyright 2021 Comcast Cable Communications Management, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docs

// swagger:parameters putRoute postRoute getRoute deleteRoute putTenant getTenant deleteTenant postRouteEvent
type appIdParamWrapper struct {
	// App ID
	// in: path
	// required: true
	AppId string `json:"appId"`
}

// swagger:parameters putRoute postRoute getRoute deleteRoute putTenant getTenant deleteTenant postRouteEvent
type orgIdParamWrapper struct {
	// Org ID
	// in: path
	// required: true
	OrgId string `json:"orgId"`
}
