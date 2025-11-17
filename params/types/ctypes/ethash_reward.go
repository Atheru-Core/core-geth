// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package ctypes

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params/vars"
	"github.com/holiman/uint256"
)

func EthashBlockReward(c ChainConfigurator, n *big.Int) *uint256.Int {
	// Custom block reward schedule for chain ID 192
	// Initial reward: 4 tokens per block
	// Halving every 8,307,692 blocks
	if c != nil && n != nil && c.GetChainID().Cmp(big.NewInt(192)) == 0 {
		halvingPeriod := big.NewInt(8_307_692)
		initialReward := uint256.NewInt(4e+18) // 4 tokens in wei
		
		// Calculate which halving period we're in
		halvingCount := new(big.Int).Div(n, halvingPeriod)
		
		// Calculate reward: initialReward / (2^halvingCount)
		reward := new(uint256.Int).Set(initialReward)
		for i := big.NewInt(0); i.Cmp(halvingCount) < 0; i.Add(i, big.NewInt(1)) {
			reward.Div(reward, uint256.NewInt(2))
		}
		
		// Minimum reward is 0 (will stop after enough halvings)
		if reward.IsZero() {
			return uint256.NewInt(0)
		}
		return reward
	}
	
	// Select the correct block reward based on chain progression
	blockReward := vars.FrontierBlockReward
	if c == nil || n == nil {
		return blockReward
	}

	if c.IsEnabled(c.GetEthashEIP1234Transition, n) {
		return vars.EIP1234FBlockReward
	} else if c.IsEnabled(c.GetEthashEIP649Transition, n) {
		return vars.EIP649FBlockReward
	} else if len(c.GetEthashBlockRewardSchedule()) > 0 {
		// Because the map is not necessarily sorted low-high, we
		// have to ensure that we're walking upwards only.
		var lastActivation uint64
		for activation, reward := range c.GetEthashBlockRewardSchedule() {
			if activation <= n.Uint64() { // Is forked
				if activation >= lastActivation {
					lastActivation = activation
					blockReward = reward
				}
			}
		}
	}

	return blockReward
}
