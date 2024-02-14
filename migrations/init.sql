CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS urls
(
    id   UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    path TEXT                            NOT NULL
);