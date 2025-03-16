DROP TABLE IF EXISTS tickets;

CREATE TABLE tickets(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "signature" VARCHAR(255) NOT NULL,
    "user_id" INTEGER NOT NULL,
    "ticket_type" VARCHAR(255) NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "used" BOOLEAN NOT NULL DEFAULT FALSE,
    "used_at" TIMESTAMP,
    "game_id" TEXT,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE INDEX tickets_user_id ON tickets(user_id);
CREATE INDEX tickets_signature ON tickets(signature);