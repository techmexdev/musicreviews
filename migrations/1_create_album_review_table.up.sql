CREATE TABLE artist(
  id UUID NOT NULL UNIQUE,
  name VARCHAR NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE album_review(
  id UUID NOT NULL UNIQUE,
  name VARCHAR NOT NULL UNIQUE,
  artist_id UUID references artist(id),
  rating INTEGER,
  PRIMARY KEY (id)
);
