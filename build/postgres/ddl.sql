-- NEW ERA
-- Таблица Role
CREATE TABLE IF NOT EXISTS roles (
    id VARCHAR PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL
);

-- Таблица User
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    role_id varchar NOT NULL DEFAULT 3,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE ON UPDATE CASCADE
);

--ROLE
INSERT INTO public.roles
(id, created_at, updated_at, "name", description)
VALUES('3', now(), now(), 'default_role', 'inspect_all');