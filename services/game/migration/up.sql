-- Drop the games table if it already exists to avoid conflicts
DROP TABLE IF EXISTS "user_choices";

DROP TABLE IF EXISTS "games";

-- Create the games table to store details about different lottery games
CREATE TABLE "games"(
    "id" VARCHAR(64) PRIMARY KEY NOT NULL, -- Unique identifier for each game
    "name" VARCHAR(64) NOT NULL, -- Name of the game (e.g., Lotto, Oslotto, Powerball, American Powerball)
    "game_type" VARCHAR(32) NOT NULL, -- Type of game (e.g., lotto, powerball)
    "num_main_numbers" INTEGER NOT NULL, -- Number of main numbers to pick in the game
    "num_bonus_numbers" INTEGER DEFAULT NULL, -- Number of bonus numbers (optional, for games like Powerball)
    "main_number_range" INTEGER NOT NULL, -- Range of the main numbers (e.g., 1-49 for Lotto)
    "bonus_number_range" INTEGER DEFAULT NULL, -- Range of bonus numbers if applicable (e.g., 1-26 for Powerball)
    "start_time" TIMESTAMP NOT NULL, -- The start time for the game's drawing window
    "end_time" TIMESTAMP NOT NULL, -- The end time for the game's drawing window
    "creator_id" INTEGER NOT NULL, -- ID of the user/admin who created the game
    "result" VARCHAR(255), -- Storing the result as a string for simplicity (e.g., "5,12,23,34,45 + 2" for Powerball)
    "created_at" TIMESTAMP NOT NULL -- Timestamp when the game was created
);

-- Drop the user_choices table if it already exists

CREATE TABLE "user_choices"(
    "id" VARCHAR(64) PRIMARY KEY NOT NULL, -- Unique identifier for each choice record
    "user_id" INTEGER NOT NULL, -- Reference to the user who made the choice
    "game_id" VARCHAR(64) NOT NULL, -- Reference to the game being played
    "chosen_main_numbers" INTEGER[][] NOT NULL, -- 2D array for multiple sets of main numbers
    "chosen_bonus_numbers" INTEGER[][], -- 2D array for multiple sets of bonus numbers (if applicable)
    "created_at" TIMESTAMP NOT NULL, -- Time when the user made the selection
    FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON DELETE CASCADE -- Ensures valid game reference
);

CREATE INDEX idx_user_choices_game_id ON user_choices(game_id);


CREATE TABLE winners (
    "id" SERIAL PRIMARY KEY,          -- Unique ID for each record
    "game_id" VARCHAR(64) NOT NULL,          -- The ID of the related game
    "divisions" JSONB NOT NULL,       -- JSON array to store user_id, match_count, and bonus
    "result_number" VARCHAR(255) NOT NULL, -- The result number associated with the win
    "created_at" TIMESTAMP DEFAULT NOW()  -- Timestamp when the record was created
);

-- Index to optimize lookups based on game_id (optional)
CREATE INDEX idx_winners_game_id ON winners(game_id);


-- INSERT INTO "games" (
--     "id", 
--     "name", 
--     "game_type", 
--     "num_main_numbers", 
--     "num_bonus_numbers", 
--     "main_number_range", 
--     "bonus_number_range", 
--     "start_time", 
--     "end_time", 
--     "creator_id", 
--     "result", 
--     "created_at"
-- ) VALUES (
--     'game_powerball_001',   -- Unique ID for the game
--     'Powerball',            -- Name of the game
--     'powerball',            -- Type of the game
--     5,                      -- Players pick 5 main numbers
--     1,                      -- Players pick 1 bonus number
--     69,                     -- Main numbers are picked from 1 to 69
--     26,                     -- Bonus numbers are picked from 1 to 26
--     '2024-09-21 08:00',     -- Start time of the game
--     '2024-09-21 20:00',     -- End time of the game
--     1,                      -- Creator ID (admin)
--     NULL,                   -- Result (set after the game is drawn)
--     '2024-09-19 10:00'      -- Game created on this date
-- );


