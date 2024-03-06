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

// SceneMetrics is a nested struct in pts response
type SceneMetrics struct {
	FailCountBiz   int64   `json:"FailCountBiz" xml:"FailCountBiz"`
	AllCount       int64   `json:"AllCount" xml:"AllCount"`
	SuccessRateBiz float64 `json:"SuccessRateBiz" xml:"SuccessRateBiz"`
	AvgRt          float64 `json:"AvgRt" xml:"AvgRt"`
	FailCountReq   int64   `json:"FailCountReq" xml:"FailCountReq"`
	AvgTps         float64 `json:"AvgTps" xml:"AvgTps"`
	Seg99Rt        float64 `json:"Seg99Rt" xml:"Seg99Rt"`
	SuccessRateReq float64 `json:"SuccessRateReq" xml:"SuccessRateReq"`
	Seg90Rt        float64 `json:"Seg90Rt" xml:"Seg90Rt"`
}
