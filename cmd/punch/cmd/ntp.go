package cmd

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"sync"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
	"github.com/tomsteele/xplode"
)

var (
	flNTPServerPort string
	flNTPTimeout    int
	flNTPWorkers    int
)

func init() {
	beaconCmd.AddCommand(ntpCmd)
	ntpCmd.PersistentFlags().StringVar(&flNTPServerPort, "p", "1-1024", "NMap style port string.")
	ntpCmd.PersistentFlags().IntVar(&flNTPTimeout, "timeout", 500, "Timeout in milliseconds.")
	ntpCmd.PersistentFlags().IntVar(&flNTPWorkers, "c", 50, "Max number of concurrent requests.")
}

func ntp(cmd *cobra.Command, args []string) {
	ports, err := xplode.Parse(flNTPServerPort)
	if err != nil {
		fmt.Println("There was an error parsing the port string.")
		fmt.Println(err)
		os.Exit(1)
	}
	results := []wp.BeaconResult{}

	mutex := sync.Mutex{}
	portChan := make(chan int)
	doneChan := make(chan bool)

	w := wow.New(os.Stdout, spin.Get(spin.Pipe), "Working")
	w.Start()
	for i := 0; i < flNTPWorkers; i++ {
		go func() {
			for p := range portChan {
				b := wp.BeaconNTP{
					Timeout: flTCPTimeout,
				}
				opts := wp.BeaconOptions{
					DestinationServerAddress: fmt.Sprintf("%s:%d", flBeaconServerAddr, p),
				}
				ok, err := wp.RunBeacon(&b, &opts)
				result := wp.MakeBeaconResult(ok, err, &b)
				mutex.Lock()
				results = append(results, result)
				mutex.Unlock()
				doneChan <- true
			}
		}()
	}
	go func() {
		for _, p := range ports {
			portChan <- p
		}
	}()
	for i := 0; i < len(ports); i++ {
		<-doneChan
	}
	close(portChan)
	close(doneChan)
	w.Stop()
	fmt.Println()

	sort.Slice(results, func(i, j int) bool {
		iurl, _ := url.Parse(results[i].Destination)
		jurl, _ := url.Parse(results[j].Destination)
		i, _ = strconv.Atoi(iurl.Port())
		j, _ = strconv.Atoi(jurl.Port())
		return i < j
	})
	wp.WriteTableBeaconResults(os.Stdout, results, flBeaconFilterFalse)
}

var ntpCmd = &cobra.Command{
	Use:   "ntp",
	Short: "Initiates NTP connections to the destination server.",
	Run:   ntp,
}
