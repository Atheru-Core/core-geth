//go:build windows
// +build windows

package vm

// Stub functions for Windows - EVMC is not available, so always use default interpreter

func getEvmcCapabilityEVM1() int {
	return evmcCapabilityEVM1
}

func getEvmcCapabilityEWASM() int {
	return evmcCapabilityEWASM
}

// On Windows, EVMC is not available, so always return the default interpreter
func newEVMCInterpreter(instance interface{}, env *EVM, cap int, readOnly bool) Interpreter {
	// EVMC not supported on Windows - fall back to default interpreter
	return NewEVMInterpreter(env)
}
