CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE TABLE
    IF NOT EXISTS public.users (
        id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
        nickname varchar(32) UNIQUE NOT NULL,
        "name" varchar(100) NOT NULL,
        birthdate date NOT NULL,
        stack text NULL,
        search text NOT NULL
    );

-- CREATE INDEX
--     CONCURRENTLY IF NOT EXISTS idx_user_trigram ON public.users USING gist (
--         search gist_trgm_ops(siglen = 64)
--     );