CREATE TABLE IF NOT EXISTS public.dental_formulas (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER REFERENCES public.patients (id),
    doctor_id INTEGER REFERENCES public.doctors (id),
    created_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    deleted_date TIMESTAMP NULL
);
