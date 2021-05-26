// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package statemgr

import (
	"time"
)

const (
	// period of state pull retry
	defaultPullStateRetryConst = 1 * time.Second
	// how long delay state pull after state candidate received
	defaultPullStateAfterStateCandidateDelayConst = 1 * time.Second
	defaultGetBlockRetryConst                     = 3 * time.Second
)

type Timers struct {
	PullStateRetry         *time.Duration
	PullStateNewBlockDelay *time.Duration
	GetBlockRetry          *time.Duration
}

func (tT Timers) SetPullStateRetry(pullStateRetry time.Duration) Timers {
	tT.PullStateRetry = &pullStateRetry
	return tT
}

func (tT Timers) SetPullStateNewBlockDelay(pullStateNewBlockDelay time.Duration) Timers {
	tT.PullStateNewBlockDelay = &pullStateNewBlockDelay
	return tT
}

func (tT Timers) SetGetBlockRetry(getBlockRetry time.Duration) Timers {
	tT.GetBlockRetry = &getBlockRetry
	return tT
}

func (tT Timers) getPullStateRetry() time.Duration {
	if tT.PullStateRetry == nil {
		return defaultPullStateRetryConst
	}
	return *tT.PullStateRetry
}

func (tT Timers) getPullStateNewBlockDelay() time.Duration {
	if tT.PullStateNewBlockDelay == nil {
		return defaultPullStateAfterStateCandidateDelayConst
	}
	return *tT.PullStateNewBlockDelay
}

func (tT Timers) getGetBlockRetry() time.Duration {
	if tT.GetBlockRetry == nil {
		return defaultGetBlockRetryConst
	}
	return *tT.GetBlockRetry
}
