// Copyright © 2019 Weald Technology Trading
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

package erc1820

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistryImplementer(t *testing.T) {
	tests := []struct {
		iface       string
		addr        common.Address
		implementer common.Address
		err         error
	}{
		{
			iface:       "ERC777TokensSender",
			addr:        common.HexToAddress("7598ac132c987A2eEa0106e2E5B6e67244349071"),
			implementer: common.HexToAddress("5182f98e2857e64a8e177404cf2b23ace4aa7967"),
		},
		{
			iface:       "0x80ac58cd",
			addr:        common.HexToAddress("418dbc78b50f52eaa5b029241c430da44133a617"),
			implementer: common.HexToAddress("418dbc78b50f52eaa5b029241c430da44133a617"),
		},
	}

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.Nil(t, err, "failed to acces Infura")
	registry, err := NewRegistry(client)
	require.Nil(t, err, "failed to acces ERC-1820 registry")

	for i, test := range tests {
		implementer, err := registry.InterfaceImplementer(test.iface, &test.addr)
		if test.err != nil {
			assert.Equal(t, test.err, err, fmt.Sprintf("Incorrect error at test %d", i))
		} else {
			assert.Nil(t, test.err, fmt.Sprintf("Unexpected error at test %d", i))
			assert.Equal(t, test.implementer, *implementer, fmt.Sprintf("Unexpected value at test %d", i))
		}
	}
}

func TestManager(t *testing.T) {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.Nil(t, err, "failed to acces Infura")
	registry, err := NewRegistry(client)
	addr := common.HexToAddress("7598ac132c987A2eEa0106e2E5B6e67244349071")
	manager, err := registry.Manager(&addr)
	require.Nil(t, err, "failed to access registry")
	expected := common.HexToAddress("3b5988414d3af6e7cb7896eA1CF1C83B3cfF503F")
	assert.Equal(t, *manager, expected, "failed to obtain correct manager address")
}
