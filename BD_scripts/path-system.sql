CREATE DATABASE path_system;

-- DROP TABLE IF EXISTS public.section CASCADE;
CREATE TABLE public.section (
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT BY 1 MINVALUE 0 MAXVALUE 32767 START WITH 1 CACHE 1 ),
	id_stage integer,
	title varchar(250) NOT NULL,
	description text NOT NULL,
	CONSTRAINT section_pk PRIMARY KEY (id)
);

-- DROP TABLE IF EXISTS public.stage CASCADE;
CREATE TABLE public.stage (
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT BY 1 MINVALUE 0 MAXVALUE 32767 START WITH 1 CACHE 1 ),
	title varchar(250),
	CONSTRAINT stage_pk PRIMARY KEY (id)
);

-- DROP TABLE IF EXISTS public.category CASCADE;
CREATE TABLE public.category (
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT BY 1 MINVALUE 0 MAXVALUE 32767 START WITH 1 CACHE 1 ),
	name varchar(250),
	CONSTRAINT category_pk PRIMARY KEY (id)
);

-- ALTER TABLE public.section DROP CONSTRAINT IF EXISTS stage_fk CASCADE;
ALTER TABLE public.section ADD CONSTRAINT stage_fk FOREIGN KEY (id_stage)
	REFERENCES public.stage (id) MATCH FULL	ON DELETE CASCADE ON UPDATE CASCADE;

-- DROP TABLE IF EXISTS public.many_category_has_many_section CASCADE;
CREATE TABLE public.many_category_has_many_section (
	id_category integer NOT NULL,
	id_section integer NOT NULL,
	CONSTRAINT many_category_has_many_section_pk PRIMARY KEY (id_category,id_section)
);

-- ALTER TABLE public.many_category_has_many_section DROP CONSTRAINT IF EXISTS category_fk CASCADE;
ALTER TABLE public.many_category_has_many_section ADD CONSTRAINT category_fk FOREIGN KEY (id_category)
	REFERENCES public.category (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;

-- ALTER TABLE public.many_category_has_many_section DROP CONSTRAINT IF EXISTS section_fk CASCADE;
ALTER TABLE public.many_category_has_many_section ADD CONSTRAINT section_fk FOREIGN KEY (id_section)
	REFERENCES public.section (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;