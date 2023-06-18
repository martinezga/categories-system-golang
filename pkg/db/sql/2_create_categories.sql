CREATE TABLE categories.categories (
    id varchar(16) NOT NULL UNIQUE PRIMARY KEY,
    name varchar(64) NOT NULL UNIQUE,
    description varchar(256)
);

-- DROP TABLE categories;
