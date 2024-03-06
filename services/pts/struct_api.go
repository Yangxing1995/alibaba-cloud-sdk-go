package pts

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

// Api is a nested struct in pts response
type Api struct {
	ApiId              string       `json:"ApiId" xml:"ApiId"`
	TimeoutInSecond    int          `json:"TimeoutInSecond" xml:"TimeoutInSecond"`
	Method             string       `json:"Method" xml:"Method"`
	RedirectCountLimit int          `json:"RedirectCountLimit" xml:"RedirectCountLimit"`
	ApiName            string       `json:"ApiName" xml:"ApiName"`
	Url                string       `json:"Url" xml:"Url"`
	Body               Body         `json:"Body" xml:"Body"`
	ExportList         []Export     `json:"ExportList" xml:"ExportList"`
	CheckPointList     []CheckPoint `json:"CheckPointList" xml:"CheckPointList"`
	HeaderList         []Header     `json:"HeaderList" xml:"HeaderList"`
}
