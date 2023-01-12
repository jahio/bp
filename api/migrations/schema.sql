--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Homebrew)
-- Dumped by pg_dump version 15.1 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: entries; Type: TABLE; Schema: public; Owner: dev
--

CREATE TABLE public.entries (
    id uuid NOT NULL,
    created_at integer,
    updated_at integer,
    systolic integer,
    diastolic integer,
    heartrate integer
);


ALTER TABLE public.entries OWNER TO dev;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: dev
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO dev;

--
-- Name: entries entries_pkey; Type: CONSTRAINT; Schema: public; Owner: dev
--

ALTER TABLE ONLY public.entries
    ADD CONSTRAINT entries_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: dev
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: entries_created_at_idx; Type: INDEX; Schema: public; Owner: dev
--

CREATE INDEX entries_created_at_idx ON public.entries USING btree (created_at);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: dev
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

