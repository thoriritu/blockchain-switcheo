package types

const (
	// ModuleName defines the module name
	ModuleName = "emarket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_emarket"
)

var (
	ParamsKey = []byte("p_emarket")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
