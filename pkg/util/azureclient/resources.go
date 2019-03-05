package azureclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

// DeploymentsClient is a minimal interface for azure DeploymentsClient
type DeploymentsClient interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, parameters resources.Deployment) (result resources.DeploymentsCreateOrUpdateFuture, err error)
	Client
	DeploymentClient() resources.DeploymentsClient
}

type deploymentsClient struct {
	resources.DeploymentsClient
}

var _ DeploymentsClient = &deploymentsClient{}

// NewDeploymentsClient creates a new DeploymentsClient
func NewDeploymentsClient(ctx context.Context, subscriptionID string, authorizer autorest.Authorizer) DeploymentsClient {
	client := resources.NewDeploymentsClient(subscriptionID)
	setupClient(ctx, &client.Client, authorizer)

	return &deploymentsClient{
		DeploymentsClient: client,
	}
}

func (c *deploymentsClient) Client() autorest.Client {
	return c.DeploymentsClient.Client
}

func (c *deploymentsClient) DeploymentClient() resources.DeploymentsClient {
	return c.DeploymentsClient
}

// ResourcesClient is a minimal interface for azure Resources Client
type ResourcesClient interface {
	DeleteByID(ctx context.Context, resourceID string) (result resources.DeleteByIDFuture, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
}

type resourcesClient struct {
	resources.Client
}

var _ ResourcesClient = &resourcesClient{}

// NewResourcesClient creates a new ResourcesClient
func NewResourcesClient(ctx context.Context, subscriptionID string, authorizer autorest.Authorizer) ResourcesClient {
	client := resources.NewClient(subscriptionID)
	setupClient(ctx, &client.Client, authorizer)

	return &resourcesClient{
		Client: client,
	}
}

// GroupsClient is a minimal interface for azure Resources Client
type GroupsClient interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters resources.Group) (result resources.Group, err error)
	List(ctx context.Context, filter string, top *int32) (result resources.GroupListResultPage, err error)
	Delete(ctx context.Context, resourceGroupName string) (result resources.GroupsDeleteFuture, err error)
	Client
}

type groupsClient struct {
	resources.GroupsClient
}

var _ GroupsClient = &groupsClient{}

// NewGroupsClient creates a new ResourcesClient
func NewGroupsClient(ctx context.Context, subscriptionID string, authorizer autorest.Authorizer) GroupsClient {
	client := resources.NewGroupsClient(subscriptionID)
	setupClient(ctx, &client.Client, authorizer)

	return &groupsClient{
		GroupsClient: client,
	}
}

func (c *groupsClient) Client() autorest.Client {
	return c.GroupsClient.Client
}
