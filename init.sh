DB_NAME="testdb2"

psql -U postgres -d ${DB_NAME} << EOF
DROP TABLE IF EXISTS orders;
EOF

psql -U postgres -d ${DB_NAME} << EOF

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    item_id INTEGER,
    price NUMERIC(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

EOF

echo "Tables created successfully."