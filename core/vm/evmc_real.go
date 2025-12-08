//go:build !windows
// +build !windows

package vm

import "github.com/ethereum/evmc/v7/bindings/go/evmc"

// Real evmc capability functions for non-Windows platforms

func getEvmcCapabilityEVM1() int {
	return int(evmc.CapabilityEVM1)
}

func getEvmcCapabilityEWASM() int {
	return int(evmc.CapabilityEWASM)
}

// On non-Windows, create real EVMC interpreter
func newEVMCInterpreter(instance interface{}, env *EVM, cap int, readOnly bool) Interpreter {
	return &EVMC{
		instance: instance.(*evmc.VM),
		env:      env,
		cap:      evmc.Capability(cap),
		readOnly: readOnly,
	}
}
