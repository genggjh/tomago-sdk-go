/*
Copyright ArxanFintech Technology Ltd. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	restapi "github.com/arxanchain/go-common/rest/api"
	"github.com/arxanchain/go-common/structs"
)

type TomagoClient struct {
	c            *restapi.Client
	entityClient *EntityClient
	assetClient  *AssetClient
	ccoinClient  *CCoinClient
}

// TomagoClient returns a handle to the agent endpoints
func NewTomagoClient(config *restapi.Config) (*TomagoClient, error) {
	c, err := restapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &TomagoClient{c: c}, nil
}

func (t *TomagoClient) GetEntityClient() structs.IEntityClient {
	if t.entityClient == nil {
		t.entityClient = &EntityClient{c: t.c}
	}
	return t.entityClient
}

func (t *TomagoClient) GetAssetClient() structs.IAssetClient {
	if t.assetClient == nil {
		t.assetClient = &AssetClient{c: t.c}
	}
	return t.assetClient
}

func (t *TomagoClient) GetCCoinClient() structs.ICCoinClient {
	if t.ccoinClient == nil {
		t.ccoinClient = &CCoinClient{c: t.c}
	}
	return t.ccoinClient
}