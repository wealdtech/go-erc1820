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
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImplementerInterface(t *testing.T) {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.Nil(t, err, "failed to acces Infura")
	registry, err := NewRegistry(client)
	require.Nil(t, err, "failed to acces ERC-1820 registry")
	addr := common.HexToAddress("907b4EB76F423595408C49c0BbB2bb117C91c594")
	implementer, err := registry.InterfaceImplementer("ERC777Token", &addr)
	require.Nil(t, err, "failed to access registry")
	expected := common.HexToAddress("907b4EB76F423595408C49c0BbB2bb117C91c594")
	assert.Equal(t, *implementer, expected, "failed to obtain correct implementer address")
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
