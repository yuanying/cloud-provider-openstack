/*
Copyright 2018 The Kubernetes Authors.

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

package sharebackends

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/shares"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/cloud-provider-openstack/pkg/share/manila/shareoptions"
)

// BuildSourceArgs contains arguments for ShareBackend.BuildSource()
type BuildSourceArgs struct {
	Share       *shares.Share
	Options     *shareoptions.ShareOptions
	Location    *shares.ExportLocation
	Clientset   clientset.Interface
	AccessRight *shares.AccessRight
}

// GrantAccessArgs contains arguments for ShareBackend.GrantAccess()
type GrantAccessArgs struct {
	Share     *shares.Share
	Options   *shareoptions.ShareOptions
	Clientset clientset.Interface
	Client    *gophercloud.ServiceClient
}

// RevokeAccessArgs contains arguments for ShareBaceknd.RevokeAccess()
type RevokeAccessArgs struct {
	ShareID         string
	SecretNamespace string
	Clientset       clientset.Interface
}
