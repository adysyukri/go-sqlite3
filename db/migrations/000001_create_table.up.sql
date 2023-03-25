CREATE TABLE IF NOT EXISTS accounts (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	balance REAL NOT NULL
);

INSERT INTO accounts (name, balance) VALUES ('abu', 1000);
INSERT INTO accounts (name, balance) VALUES ('boba', 100);
