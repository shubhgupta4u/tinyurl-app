CREATE TABLE IF NOT EXISTS organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID REFERENCES organizations(id),
    email TEXT UNIQUE NOT NULL,
    mobile TEXT UNIQUE,           
    password_hash TEXT NOT NULL,
    name TEXT NOT NULL
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS links (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    org_id UUID REFERENCES organizations(id),
    short_code TEXT NOT NULL,
    target_url TEXT NOT NULL,
    domain TEXT NOT NULL DEFAULT 'tiny.example.com',
    custom_alias BOOLEAN DEFAULT FALSE,
    title TEXT,
    tags TEXT[],
    campaign_id UUID NULL,
    expires_at TIMESTAMPTZ,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS clicks (
    id BIGSERIAL PRIMARY KEY,
    link_id BIGINT REFERENCES links(id),
    short_code TEXT NOT NULL,
    domain TEXT NOT NULL,
    ip INET,
    user_agent TEXT,
    referrer TEXT,
    country TEXT,
    device_type TEXT,
    campaign_id UUID NULL,
    utm JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS campaigns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    org_id UUID REFERENCES organizations(id),
    name TEXT NOT NULL,
    description TEXT,
    utm_defaults JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS custom_domains (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID REFERENCES organizations(id),
    domain TEXT UNIQUE NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    verification_token TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS qr_codes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    link_id BIGINT REFERENCES links(id),
    format TEXT DEFAULT 'png',
    size INT DEFAULT 256,
    path TEXT, -- path to S3 / storage
    created_at TIMESTAMPTZ DEFAULT now()
);
