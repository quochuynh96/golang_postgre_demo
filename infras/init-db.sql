CREATE DATABASE demo;

\c demo

---------------------------------------------
-- Create and init mock data for hub table --
---------------------------------------------
DROP TABLE IF EXISTS "hub";

CREATE TABLE "hub" (
    id SERIAL PRIMARY KEY,
    name varchar(50) NOT NULL,
    short_name varchar(50) NOT NULL
);

INSERT INTO "hub"
    ("name", "short_name")
VALUES
    ('Viet Nam', 'VN'),
    ('Singapore', 'SIN'),
    ('Thailand', 'TL');

SELECT * FROM "hub";

----------------------------------------------
-- Create and init mock data for team table --
----------------------------------------------
DROP TABLE IF EXISTS "team";

CREATE TABLE "team" (
    id SERIAL PRIMARY KEY,
    name varchar(50) NOT NULL UNIQUE ,
    hub_id INT,
    CONSTRAINT fk_team_hub
      FOREIGN KEY(hub_id)
	  REFERENCES hub(id)
);

INSERT INTO "team"
    ("name", "hub_id")
VALUES
    ('team1', 1),
    ('team2', 1),
    ('team3', 2);

SELECT * FROM "team";

----------------------------------------------
-- Create and init mock data for user table --
----------------------------------------------
DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username varchar(50) NOT NULL UNIQUE ,
    password varchar(255) NOT NULL,
    team_id INT,
    CONSTRAINT fk_user_team
      FOREIGN KEY(team_id)
	  REFERENCES team(id)
);

INSERT INTO "user"
    ("username", "password", "team_id")
VALUES
    ('user1', 's5cr5tPassword', 1),
    ('user2', 's5cr5tPassword', 1),
    ('user3', 's5cr5tPassword', 1),
    ('user4', 's5cr5tPassword', 3);

SELECT * FROM "user";

---------------------------------------------
-- Query mock data --
---------------------------------------------
SELECT u.username, t.name, h.name
FROM hub h JOIN team t on h.id = t.hub_id JOIN "user" u on t.id = u.team_id
ORDER BY u.username

-- RESULT
--| user1 | team1 | Viet Nam |
--| user2 | team1 | Viet Nam |
--| user3 | team1 | Viet Nam |
--| user4 | team3 | Singapore|


