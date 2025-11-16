-- Index for org users lookup
CREATE INDEX IF NOT EXISTS idx_users_org_id ON users(org_id);

-- Unique constraint to prevent duplicate short codes per domain
CREATE UNIQUE INDEX IF NOT EXISTS idx_links_domain_short_code ON links(domain, short_code);

-- Index for fast lookups by user
CREATE INDEX IF NOT EXISTS idx_links_user_id ON links(user_id);

-- Index for expiry cleanup
CREATE INDEX IF NOT EXISTS idx_links_expires_at ON links(expires_at);

-- Index for redirect performance aggregation (clicks per link)
CREATE INDEX IF NOT EXISTS idx_clicks_link_id_created_at ON clicks(link_id, created_at DESC);

-- Index for analytics queries by campaign
CREATE INDEX IF NOT EXISTS idx_clicks_campaign_id_created_at ON clicks(campaign_id, created_at DESC);

-- Optional: partial index for last 30 days analytics
CREATE INDEX IF NOT EXISTS idx_clicks_recent ON clicks(created_at DESC) WHERE created_at > now() - interval '30 days';

-- Index for campaigns table
CREATE INDEX IF NOT EXISTS idx_campaigns_org_id ON campaigns(org_id);

-- Index to quickly find verified domains
CREATE INDEX IF NOT EXISTS idx_custom_domains_verified ON custom_domains(verified);

-- Index for QR codes table
CREATE INDEX IF NOT EXISTS idx_qr_link_id ON qr_codes(link_id);
