// +build !android

package main

import kcp "github.com/xtaci/kcp-go/v5"

func DialKCP(raddr string, block kcp.BlockCrypt, dataShards, parityShards int) (*kcp.UDPSession, error) {
	return kcp.DialWithOptions(raddr, block, dataShards, parityShards)
}
