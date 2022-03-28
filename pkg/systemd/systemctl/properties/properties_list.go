package properties

import (
	systemctl2 "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/systemd/systemctl"
)

var Properties = []Property{
	ActiveEnterTimestamp,
	ActiveEnterTimestampMonotonic,
	ActiveExitTimestampMonotonic,
	systemctl2.ActiveState,
	After,
	AllowIsolate,
	AssertResult,
	AssertTimestamp,
	AssertTimestampMonotonic,
	Before,
	BlockIOAccounting,
	BlockIOWeight,
	CPUAccounting,
	CPUAffinityFromNUMA,
	CPUQuotaPerSecUSec,
	CPUQuotaPeriodUSec,
	CPUSchedulingPolicy,
	CPUSchedulingPriority,
	CPUSchedulingResetOnFork,
	CPUShares,
	CPUUsageNSec,
	CPUWeight,
	CacheDirectoryMode,
	CanFreeze,
	CanIsolate,
	CanReload,
	CanStart,
	CanStop,
	CapabilityBoundingSet,
	CleanResult,
	CollectMode,
	ConditionResult,
	ConditionTimestamp,
	ConditionTimestampMonotonic,
	ConfigurationDirectoryMode,
	Conflicts,
	ControlGroup,
	ControlPID,
	CoredumpFilter,
	DefaultDependencies,
	DefaultMemoryLow,
	DefaultMemoryMin,
	Delegate,
	Description,
	DevicePolicy,
	DynamicUser,
	EffectiveCPUs,
	EffectiveMemoryNodes,
	ExecMainCode,
	ExecMainExitTimestampMonotonic,
	ExecMainPID,
	ExecMainStartTimestamp,
	ExecMainStartTimestampMonotonic,
	ExecMainStatus,
	ExecReload,
	ExecReloadEx,
	ExecStart,
	ExecStartEx,
	FailureAction,
	FileDescriptorStoreMax,
	FinalKillSignal,
	FragmentPath,
	FreezerState,
	GID,
	GuessMainPID,
	IOAccounting,
	IOReadBytes,
	IOReadOperations,
	IOSchedulingClass,
	IOSchedulingPriority,
	IOWeight,
	IOWriteBytes,
	IOWriteOperations,
	IPAccounting,
	IPEgressBytes,
	IPEgressPackets,
	IPIngressBytes,
	IPIngressPackets,
	Id,
	IgnoreOnIsolate,
	IgnoreSIGPIPE,
	InactiveEnterTimestampMonotonic,
	InactiveExitTimestamp,
	InactiveExitTimestampMonotonic,
	InvocationID,
	JobRunningTimeoutUSec,
	JobTimeoutAction,
	JobTimeoutUSec,
	KeyringMode,
	KillMode,
	KillSignal,
	LimitAS,
	LimitASSoft,
	LimitCORE,
	LimitCORESoft,
	LimitCPU,
	LimitCPUSoft,
	LimitDATA,
	LimitDATASoft,
	LimitFSIZE,
	LimitFSIZESoft,
	LimitLOCKS,
	LimitLOCKSSoft,
	LimitMEMLOCK,
	LimitMEMLOCKSoft,
	LimitMSGQUEUE,
	LimitMSGQUEUESoft,
	LimitNICE,
	LimitNICESoft,
	LimitNOFILE,
	LimitNOFILESoft,
	LimitNPROC,
	LimitNPROCSoft,
	LimitRSS,
	LimitRSSSoft,
	LimitRTPRIO,
	LimitRTPRIOSoft,
	LimitRTTIME,
	LimitRTTIMESoft,
	LimitSIGPENDING,
	LimitSIGPENDINGSoft,
	LimitSTACK,
	LimitSTACKSoft,
	LoadState,
	LockPersonality,
	LogLevelMax,
	LogRateLimitBurst,
	LogRateLimitIntervalUSec,
	LogsDirectoryMode,
	MainPID,
	ManagedOOMMemoryPressure,
	ManagedOOMMemoryPressureLimit,
	ManagedOOMPreference,
	ManagedOOMSwap,
	MemoryAccounting,
	MemoryCurrent,
	MemoryDenyWriteExecute,
	MemoryHigh,
	MemoryLimit,
	MemoryLow,
	MemoryMax,
	MemoryMin,
	MemorySwapMax,
	MountAPIVFS,
	NFileDescriptorStore,
	NRestarts,
	NUMAPolicy,
	Names,
	NeedDaemonReload,
	Nice,
	NoNewPrivileges,
	NonBlocking,
	NotifyAccess,
	OOMPolicy,
	OOMScoreAdjust,
	OnFailureJobMode,
	PIDFile,
	Perpetual,
	PrivateDevices,
	PrivateIPC,
	PrivateMounts,
	PrivateNetwork,
	PrivateTmp,
	PrivateUsers,
	ProcSubset,
	ProtectClock,
	ProtectControlGroups,
	ProtectHome,
	ProtectHostname,
	ProtectKernelLogs,
	ProtectKernelModules,
	ProtectKernelTunables,
	ProtectProc,
	ProtectSystem,
	RefuseManualStart,
	RefuseManualStop,
	ReloadResult,
	RemainAfterExit,
	RemoveIPC,
	Requires,
	systemctl2.Restart,
	RestartKillSignal,
	RestartUSec,
	RestrictNamespaces,
	RestrictRealtime,
	RestrictSUIDSGID,
	Result,
	RootDirectoryStartOnly,
	RuntimeDirectoryMode,
	RuntimeDirectoryPreserve,
	RuntimeMaxUSec,
	SameProcessGroup,
	SecureBits,
	SendSIGHUP,
	SendSIGKILL,
	Slice,
	StandardError,
	StandardInput,
	StandardOutput,
	StartLimitAction,
	StartLimitBurst,
	StartLimitIntervalUSec,
	StartupBlockIOWeight,
	StartupCPUShares,
	StartupCPUWeight,
	StartupIOWeight,
	StateChangeTimestamp,
	StateChangeTimestampMonotonic,
	StateDirectoryMode,
	StatusErrno,
	StopWhenUnneeded,
	systemctl2.SubState,
	SuccessAction,
	SyslogFacility,
	SyslogLevel,
	SyslogLevelPrefix,
	SyslogPriority,
	SystemCallErrorNumber,
	TTYReset,
	TTYVHangup,
	TTYVTDisallocate,
	TasksAccounting,
	TasksCurrent,
	TasksMax,
	TimeoutAbortUSec,
	TimeoutCleanUSec,
	TimeoutStartFailureMode,
	TimeoutStartUSec,
	TimeoutStopFailureMode,
	TimeoutStopUSec,
	TimerSlackNSec,
	Transient,
	Type,
	UID,
	UMask,
	UnitFilePreset,
	systemctl2.UnitFileState,
	UtmpMode,
	WantedBy,
	WatchdogSignal,
	WatchdogTimestampMonotonic,
	WatchdogUSec,
}
