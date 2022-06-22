CREATE TABLE IF NOT EXISTS activities (
                                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                                          name TEXT NOT NULL,
                                          dtime TEXT NOT NULL,
                                          duration TEXT NOT NULL,
                                          category TEXT NOT NULL
);

DELETE FROM activities WHERE id = 1 OR id = 2;

INSERT INTO activities (id, name, dtime, duration, category)
VALUES(1, 'BJJ', '2022.06.22 19:29', '1h30m0s', 'Jiu Jitsu');

INSERT INTO activities (id, name, dtime, duration, category)
VALUES(2, 'Run', '2022.06.22 20:16', '30m', 'Run');

SELECT * FROM activities;
