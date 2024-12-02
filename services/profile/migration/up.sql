DROP TABLE IF EXISTS "profiles";
CREATE TABLE "profiles"(
    "user_id" SERIAL PRIMARY KEY NOT NULL,
    "sid" VARCHAR(30) NOT NULL,
    "username" VARCHAR(64) NOT NULL,
    "score" double precision NOT NULL,
    "impression" INTEGER NOT NULL,
    "rank" INTEGER NOT NULL,
    "games_quantity" INTEGER NOT NULL,
    "won_games" INTEGER NOT NULL,
    "lost_games" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);