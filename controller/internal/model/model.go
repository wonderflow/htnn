// Copyright The HTNN Authors.
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

package model

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
)

type GatewaySection struct {
	// Fields here can't be pointer because we use GatewaySection as map key
	NsName      types.NamespacedName
	SectionName string
}

func (g GatewaySection) String() string {
	return fmt.Sprintf("%s/%s", g.NsName.String(), g.SectionName)
}

type Gateway struct {
	GatewaySection *GatewaySection
	// HasHCM shows if the HCM HTTP filter is present in the gateway
	HasHCM bool
}

type VirtualHost struct {
	GatewaySection *GatewaySection
	// NsName is the namespace and name of the k8s resource which generates VirtualHost
	NsName *types.NamespacedName
	// Name is the name of VirtualHost in the RDS
	Name string
	// ECDSResourceName is the name of ECDS which is used to configure the Gateway attached by this VirtualHost
	ECDSResourceName string
}

const (
	CategoryECDSGolang   = "ecds_golang"
	CategoryECDSListener = "ecds_listener"
	CategoryECDSNetwork  = "ecds_network"
	CategoryListener     = "listener"

	CategoryRoute       = "route"
	CategoryRouteFilter = "route_filter"

	// This constant is used in the resource name which doesn't support '_' in the name
	ECDSGolangPlugins = "golang-filter"
)
