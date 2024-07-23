CREATE EXTENSION IF NOT EXISTS pgcrypto;

\connect neffable_db

CREATE TABLE IF NOT EXISTS soul(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    display_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS soulConnection(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_soul UUID REFERENCES soul(id) ON DELETE CASCADE,
    second_soul UUID REFERENCES soul(id) ON DELETE CASCADE,
    connection_code TEXT NOT NULL,
    connection_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);