package storage_market

import (
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"

	"github.com/filecoin-project/specs-actors/actors/abi"
	. "github.com/filecoin-project/specs-actors/actors/util/adt"
)

// A specialization of a map of addresses to token amounts.
// It is an error to query for a key that doesn't exist.
type DealArray struct {
	*Array
}

// Interprets a store as balance table with root `r`.
func AsDealArray(s Store, r cid.Cid) *DealArray {
	return &DealArray{AsArray(s, r)}
}

// Returns the root cid of underlying AMT.
func (t *DealArray) Root() cid.Cid {
	return t.Array.Root()
}

// Gets the deal for a key. The entry must have been previously initialized.
func (t *DealArray) Get(id abi.DealID) (*OnChainDeal, error) {
	var value OnChainDeal
	found, err := t.Array.Get(uint64(id), &value)
	if err != nil {
		return nil, err // The errors from Map carry good information, no need to wrap here.
	}
	if !found {
		return nil, errors.Errorf("deal %d not found", id)
	}
	return &value, nil
}

func (t *DealArray) Set(k abi.DealID, value *OnChainDeal) error {
	return t.Array.Set(uint64(k), value)
}

func (t *DealArray) Delete(key uint64) error {
	return t.Array.Delete(key)
}