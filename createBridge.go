package main


import (
    "fmt"
    "net"
    "os"
    "syscall"

    "github.com/vishvananda/netlink"
    )


func setupBridge(){

    if k,_ = netlink.LinkByName(bridge), 
    la:= netlink.NewLinkAttrs()
    la.Name = "foo"
    mybridge:= &netlink.Bridge{LinkAttrs: la}
    err := netlink.LinkAdd(mybridge)
    if err != nil {
        fmt.Printf("Could not add %s: %v\n", la.Name, err)

    }
    eth1, _ := netlink.LinkByName("eth1")
    netlink.LinkSetMaster(eth1,mybridge)

}
