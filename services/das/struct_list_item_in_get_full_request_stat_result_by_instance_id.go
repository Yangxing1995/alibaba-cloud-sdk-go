package das

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

// ListItemInGetFullRequestStatResultByInstanceId is a nested struct in das response
type ListItemInGetFullRequestStatResultByInstanceId struct {
	SqlId                       string   `json:"SqlId" xml:"SqlId"`
	AvgFetchRows                int64    `json:"AvgFetchRows" xml:"AvgFetchRows"`
	AvgLockWaitTime             string   `json:"AvgLockWaitTime" xml:"AvgLockWaitTime"`
	AvgLogicalRead              string   `json:"AvgLogicalRead" xml:"AvgLogicalRead"`
	AvgPhysicalAsyncRead        int64    `json:"AvgPhysicalAsyncRead" xml:"AvgPhysicalAsyncRead"`
	AvgPhysicalSyncRead         int64    `json:"AvgPhysicalSyncRead" xml:"AvgPhysicalSyncRead"`
	AvgExaminedRows             string   `json:"AvgExaminedRows" xml:"AvgExaminedRows"`
	AvgReturnedRows             string   `json:"AvgReturnedRows" xml:"AvgReturnedRows"`
	AvgUpdatedRows              int64    `json:"AvgUpdatedRows" xml:"AvgUpdatedRows"`
	AvgRt                       string   `json:"AvgRt" xml:"AvgRt"`
	AvgSqlCount                 int64    `json:"AvgSqlCount" xml:"AvgSqlCount"`
	Count                       int64    `json:"Count" xml:"Count"`
	CountRate                   string   `json:"CountRate" xml:"CountRate"`
	Database                    string   `json:"Database" xml:"Database"`
	ErrorCount                  int64    `json:"ErrorCount" xml:"ErrorCount"`
	FetchRows                   int64    `json:"FetchRows" xml:"FetchRows"`
	Ip                          string   `json:"Ip" xml:"Ip"`
	LockWaitTime                string   `json:"LockWaitTime" xml:"LockWaitTime"`
	LogicalRead                 int64    `json:"LogicalRead" xml:"LogicalRead"`
	PhysicalAsyncRead           int64    `json:"PhysicalAsyncRead" xml:"PhysicalAsyncRead"`
	PhysicalSyncRead            int64    `json:"PhysicalSyncRead" xml:"PhysicalSyncRead"`
	Port                        int64    `json:"Port" xml:"Port"`
	Psql                        string   `json:"Psql" xml:"Psql"`
	Rows                        int64    `json:"Rows" xml:"Rows"`
	ExaminedRows                int64    `json:"ExaminedRows" xml:"ExaminedRows"`
	RtGreaterThanOneSecondCount int64    `json:"RtGreaterThanOneSecondCount" xml:"RtGreaterThanOneSecondCount"`
	RtRate                      string   `json:"RtRate" xml:"RtRate"`
	SqlCount                    int64    `json:"SqlCount" xml:"SqlCount"`
	SumUpdatedRows              int64    `json:"SumUpdatedRows" xml:"SumUpdatedRows"`
	Version                     int64    `json:"Version" xml:"Version"`
	VpcId                       string   `json:"VpcId" xml:"VpcId"`
	Tables                      []string `json:"Tables" xml:"Tables"`
}
