CREATE TABLE IF NOT EXISTS t_currency (
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT        NOT NULL DEFAULT '',
    char_code TEXT        NOT NULL DEFAULT '',
    value     DECIMAL     NOT NULL DEFAULT 0.0,
    date      TIMESTAMP   NOT NULL
);
