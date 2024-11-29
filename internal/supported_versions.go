// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package internal

import (
	"github.com/containernetworking/cni/pkg/version"
)

func SupportedVersions() []string {
	// support CNI versions that support plugin chaining
	supported := []string{}
	unsupported := map[string]bool{"0.1.0": true, "0.2.0": true}

	for _, v := range version.All.SupportedVersions() {
		if _, ok := unsupported[v]; !ok {
			supported = append(supported, v)
		}
	}

	return supported
}
