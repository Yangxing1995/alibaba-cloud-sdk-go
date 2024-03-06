package live

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

// TextLayerContent is a nested struct in live response
type TextLayerContent struct {
	Color                 string  `json:"Color" xml:"Color"`
	BorderColor           string  `json:"BorderColor" xml:"BorderColor"`
	BorderWidthNormalized float64 `json:"BorderWidthNormalized" xml:"BorderWidthNormalized"`
	Text                  string  `json:"Text" xml:"Text"`
	SizeNormalized        float64 `json:"SizeNormalized" xml:"SizeNormalized"`
	FontName              string  `json:"FontName" xml:"FontName"`
}
