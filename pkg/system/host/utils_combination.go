package host

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/arr"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/admin"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/utilstime"
)

func GetCombinationData(debug bool) Combination {
	var comb Combination
	chServer := make(chan string)
	chTime := make(chan *utilstime.Time)
	chUptime := make(chan admin.Details)
	chMem := make(chan *arr.Array)
	chKernel := make(chan KernelInfo)
	chPro := make(chan ProgressInfo)
	chDisk := make(chan DiskInfoDetail)

	go func() { chServer <- getServerInfo(debug) }()
	go func() { chTime <- utilstime.SystemTime() }()
	go func() { chUptime <- admin.Info() }()
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
