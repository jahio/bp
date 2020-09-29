--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

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
-- Name: entries; Type: TABLE; Schema: public; Owner: jah
--

CREATE TABLE public.entries (
    id uuid NOT NULL,
    created_at integer,
    updated_at integer,
    systolic integer,
    diastolic integer,
    heartrate integer
);


ALTER TABLE public.entries OWNER TO jah;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: jah
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO jah;

--
-- Name: entries entries_pkey; Type: CONSTRAINT; Schema: public; Owner: jah
--

ALTER TABLE ONLY public.entries
    ADD CONSTRAINT entries_pkey PRIMARY KEY (id);


--
-- Name: entries_created_at_idx; Type: INDEX; Schema: public; Owner: jah
--

CREATE INDEX entries_created_at_idx ON public.entries USING btree (created_at);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: jah
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

