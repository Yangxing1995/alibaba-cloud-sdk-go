package dms_enterprise

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

// Column is a nested struct in dms_enterprise response
type Column struct {
	Charset        string `json:"Charset" xml:"Charset"`
	AutoIncrement  bool   `json:"AutoIncrement" xml:"AutoIncrement"`
	ColumnId       string `json:"ColumnId" xml:"ColumnId"`
	ColumnPosition int    `json:"ColumnPosition" xml:"ColumnPosition"`
	PrimaryKey     string `json:"PrimaryKey" xml:"PrimaryKey"`
	DefaultValue   string `json:"DefaultValue" xml:"DefaultValue"`
	Sensitive      bool   `json:"Sensitive" xml:"Sensitive"`
	Nullable       bool   `json:"Nullable" xml:"Nullable"`
	Fictive        bool   `json:"Fictive" xml:"Fictive"`
	ColumnName     string `json:"ColumnName" xml:"ColumnName"`
	SecurityLevel  string `json:"SecurityLevel" xml:"SecurityLevel"`
	ColumnType     string `json:"ColumnType" xml:"ColumnType"`
	Position       int    `json:"Position" xml:"Position"`
	DataLength     int64  `json:"DataLength" xml:"DataLength"`
	FunctionType   string `json:"FunctionType" xml:"FunctionType"`
	Description    string `json:"Description" xml:"Description"`
	DataScale      int    `json:"DataScale" xml:"DataScale"`
	DataPrecision  int    `json:"DataPrecision" xml:"DataPrecision"`
}
