DROP TABLE IF EXISTS "profiles";
DROP TABLE IF EXISTS "rank_logs";
CREATE TABLE "profiles"(
    "user_id" SERIAL PRIMARY KEY NOT NULL,
    "sid" VARCHAR(30) NOT NULL,
    "username" VARCHAR(64) NOT NULL,
    "score" double precision NOT NULL,
    "impression" INTEGER NOT NULL,
    "d_coin" INTEGER NOT NULL,
    "rank" INTEGER NOT NULL,
    "games_quantity" INTEGER NOT NULL,
    "won_games" INTEGER NOT NULL,
    "lost_games" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE "rank_logs"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" SERIAL NOT NULL REFERENCES "profiles"("user_id"),
    "rank" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);