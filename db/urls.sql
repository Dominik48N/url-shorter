CREATE TABLE urls (
    link VARCHAR(12) UNIQUE,
    redirect_url TEXT,
    owner VARCHAR(16)
);
