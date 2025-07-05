CREATE TABLE IF NOT EXISTS public.appointments (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER REFERENCES public.patients (id),
    doctor_id INTEGER REFERENCES public.doctors (id),
    appointment_date TIMESTAMP NOT NULL,
    duration INTERVAL NOT NULL,
    status VARCHAR(255) NOT NULL,
    comment VARCHAR(255),
    created_date TIMESTAMP DEFAULT NOW(),
    updated_date TIMESTAMP DEFAULT NOW(),
    deleted_date TIMESTAMP NULL
    );
