package vod

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

// MaterialInGetEditingProjectMaterials is a nested struct in vod response
type MaterialInGetEditingProjectMaterials struct {
	Status       string                                `json:"Status" xml:"Status"`
	CreationTime string                                `json:"CreationTime" xml:"CreationTime"`
	CateId       int                                   `json:"CateId" xml:"CateId"`
	MaterialType string                                `json:"MaterialType" xml:"MaterialType"`
	Tags         string                                `json:"Tags" xml:"Tags"`
	SpriteConfig string                                `json:"SpriteConfig" xml:"SpriteConfig"`
	Source       string                                `json:"Source" xml:"Source"`
	CateName     string                                `json:"CateName" xml:"CateName"`
	ModifiedTime string                                `json:"ModifiedTime" xml:"ModifiedTime"`
	Description  string                                `json:"Description" xml:"Description"`
	MaterialId   string                                `json:"MaterialId" xml:"MaterialId"`
	Size         int64                                 `json:"Size" xml:"Size"`
	CoverURL     string                                `json:"CoverURL" xml:"CoverURL"`
	Duration     float64                               `json:"Duration" xml:"Duration"`
	Title        string                                `json:"Title" xml:"Title"`
	Sprites      SpritesInGetEditingProjectMaterials   `json:"Sprites" xml:"Sprites"`
	Snapshots    SnapshotsInGetEditingProjectMaterials `json:"Snapshots" xml:"Snapshots"`
}
