//
// Copyright 2021, Offchain Labs, Inc. All rights reserved.
//

package precompiles

import (
	"errors"
	"math/big"
)

type ArbFunctionTable struct {
	Address addr
}

func (con ArbFunctionTable) Upload(c ctx, evm mech, buf []byte) error {
	return nil
}

func (con ArbFunctionTable) Size(c ctx, evm mech, addr addr) (huge, error) {
	return big.NewInt(0), nil
}

func (con ArbFunctionTable) Get(c ctx, evm mech, addr addr, index huge) (huge, bool, huge, error) {
	return nil, false, nil, errors.New("table is empty")
}
