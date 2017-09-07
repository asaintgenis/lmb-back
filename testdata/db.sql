DROP TABLE IF EXISTS beerBar;
DROP TABLE IF EXISTS beer;
DROP TABLE IF EXISTS bar;
DROP TABLE IF EXISTS beerStyle;

CREATE TABLE beerStyle
(
  id SERIAL PRIMARY KEY,
  srm INTEGER,
  name VARCHAR(100) NOT NULL,
  color VARCHAR(7),
  ebc INTEGER
);

CREATE TABLE beer
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    content VARCHAR(5000),
    pic VARCHAR(5000),
    beerstyle_id INTEGER NOT NULL,
    FOREIGN KEY (beerstyle_id) REFERENCES beerStyle (id) ON DELETE CASCADE
);

CREATE TABLE bar
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(160)  NOT NULL
);

CREATE TABLE beerBar
(
  id SERIAL PRIMARY KEY,
  beer_id INTEGER NOT NULL,
  bar_id INTEGER NOT NULL,
  FOREIGN KEY (beer_id) REFERENCES beer (id) ON DELETE CASCADE,
  FOREIGN KEY (bar_id) REFERENCES bar (id) ON DELETE CASCADE
);

INSERT INTO beer (name, content, pic, beerstyle_id) VALUES ('1664', 'Biere de base, dite de type pisse','base64',1);
INSERT INTO beer (name, content, pic, beerstyle_id) VALUES ('Brooklyn', 'Bonne biere des familles', 'base64', 2);
INSERT INTO beer (name, content, pic, beerstyle_id) VALUES ('Guinness', 'la legende, la classe',3);

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

INSERT INTO beerStyle (srm, name, color, ebc) VALUES (3,'Blonde Ale','#F6F513',6);
INSERT INTO beerStyle (srm, name, color, ebc) VALUES (6, 'India Pale Ale', '#D5BC26', 12);
INSERT INTO beerStyle (srm, name, color, ebc) VALUES (29, 'Stout', '#0F0B0A', 57);
