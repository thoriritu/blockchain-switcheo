package types

const (
	// ModuleName defines the module name
	ModuleName = "switcheo"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_switcheo"
)

var (
	ParamsKey = []byte("p_switcheo")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ItemKey      = "Item/value/"
	ItemCountKey = "Item/count/"
)
