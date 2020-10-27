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
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/wealdtech/go-erc1820/contracts"
	"golang.org/x/crypto/sha3"
)

// _contractAddress is the well-known address of the ERC-1820 contract on all networks
var _contractAddress = common.HexToAddress("1820a4b7618bde71dce8cdc73aab6c95905fad24")

// Registry is the struct that holds information about the ERC-1820 registry
type Registry struct {
	backend  bind.ContractBackend
	contract *contracts.Erc1820Registry
}

// NewRegistry creates a new ERC-1820 registry client
func NewRegistry(backend bind.ContractBackend) (*Registry, error) {
	contract, err := contracts.NewErc1820Registry(_contractAddress, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{
		backend:  backend,
		contract: contract,
	}, nil
}

// InterfaceImplementer returns the address of the implementer of the given interface for the given address
func (r *Registry) InterfaceImplementer(iface string, address *common.Address) (*common.Address, error) {
	var hash [32]byte
	if len(iface) == 10 && strings.HasPrefix(iface, "0x") {
		// Looks like an ERC-165 interface ID
		bytes, err := hex.DecodeString(iface[2:])
		if err != nil {
			return nil, err
		}
		copy(hash[:], bytes)
	} else {
		keccak := sha3.NewLegacyKeccak256()
		_, err := keccak.Write([]byte(iface))
		if err != nil {
			return nil, errors.Wrap(err, "failed to write bytes to hash")
		}
		copy(hash[:], keccak.Sum(nil))
	}
	implementer, err := r.contract.GetInterfaceImplementer(nil, *address, hash)
	return &implementer, err
}

// SetInterfaceImplementer sets the address of the implementer of the given interface for the given address
func (r *Registry) SetInterfaceImplementer(opts *bind.TransactOpts, iface string, address *common.Address, implementer *common.Address) (*types.Transaction, error) {
	keccak := sha3.NewLegacyKeccak256()
	_, err := keccak.Write([]byte(iface))
	if err != nil {
		return nil, errors.Wrap(err, "failed to write bytes to hash")
	}
	var hash [32]byte
	copy(hash[:], keccak.Sum(nil))
	tx, err := r.contract.SetInterfaceImplementer(opts, *address, hash, *implementer)
	if err == nil {
		return tx, nil
	}

	// An error occurred.  The most common reason for this is that the implementer doesn't implement the interface, so check for
	// this situation and if so return an appropriate error
	implementerContract, err2 := NewImplementer(r.backend, implementer)
	if err2 != nil {
		return nil, err
	}
	implements, err2 := implementerContract.ImplementsInterface(iface, address)
	if err2 != nil {
		return nil, err
	}
	if !implements {
		return nil, fmt.Errorf("%s does not implement %s", implementer.Hex(), iface)
	}
	// Some other error; just return it
	return nil, err
}

// Manager returns the manager of the given address
func (r *Registry) Manager(address *common.Address) (*common.Address, error) {
	manager, err := r.contract.GetManager(nil, *address)
	return &manager, err
}

// SetManager sets the manager of the given address
func (r *Registry) SetManager(opts *bind.TransactOpts, address *common.Address, manager *common.Address) (*types.Transaction, error) {
	tx, err := r.contract.SetManager(opts, *address, *manager)
	if err == nil {
		return tx, nil
	}

	// An error occurred.  Try to work out what the problem is and report a more useful error
	existingManager, err2 := r.Manager(address)
	if err2 != nil {
		return nil, err
	}
	if bytes.Equal(existingManager.Bytes(), address.Bytes()) {
		return nil, fmt.Errorf("new manager can only be set by existing manager %s", existingManager.Hex())
	}
	return nil, err
}
