package ecd

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

// Snapshot is a nested struct in ecd response
type Snapshot struct {
	CreationTime            string `json:"CreationTime" xml:"CreationTime"`
	Status                  string `json:"Status" xml:"Status"`
	SnapshotType            string `json:"SnapshotType" xml:"SnapshotType"`
	SnapshotName            string `json:"SnapshotName" xml:"SnapshotName"`
	Progress                string `json:"Progress" xml:"Progress"`
	Description             string `json:"Description" xml:"Description"`
	SnapshotId              string `json:"SnapshotId" xml:"SnapshotId"`
	RemainTime              int    `json:"RemainTime" xml:"RemainTime"`
	SourceDiskSize          string `json:"SourceDiskSize" xml:"SourceDiskSize"`
	SourceDiskType          string `json:"SourceDiskType" xml:"SourceDiskType"`
	DesktopId               string `json:"DesktopId" xml:"DesktopId"`
	VolumeEncryptionEnabled bool   `json:"VolumeEncryptionEnabled" xml:"VolumeEncryptionEnabled"`
	VolumeEncryptionKey     string `json:"VolumeEncryptionKey" xml:"VolumeEncryptionKey"`
}
