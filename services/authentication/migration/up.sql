DROP TABLE IF EXISTS "tokens";
DROP TABLE IF EXISTS "passwords";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "login_history";

-- Users table
CREATE TABLE "users" (
    "user_id" SERIAL NOT NULL PRIMARY KEY,
    "phone_number" TEXT,
    "email" TEXT,
    "account_username" TEXT,
    "user_role" TEXT NOT NULL,
    "user_status" TEXT NOT NULL,
    "provider" TEXT DEFAULT 'local',      -- 'local' for normal login, 'google' or other for SSO
    "is_sso" BOOLEAN DEFAULT false,       -- Indicates if the user signed up via SSO
    "created_at" TIMESTAMPTZ NOT NULL
);

-- Tokens table
CREATE TABLE "tokens" (
    "access_token" TEXT NOT NULL PRIMARY KEY,
    "refresh_token" TEXT NOT NULL,
    "user_id" INT NOT NULL,
    "user_role" TEXT NOT NULL,
    "session_id" SMALLINT NOT NULL,
    "token_status" TEXT NOT NULL,
    "ip" TEXT NOT NULL,
    "agent" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "access_token_expire_at" TIMESTAMPTZ NOT NULL,
    "refresh_token_expire_at" TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

-- Passwords table (only for non-SSO users)
CREATE TABLE "passwords" (
    "user_id" SERIAL NOT NULL PRIMARY KEY REFERENCES users (user_id),
    "password" TEXT NOT NULL
);

-- Login history table
CREATE TABLE "login_history"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL
);
