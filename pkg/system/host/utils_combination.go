package host

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/arr"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/remote/v1/remote"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/times/utilstime"
)

func GetCombinationData(debug bool) Combination {
	var comb Combination
	chServer := make(chan string)
	chTime := make(chan *utilstime.Time)
	chUptime := make(chan remote.Details)
	chMem := make(chan *arr.Array)
	chKernel := make(chan KernelInfo)
	chPro := make(chan ProgressInfo)
	chDisk := make(chan DiskInfoDetail)

	go func() { chServer <- getServerInfo(debug) }()
	go func() { chTime <- utilstime.SystemTime() }()
	go func() { chUptime <- remote.Info() }()
	go func() { chMem <- GetMemory() }()
	go func() { chKernel <- getKernelData(debug) }()
	go func() { chPro <- getProgressData(debug) }()
	go func() { chDisk <- getDiskInfoDetail(debug) }()

	comb.ServerInfo = <-chServer
	comb.SystemTime = <-chTime
	comb.Uptime = <-chUptime
	comb.MemInfo = <-chMem
	comb.KernelInfo = <-chKernel
	comb.ProgressInfo = <-chPro
	comb.DiskInfo = <-chDisk

	return comb
}
