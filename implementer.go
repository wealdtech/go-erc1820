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
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-erc1820/contracts"
	"golang.org/x/crypto/sha3"
)

// Implementer is the struct that holds information about an ERC-1820 interface implementer
type Implementer struct {
	contract *contracts.Erc1820Implementer
}

// NewImplementer creates a new ERC-1820 registry implementer client
func NewImplementer(backend bind.ContractBackend, address *common.Address) (*Implementer, error) {
	contract, err := contracts.NewErc1820Implementer(*address, backend)
	if err != nil {
		return nil, err
	}
	return &Implementer{
		contract: contract,
	}, nil
}

// _acceptMagic = keccak256("ERC1820_ACCEPT_MAGIC")
var _acceptMagic = common.HexToHash("a2ef4600d742022d532d4747cb3547474667d6f13804902513b2ec01c848f4b4").Bytes()

// ImplementsInterface returns true if the address implements the given interface for the given address
func (i *Implementer) ImplementsInterface(iface string, address *common.Address) (bool, error) {
	keccak := sha3.NewLegacyKeccak256()
	keccak.Write([]byte(iface))
	var hash [32]byte
	copy(hash[:], keccak.Sum(nil))
	result, err := i.contract.CanImplementInterfaceForAddress(nil, hash, *address)
	if err != nil {
		if err.Error() == "no contract code at given address" {
			// address is not a contract so cannot implement iface
			return false, nil
		}
		if err.Error() == "abi: unmarshalling empty output" {
			// address does not implement iface
			return false, nil
		}
		// Some other error
		return false, err
	}
	return bytes.Compare(result[:], _acceptMagic) == 0, nil
}
