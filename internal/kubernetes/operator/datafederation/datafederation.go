// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datafederation

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas-sdk/v20231001002/admin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DeletingState = "DELETING"
	DeletedState  = "DELETED"
)

func BuildAtlasDataFederation(dataFederationStore atlas.DataFederationStore, dataFederationName, projectID, projectName, operatorVersion, targetNamespace string, dictionary map[string]string) (*atlasV1.AtlasDataFederation, error) {
	dataFederation, err := dataFederationStore.DataFederation(projectID, dataFederationName)
	if err != nil {
		return nil, err
	}
	if !isDataFederationExportable(dataFederation) {
		return nil, nil
	}
	atlasDataFederation := &atlasV1.AtlasDataFederation{
		TypeMeta: v1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasDataFederation",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: getDataFederationSpec(dataFederation, targetNamespace, projectName),
		Status: status.DataFederationStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
	return atlasDataFederation, nil
}

func isDataFederationExportable(dataFederation *admin.DataLakeTenant) bool {
	state := dataFederation.GetState()
	return state != DeletingState && state != DeletedState
}

func getDataFederationSpec(dataFederationSpec *admin.DataLakeTenant, targetNamespace, projectName string) atlasV1.DataFederationSpec {
	return atlasV1.DataFederationSpec{
		Project:             common.ResourceRefNamespaced{Name: projectName, Namespace: targetNamespace},
		Name:                dataFederationSpec.GetName(),
		CloudProviderConfig: getCloudProviderConfig(dataFederationSpec.CloudProviderConfig),
		DataProcessRegion:   getDataProcessRegion(dataFederationSpec.DataProcessRegion),
		Storage:             getStorage(dataFederationSpec.Storage),
	}
}

func getCloudProviderConfig(cloudProviderConfig *admin.DataLakeCloudProviderConfig) *atlasV1.CloudProviderConfig {
	if cloudProviderConfig == nil {
		return &atlasV1.CloudProviderConfig{}
	}
	return &atlasV1.CloudProviderConfig{
		AWS: &atlasV1.AWSProviderConfig{
			RoleID:       cloudProviderConfig.Aws.RoleId,
			TestS3Bucket: cloudProviderConfig.Aws.TestS3Bucket,
		},
	}
}

func getDataProcessRegion(dataProcessRegion *admin.DataLakeDataProcessRegion) *atlasV1.DataProcessRegion {
	if dataProcessRegion == nil {
		return &atlasV1.DataProcessRegion{}
	}
	return &atlasV1.DataProcessRegion{
		CloudProvider: dataProcessRegion.CloudProvider,
		Region:        dataProcessRegion.Region,
	}
}

func getStorage(storage *admin.DataLakeStorage) *atlasV1.Storage {
	if storage == nil {
		return &atlasV1.Storage{}
	}
	return &atlasV1.Storage{
		Databases: getDatabases(storage.Databases),
		Stores:    getStores(storage.Stores),
	}
}

func getDatabases(database []admin.DataLakeDatabaseInstance) []atlasV1.Database {
	if database == nil {
		return []atlasV1.Database{}
	}
	result := make([]atlasV1.Database, 0, len(database))

	for _, obj := range database {
		result = append(result, atlasV1.Database{
			Collections:            getCollection(obj.GetCollections()),
			MaxWildcardCollections: obj.GetMaxWildcardCollections(),
			Name:                   obj.GetName(),
			Views:                  getViews(obj.GetViews()),
		})
	}
	return result
}

func getCollection(collections []admin.DataLakeDatabaseCollection) []atlasV1.Collection {
	if collections == nil {
		return []atlasV1.Collection{}
	}
	result := make([]atlasV1.Collection, 0, len(collections))

	for _, obj := range collections {
		result = append(result, atlasV1.Collection{
			DataSources: getDataSources(obj.GetDataSources()),
			Name:        obj.GetName(),
		})
	}
	return result
}

func getDataSources(dataSources []admin.DataLakeDatabaseDataSourceSettings) []atlasV1.DataSource {
	if dataSources == nil {
		return []atlasV1.DataSource{}
	}
	result := make([]atlasV1.DataSource, 0, len(dataSources))

	for _, obj := range dataSources {
		result = append(result, atlasV1.DataSource{
			AllowInsecure:       obj.GetAllowInsecure(),
			Collection:          obj.GetCollection(),
			CollectionRegex:     obj.GetCollectionRegex(),
			Database:            obj.GetDatabase(),
			DatabaseRegex:       obj.GetDatabaseRegex(),
			DefaultFormat:       obj.GetDefaultFormat(),
			Path:                obj.GetPath(),
			ProvenanceFieldName: obj.GetProvenanceFieldName(),
			StoreName:           obj.GetStoreName(),
			Urls:                obj.GetUrls(),
		})
	}
	return result
}

func getViews(views []admin.DataLakeApiBase) []atlasV1.View {
	if views == nil {
		return []atlasV1.View{}
	}
	result := make([]atlasV1.View, 0, len(views))

	for _, obj := range views {
		result = append(result, atlasV1.View{
			Name:     obj.GetName(),
			Pipeline: obj.GetPipeline(),
			Source:   obj.GetSource(),
		})
	}
	return result
}

func getStores(stores []admin.DataLakeStoreSettings) []atlasV1.Store {
	if stores == nil {
		return []atlasV1.Store{}
	}
	result := make([]atlasV1.Store, 0, len(stores))

	for _, obj := range stores {
		result = append(result, atlasV1.Store{
			Name:                     obj.GetName(),
			Provider:                 obj.Provider,
			AdditionalStorageClasses: obj.GetAdditionalStorageClasses(),
			Bucket:                   obj.GetBucket(),
			Delimiter:                obj.GetDelimiter(),
			IncludeTags:              obj.GetIncludeTags(),
			Prefix:                   obj.GetPrefix(),
			Public:                   obj.GetPublic(),
			Region:                   obj.GetRegion(),
		})
	}
	return result
}
