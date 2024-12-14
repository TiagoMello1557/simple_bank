-- Table: account
CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    owner VARCHAR NOT NULL,
    balance BIGINT NOT NULL,
    currency VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Create index on owner
CREATE INDEX idx_account_owner ON accounts (owner);

-- Table: entries
CREATE TABLE entries (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Create index on account_id
CREATE INDEX idx_entries_account_id ON entries (account_id);

-- Table: transfers
CREATE TABLE transfers (
    id BIGSERIAL PRIMARY KEY,
    from_account_id BIGINT NOT NULL REFERENCES accounts(id),
    to_account_id BIGINT NOT NULL REFERENCES accounts(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Create indexes on from_account_id, to_account_id, and their combination
CREATE INDEX idx_transfers_from_account_id ON transfers (from_account_id);
CREATE INDEX idx_transfers_to_account_id ON transfers (to_account_id);
CREATE INDEX idx_transfers_from_to_account_id ON transfers (from_account_id, to_account_id);