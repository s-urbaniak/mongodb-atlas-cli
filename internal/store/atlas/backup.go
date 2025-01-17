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

package atlas

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

type CompliancePolicyDescriber interface {
	DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
}

//go:generate mockgen -destination=../../mocks/atlas/mock_backup.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CompliancePolicyDescriber,CompliancePolicyUpdater,CompliancePolicyEncryptionAtRestEnabler,CompliancePolicyEncryptionAtRestDisabler,CompliancePolicyEnabler,CompliancePolicyCopyProtectionEnabler,CompliancePolicyCopyProtectionDisabler,CompliancePolicyPointInTimeRestoresEnabler

type CompliancePolicyPointInTimeRestoresEnabler interface {
	EnablePointInTimeRestore(projectID string, restoreWindowDays int) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyEnabler interface {
	EnableCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyCopyProtectionEnabler interface {
	EnableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}
type CompliancePolicyCopyProtectionDisabler interface {
	DisableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}
type CompliancePolicyEncryptionAtRestUpdater interface {
	UpdateEncryptionAtRest(projectID string, enable bool) (*atlasv2.DataProtectionSettings20231001, error)
}
type CompliancePolicyUpdater interface {
	CompliancePolicyDescriber
	UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error)
}

type CompliancePolicyEncryptionAtRestEnabler interface {
	EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyEncryptionAtRestDisabler interface {
	DisableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

func (s *Store) DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	return result, err
}

func (s *Store) UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, opts).Execute()
	return result, err
}

func (s *Store) EnablePointInTimeRestore(projectID string, restoreWindowDays int) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetRestoreWindowDays(restoreWindowDays)
	compliancePolicy.SetPitEnabled(true)
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *Store) EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetEncryptionAtRestEnabled(true)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Store) EnableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetCopyProtectionEnabled(true)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *Store) DisableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetEncryptionAtRestEnabled(false)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Store) DisableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetCopyProtectionEnabled(false)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *Store) EnableCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy := newEmptyCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func newEmptyCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) *atlasv2.DataProtectionSettings20231001 {
	policy := atlasv2.NewDataProtectionSettings20231001(authorizedEmail, authorizedFirstName, authorizedLastName)
	policy.SetProjectId(projectID)
	return policy
}
