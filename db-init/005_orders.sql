CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    status TEXT NOT NULL,
    total_amount NUMERIC NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
