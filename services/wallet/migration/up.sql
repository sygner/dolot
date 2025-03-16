DROP TABLE "transactions";
DROP TABLE "wallets";
DROP TABLE "coins";

-- Creating a table for coins with columns: id, currency_name, currency_id
CREATE TABLE "coins" (
    "id" SERIAL PRIMARY KEY,
    "currency_name" VARCHAR(32) NOT NULL,
    "currency_symbol" VARCHAR(32) NOT NULL UNIQUE
);

-- Insert the default coin 'LUNA'
INSERT INTO "coins" ("currency_name", "currency_symbol") VALUES ('LUNA', 'luna');

-- Creating a table for wallets with columns: id, sid, coin_id, balance, address, public_key, private_key, created_at, updated_at
CREATE TABLE "wallets" (
    "id" SERIAL PRIMARY KEY,
    "sid" VARCHAR(64) NOT NULL,
    "user_id" INT NOT NULL,
    "coin_id" INT NOT NULL,
    "balance" DOUBLE PRECISION NOT NULL,
    "address" TEXT NOT NULL UNIQUE,
    "public_key" TEXT NOT NULL UNIQUE,
    "private_key" TEXT NOT NULL,
    "mnemonic" TEXT,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_coin FOREIGN KEY ("coin_id") REFERENCES "coins" ("id") ON DELETE CASCADE
);

-- Adding an index on the "sid" column for faster lookup in the wallets table
CREATE INDEX "wallets_sid_index" ON "wallets" ("sid");

CREATE TABLE "transactions" (
    "tx_id" TEXT NOT NULL PRIMARY KEY,
    "currency_id" TEXT NOT NULL,
    "currency_name" TEXT NOT NULL,
    "from_address" TEXT NOT NULL,
    "to_address" TEXT NOT NULL,
    "from_wallet_id" INT NOT NULL,
    "from_public_key" TEXT NOT NULL,
    "amount" DOUBLE PRECISION NOT NULL,
    "transaction_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_wallet FOREIGN KEY ("from_wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE -- Optional, if "from_wallet_id" references the "wallets" table
);

-- Optional: Add indexes on columns that will be queried frequently
CREATE INDEX "transactions_from_wallet_id_index" ON "transactions" ("from_wallet_id");
