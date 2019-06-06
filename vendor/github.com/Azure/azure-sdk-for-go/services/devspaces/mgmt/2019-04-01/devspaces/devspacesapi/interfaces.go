package devspacesapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/devspaces/mgmt/2019-04-01/devspaces"
)

// ContainerHostMappingsClientAPI contains the set of methods on the ContainerHostMappingsClient type.
type ContainerHostMappingsClientAPI interface {
	GetContainerHostMapping(ctx context.Context, containerHostMapping devspaces.ContainerHostMapping, resourceGroupName string, location string) (result devspaces.ContainerHostMapping, err error)
}

var _ ContainerHostMappingsClientAPI = (*devspaces.ContainerHostMappingsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result devspaces.ResourceProviderOperationListPage, err error)
}

var _ OperationsClientAPI = (*devspaces.OperationsClient)(nil)

// ControllersClientAPI contains the set of methods on the ControllersClient type.
type ControllersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, name string, controller devspaces.Controller) (result devspaces.ControllersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result devspaces.ControllersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, name string) (result devspaces.Controller, err error)
	List(ctx context.Context) (result devspaces.ControllerListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result devspaces.ControllerListPage, err error)
	ListConnectionDetails(ctx context.Context, resourceGroupName string, name string, listConnectionDetailsParameters devspaces.ListConnectionDetailsParameters) (result devspaces.ControllerConnectionDetailsList, err error)
	Update(ctx context.Context, resourceGroupName string, name string, controllerUpdateParameters devspaces.ControllerUpdateParameters) (result devspaces.Controller, err error)
}

var _ ControllersClientAPI = (*devspaces.ControllersClient)(nil)