-- Example 1: Lotto
-- sql
-- Copy code
-- INSERT INTO "games" (
--     "id", "name", "game_type", "num_main_numbers", 
--     "num_bonus_numbers", "main_number_range", "bonus_number_range", 
--     "start_time", "end_time", "creator_id", "result", "created_at"
-- ) VALUES (
--     '1', 'Lotto', 'lotto', 
--     6, -- Number of main numbers
--     0, -- No bonus numbers for Lotto
--     49, -- Main numbers range from 1 to 49
--     NULL, -- No bonus number range
--     '2024-09-21 10:00:00', '2024-09-21 23:59:00', 
--     123, -- creator_id
--     NULL, -- result will be set later
--     '2024-09-20 12:00:00' -- Created timestamp
-- );
-- Example 2: Ozlotto
-- sql
-- Copy code
-- INSERT INTO "games" (
--     "id", "name", "game_type", "num_main_numbers", 
--     "num_bonus_numbers", "main_number_range", "bonus_number_range", 
--     "start_time", "end_time", "creator_id", "result", "created_at"
-- ) VALUES (
--     '2', 'Ozlotto', 'ozlotto', 
--     7, -- Number of main numbers
--     0, -- No bonus numbers for Ozlotto
--     45, -- Main numbers range from 1 to 45
--     NULL, -- No bonus number range
--     '2024-09-22 10:00:00', '2024-09-22 23:59:00', 
--     124, -- creator_id
--     NULL, -- result will be set later
--     '2024-09-20 12:00:00' -- Created timestamp
-- );
-- Example 3: Powerball
-- sql
-- Copy code
-- INSERT INTO "games" (
--     "id", "name", "game_type", "num_main_numbers", 
--     "num_bonus_numbers", "main_number_range", "bonus_number_range", 
--     "start_time", "end_time", "creator_id", "result", "created_at"
-- ) VALUES (
--     '3', 'Powerball', 'powerball', 
--     5, -- Number of main numbers
--     1, -- One bonus number (Powerball)
--     69, -- Main numbers range from 1 to 69
--     26, -- Bonus number range from 1 to 26
--     '2024-09-23 10:00:00', '2024-09-23 23:59:00', 
--     125, -- creator_id
--     NULL, -- result will be set later
--     '2024-09-20 12:00:00' -- Created timestamp
-- );
-- Example 4: American Powerball
-- sql
-- Copy code
-- INSERT INTO "games" (
--     "id", "name", "game_type", "num_main_numbers", 
--     "num_bonus_numbers", "main_number_range", "bonus_number_range", 
--     "start_time", "end_time", "creator_id", "result", "created_at"
-- ) VALUES (
--     '4', 'American Powerball', 'american_powerball', 
--     5, -- Number of main numbers
--     1, -- One bonus number (Powerball)
--     69, -- Main numbers range from 1 to 69
--     26, -- Bonus number range from 1 to 26
--     '2024-09-24 10:00:00', '2024-09-24 23:59:00', 
--     126, -- creator_id
--     NULL, -- result will be set later
--     '2024-09-20 12:00:00' -- Created timestamp
-- );

-- INSERT INTO games(id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, created_at)
-- VALUES 
-- ('1', 'Lotto', 'lotto', 6, 0, 49, NULL, '2024-09-25 10:00:00', '2024-09-25 11:00:00', 1, '5, 12, 23, 34, 45, 48', NOW()),
-- ('2', 'Ozlotto', 'ozlotto', 7, 0, 45, NULL, '2024-09-25 12:00:00', '2024-09-25 13:00:00', 2, '3, 8, 19, 27, 32, 37, 44', NOW()),
-- ('3', 'Powerball', 'powerball', 5, 1, 69, 26, '2024-09-26 14:00:00', '2024-09-26 15:00:00', 3, '7, 13, 22, 35, 48 + 9', NOW()),
-- ('4', 'American Powerball', 'american_powerball', 5, 1, 69, 26, '2024-09-27 16:00:00', '2024-09-27 17:00:00', 4, '11, 28, 41, 50, 59 + 3', NOW());
