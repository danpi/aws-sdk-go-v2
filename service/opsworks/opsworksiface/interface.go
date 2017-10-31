// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworksiface provides an interface to enable mocking the AWS OpsWorks service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opsworksiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

// OpsWorksAPI provides an interface to enable mocking the
// opsworks.OpsWorks service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS OpsWorks.
//    func myFunc(svc opsworksiface.OpsWorksAPI) bool {
//        // Make svc.AssignInstance request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := opsworks.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOpsWorksClient struct {
//        opsworksiface.OpsWorksAPI
//    }
//    func (m *mockOpsWorksClient) AssignInstance(input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOpsWorksClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type OpsWorksAPI interface {
	AssignInstance(*opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)
	AssignInstanceWithContext(aws.Context, *opsworks.AssignInstanceInput, ...aws.Option) (*opsworks.AssignInstanceOutput, error)
	AssignInstanceRequest(*opsworks.AssignInstanceInput) (*aws.Request, *opsworks.AssignInstanceOutput)

	AssignVolume(*opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)
	AssignVolumeWithContext(aws.Context, *opsworks.AssignVolumeInput, ...aws.Option) (*opsworks.AssignVolumeOutput, error)
	AssignVolumeRequest(*opsworks.AssignVolumeInput) (*aws.Request, *opsworks.AssignVolumeOutput)

	AssociateElasticIp(*opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)
	AssociateElasticIpWithContext(aws.Context, *opsworks.AssociateElasticIpInput, ...aws.Option) (*opsworks.AssociateElasticIpOutput, error)
	AssociateElasticIpRequest(*opsworks.AssociateElasticIpInput) (*aws.Request, *opsworks.AssociateElasticIpOutput)

	AttachElasticLoadBalancer(*opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)
	AttachElasticLoadBalancerWithContext(aws.Context, *opsworks.AttachElasticLoadBalancerInput, ...aws.Option) (*opsworks.AttachElasticLoadBalancerOutput, error)
	AttachElasticLoadBalancerRequest(*opsworks.AttachElasticLoadBalancerInput) (*aws.Request, *opsworks.AttachElasticLoadBalancerOutput)

	CloneStack(*opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)
	CloneStackWithContext(aws.Context, *opsworks.CloneStackInput, ...aws.Option) (*opsworks.CloneStackOutput, error)
	CloneStackRequest(*opsworks.CloneStackInput) (*aws.Request, *opsworks.CloneStackOutput)

	CreateApp(*opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)
	CreateAppWithContext(aws.Context, *opsworks.CreateAppInput, ...aws.Option) (*opsworks.CreateAppOutput, error)
	CreateAppRequest(*opsworks.CreateAppInput) (*aws.Request, *opsworks.CreateAppOutput)

	CreateDeployment(*opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)
	CreateDeploymentWithContext(aws.Context, *opsworks.CreateDeploymentInput, ...aws.Option) (*opsworks.CreateDeploymentOutput, error)
	CreateDeploymentRequest(*opsworks.CreateDeploymentInput) (*aws.Request, *opsworks.CreateDeploymentOutput)

	CreateInstance(*opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)
	CreateInstanceWithContext(aws.Context, *opsworks.CreateInstanceInput, ...aws.Option) (*opsworks.CreateInstanceOutput, error)
	CreateInstanceRequest(*opsworks.CreateInstanceInput) (*aws.Request, *opsworks.CreateInstanceOutput)

	CreateLayer(*opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)
	CreateLayerWithContext(aws.Context, *opsworks.CreateLayerInput, ...aws.Option) (*opsworks.CreateLayerOutput, error)
	CreateLayerRequest(*opsworks.CreateLayerInput) (*aws.Request, *opsworks.CreateLayerOutput)

	CreateStack(*opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)
	CreateStackWithContext(aws.Context, *opsworks.CreateStackInput, ...aws.Option) (*opsworks.CreateStackOutput, error)
	CreateStackRequest(*opsworks.CreateStackInput) (*aws.Request, *opsworks.CreateStackOutput)

	CreateUserProfile(*opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)
	CreateUserProfileWithContext(aws.Context, *opsworks.CreateUserProfileInput, ...aws.Option) (*opsworks.CreateUserProfileOutput, error)
	CreateUserProfileRequest(*opsworks.CreateUserProfileInput) (*aws.Request, *opsworks.CreateUserProfileOutput)

	DeleteApp(*opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)
	DeleteAppWithContext(aws.Context, *opsworks.DeleteAppInput, ...aws.Option) (*opsworks.DeleteAppOutput, error)
	DeleteAppRequest(*opsworks.DeleteAppInput) (*aws.Request, *opsworks.DeleteAppOutput)

	DeleteInstance(*opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)
	DeleteInstanceWithContext(aws.Context, *opsworks.DeleteInstanceInput, ...aws.Option) (*opsworks.DeleteInstanceOutput, error)
	DeleteInstanceRequest(*opsworks.DeleteInstanceInput) (*aws.Request, *opsworks.DeleteInstanceOutput)

	DeleteLayer(*opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)
	DeleteLayerWithContext(aws.Context, *opsworks.DeleteLayerInput, ...aws.Option) (*opsworks.DeleteLayerOutput, error)
	DeleteLayerRequest(*opsworks.DeleteLayerInput) (*aws.Request, *opsworks.DeleteLayerOutput)

	DeleteStack(*opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)
	DeleteStackWithContext(aws.Context, *opsworks.DeleteStackInput, ...aws.Option) (*opsworks.DeleteStackOutput, error)
	DeleteStackRequest(*opsworks.DeleteStackInput) (*aws.Request, *opsworks.DeleteStackOutput)

	DeleteUserProfile(*opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)
	DeleteUserProfileWithContext(aws.Context, *opsworks.DeleteUserProfileInput, ...aws.Option) (*opsworks.DeleteUserProfileOutput, error)
	DeleteUserProfileRequest(*opsworks.DeleteUserProfileInput) (*aws.Request, *opsworks.DeleteUserProfileOutput)

	DeregisterEcsCluster(*opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)
	DeregisterEcsClusterWithContext(aws.Context, *opsworks.DeregisterEcsClusterInput, ...aws.Option) (*opsworks.DeregisterEcsClusterOutput, error)
	DeregisterEcsClusterRequest(*opsworks.DeregisterEcsClusterInput) (*aws.Request, *opsworks.DeregisterEcsClusterOutput)

	DeregisterElasticIp(*opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)
	DeregisterElasticIpWithContext(aws.Context, *opsworks.DeregisterElasticIpInput, ...aws.Option) (*opsworks.DeregisterElasticIpOutput, error)
	DeregisterElasticIpRequest(*opsworks.DeregisterElasticIpInput) (*aws.Request, *opsworks.DeregisterElasticIpOutput)

	DeregisterInstance(*opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)
	DeregisterInstanceWithContext(aws.Context, *opsworks.DeregisterInstanceInput, ...aws.Option) (*opsworks.DeregisterInstanceOutput, error)
	DeregisterInstanceRequest(*opsworks.DeregisterInstanceInput) (*aws.Request, *opsworks.DeregisterInstanceOutput)

	DeregisterRdsDbInstance(*opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)
	DeregisterRdsDbInstanceWithContext(aws.Context, *opsworks.DeregisterRdsDbInstanceInput, ...aws.Option) (*opsworks.DeregisterRdsDbInstanceOutput, error)
	DeregisterRdsDbInstanceRequest(*opsworks.DeregisterRdsDbInstanceInput) (*aws.Request, *opsworks.DeregisterRdsDbInstanceOutput)

	DeregisterVolume(*opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)
	DeregisterVolumeWithContext(aws.Context, *opsworks.DeregisterVolumeInput, ...aws.Option) (*opsworks.DeregisterVolumeOutput, error)
	DeregisterVolumeRequest(*opsworks.DeregisterVolumeInput) (*aws.Request, *opsworks.DeregisterVolumeOutput)

	DescribeAgentVersions(*opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)
	DescribeAgentVersionsWithContext(aws.Context, *opsworks.DescribeAgentVersionsInput, ...aws.Option) (*opsworks.DescribeAgentVersionsOutput, error)
	DescribeAgentVersionsRequest(*opsworks.DescribeAgentVersionsInput) (*aws.Request, *opsworks.DescribeAgentVersionsOutput)

	DescribeApps(*opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)
	DescribeAppsWithContext(aws.Context, *opsworks.DescribeAppsInput, ...aws.Option) (*opsworks.DescribeAppsOutput, error)
	DescribeAppsRequest(*opsworks.DescribeAppsInput) (*aws.Request, *opsworks.DescribeAppsOutput)

	DescribeCommands(*opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)
	DescribeCommandsWithContext(aws.Context, *opsworks.DescribeCommandsInput, ...aws.Option) (*opsworks.DescribeCommandsOutput, error)
	DescribeCommandsRequest(*opsworks.DescribeCommandsInput) (*aws.Request, *opsworks.DescribeCommandsOutput)

	DescribeDeployments(*opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)
	DescribeDeploymentsWithContext(aws.Context, *opsworks.DescribeDeploymentsInput, ...aws.Option) (*opsworks.DescribeDeploymentsOutput, error)
	DescribeDeploymentsRequest(*opsworks.DescribeDeploymentsInput) (*aws.Request, *opsworks.DescribeDeploymentsOutput)

	DescribeEcsClusters(*opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)
	DescribeEcsClustersWithContext(aws.Context, *opsworks.DescribeEcsClustersInput, ...aws.Option) (*opsworks.DescribeEcsClustersOutput, error)
	DescribeEcsClustersRequest(*opsworks.DescribeEcsClustersInput) (*aws.Request, *opsworks.DescribeEcsClustersOutput)

	DescribeEcsClustersPages(*opsworks.DescribeEcsClustersInput, func(*opsworks.DescribeEcsClustersOutput, bool) bool) error
	DescribeEcsClustersPagesWithContext(aws.Context, *opsworks.DescribeEcsClustersInput, func(*opsworks.DescribeEcsClustersOutput, bool) bool, ...aws.Option) error

	DescribeElasticIps(*opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)
	DescribeElasticIpsWithContext(aws.Context, *opsworks.DescribeElasticIpsInput, ...aws.Option) (*opsworks.DescribeElasticIpsOutput, error)
	DescribeElasticIpsRequest(*opsworks.DescribeElasticIpsInput) (*aws.Request, *opsworks.DescribeElasticIpsOutput)

	DescribeElasticLoadBalancers(*opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)
	DescribeElasticLoadBalancersWithContext(aws.Context, *opsworks.DescribeElasticLoadBalancersInput, ...aws.Option) (*opsworks.DescribeElasticLoadBalancersOutput, error)
	DescribeElasticLoadBalancersRequest(*opsworks.DescribeElasticLoadBalancersInput) (*aws.Request, *opsworks.DescribeElasticLoadBalancersOutput)

	DescribeInstances(*opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)
	DescribeInstancesWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.Option) (*opsworks.DescribeInstancesOutput, error)
	DescribeInstancesRequest(*opsworks.DescribeInstancesInput) (*aws.Request, *opsworks.DescribeInstancesOutput)

	DescribeLayers(*opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)
	DescribeLayersWithContext(aws.Context, *opsworks.DescribeLayersInput, ...aws.Option) (*opsworks.DescribeLayersOutput, error)
	DescribeLayersRequest(*opsworks.DescribeLayersInput) (*aws.Request, *opsworks.DescribeLayersOutput)

	DescribeLoadBasedAutoScaling(*opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
	DescribeLoadBasedAutoScalingWithContext(aws.Context, *opsworks.DescribeLoadBasedAutoScalingInput, ...aws.Option) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
	DescribeLoadBasedAutoScalingRequest(*opsworks.DescribeLoadBasedAutoScalingInput) (*aws.Request, *opsworks.DescribeLoadBasedAutoScalingOutput)

	DescribeMyUserProfile(*opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)
	DescribeMyUserProfileWithContext(aws.Context, *opsworks.DescribeMyUserProfileInput, ...aws.Option) (*opsworks.DescribeMyUserProfileOutput, error)
	DescribeMyUserProfileRequest(*opsworks.DescribeMyUserProfileInput) (*aws.Request, *opsworks.DescribeMyUserProfileOutput)

	DescribePermissions(*opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)
	DescribePermissionsWithContext(aws.Context, *opsworks.DescribePermissionsInput, ...aws.Option) (*opsworks.DescribePermissionsOutput, error)
	DescribePermissionsRequest(*opsworks.DescribePermissionsInput) (*aws.Request, *opsworks.DescribePermissionsOutput)

	DescribeRaidArrays(*opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)
	DescribeRaidArraysWithContext(aws.Context, *opsworks.DescribeRaidArraysInput, ...aws.Option) (*opsworks.DescribeRaidArraysOutput, error)
	DescribeRaidArraysRequest(*opsworks.DescribeRaidArraysInput) (*aws.Request, *opsworks.DescribeRaidArraysOutput)

	DescribeRdsDbInstances(*opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)
	DescribeRdsDbInstancesWithContext(aws.Context, *opsworks.DescribeRdsDbInstancesInput, ...aws.Option) (*opsworks.DescribeRdsDbInstancesOutput, error)
	DescribeRdsDbInstancesRequest(*opsworks.DescribeRdsDbInstancesInput) (*aws.Request, *opsworks.DescribeRdsDbInstancesOutput)

	DescribeServiceErrors(*opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)
	DescribeServiceErrorsWithContext(aws.Context, *opsworks.DescribeServiceErrorsInput, ...aws.Option) (*opsworks.DescribeServiceErrorsOutput, error)
	DescribeServiceErrorsRequest(*opsworks.DescribeServiceErrorsInput) (*aws.Request, *opsworks.DescribeServiceErrorsOutput)

	DescribeStackProvisioningParameters(*opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)
	DescribeStackProvisioningParametersWithContext(aws.Context, *opsworks.DescribeStackProvisioningParametersInput, ...aws.Option) (*opsworks.DescribeStackProvisioningParametersOutput, error)
	DescribeStackProvisioningParametersRequest(*opsworks.DescribeStackProvisioningParametersInput) (*aws.Request, *opsworks.DescribeStackProvisioningParametersOutput)

	DescribeStackSummary(*opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)
	DescribeStackSummaryWithContext(aws.Context, *opsworks.DescribeStackSummaryInput, ...aws.Option) (*opsworks.DescribeStackSummaryOutput, error)
	DescribeStackSummaryRequest(*opsworks.DescribeStackSummaryInput) (*aws.Request, *opsworks.DescribeStackSummaryOutput)

	DescribeStacks(*opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)
	DescribeStacksWithContext(aws.Context, *opsworks.DescribeStacksInput, ...aws.Option) (*opsworks.DescribeStacksOutput, error)
	DescribeStacksRequest(*opsworks.DescribeStacksInput) (*aws.Request, *opsworks.DescribeStacksOutput)

	DescribeTimeBasedAutoScaling(*opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
	DescribeTimeBasedAutoScalingWithContext(aws.Context, *opsworks.DescribeTimeBasedAutoScalingInput, ...aws.Option) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
	DescribeTimeBasedAutoScalingRequest(*opsworks.DescribeTimeBasedAutoScalingInput) (*aws.Request, *opsworks.DescribeTimeBasedAutoScalingOutput)

	DescribeUserProfiles(*opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)
	DescribeUserProfilesWithContext(aws.Context, *opsworks.DescribeUserProfilesInput, ...aws.Option) (*opsworks.DescribeUserProfilesOutput, error)
	DescribeUserProfilesRequest(*opsworks.DescribeUserProfilesInput) (*aws.Request, *opsworks.DescribeUserProfilesOutput)

	DescribeVolumes(*opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)
	DescribeVolumesWithContext(aws.Context, *opsworks.DescribeVolumesInput, ...aws.Option) (*opsworks.DescribeVolumesOutput, error)
	DescribeVolumesRequest(*opsworks.DescribeVolumesInput) (*aws.Request, *opsworks.DescribeVolumesOutput)

	DetachElasticLoadBalancer(*opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)
	DetachElasticLoadBalancerWithContext(aws.Context, *opsworks.DetachElasticLoadBalancerInput, ...aws.Option) (*opsworks.DetachElasticLoadBalancerOutput, error)
	DetachElasticLoadBalancerRequest(*opsworks.DetachElasticLoadBalancerInput) (*aws.Request, *opsworks.DetachElasticLoadBalancerOutput)

	DisassociateElasticIp(*opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)
	DisassociateElasticIpWithContext(aws.Context, *opsworks.DisassociateElasticIpInput, ...aws.Option) (*opsworks.DisassociateElasticIpOutput, error)
	DisassociateElasticIpRequest(*opsworks.DisassociateElasticIpInput) (*aws.Request, *opsworks.DisassociateElasticIpOutput)

	GetHostnameSuggestion(*opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)
	GetHostnameSuggestionWithContext(aws.Context, *opsworks.GetHostnameSuggestionInput, ...aws.Option) (*opsworks.GetHostnameSuggestionOutput, error)
	GetHostnameSuggestionRequest(*opsworks.GetHostnameSuggestionInput) (*aws.Request, *opsworks.GetHostnameSuggestionOutput)

	GrantAccess(*opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)
	GrantAccessWithContext(aws.Context, *opsworks.GrantAccessInput, ...aws.Option) (*opsworks.GrantAccessOutput, error)
	GrantAccessRequest(*opsworks.GrantAccessInput) (*aws.Request, *opsworks.GrantAccessOutput)

	ListTags(*opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *opsworks.ListTagsInput, ...aws.Option) (*opsworks.ListTagsOutput, error)
	ListTagsRequest(*opsworks.ListTagsInput) (*aws.Request, *opsworks.ListTagsOutput)

	RebootInstance(*opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)
	RebootInstanceWithContext(aws.Context, *opsworks.RebootInstanceInput, ...aws.Option) (*opsworks.RebootInstanceOutput, error)
	RebootInstanceRequest(*opsworks.RebootInstanceInput) (*aws.Request, *opsworks.RebootInstanceOutput)

	RegisterEcsCluster(*opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)
	RegisterEcsClusterWithContext(aws.Context, *opsworks.RegisterEcsClusterInput, ...aws.Option) (*opsworks.RegisterEcsClusterOutput, error)
	RegisterEcsClusterRequest(*opsworks.RegisterEcsClusterInput) (*aws.Request, *opsworks.RegisterEcsClusterOutput)

	RegisterElasticIp(*opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)
	RegisterElasticIpWithContext(aws.Context, *opsworks.RegisterElasticIpInput, ...aws.Option) (*opsworks.RegisterElasticIpOutput, error)
	RegisterElasticIpRequest(*opsworks.RegisterElasticIpInput) (*aws.Request, *opsworks.RegisterElasticIpOutput)

	RegisterInstance(*opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)
	RegisterInstanceWithContext(aws.Context, *opsworks.RegisterInstanceInput, ...aws.Option) (*opsworks.RegisterInstanceOutput, error)
	RegisterInstanceRequest(*opsworks.RegisterInstanceInput) (*aws.Request, *opsworks.RegisterInstanceOutput)

	RegisterRdsDbInstance(*opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)
	RegisterRdsDbInstanceWithContext(aws.Context, *opsworks.RegisterRdsDbInstanceInput, ...aws.Option) (*opsworks.RegisterRdsDbInstanceOutput, error)
	RegisterRdsDbInstanceRequest(*opsworks.RegisterRdsDbInstanceInput) (*aws.Request, *opsworks.RegisterRdsDbInstanceOutput)

	RegisterVolume(*opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)
	RegisterVolumeWithContext(aws.Context, *opsworks.RegisterVolumeInput, ...aws.Option) (*opsworks.RegisterVolumeOutput, error)
	RegisterVolumeRequest(*opsworks.RegisterVolumeInput) (*aws.Request, *opsworks.RegisterVolumeOutput)

	SetLoadBasedAutoScaling(*opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)
	SetLoadBasedAutoScalingWithContext(aws.Context, *opsworks.SetLoadBasedAutoScalingInput, ...aws.Option) (*opsworks.SetLoadBasedAutoScalingOutput, error)
	SetLoadBasedAutoScalingRequest(*opsworks.SetLoadBasedAutoScalingInput) (*aws.Request, *opsworks.SetLoadBasedAutoScalingOutput)

	SetPermission(*opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)
	SetPermissionWithContext(aws.Context, *opsworks.SetPermissionInput, ...aws.Option) (*opsworks.SetPermissionOutput, error)
	SetPermissionRequest(*opsworks.SetPermissionInput) (*aws.Request, *opsworks.SetPermissionOutput)

	SetTimeBasedAutoScaling(*opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)
	SetTimeBasedAutoScalingWithContext(aws.Context, *opsworks.SetTimeBasedAutoScalingInput, ...aws.Option) (*opsworks.SetTimeBasedAutoScalingOutput, error)
	SetTimeBasedAutoScalingRequest(*opsworks.SetTimeBasedAutoScalingInput) (*aws.Request, *opsworks.SetTimeBasedAutoScalingOutput)

	StartInstance(*opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)
	StartInstanceWithContext(aws.Context, *opsworks.StartInstanceInput, ...aws.Option) (*opsworks.StartInstanceOutput, error)
	StartInstanceRequest(*opsworks.StartInstanceInput) (*aws.Request, *opsworks.StartInstanceOutput)

	StartStack(*opsworks.StartStackInput) (*opsworks.StartStackOutput, error)
	StartStackWithContext(aws.Context, *opsworks.StartStackInput, ...aws.Option) (*opsworks.StartStackOutput, error)
	StartStackRequest(*opsworks.StartStackInput) (*aws.Request, *opsworks.StartStackOutput)

	StopInstance(*opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)
	StopInstanceWithContext(aws.Context, *opsworks.StopInstanceInput, ...aws.Option) (*opsworks.StopInstanceOutput, error)
	StopInstanceRequest(*opsworks.StopInstanceInput) (*aws.Request, *opsworks.StopInstanceOutput)

	StopStack(*opsworks.StopStackInput) (*opsworks.StopStackOutput, error)
	StopStackWithContext(aws.Context, *opsworks.StopStackInput, ...aws.Option) (*opsworks.StopStackOutput, error)
	StopStackRequest(*opsworks.StopStackInput) (*aws.Request, *opsworks.StopStackOutput)

	TagResource(*opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *opsworks.TagResourceInput, ...aws.Option) (*opsworks.TagResourceOutput, error)
	TagResourceRequest(*opsworks.TagResourceInput) (*aws.Request, *opsworks.TagResourceOutput)

	UnassignInstance(*opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)
	UnassignInstanceWithContext(aws.Context, *opsworks.UnassignInstanceInput, ...aws.Option) (*opsworks.UnassignInstanceOutput, error)
	UnassignInstanceRequest(*opsworks.UnassignInstanceInput) (*aws.Request, *opsworks.UnassignInstanceOutput)

	UnassignVolume(*opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)
	UnassignVolumeWithContext(aws.Context, *opsworks.UnassignVolumeInput, ...aws.Option) (*opsworks.UnassignVolumeOutput, error)
	UnassignVolumeRequest(*opsworks.UnassignVolumeInput) (*aws.Request, *opsworks.UnassignVolumeOutput)

	UntagResource(*opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *opsworks.UntagResourceInput, ...aws.Option) (*opsworks.UntagResourceOutput, error)
	UntagResourceRequest(*opsworks.UntagResourceInput) (*aws.Request, *opsworks.UntagResourceOutput)

	UpdateApp(*opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)
	UpdateAppWithContext(aws.Context, *opsworks.UpdateAppInput, ...aws.Option) (*opsworks.UpdateAppOutput, error)
	UpdateAppRequest(*opsworks.UpdateAppInput) (*aws.Request, *opsworks.UpdateAppOutput)

	UpdateElasticIp(*opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)
	UpdateElasticIpWithContext(aws.Context, *opsworks.UpdateElasticIpInput, ...aws.Option) (*opsworks.UpdateElasticIpOutput, error)
	UpdateElasticIpRequest(*opsworks.UpdateElasticIpInput) (*aws.Request, *opsworks.UpdateElasticIpOutput)

	UpdateInstance(*opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)
	UpdateInstanceWithContext(aws.Context, *opsworks.UpdateInstanceInput, ...aws.Option) (*opsworks.UpdateInstanceOutput, error)
	UpdateInstanceRequest(*opsworks.UpdateInstanceInput) (*aws.Request, *opsworks.UpdateInstanceOutput)

	UpdateLayer(*opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)
	UpdateLayerWithContext(aws.Context, *opsworks.UpdateLayerInput, ...aws.Option) (*opsworks.UpdateLayerOutput, error)
	UpdateLayerRequest(*opsworks.UpdateLayerInput) (*aws.Request, *opsworks.UpdateLayerOutput)

	UpdateMyUserProfile(*opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)
	UpdateMyUserProfileWithContext(aws.Context, *opsworks.UpdateMyUserProfileInput, ...aws.Option) (*opsworks.UpdateMyUserProfileOutput, error)
	UpdateMyUserProfileRequest(*opsworks.UpdateMyUserProfileInput) (*aws.Request, *opsworks.UpdateMyUserProfileOutput)

	UpdateRdsDbInstance(*opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)
	UpdateRdsDbInstanceWithContext(aws.Context, *opsworks.UpdateRdsDbInstanceInput, ...aws.Option) (*opsworks.UpdateRdsDbInstanceOutput, error)
	UpdateRdsDbInstanceRequest(*opsworks.UpdateRdsDbInstanceInput) (*aws.Request, *opsworks.UpdateRdsDbInstanceOutput)

	UpdateStack(*opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)
	UpdateStackWithContext(aws.Context, *opsworks.UpdateStackInput, ...aws.Option) (*opsworks.UpdateStackOutput, error)
	UpdateStackRequest(*opsworks.UpdateStackInput) (*aws.Request, *opsworks.UpdateStackOutput)

	UpdateUserProfile(*opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)
	UpdateUserProfileWithContext(aws.Context, *opsworks.UpdateUserProfileInput, ...aws.Option) (*opsworks.UpdateUserProfileOutput, error)
	UpdateUserProfileRequest(*opsworks.UpdateUserProfileInput) (*aws.Request, *opsworks.UpdateUserProfileOutput)

	UpdateVolume(*opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
	UpdateVolumeWithContext(aws.Context, *opsworks.UpdateVolumeInput, ...aws.Option) (*opsworks.UpdateVolumeOutput, error)
	UpdateVolumeRequest(*opsworks.UpdateVolumeInput) (*aws.Request, *opsworks.UpdateVolumeOutput)

	WaitUntilAppExists(*opsworks.DescribeAppsInput) error
	WaitUntilAppExistsWithContext(aws.Context, *opsworks.DescribeAppsInput, ...aws.WaiterOption) error

	WaitUntilDeploymentSuccessful(*opsworks.DescribeDeploymentsInput) error
	WaitUntilDeploymentSuccessfulWithContext(aws.Context, *opsworks.DescribeDeploymentsInput, ...aws.WaiterOption) error

	WaitUntilInstanceOnline(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceOnlineWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceRegistered(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceRegisteredWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceStopped(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceStoppedWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceTerminated(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceTerminatedWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error
}

var _ OpsWorksAPI = (*opsworks.OpsWorks)(nil)