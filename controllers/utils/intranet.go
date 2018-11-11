package utils

import (
    "net"
    "net/http"
    "strings"
)

// IsRequestFromIntranet - checks whether incomming request is from intranet or not
func IsRequestFromIntranet(r *http.Request) bool {
    isPrivate := false
    decomposedAddr := strings.Split(r.RemoteAddr, ":")
    pureIP := decomposedAddr[0]
    IP := net.ParseIP(pureIP)

    _, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
    _, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
    _, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")

    isPrivate = private24BitBlock.Contains(IP) || private20BitBlock.Contains(IP) || private16BitBlock.Contains(IP)

    return isPrivate
}

