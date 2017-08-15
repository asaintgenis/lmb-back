DROP TABLE IF EXISTS beer;
DROP TABLE IF EXISTS bar;

CREATE TABLE beer
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120)
);

CREATE TABLE bar
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(160)  NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artist (id) ON DELETE CASCADE
);

CREATE TABLE beerBar
(
  id SERIAL PRIMARY KEY,
  beer_id INTEGER NOT NULL,
  bar_id INTEGER NOT NULL,
  FOREIGN KEY (beer_id) REFERENCES beer (id) ON DELETE CASCADE
  FOREIGN KEY (bar_id) REFERENCES bar (id) ON DELETE CASCADE

)

INSERT INTO beer (name) VALUES ('1664');
INSERT INTO beer (name) VALUES ('Kronenbourg');
INSERT INTO beer (name) VALUES ('Guinness');

INSERT INTO bar (title) VALUES ('La Kolok');
INSERT INTO bar (title) VALUES ('Guinness tavern');
INSERT INTO bar (title) VALUES ('King georges');

INSERT INTO beerBar (beer_id, bar_id) VALUES (1, 1);
INSERT INTO beerBar (beer_id, bar_id) VALUES (1, 2);
INSERT INTO beerBar (beer_id, bar_id) VALUES (2, 1);
INSERT INTO beerBar (beer_id, bar_id) VALUES (2, 2);
INSERT INTO beerBar (beer_id, bar_id) VALUES (2, 3);
INSERT INTO beerBar (beer_id, bar_id) VALUES (3, 1);
INSERT INTO beerBar (beer_id, bar_id) VALUES (3, 3);

