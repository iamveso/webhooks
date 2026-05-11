CREATE TABLE IF NOT EXISTS usernames(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR(128) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS webhook_config(
    id CHAR(26) PRIMARY KEY,
    username VARCHAR(128) UNIQUE REFERENCES usernames(username),
    private_key TEXT NOT NULL,
    public_key TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS webhook_subcriptions (
    id CHAR(26) PRIMARY KEY,
    config_id CHAR(26) NOT NULL REFERENCES webhook_config(id),
    username VARCHAR(128) NOT NULL,
    event VARCHAR(128) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- table for event data
-- CREATE TABLE IF NOT EXISTS outbound_webhooks();

-- CREATE TABLE IF NOT EXISTS inbound_webhooks();
-- -- table for delivery data
-- CREATE TABLE IF NOT EXISTS outbound_delivery();

-- CREATE TABLE IF NOT EXISTS inbound_delivery();
