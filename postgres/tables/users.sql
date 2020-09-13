BEGIN TRANSACTION;

CREATE TABLE Users (
    ID serial PRIMARY Key, 
    Email VARCHAR(100) UNIQUE NOT NULL, 
    Password VARCHAR(1000) NOT NULL
);

COMMIT;