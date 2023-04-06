package types

import (
	"fmt"
	"time"

	"github.com/filecoin-project/boost/storagemarket/types/dealcheckpoints"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
)

// ProviderDealState is the local state tracked for a deal by the StorageProvider.
type ProviderDealState struct {
	// DealUuid is an unique uuid generated by client for the deal.
	DealUuid uuid.UUID
	// CreatedAt is the time at which the deal was stored
	CreatedAt time.Time
	// ClientDealProposal is the deal proposal sent by the client.
	ClientDealProposal market.ClientDealProposal
	// IsOffline is true for offline deals i.e. deals where the actual data to be stored by the SP is sent out of band
	// and not via an online data transfer.
	IsOffline bool
	// CleanupData indicates whether to remove the data for a deal after the deal has been added to a sector.
	// This is always true for online deals, and can be set as a flag for offline deals.
	CleanupData bool

	// ClientPeerID is the Clients libp2p Peer ID.
	ClientPeerID peer.ID

	// DealDataRoot is the root of the IPLD DAG that the client wants to store.
	DealDataRoot cid.Cid

	// InboundCARPath is the file-path where the storage provider will persist the CAR file sent by the client.
	InboundFilePath string

	// Transfer has the parameters for the data transfer
	Transfer Transfer

	// Chain Vars
	ChainDealID abi.DealID
	PublishCID  *cid.Cid

	// sector packing info
	SectorID abi.SectorNumber
	Offset   abi.PaddedPieceSize
	Length   abi.PaddedPieceSize

	// deal checkpoint in DB.
	Checkpoint dealcheckpoints.Checkpoint
	// CheckpointAt is the time at which the deal entered in the last state
	CheckpointAt time.Time

	// set if there's an error
	Err string
	// if there was an error, indicates whether and how to retry (auto / manual)
	Retry DealRetryType

	// NBytesReceived is the number of bytes Received for this deal
	NBytesReceived int64

	// Keep unsealed copy of the data
	FastRetrieval bool

	//Announce deal to the IPNI(Index Provider)
	AnnounceToIPNI bool
}

func (d *ProviderDealState) String() string {
	return fmt.Sprintf("%+v", *d)
}

func (d *ProviderDealState) SignedProposalCid() (cid.Cid, error) {
	propnd, err := cborutil.AsIpld(&d.ClientDealProposal)
	if err != nil {
		return cid.Undef, fmt.Errorf("failed to compute signed deal proposal ipld node: %w", err)
	}

	return propnd.Cid(), nil
}

type DealRetryType string

const (
	// DealRetryAuto means that when boost restarts, it will automatically
	// retry the deal
	DealRetryAuto DealRetryType = "auto"
	// DealRetryManual means that boost will not automatically retry the
	// deal, it must be manually retried by the user
	DealRetryManual DealRetryType = "manual"
	// DealRetryFatal means that the deal will fail immediately and permanently
	DealRetryFatal DealRetryType = "fatal"
)
