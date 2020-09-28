// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type Snapshot struct {

    /* 云硬盘快照ID (Optional) */
    SnapshotId string `json:"snapshotId"`

    /* 快照来源 可以有self，others两种来源 (Optional) */
    SnapshotSource string `json:"snapshotSource"`

    /* 创建快照的云硬盘ID(snapshotSource为others时不展示) (Optional) */
    DiskId string `json:"diskId"`

    /* 快照大小，单位为GiB (Optional) */
    SnapshotSizeGB int `json:"snapshotSizeGB"`

    /* 快照关联的所有镜像ID(snapshotSource为others时不展示) (Optional) */
    Images []string `json:"images"`

    /* 快照名称 (Optional) */
    Name string `json:"name"`

    /* 快照描述 (Optional) */
    Description string `json:"description"`

    /* 快照状态，取值为 creating、available、in-use、deleting、error_create、error_delete 之一 (Optional) */
    Status string `json:"status"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 共享信息 (Optional) */
    SharInfo []ShareInfo `json:"sharInfo"`

    /* 快照是否为加密盘的快照 (Optional) */
    Encrypted bool `json:"encrypted"`
}
