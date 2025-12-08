//go:build windows
// +build windows

// Stub implementation of EVMC for Windows builds
// EVMC is not supported on Windows due to build constraints in the evmc package

package vm

import (
	"sync"
)

// Stub types for Windows - EVMC is not available on Windows
type evmcVM struct{}

// Stub capability constants to match evmc package
const (
	evmcCapabilityEVM1  = 0
	evmcCapabilityEWASM = 1
)

// EVMC represents a stub EVMC interpreter (not available on Windows)
// This must match the struct in evmc.go for compatibility
// On Windows, we use compatible types since evmc package is not available
type EVMC struct {
	instance interface{} // Stub - *evmc.VM on non-Windows, *evmcVM on Windows
	env      *EVM
	cap      int // Stub capability - evmc.Capability on non-Windows (which is also int-based)
	readOnly bool
}

var (
	evmModule   *evmcVM
	ewasmModule *evmcVM
	evmcMux     sync.Mutex
)

// InitEVMCEVM is a stub function for Windows
func InitEVMCEVM(config string) {
	// EVMC not supported on Windows - this is a no-op
}

// InitEVMCEwasm is a stub function for Windows
func InitEVMCEwasm(config string) {
	// EVMC not supported on Windows - this is a no-op
}
