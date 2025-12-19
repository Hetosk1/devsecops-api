CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    product_id UUID REFERENCES products(id),
    rating INT NOT NULL,
    comment TEXT,
    created_at TIMESTAMP DEFAULT now()
);
