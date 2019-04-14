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

func TestImplementer(t *testing.T) {
	tests := []struct {
		contractAddr common.Address
		iface        string
		forAddr      common.Address
		res          bool
		err          error
	}{
		{
			contractAddr: common.HexToAddress("62284ed69b907af90ecba2feef0bf12a99563563"),
			iface:        "ERC777TokensSender",
			forAddr:      common.HexToAddress("00"),
			res:          true,
		},
		{
			contractAddr: common.HexToAddress("62284ed69b907af90ecba2feef0bf12a99563563"),
			iface:        "ERC777TokensFlanger",
			forAddr:      common.HexToAddress("00"),
			res:          false,
		},
		{
			contractAddr: common.HexToAddress("00"),
			iface:        "ERC777TokensSender",
			forAddr:      common.HexToAddress("00"),
			res:          false,
		},
	}

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.Nil(t, err, "failed to acces Infura")

	for i, test := range tests {
		contract, err := NewImplementer(client, &test.contractAddr)
		require.Nil(t, err, fmt.Sprintf("failed to access implementer at test %d", i))
		res, err := contract.ImplementsInterface(test.iface, &test.forAddr)
		if test.err != nil {
			require.Equal(t, test.err, err, fmt.Sprintf("Incorrect error at test %d", i))
		} else {
			require.Nil(t, err, fmt.Sprintf("Unexpected error at test %d", i))
			assert.Equal(t, test.res, res, fmt.Sprintf("Incorrect result at test %d", i))
		}
	}
}
