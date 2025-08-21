package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"koizumib/ipinfo/internal/netcalc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <CIDR> [<CIDR> ...]\n", os.Args[0])
		os.Exit(1)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "IPAddress\tSubnetMask\tNetworkAddress\tBroadcastAddress\tHostRange\tHosts")
	for _, arg := range os.Args[1:] {
		row, err := netcalc.RowFromCIDR(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", arg, err)
			continue
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\n",
			row.IPAddress, row.SubnetMask, row.NetworkAddress,
			row.BroadcastAddress, row.HostRange, row.Hosts)
	}
	w.Flush()
}
