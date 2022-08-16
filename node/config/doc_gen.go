// Code generated by github.com/filecoin-project/lotus/node/config/cfgdocgen. DO NOT EDIT.

package config

type DocField struct {
	Name    string
	Type    string
	Comment string
}

var Doc = map[string][]DocField{
	"Backup": []DocField{
		{
			Name: "DisableMetadataLog",
			Type: "bool",

			Comment: `Note that in case of metadata corruption it might be much harder to recover
your node if metadata log is disabled`,
		},
	},
	"Boost": []DocField{
		{
			Name: "ConfigVersion",
			Type: "int",

			Comment: `The version of the config file (used for migrations)`,
		},
		{
			Name: "Storage",
			Type: "StorageConfig",

			Comment: ``,
		},
		{
			Name: "SealerApiInfo",
			Type: "string",

			Comment: `The connect string for the sealing RPC API (lotus miner)`,
		},
		{
			Name: "SectorIndexApiInfo",
			Type: "string",

			Comment: `The connect string for the sector index RPC API (lotus miner)`,
		},
		{
			Name: "Dealmaking",
			Type: "DealmakingConfig",

			Comment: ``,
		},
		{
			Name: "Wallets",
			Type: "WalletsConfig",

			Comment: ``,
		},
		{
			Name: "Graphql",
			Type: "GraphqlConfig",

			Comment: ``,
		},
		{
			Name: "LotusDealmaking",
			Type: "lotus_config.DealmakingConfig",

			Comment: `Lotus configs`,
		},
		{
			Name: "LotusFees",
			Type: "FeeConfig",

			Comment: ``,
		},
		{
			Name: "DAGStore",
			Type: "lotus_config.DAGStoreConfig",

			Comment: ``,
		},
		{
			Name: "IndexProvider",
			Type: "lotus_config.IndexProviderConfig",

			Comment: ``,
		},
	},
	"Common": []DocField{
		{
			Name: "API",
			Type: "lotus_config.API",

			Comment: ``,
		},
		{
			Name: "Backup",
			Type: "lotus_config.Backup",

			Comment: ``,
		},
		{
			Name: "Libp2p",
			Type: "lotus_config.Libp2p",

			Comment: ``,
		},
		{
			Name: "Pubsub",
			Type: "lotus_config.Pubsub",

			Comment: ``,
		},
	},
	"DealmakingConfig": []DocField{
		{
			Name: "ConsiderOnlineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept online deals`,
		},
		{
			Name: "ConsiderOfflineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline deals`,
		},
		{
			Name: "ConsiderOnlineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept retrieval deals`,
		},
		{
			Name: "ConsiderOfflineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline retrieval deals`,
		},
		{
			Name: "ConsiderVerifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept verified deals`,
		},
		{
			Name: "ConsiderUnverifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept unverified deals`,
		},
		{
			Name: "PieceCidBlocklist",
			Type: "[]cid.Cid",

			Comment: `A list of Data CIDs to reject when making deals`,
		},
		{
			Name: "ExpectedSealDuration",
			Type: "Duration",

			Comment: `Maximum expected amount of time getting the deal into a sealed sector will take
This includes the time the deal will need to get transferred and published
before being assigned to a sector`,
		},
		{
			Name: "MaxDealStartDelay",
			Type: "Duration",

			Comment: `Maximum amount of time proposed deal StartEpoch can be in future`,
		},
		{
			Name: "MaxProviderCollateralMultiplier",
			Type: "uint64",

			Comment: `The maximum collateral that the provider will put up against a deal,
as a multiplier of the minimum collateral bound`,
		},
		{
			Name: "MaxStagingDealsBytes",
			Type: "int64",

			Comment: `The maximum allowed disk usage size in bytes of staging deals not yet
passed to the sealing node by the markets service. 0 is unlimited.`,
		},
		{
			Name: "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for storage deals`,
		},
		{
			Name: "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for retrieval deals`,
		},
		{
			Name: "StartEpochSealingBuffer",
			Type: "uint64",

			Comment: `Minimum start epoch buffer to give time for sealing of sector with deal.`,
		},
		{
			Name: "DealProposalLogDuration",
			Type: "Duration",

			Comment: `The amount of time to keep deal proposal logs for before cleaning them up.`,
		},
		{
			Name: "Filter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of storage deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalFilter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of retrieval deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalPricing",
			Type: "*lotus_config.RetrievalPricing",

			Comment: ``,
		},
		{
			Name: "MaxTransferDuration",
			Type: "Duration",

			Comment: `The maximum amount of time a transfer can take before it fails`,
		},
	},
	"FeeConfig": []DocField{
		{
			Name: "MaxPublishDealsFee",
			Type: "types.FIL",

			Comment: `The maximum fee to pay when sending the PublishStorageDeals message`,
		},
		{
			Name: "MaxMarketBalanceAddFee",
			Type: "types.FIL",

			Comment: `The maximum fee to pay when sending the AddBalance message (used by legacy markets)`,
		},
	},
	"GraphqlConfig": []DocField{
		{
			Name: "Port",
			Type: "uint64",

			Comment: `The port that the graphql server listens on`,
		},
	},
	"LotusDealmakingConfig": []DocField{
		{
			Name: "PieceCidBlocklist",
			Type: "[]cid.Cid",

			Comment: `A list of Data CIDs to reject when making deals`,
		},
		{
			Name: "ExpectedSealDuration",
			Type: "Duration",

			Comment: `Maximum expected amount of time getting the deal into a sealed sector will take
This includes the time the deal will need to get transferred and published
before being assigned to a sector`,
		},
		{
			Name: "MaxDealStartDelay",
			Type: "Duration",

			Comment: `Maximum amount of time proposed deal StartEpoch can be in future`,
		},
		{
			Name: "PublishMsgPeriod",
			Type: "Duration",

			Comment: `When a deal is ready to publish, the amount of time to wait for more
deals to be ready to publish before publishing them all as a batch`,
		},
		{
			Name: "MaxDealsPerPublishMsg",
			Type: "uint64",

			Comment: `The maximum number of deals to include in a single PublishStorageDeals
message`,
		},
		{
			Name: "MaxProviderCollateralMultiplier",
			Type: "uint64",

			Comment: `The maximum collateral that the provider will put up against a deal,
as a multiplier of the minimum collateral bound`,
		},
		{
			Name: "MaxStagingDealsBytes",
			Type: "int64",

			Comment: `The maximum allowed disk usage size in bytes of staging deals not yet
passed to the sealing node by the markets service. 0 is unlimited.`,
		},
		{
			Name: "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for storage deals`,
		},
		{
			Name: "SimultaneousTransfersForStoragePerClient",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers from any single client
for storage deals.
Unset by default (0), and values higher than SimultaneousTransfersForStorage
will have no effect; i.e. the total number of simultaneous data transfers
across all storage clients is bound by SimultaneousTransfersForStorage
regardless of this number.`,
		},
		{
			Name: "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for retrieval deals`,
		},
		{
			Name: "StartEpochSealingBuffer",
			Type: "uint64",

			Comment: `Minimum start epoch buffer to give time for sealing of sector with deal.`,
		},
		{
			Name: "Filter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of storage deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalFilter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of retrieval deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalPricing",
			Type: "*lotus_config.RetrievalPricing",

			Comment: ``,
		},
	},
	"StorageConfig": []DocField{
		{
			Name: "ParallelFetchLimit",
			Type: "int",

			Comment: `The maximum number of concurrent fetch operations to the storage subsystem`,
		},
	},
	"WalletsConfig": []DocField{
		{
			Name: "Miner",
			Type: "string",

			Comment: `The "owner" address of the miner`,
		},
		{
			Name: "PublishStorageDeals",
			Type: "string",

			Comment: `The wallet used to send PublishStorageDeals messages.
Must be a control or worker address of the miner.`,
		},
		{
			Name: "DealCollateral",
			Type: "string",

			Comment: `The wallet used as the source for storage deal collateral`,
		},
		{
			Name: "PledgeCollateral",
			Type: "string",

			Comment: `Deprecated: Renamed to DealCollateral`,
		},
	},
	"lotus_config.API": []DocField{
		{
			Name: "ListenAddress",
			Type: "string",

			Comment: `Binding address for the Lotus API`,
		},
		{
			Name: "RemoteListenAddress",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "Timeout",
			Type: "Duration",

			Comment: ``,
		},
	},
	"lotus_config.Backup": []DocField{
		{
			Name: "DisableMetadataLog",
			Type: "bool",

			Comment: `Note that in case of metadata corruption it might be much harder to recover
your node if metadata log is disabled`,
		},
	},
	"lotus_config.BatchFeeConfig": []DocField{
		{
			Name: "Base",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "PerSector",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"lotus_config.Chainstore": []DocField{
		{
			Name: "EnableSplitstore",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "Splitstore",
			Type: "Splitstore",

			Comment: ``,
		},
	},
	"lotus_config.Client": []DocField{
		{
			Name: "UseIpfs",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "IpfsOnlineMode",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "IpfsMAddr",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "IpfsUseForRetrieval",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers for storage deals`,
		},
		{
			Name: "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers for retrieval deals`,
		},
		{
			Name: "OffChainRetrieval",
			Type: "bool",

			Comment: `Require that retrievals perform no on-chain operations. Paid retrievals
without existing payment channels with available funds will fail instead
of automatically performing on-chain operations.`,
		},
	},
	"lotus_config.Common": []DocField{
		{
			Name: "API",
			Type: "API",

			Comment: ``,
		},
		{
			Name: "Backup",
			Type: "Backup",

			Comment: ``,
		},
		{
			Name: "Logging",
			Type: "Logging",

			Comment: ``,
		},
		{
			Name: "Libp2p",
			Type: "Libp2p",

			Comment: ``,
		},
		{
			Name: "Pubsub",
			Type: "Pubsub",

			Comment: ``,
		},
	},
	"lotus_config.DAGStoreConfig": []DocField{
		{
			Name: "RootDir",
			Type: "string",

			Comment: `Path to the dagstore root directory. This directory contains three
subdirectories, which can be symlinked to alternative locations if
need be:
- ./transients: caches unsealed deals that have been fetched from the
storage subsystem for serving retrievals.
- ./indices: stores shard indices.
- ./datastore: holds the KV store tracking the state of every shard
known to the DAG store.
Default value: <LOTUS_MARKETS_PATH>/dagstore (split deployment) or
<LOTUS_MINER_PATH>/dagstore (monolith deployment)`,
		},
		{
			Name: "MaxConcurrentIndex",
			Type: "int",

			Comment: `The maximum amount of indexing jobs that can run simultaneously.
0 means unlimited.
Default value: 5.`,
		},
		{
			Name: "MaxConcurrentReadyFetches",
			Type: "int",

			Comment: `The maximum amount of unsealed deals that can be fetched simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Name: "MaxConcurrentUnseals",
			Type: "int",

			Comment: `The maximum amount of unseals that can be processed simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Name: "MaxConcurrencyStorageCalls",
			Type: "int",

			Comment: `The maximum number of simultaneous inflight API calls to the storage
subsystem.
Default value: 100.`,
		},
		{
			Name: "GCInterval",
			Type: "Duration",

			Comment: `The time between calls to periodic dagstore GC, in time.Duration string
representation, e.g. 1m, 5m, 1h.
Default value: 1 minute.`,
		},
	},
	"lotus_config.DealmakingConfig": []DocField{
		{
			Name: "ConsiderOnlineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept online deals`,
		},
		{
			Name: "ConsiderOfflineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline deals`,
		},
		{
			Name: "ConsiderOnlineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept retrieval deals`,
		},
		{
			Name: "ConsiderOfflineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline retrieval deals`,
		},
		{
			Name: "ConsiderVerifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept verified deals`,
		},
		{
			Name: "ConsiderUnverifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept unverified deals`,
		},
		{
			Name: "PieceCidBlocklist",
			Type: "[]cid.Cid",

			Comment: `A list of Data CIDs to reject when making deals`,
		},
		{
			Name: "ExpectedSealDuration",
			Type: "Duration",

			Comment: `Maximum expected amount of time getting the deal into a sealed sector will take
This includes the time the deal will need to get transferred and published
before being assigned to a sector`,
		},
		{
			Name: "MaxDealStartDelay",
			Type: "Duration",

			Comment: `Maximum amount of time proposed deal StartEpoch can be in future`,
		},
		{
			Name: "PublishMsgPeriod",
			Type: "Duration",

			Comment: `When a deal is ready to publish, the amount of time to wait for more
deals to be ready to publish before publishing them all as a batch`,
		},
		{
			Name: "MaxDealsPerPublishMsg",
			Type: "uint64",

			Comment: `The maximum number of deals to include in a single PublishStorageDeals
message`,
		},
		{
			Name: "MaxProviderCollateralMultiplier",
			Type: "uint64",

			Comment: `The maximum collateral that the provider will put up against a deal,
as a multiplier of the minimum collateral bound`,
		},
		{
			Name: "MaxStagingDealsBytes",
			Type: "int64",

			Comment: `The maximum allowed disk usage size in bytes of staging deals not yet
passed to the sealing node by the markets service. 0 is unlimited.`,
		},
		{
			Name: "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for storage deals`,
		},
		{
			Name: "SimultaneousTransfersForStoragePerClient",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers from any single client
for storage deals.
Unset by default (0), and values higher than SimultaneousTransfersForStorage
will have no effect; i.e. the total number of simultaneous data transfers
across all storage clients is bound by SimultaneousTransfersForStorage
regardless of this number.`,
		},
		{
			Name: "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for retrieval deals`,
		},
		{
			Name: "StartEpochSealingBuffer",
			Type: "uint64",

			Comment: `Minimum start epoch buffer to give time for sealing of sector with deal.`,
		},
		{
			Name: "Filter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of storage deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalFilter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of retrieval deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalPricing",
			Type: "*RetrievalPricing",

			Comment: ``,
		},
	},
	"lotus_config.FeeConfig": []DocField{
		{
			Name: "DefaultMaxFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"lotus_config.FullNode": []DocField{
		{
			Name: "Client",
			Type: "Client",

			Comment: ``,
		},
		{
			Name: "Wallet",
			Type: "Wallet",

			Comment: ``,
		},
		{
			Name: "Fees",
			Type: "FeeConfig",

			Comment: ``,
		},
		{
			Name: "Chainstore",
			Type: "Chainstore",

			Comment: ``,
		},
	},
	"lotus_config.IndexProviderConfig": []DocField{
		{
			Name: "Enable",
			Type: "bool",

			Comment: `Enable set whether to enable indexing announcement to the network and expose endpoints that
allow indexer nodes to process announcements. Enabled by default.`,
		},
		{
			Name: "EntriesCacheCapacity",
			Type: "int",

			Comment: `EntriesCacheCapacity sets the maximum capacity to use for caching the indexing advertisement
entries. Defaults to 1024 if not specified. The cache is evicted using LRU policy. The
maximum storage used by the cache is a factor of EntriesCacheCapacity, EntriesChunkSize and
the length of multihashes being advertised. For example, advertising 128-bit long multihashes
with the default EntriesCacheCapacity, and EntriesChunkSize means the cache size can grow to
256MiB when full.`,
		},
		{
			Name: "EntriesChunkSize",
			Type: "int",

			Comment: `EntriesChunkSize sets the maximum number of multihashes to include in a single entries chunk.
Defaults to 16384 if not specified. Note that chunks are chained together for indexing
advertisements that include more multihashes than the configured EntriesChunkSize.`,
		},
		{
			Name: "TopicName",
			Type: "string",

			Comment: `TopicName sets the topic name on which the changes to the advertised content are announced.
If not explicitly specified, the topic name is automatically inferred from the network name
in following format: '/indexer/ingest/<network-name>'
Defaults to empty, which implies the topic name is inferred from network name.`,
		},
		{
			Name: "PurgeCacheOnStart",
			Type: "bool",

			Comment: `PurgeCacheOnStart sets whether to clear any cached entries chunks when the provider engine
starts. By default, the cache is rehydrated from previously cached entries stored in
datastore if any is present.`,
		},
	},
	"lotus_config.Libp2p": []DocField{
		{
			Name: "ListenAddresses",
			Type: "[]string",

			Comment: `Binding address for the libp2p host - 0 means random port.
Format: multiaddress; see https://multiformats.io/multiaddr/`,
		},
		{
			Name: "AnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to explicitally announce to other peers. If not specified,
all interface addresses are announced
Format: multiaddress`,
		},
		{
			Name: "NoAnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to not announce
Format: multiaddress`,
		},
		{
			Name: "BootstrapPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "ProtectedPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DisableNatPortMap",
			Type: "bool",

			Comment: `When not disabled (default), lotus asks NAT devices (e.g., routers), to
open up an external port and forward it to the port lotus is running on.
When this works (i.e., when your router supports NAT port forwarding),
it makes the local lotus node accessible from the public internet`,
		},
		{
			Name: "ConnMgrLow",
			Type: "uint",

			Comment: `ConnMgrLow is the number of connections that the basic connection manager
will trim down to.`,
		},
		{
			Name: "ConnMgrHigh",
			Type: "uint",

			Comment: `ConnMgrHigh is the number of connections that, when exceeded, will trigger
a connection GC operation. Note: protected/recently formed connections don't
count towards this limit.`,
		},
		{
			Name: "ConnMgrGrace",
			Type: "Duration",

			Comment: `ConnMgrGrace is a time duration that new connections are immune from being
closed by the connection manager.`,
		},
	},
	"lotus_config.Logging": []DocField{
		{
			Name: "SubsystemLevels",
			Type: "map[string]string",

			Comment: `SubsystemLevels specify per-subsystem log levels`,
		},
	},
	"lotus_config.MinerAddressConfig": []DocField{
		{
			Name: "PreCommitControl",
			Type: "[]string",

			Comment: `Addresses to send PreCommit messages from`,
		},
		{
			Name: "CommitControl",
			Type: "[]string",

			Comment: `Addresses to send Commit messages from`,
		},
		{
			Name: "TerminateControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DealPublishControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DisableOwnerFallback",
			Type: "bool",

			Comment: `DisableOwnerFallback disables usage of the owner address for messages
sent automatically`,
		},
		{
			Name: "DisableWorkerFallback",
			Type: "bool",

			Comment: `DisableWorkerFallback disables usage of the worker address for messages
sent automatically, if control addresses are configured.
A control address that doesn't have enough funds will still be chosen
over the worker address if this flag is set.`,
		},
	},
	"lotus_config.MinerFeeConfig": []DocField{
		{
			Name: "MaxPreCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxPreCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: `maxBatchFee = maxBase + maxPerSector * nSectors`,
		},
		{
			Name: "MaxCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: ``,
		},
		{
			Name: "MaxTerminateGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxWindowPoStGasFee",
			Type: "types.FIL",

			Comment: `WindowPoSt is a high-value operation, so the default fee should be high.`,
		},
		{
			Name: "MaxPublishDealsFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxMarketBalanceAddFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"lotus_config.MinerSubsystemConfig": []DocField{
		{
			Name: "EnableMining",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableSealing",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableSectorStorage",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableMarkets",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "SealerApiInfo",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "SectorIndexApiInfo",
			Type: "string",

			Comment: ``,
		},
	},
	"lotus_config.ProvingConfig": []DocField{
		{
			Name: "ParallelCheckLimit",
			Type: "int",

			Comment: `After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Name: "DisableBuiltinWindowPoSt",
			Type: "bool",

			Comment: `After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Name: "DisableBuiltinWinningPoSt",
			Type: "bool",

			Comment: `WARNING: If no WinningPoSt workers are connected, Winning PoSt WILL FAIL resulting in lost block rewards.
Before enabling this option, make sure your PoSt workers work correctly.`,
		},
		{
			Name: "DisableWDPoStPreChecks",
			Type: "bool",

			Comment: `After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Name: "MaxPartitionsPerPoStMessage",
			Type: "int",

			Comment: `Setting this value above the network limit has no effect`,
		},
		{
			Name: "MaxPartitionsPerRecoveryMessage",
			Type: "int",

			Comment: `In some cases when submitting DeclareFaultsRecovered messages,
there may be too many recoveries to fit in a BlockGasLimit.
In those cases it may be necessary to set this value to something low (eg 1);
Note that setting this value lower may result in less efficient gas use - more messages will be sent than needed,
resulting in more total gas use (but each message will have lower gas limit)`,
		},
	},
	"lotus_config.Pubsub": []DocField{
		{
			Name: "Bootstrapper",
			Type: "bool",

			Comment: `Run the node in bootstrap-node mode`,
		},
		{
			Name: "DirectPeers",
			Type: "[]string",

			Comment: `DirectPeers specifies peers with direct peering agreements. These peers are
connected outside of the mesh, with all (valid) message unconditionally
forwarded to them. The router will maintain open connections to these peers.
Note that the peering agreement should be reciprocal with direct peers
symmetrically configured at both ends.
Type: Array of multiaddress peerinfo strings, must include peerid (/p2p/12D3K...`,
		},
		{
			Name: "IPColocationWhitelist",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "RemoteTracer",
			Type: "string",

			Comment: ``,
		},
	},
	"lotus_config.RetrievalPricing": []DocField{
		{
			Name: "Strategy",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "Default",
			Type: "*RetrievalPricingDefault",

			Comment: ``,
		},
		{
			Name: "External",
			Type: "*RetrievalPricingExternal",

			Comment: ``,
		},
	},
	"lotus_config.RetrievalPricingDefault": []DocField{
		{
			Name: "VerifiedDealsFreeTransfer",
			Type: "bool",

			Comment: `VerifiedDealsFreeTransfer configures zero fees for data transfer for a retrieval deal
of a payloadCid that belongs to a verified storage deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "default".
default value is true`,
		},
	},
	"lotus_config.RetrievalPricingExternal": []DocField{
		{
			Name: "Path",
			Type: "string",

			Comment: `Path of the external script that will be run to price a retrieval deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "external".`,
		},
	},
	"lotus_config.SealerConfig": []DocField{
		{
			Name: "ParallelFetchLimit",
			Type: "int",

			Comment: ``,
		},
		{
			Name: "AllowAddPiece",
			Type: "bool",

			Comment: `Local worker config`,
		},
		{
			Name: "AllowPreCommit1",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowPreCommit2",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowCommit",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowUnseal",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowReplicaUpdate",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowProveReplicaUpdate2",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowRegenSectorKey",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "Assigner",
			Type: "string",

			Comment: `Assigner specifies the worker assigner to use when scheduling tasks.
"utilization" (default) - assign tasks to workers with lowest utilization.
"spread" - assign tasks to as many distinct workers as possible.`,
		},
		{
			Name: "DisallowRemoteFinalize",
			Type: "bool",

			Comment: `DisallowRemoteFinalize when set to true will force all Finalize tasks to
run on workers with local access to both long-term storage and the sealing
path containing the sector.
--
WARNING: Only set this if all workers have access to long-term storage
paths. If this flag is enabled, and there are workers without long-term
storage access, sectors will not be moved from them, and Finalize tasks
will appear to be stuck.
--
If you see stuck Finalize tasks after enabling this setting, check
'lotus-miner sealing sched-diag' and 'lotus-miner storage find [sector num]'`,
		},
		{
			Name: "ResourceFiltering",
			Type: "sealer.ResourceFilteringStrategy",

			Comment: `ResourceFiltering instructs the system which resource filtering strategy
to use when evaluating tasks against this worker. An empty value defaults
to "hardware".`,
		},
	},
	"lotus_config.SealingConfig": []DocField{
		{
			Name: "MaxWaitDealsSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be waiting for more deals to be packed in it before it begins sealing at any given time.
If the miner is accepting multiple deals in parallel, up to MaxWaitDealsSectors of new sectors will be created.
If more than MaxWaitDealsSectors deals are accepted in parallel, only MaxWaitDealsSectors deals will be processed in parallel
Note that setting this number too high in relation to deal ingestion rate may result in poor sector packing efficiency
0 = no limit`,
		},
		{
			Name: "MaxSealingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when creating new CC sectors (0 = unlimited)`,
		},
		{
			Name: "MaxSealingSectorsForDeals",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when creating new sectors with deals (0 = unlimited)`,
		},
		{
			Name: "PreferNewSectorsForDeals",
			Type: "bool",

			Comment: `Prefer creating new sectors even if there are sectors Available for upgrading.
This setting combined with MaxUpgradingSectors set to a value higher than MaxSealingSectorsForDeals makes it
possible to use fast sector upgrades to handle high volumes of storage deals, while still using the simple sealing
flow when the volume of storage deals is lower.`,
		},
		{
			Name: "MaxUpgradingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when upgrading CC sectors with deals (0 = MaxSealingSectorsForDeals)`,
		},
		{
			Name: "CommittedCapacitySectorLifetime",
			Type: "Duration",

			Comment: `CommittedCapacitySectorLifetime is the duration a Committed Capacity (CC) sector will
live before it must be extended or converted into sector containing deals before it is
terminated. Value must be between 180-540 days inclusive`,
		},
		{
			Name: "WaitDealsDelay",
			Type: "Duration",

			Comment: `Period of time that a newly created sector will wait for more deals to be packed in to before it starts to seal.
Sectors which are fully filled will start sealing immediately`,
		},
		{
			Name: "AlwaysKeepUnsealedCopy",
			Type: "bool",

			Comment: `Whether to keep unsealed copies of deal data regardless of whether the client requested that. This lets the miner
avoid the relatively high cost of unsealing the data later, at the cost of more storage space`,
		},
		{
			Name: "FinalizeEarly",
			Type: "bool",

			Comment: `Run sector finalization before submitting sector proof to the chain`,
		},
		{
			Name: "MakeNewSectorForDeals",
			Type: "bool",

			Comment: `Whether new sectors are created to pack incoming deals
When this is set to false no new sectors will be created for sealing incoming deals
This is useful for forcing all deals to be assigned as snap deals to sectors marked for upgrade`,
		},
		{
			Name: "MakeCCSectorsAvailable",
			Type: "bool",

			Comment: `After sealing CC sectors, make them available for upgrading with deals`,
		},
		{
			Name: "CollateralFromMinerBalance",
			Type: "bool",

			Comment: `Whether to use available miner balance for sector collateral instead of sending it with each message`,
		},
		{
			Name: "AvailableBalanceBuffer",
			Type: "types.FIL",

			Comment: `Minimum available balance to keep in the miner actor before sending it with messages`,
		},
		{
			Name: "DisableCollateralFallback",
			Type: "bool",

			Comment: `Don't send collateral with messages even if there is no available balance in the miner actor`,
		},
		{
			Name: "BatchPreCommits",
			Type: "bool",

			Comment: `enable / disable precommit batching (takes effect after nv13)`,
		},
		{
			Name: "MaxPreCommitBatch",
			Type: "int",

			Comment: `maximum precommit batch size - batches will be sent immediately above this size`,
		},
		{
			Name: "PreCommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Name: "PreCommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deal in batch would start expiring`,
		},
		{
			Name: "AggregateCommits",
			Type: "bool",

			Comment: `enable / disable commit aggregation (takes effect after nv13)`,
		},
		{
			Name: "MinCommitBatch",
			Type: "int",

			Comment: `minimum batched commit size - batches above this size will eventually be sent on a timeout`,
		},
		{
			Name: "MaxCommitBatch",
			Type: "int",

			Comment: `maximum batched commit size - batches will be sent immediately above this size`,
		},
		{
			Name: "CommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Name: "CommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deals in batch would start expiring`,
		},
		{
			Name: "BatchPreCommitAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing precommit batching, instead
sending precommit messages to the chain individually`,
		},
		{
			Name: "AggregateAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing commit aggregation, instead
submitting proofs to the chain individually`,
		},
		{
			Name: "TerminateBatchMax",
			Type: "uint64",

			Comment: ``,
		},
		{
			Name: "TerminateBatchMin",
			Type: "uint64",

			Comment: ``,
		},
		{
			Name: "TerminateBatchWait",
			Type: "Duration",

			Comment: ``,
		},
	},
	"lotus_config.Splitstore": []DocField{
		{
			Name: "ColdStoreType",
			Type: "string",

			Comment: `ColdStoreType specifies the type of the coldstore.
It can be "universal" (default) or "discard" for discarding cold blocks.`,
		},
		{
			Name: "HotStoreType",
			Type: "string",

			Comment: `HotStoreType specifies the type of the hotstore.
Only currently supported value is "badger".`,
		},
		{
			Name: "MarkSetType",
			Type: "string",

			Comment: `MarkSetType specifies the type of the markset.
It can be "map" for in memory marking or "badger" (default) for on-disk marking.`,
		},
		{
			Name: "HotStoreMessageRetention",
			Type: "uint64",

			Comment: `HotStoreMessageRetention specifies the retention policy for messages, in finalities beyond
the compaction boundary; default is 0.`,
		},
		{
			Name: "HotStoreFullGCFrequency",
			Type: "uint64",

			Comment: `HotStoreFullGCFrequency specifies how often to perform a full (moving) GC on the hotstore.
A value of 0 disables, while a value 1 will do full GC in every compaction.
Default is 20 (about once a week).`,
		},
	},
	"lotus_config.StorageMiner": []DocField{
		{
			Name: "Subsystems",
			Type: "MinerSubsystemConfig",

			Comment: ``,
		},
		{
			Name: "Dealmaking",
			Type: "DealmakingConfig",

			Comment: ``,
		},
		{
			Name: "IndexProvider",
			Type: "IndexProviderConfig",

			Comment: ``,
		},
		{
			Name: "Proving",
			Type: "ProvingConfig",

			Comment: ``,
		},
		{
			Name: "Sealing",
			Type: "SealingConfig",

			Comment: ``,
		},
		{
			Name: "Storage",
			Type: "SealerConfig",

			Comment: ``,
		},
		{
			Name: "Fees",
			Type: "MinerFeeConfig",

			Comment: ``,
		},
		{
			Name: "Addresses",
			Type: "MinerAddressConfig",

			Comment: ``,
		},
		{
			Name: "DAGStore",
			Type: "DAGStoreConfig",

			Comment: ``,
		},
	},
	"lotus_config.Wallet": []DocField{
		{
			Name: "RemoteBackend",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "EnableLedger",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "DisableLocal",
			Type: "bool",

			Comment: ``,
		},
	},
}
