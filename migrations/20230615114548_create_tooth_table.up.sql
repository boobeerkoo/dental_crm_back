CREATE TABLE IF NOT EXISTS public.tooth (
    id SERIAL PRIMARY KEY,
    dental_formula_id INTEGER REFERENCES public.dental_formulas (id),
    tooth_number INTEGER NOT NULL,
    tooth_name TEXT,
    type TEXT,
    damage TEXT,
    parodont TEXT,
    endo TEXT,
    constructions TEXT,
    created_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    deleted_date TIMESTAMP NULL
);
