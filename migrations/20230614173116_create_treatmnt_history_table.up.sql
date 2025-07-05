CREATE TABLE IF NOT EXISTS public.treatments (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER REFERENCES public.patients (id),
    doctor_id INTEGER REFERENCES public.doctors (id),
    type VARCHAR(255) NOT NULL,
    comment TEXT,
    /*dental_formula_id INTEGER REFERENCES public.dental_formulas (id)*/
    created_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    deleted_date TIMESTAMP NULL
);
