package qualitycheck

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

// RuleInfo is a nested struct in qualitycheck response
type RuleInfo struct {
	Status                   int                               `json:"Status" xml:"Status"`
	Type                     int                               `json:"Type" xml:"Type"`
	ScoreSubId               int                               `json:"ScoreSubId" xml:"ScoreSubId"`
	IsOnline                 int                               `json:"IsOnline" xml:"IsOnline"`
	CreateTime               string                            `json:"CreateTime" xml:"CreateTime"`
	CreateEmpid              string                            `json:"CreateEmpid" xml:"CreateEmpid"`
	LastUpdateEmpid          string                            `json:"LastUpdateEmpid" xml:"LastUpdateEmpid"`
	IsDelete                 int                               `json:"IsDelete" xml:"IsDelete"`
	Rid                      string                            `json:"Rid" xml:"Rid"`
	RuleScoreType            int                               `json:"RuleScoreType" xml:"RuleScoreType"`
	EndTime                  string                            `json:"EndTime" xml:"EndTime"`
	Weight                   string                            `json:"Weight" xml:"Weight"`
	StartTime                string                            `json:"StartTime" xml:"StartTime"`
	RuleLambda               string                            `json:"RuleLambda" xml:"RuleLambda"`
	ScoreSubName             string                            `json:"ScoreSubName" xml:"ScoreSubName"`
	AutoReview               int                               `json:"AutoReview" xml:"AutoReview"`
	Comments                 string                            `json:"Comments" xml:"Comments"`
	LastUpdateTime           string                            `json:"LastUpdateTime" xml:"LastUpdateTime"`
	ScoreName                string                            `json:"ScoreName" xml:"ScoreName"`
	Name                     string                            `json:"Name" xml:"Name"`
	ScoreId                  int                               `json:"ScoreId" xml:"ScoreId"`
	BusinessCategoryNameList BusinessCategoryNameListInGetRule `json:"BusinessCategoryNameList" xml:"BusinessCategoryNameList"`
}
