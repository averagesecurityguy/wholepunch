package wp

import (
	"io"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// WriteTableBeaconResults writes beacon results to a ascii table.
func WriteTableBeaconResults(writer io.Writer, results []BeaconResult, filterFalse bool) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{
		"Name",
		"Success",
		"Destination",
		"Info",
	})
	for _, r := range results {
		if filterFalse && !r.WasOk {
			continue
		}
		table.Append([]string{
			r.Name,
			strconv.FormatBool(r.WasOk),
			r.Destination,
			r.Info,
		})
	}
	table.Render()
}
