package imageprocess

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

// Data is a nested struct in imageprocess response
type Data struct {
	Text              string                 `json:"Text" xml:"Text"`
	ErrorMessage      string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	NormalProbability string                 `json:"NormalProbability" xml:"NormalProbability"`
	SessionId         string                 `json:"SessionId" xml:"SessionId"`
	NewProbability    string                 `json:"NewProbability" xml:"NewProbability"`
	Reports           map[string]interface{} `json:"Reports" xml:"Reports"`
	DUrl              string                 `json:"DUrl" xml:"DUrl"`
	ImageUrl          string                 `json:"ImageUrl" xml:"ImageUrl"`
	Mask              string                 `json:"Mask" xml:"Mask"`
	ResultsEnglish    map[string]interface{} `json:"ResultsEnglish" xml:"ResultsEnglish"`
	ErrorCode         string                 `json:"ErrorCode" xml:"ErrorCode"`
	BodyPart          string                 `json:"BodyPart" xml:"BodyPart"`
	JobId             string                 `json:"JobId" xml:"JobId"`
	OrgId             string                 `json:"OrgId" xml:"OrgId"`
	Data              string                 `json:"Data" xml:"Data"`
	Result            string                 `json:"Result" xml:"Result"`
	Question          string                 `json:"Question" xml:"Question"`
	OtherProbability  string                 `json:"OtherProbability" xml:"OtherProbability"`
	Words             int64                  `json:"Words" xml:"Words"`
	LesionRatio       string                 `json:"LesionRatio" xml:"LesionRatio"`
	QuestionType      string                 `json:"QuestionType" xml:"QuestionType"`
	AnswerType        string                 `json:"AnswerType" xml:"AnswerType"`
	NUrl              string                 `json:"NUrl" xml:"NUrl"`
	ImageType         string                 `json:"ImageType" xml:"ImageType"`
	ImageQuality      float64                `json:"ImageQuality" xml:"ImageQuality"`
	Status            string                 `json:"Status" xml:"Status"`
	Results           map[string]interface{} `json:"Results" xml:"Results"`
	ResultURL         string                 `json:"ResultURL" xml:"ResultURL"`
	OrgName           string                 `json:"OrgName" xml:"OrgName"`
	Spacing           []float64              `json:"Spacing" xml:"Spacing"`
	Options           []string               `json:"Options" xml:"Options"`
	Origin            []float64              `json:"Origin" xml:"Origin"`
	Lesion            Lesion                 `json:"Lesion" xml:"Lesion"`
	Series            []Serie                `json:"Series" xml:"Series"`
	Fractures         []FracturesItem        `json:"Fractures" xml:"Fractures"`
	Vertebras         []Vertebra             `json:"Vertebras" xml:"Vertebras"`
	Detections        []DetectionsItem       `json:"Detections" xml:"Detections"`
	KLDetections      []KLDetectionsItem     `json:"KLDetections" xml:"KLDetections"`
	KeyPoints         []KeyPointsItem        `json:"KeyPoints" xml:"KeyPoints"`
	Discs             []Disc                 `json:"Discs" xml:"Discs"`
}
