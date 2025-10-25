package commandutil

import "github.com/spf13/cobra"

type FlagRegister func(*cobra.Command)

var flagRegisters = map[RegisterToken][]FlagRegister{}

// "RegisterIntoFlagChunk" flagRegisters a flag into a flag chunk.
//
// Flag chunks can be utilized later to consistently insert flags into commands
func RegisterFlagConfiguration(token RegisterToken, registor FlagRegister) {
	if _, found := flagRegisters[token]; !found {
		flagRegisters[token] = []FlagRegister{registor}
	} else {
		flagRegisters[token] = append(flagRegisters[token], registor)
	}
}

// "RegisterFromFlagChunk" flagRegisters flags into the providen command.
func ConsumeFlagConfiguration(token RegisterToken, command *cobra.Command) {
	if FlagRegister, found := flagRegisters[token]; found {
		for _, FlagRegister := range FlagRegister {
			FlagRegister(command)
		}
	}
}
