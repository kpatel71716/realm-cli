// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/realm_client.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	api "github.com/10gen/realm-cli/api"
	auth "github.com/10gen/realm-cli/auth"
	hosting "github.com/10gen/realm-cli/hosting"
	models "github.com/10gen/realm-cli/models"
	secrets "github.com/10gen/realm-cli/secrets"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockRealmClient is a mock of RealmClient interface
type MockRealmClient struct {
	ctrl     *gomock.Controller
	recorder *MockRealmClientMockRecorder
}

// MockRealmClientMockRecorder is the mock recorder for MockRealmClient
type MockRealmClientMockRecorder struct {
	mock *MockRealmClient
}

// NewMockRealmClient creates a new mock instance
func NewMockRealmClient(ctrl *gomock.Controller) *MockRealmClient {
	mock := &MockRealmClient{ctrl: ctrl}
	mock.recorder = &MockRealmClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRealmClient) EXPECT() *MockRealmClientMockRecorder {
	return m.recorder
}

// AddSecret mocks base method
func (m *MockRealmClient) AddSecret(groupID, appID string, secret secrets.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecret", groupID, appID, secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSecret indicates an expected call of AddSecret
func (mr *MockRealmClientMockRecorder) AddSecret(groupID, appID, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecret", reflect.TypeOf((*MockRealmClient)(nil).AddSecret), groupID, appID, secret)
}

// Authenticate mocks base method
func (m *MockRealmClient) Authenticate(authProvider auth.AuthenticationProvider) (*auth.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", authProvider)
	ret0, _ := ret[0].(*auth.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockRealmClientMockRecorder) Authenticate(authProvider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockRealmClient)(nil).Authenticate), authProvider)
}

