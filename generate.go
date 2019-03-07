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

//go:generate abigen -abi erc1820registry.abi -out erc1820registry.go -pkg erc1820 -type Erc1820Registry
//go:generate abigen -abi erc1820implementer.abi -out erc1820implementer.go -pkg erc1820 -type Erc1820Implementer
