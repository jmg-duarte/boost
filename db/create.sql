CREATE TABLE IF NOT EXISTS DealLogs (
    DealUUID TEXT,
    CreatedAt DateTime,
    LogTEXT TEXT
);

CREATE TABLE IF NOT EXISTS Deals (
    ID TEXT,
    CreatedAt DateTime,
    DealProposalSignature BLOB,
    PieceCID TEXT,
    PieceSize INT,
    VerifiedDeal BOOL,
    ClientAddress TEXT,
    ProviderAddress TEXT,
    Label TEXT,
    StartEpoch INT,
    EndEpoch INT,
    StoragePricePerEpoch BLOB,
    ProviderCollateral BLOB,
    ClientCollateral BLOB,
    SelfPeerID TEXT,
    ClientPeerID TEXT,
    DealDataRoot TEXT,
    InboundFilePath TEXT,
    TransferType TEXT,
    TransferParams BLOB,
    TransferSize INT,
    ChainDealID INT,
    PublishCID TEXT,
    SectorID INT,
    Offset INT,
    Length INT,
    Checkpoint TEXT,
    Error TEXT,
    PRIMARY KEY(ID)
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS FundsLogs (
    DealUUID TEXT,
    CreatedAt DateTime,
    Amount BLOB,
    LogText TEXT
);

CREATE TABLE IF NOT EXISTS FundsTagged (
    DealUUID TEXT,
    CreatedAt DateTime,
    Collateral INT,
    PubMsg INT
);