// CopyAsset mocks base method
func (m *MockRealmClient) CopyAsset(groupID, appID, fromPath, toPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyAsset", groupID, appID, fromPath, toPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyAsset indicates an expected call of CopyAsset
func (mr *MockRealmClientMockRecorder) CopyAsset(groupID, appID, fromPath, toPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyAsset", reflect.TypeOf((*MockRealmClient)(nil).CopyAsset), groupID, appID, fromPath, toPath)
}

// CreateDraft mocks base method
func (m *MockRealmClient) CreateDraft(groupID, appID string) (*models.AppDraft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDraft", groupID, appID)
	ret0, _ := ret[0].(*models.AppDraft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDraft indicates an expected call of CreateDraft
func (mr *MockRealmClientMockRecorder) CreateDraft(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDraft", reflect.TypeOf((*MockRealmClient)(nil).CreateDraft), groupID, appID)
}

// CreateEmptyApp mocks base method
func (m *MockRealmClient) CreateEmptyApp(groupID, appName, location, deploymentModel string) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmptyApp", groupID, appName, location, deploymentModel)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmptyApp indicates an expected call of CreateEmptyApp
func (mr *MockRealmClientMockRecorder) CreateEmptyApp(groupID, appName, location, deploymentModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmptyApp", reflect.TypeOf((*MockRealmClient)(nil).CreateEmptyApp), groupID, appName, location, deploymentModel)
}

// DeleteAsset mocks base method
func (m *MockRealmClient) DeleteAsset(groupID, appID, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsset", groupID, appID, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAsset indicates an expected call of DeleteAsset
func (mr *MockRealmClientMockRecorder) DeleteAsset(groupID, appID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsset", reflect.TypeOf((*MockRealmClient)(nil).DeleteAsset), groupID, appID, path)
}

// DeployDraft mocks base method
func (m *MockRealmClient) DeployDraft(groupID, appID, draftID string) (*models.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployDraft", groupID, appID, draftID)
	ret0, _ := ret[0].(*models.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployDraft indicates an expected call of DeployDraft
func (mr *MockRealmClientMockRecorder) DeployDraft(groupID, appID, draftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployDraft", reflect.TypeOf((*MockRealmClient)(nil).DeployDraft), groupID, appID, draftID)
}

// Diff mocks base method
func (m *MockRealmClient) Diff(groupID, appID string, appData []byte, strategy string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Diff", groupID, appID, appData, strategy)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Diff indicates an expected call of Diff
func (mr *MockRealmClientMockRecorder) Diff(groupID, appID, appData, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diff", reflect.TypeOf((*MockRealmClient)(nil).Diff), groupID, appID, appData, strategy)
}

// DiscardDraft mocks base method
func (m *MockRealmClient) DiscardDraft(groupID, appID, draftID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscardDraft", groupID, appID, draftID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscardDraft indicates an expected call of DiscardDraft
func (mr *MockRealmClientMockRecorder) DiscardDraft(groupID, appID, draftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardDraft", reflect.TypeOf((*MockRealmClient)(nil).DiscardDraft), groupID, appID, draftID)
}

// DraftDiff mocks base method
func (m *MockRealmClient) DraftDiff(groupID, appID, draftID string) (*models.DraftDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DraftDiff", groupID, appID, draftID)
	ret0, _ := ret[0].(*models.DraftDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DraftDiff indicates an expected call of DraftDiff
func (mr *MockRealmClientMockRecorder) DraftDiff(groupID, appID, draftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DraftDiff", reflect.TypeOf((*MockRealmClient)(nil).DraftDiff), groupID, appID, draftID)
}

// Export mocks base method
func (m *MockRealmClient) Export(groupID, appID string, strategy api.ExportStrategy) (string, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", groupID, appID, strategy)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Export indicates an expected call of Export
func (mr *MockRealmClientMockRecorder) Export(groupID, appID, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockRealmClient)(nil).Export), groupID, appID, strategy)
}

// ExportDependencies mocks base method
func (m *MockRealmClient) ExportDependencies(groupID, appID string) (string, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportDependencies", groupID, appID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportDependencies indicates an expected call of ExportDependencies
func (mr *MockRealmClientMockRecorder) ExportDependencies(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportDependencies", reflect.TypeOf((*MockRealmClient)(nil).ExportDependencies), groupID, appID)
}

// FetchAppByClientAppID mocks base method
func (m *MockRealmClient) FetchAppByClientAppID(clientAppID string) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAppByClientAppID", clientAppID)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAppByClientAppID indicates an expected call of FetchAppByClientAppID
func (mr *MockRealmClientMockRecorder) FetchAppByClientAppID(clientAppID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAppByClientAppID", reflect.TypeOf((*MockRealmClient)(nil).FetchAppByClientAppID), clientAppID)
}

// FetchAppByGroupIDAndClientAppID mocks base method
func (m *MockRealmClient) FetchAppByGroupIDAndClientAppID(groupID, clientAppID string) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAppByGroupIDAndClientAppID", groupID, clientAppID)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAppByGroupIDAndClientAppID indicates an expected call of FetchAppByGroupIDAndClientAppID
func (mr *MockRealmClientMockRecorder) FetchAppByGroupIDAndClientAppID(groupID, clientAppID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAppByGroupIDAndClientAppID", reflect.TypeOf((*MockRealmClient)(nil).FetchAppByGroupIDAndClientAppID), groupID, clientAppID)
}

// FetchAppsByGroupID mocks base method
func (m *MockRealmClient) FetchAppsByGroupID(groupID string) ([]*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAppsByGroupID", groupID)
	ret0, _ := ret[0].([]*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAppsByGroupID indicates an expected call of FetchAppsByGroupID
func (mr *MockRealmClientMockRecorder) FetchAppsByGroupID(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAppsByGroupID", reflect.TypeOf((*MockRealmClient)(nil).FetchAppsByGroupID), groupID)
}

// GetDeployment mocks base method
func (m *MockRealmClient) GetDeployment(groupID, appID, deploymentID string) (*models.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", groupID, appID, deploymentID)
	ret0, _ := ret[0].(*models.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment
func (mr *MockRealmClientMockRecorder) GetDeployment(groupID, appID, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockRealmClient)(nil).GetDeployment), groupID, appID, deploymentID)
}

// GetDrafts mocks base method
func (m *MockRealmClient) GetDrafts(groupID, appID string) ([]models.AppDraft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDrafts", groupID, appID)
	ret0, _ := ret[0].([]models.AppDraft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDrafts indicates an expected call of GetDrafts
func (mr *MockRealmClientMockRecorder) GetDrafts(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDrafts", reflect.TypeOf((*MockRealmClient)(nil).GetDrafts), groupID, appID)
}

// Import mocks base method
func (m *MockRealmClient) Import(groupID, appID string, appData []byte, strategy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", groupID, appID, appData, strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Import indicates an expected call of Import
func (mr *MockRealmClientMockRecorder) Import(groupID, appID, appData, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockRealmClient)(nil).Import), groupID, appID, appData, strategy)
}

// InvalidateCache mocks base method
func (m *MockRealmClient) InvalidateCache(groupID, appID, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateCache", groupID, appID, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateCache indicates an expected call of InvalidateCache
func (mr *MockRealmClientMockRecorder) InvalidateCache(groupID, appID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCache", reflect.TypeOf((*MockRealmClient)(nil).InvalidateCache), groupID, appID, path)
}

// ListAssetsForAppID mocks base method
func (m *MockRealmClient) ListAssetsForAppID(groupID, appID string) ([]hosting.AssetMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAssetsForAppID", groupID, appID)
	ret0, _ := ret[0].([]hosting.AssetMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssetsForAppID indicates an expected call of ListAssetsForAppID
func (mr *MockRealmClientMockRecorder) ListAssetsForAppID(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssetsForAppID", reflect.TypeOf((*MockRealmClient)(nil).ListAssetsForAppID), groupID, appID)
}

// ListSecrets mocks base method
func (m *MockRealmClient) ListSecrets(groupID, appID string) ([]secrets.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", groupID, appID)
	ret0, _ := ret[0].([]secrets.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockRealmClientMockRecorder) ListSecrets(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockRealmClient)(nil).ListSecrets), groupID, appID)
}

// MoveAsset mocks base method
func (m *MockRealmClient) MoveAsset(groupID, appID, fromPath, toPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveAsset", groupID, appID, fromPath, toPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveAsset indicates an expected call of MoveAsset
func (mr *MockRealmClientMockRecorder) MoveAsset(groupID, appID, fromPath, toPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveAsset", reflect.TypeOf((*MockRealmClient)(nil).MoveAsset), groupID, appID, fromPath, toPath)
}

// RemoveSecretByID mocks base method
func (m *MockRealmClient) RemoveSecretByID(groupID, appID, secretID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretByID", groupID, appID, secretID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretByID indicates an expected call of RemoveSecretByID
func (mr *MockRealmClientMockRecorder) RemoveSecretByID(groupID, appID, secretID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretByID", reflect.TypeOf((*MockRealmClient)(nil).RemoveSecretByID), groupID, appID, secretID)
}

// RemoveSecretByName mocks base method
func (m *MockRealmClient) RemoveSecretByName(groupID, appID, secretName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretByName", groupID, appID, secretName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretByName indicates an expected call of RemoveSecretByName
func (mr *MockRealmClientMockRecorder) RemoveSecretByName(groupID, appID, secretName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretByName", reflect.TypeOf((*MockRealmClient)(nil).RemoveSecretByName), groupID, appID, secretName)
}

// SetAssetAttributes mocks base method
func (m *MockRealmClient) SetAssetAttributes(groupID, appID, path string, attributes ...hosting.AssetAttribute) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{groupID, appID, path}
	for _, a := range attributes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetAssetAttributes", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAssetAttributes indicates an expected call of SetAssetAttributes
func (mr *MockRealmClientMockRecorder) SetAssetAttributes(groupID, appID, path interface{}, attributes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{groupID, appID, path}, attributes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAssetAttributes", reflect.TypeOf((*MockRealmClient)(nil).SetAssetAttributes), varargs...)
}

// UpdateSecretByID mocks base method
func (m *MockRealmClient) UpdateSecretByID(groupID, appID, secretID, secretValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretByID", groupID, appID, secretID, secretValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretByID indicates an expected call of UpdateSecretByID
func (mr *MockRealmClientMockRecorder) UpdateSecretByID(groupID, appID, secretID, secretValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretByID", reflect.TypeOf((*MockRealmClient)(nil).UpdateSecretByID), groupID, appID, secretID, secretValue)
}

// UpdateSecretByName mocks base method
func (m *MockRealmClient) UpdateSecretByName(groupID, appID, secretName, secretValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretByName", groupID, appID, secretName, secretValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretByName indicates an expected call of UpdateSecretByName
func (mr *MockRealmClientMockRecorder) UpdateSecretByName(groupID, appID, secretName, secretValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretByName", reflect.TypeOf((*MockRealmClient)(nil).UpdateSecretByName), groupID, appID, secretName, secretValue)
}

// UploadAsset mocks base method
func (m *MockRealmClient) UploadAsset(groupID, appID, path, hash string, size int64, body io.Reader, attributes ...hosting.AssetAttribute) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{groupID, appID, path, hash, size, body}
	for _, a := range attributes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadAsset", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAsset indicates an expected call of UploadAsset
func (mr *MockRealmClientMockRecorder) UploadAsset(groupID, appID, path, hash, size, body interface{}, attributes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{groupID, appID, path, hash, size, body}, attributes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAsset", reflect.TypeOf((*MockRealmClient)(nil).UploadAsset), varargs...)
}

// UploadDependencies mocks base method
func (m *MockRealmClient) UploadDependencies(groupID, appID, fullPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadDependencies", groupID, appID, fullPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadDependencies indicates an expected call of UploadDependencies
func (mr *MockRealmClientMockRecorder) UploadDependencies(groupID, appID, fullPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadDependencies", reflect.TypeOf((*MockRealmClient)(nil).UploadDependencies), groupID, appID, fullPath)
}
