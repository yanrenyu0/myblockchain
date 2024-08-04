package types

const (
	// ModuleName defines the module name
	ModuleName = "myblockchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_myblockchain"

	// ResourceKey is used to uniquely identify resources within the system.
	// It will be used as the beginning of the key for each resource, followed bei their unique ID
	ResourceKey = "Resource/value/"

	// This key will be used to keep track of the ID of the latest resource added to the store.
	ResourceCountKey = "Resource/count/"
)

var (
	ParamsKey = []byte("p_myblockchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
