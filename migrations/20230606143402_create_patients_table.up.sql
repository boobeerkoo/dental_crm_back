CREATE TABLE IF NOT EXISTS public.patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone1 VARCHAR(20),
    phone2 VARCHAR(20),
    address VARCHAR(255),
    email VARCHAR(255),
    sex VARCHAR(10),
    important_info TEXT,
    comment TEXT,
    status VARCHAR(20) DEFAULT 'Новий',
    date_of_birth DATE,
    created_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    deleted_date TIMESTAMP NULL
);
