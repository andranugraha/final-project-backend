--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Ubuntu 14.6-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)

-- Started on 2023-02-05 10:38:17 WIB

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

--
-- TOC entry 5 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 243 (class 1259 OID 17833)
-- Name: carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.carts (
    id integer NOT NULL,
    user_id integer NOT NULL,
    course_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.carts OWNER TO postgres;

--
-- TOC entry 242 (class 1259 OID 17832)
-- Name: carts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.carts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.carts_id_seq OWNER TO postgres;

--
-- TOC entry 3630 (class 0 OID 0)
-- Dependencies: 242
-- Name: carts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.carts_id_seq OWNED BY public.carts.id;


--
-- TOC entry 223 (class 1259 OID 17600)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 17599)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO postgres;

--
-- TOC entry 3631 (class 0 OID 0)
-- Dependencies: 222
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- TOC entry 227 (class 1259 OID 17626)
-- Name: course_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course_tags (
    id integer NOT NULL,
    course_id integer NOT NULL,
    tag_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.course_tags OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 17625)
-- Name: course_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.course_tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_tags_id_seq OWNER TO postgres;

--
-- TOC entry 3632 (class 0 OID 0)
-- Dependencies: 226
-- Name: course_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.course_tags_id_seq OWNED BY public.course_tags.id;


--
-- TOC entry 217 (class 1259 OID 17564)
-- Name: courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.courses (
    id integer NOT NULL,
    category_id integer NOT NULL,
    title character varying NOT NULL,
    slug character varying NOT NULL,
    summary text NOT NULL,
    content text NOT NULL,
    img_url character varying NOT NULL,
    author_name character varying NOT NULL,
    status character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    img_thumbnail character varying NOT NULL,
    price double precision NOT NULL
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 17563)
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.courses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.courses_id_seq OWNER TO postgres;

--
-- TOC entry 3633 (class 0 OID 0)
-- Dependencies: 216
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- TOC entry 221 (class 1259 OID 17590)
-- Name: favorites; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.favorites (
    id integer NOT NULL,
    user_id integer NOT NULL,
    course_id integer NOT NULL,
    date timestamp without time zone DEFAULT now() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.favorites OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 17589)
-- Name: favorites_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.favorites_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.favorites_id_seq OWNER TO postgres;

--
-- TOC entry 3634 (class 0 OID 0)
-- Dependencies: 220
-- Name: favorites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.favorites_id_seq OWNED BY public.favorites.id;


--
-- TOC entry 237 (class 1259 OID 17680)
-- Name: gifts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gifts (
    id integer NOT NULL,
    name character varying NOT NULL,
    stock integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.gifts OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 17679)
-- Name: gifts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gifts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gifts_id_seq OWNER TO postgres;

--
-- TOC entry 3635 (class 0 OID 0)
-- Dependencies: 236
-- Name: gifts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gifts_id_seq OWNED BY public.gifts.id;


--
-- TOC entry 229 (class 1259 OID 17635)
-- Name: invoices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invoices (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id integer NOT NULL,
    status character varying NOT NULL,
    total double precision NOT NULL,
    payment_date timestamp without time zone,
    voucher_id integer,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    discount double precision NOT NULL,
    subtotal double precision NOT NULL,
    user_voucher_id integer
);


ALTER TABLE public.invoices OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 17634)
-- Name: invoices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.invoices_id_seq OWNER TO postgres;

--
-- TOC entry 3636 (class 0 OID 0)
-- Dependencies: 228
-- Name: invoices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.invoices_id_seq OWNED BY public.invoices.id;


--
-- TOC entry 215 (class 1259 OID 17551)
-- Name: levels; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.levels (
    id integer NOT NULL,
    name character varying NOT NULL,
    discount double precision NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    min_transaction integer NOT NULL,
    avatar_url character varying,
    point integer
);


ALTER TABLE public.levels OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 17550)
-- Name: levels_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.levels_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.levels_id_seq OWNER TO postgres;

--
-- TOC entry 3637 (class 0 OID 0)
-- Dependencies: 214
-- Name: levels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.levels_id_seq OWNED BY public.levels.id;


--
-- TOC entry 245 (class 1259 OID 17929)
-- Name: redeemables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.redeemables (
    id integer NOT NULL,
    user_id integer NOT NULL,
    point integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.redeemables OWNER TO postgres;

--
-- TOC entry 244 (class 1259 OID 17928)
-- Name: redeemables_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.redeemables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.redeemables_id_seq OWNER TO postgres;

--
-- TOC entry 3638 (class 0 OID 0)
-- Dependencies: 244
-- Name: redeemables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.redeemables_id_seq OWNED BY public.redeemables.id;


--
-- TOC entry 213 (class 1259 OID 17538)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 17537)
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO postgres;

--
-- TOC entry 3639 (class 0 OID 0)
-- Dependencies: 212
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- TOC entry 225 (class 1259 OID 17613)
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 17612)
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO postgres;

--
-- TOC entry 3640 (class 0 OID 0)
-- Dependencies: 224
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- TOC entry 239 (class 1259 OID 17693)
-- Name: tracks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tracks (
    id integer NOT NULL,
    user_id integer NOT NULL,
    status character varying NOT NULL,
    depart_date timestamp without time zone DEFAULT now() NOT NULL,
    arrive_date timestamp without time zone,
    name character varying NOT NULL,
    description text NOT NULL,
    total integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.tracks OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 17692)
-- Name: tracks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tracks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tracks_id_seq OWNER TO postgres;

--
-- TOC entry 3641 (class 0 OID 0)
-- Dependencies: 238
-- Name: tracks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tracks_id_seq OWNED BY public.tracks.id;


--
-- TOC entry 231 (class 1259 OID 17647)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    invoice_id uuid NOT NULL,
    course_id integer NOT NULL,
    price double precision NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 17646)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 3642 (class 0 OID 0)
-- Dependencies: 230
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 219 (class 1259 OID 17579)
-- Name: user_courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_courses (
    id integer NOT NULL,
    user_id integer NOT NULL,
    course_id integer NOT NULL,
    status character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_courses OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 17578)
-- Name: user_courses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_courses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_courses_id_seq OWNER TO postgres;

--
-- TOC entry 3643 (class 0 OID 0)
-- Dependencies: 218
-- Name: user_courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_courses_id_seq OWNED BY public.user_courses.id;


--
-- TOC entry 241 (class 1259 OID 17705)
-- Name: user_gifts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_gifts (
    id integer NOT NULL,
    user_id integer NOT NULL,
    gift_id integer NOT NULL,
    track_id integer NOT NULL,
    subtotal integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_gifts OWNER TO postgres;

--
-- TOC entry 240 (class 1259 OID 17704)
-- Name: user_gifts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_gifts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_gifts_id_seq OWNER TO postgres;

--
-- TOC entry 3644 (class 0 OID 0)
-- Dependencies: 240
-- Name: user_gifts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_gifts_id_seq OWNED BY public.user_gifts.id;


--
-- TOC entry 235 (class 1259 OID 17669)
-- Name: user_vouchers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_vouchers (
    id integer NOT NULL,
    user_id integer NOT NULL,
    voucher_id integer NOT NULL,
    expiry_date timestamp without time zone NOT NULL,
    is_consumed boolean NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.user_vouchers OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 17668)
-- Name: user_vouchers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_vouchers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_vouchers_id_seq OWNER TO postgres;

--
-- TOC entry 3645 (class 0 OID 0)
-- Dependencies: 234
-- Name: user_vouchers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_vouchers_id_seq OWNED BY public.user_vouchers.id;


--
-- TOC entry 211 (class 1259 OID 17521)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying NOT NULL,
    password text NOT NULL,
    role_id integer NOT NULL,
    username character varying NOT NULL,
    fullname character varying NOT NULL,
    address text NOT NULL,
    phone_no character varying NOT NULL,
    level_id integer NOT NULL,
    referral character varying NOT NULL,
    ref_referral character varying,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 17520)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3646 (class 0 OID 0)
-- Dependencies: 210
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 233 (class 1259 OID 17656)
-- Name: vouchers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vouchers (
    id integer NOT NULL,
    name character varying NOT NULL,
    code character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    amount double precision NOT NULL,
    min_amount double precision
);


ALTER TABLE public.vouchers OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 17655)
-- Name: vouchers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vouchers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vouchers_id_seq OWNER TO postgres;

--
-- TOC entry 3647 (class 0 OID 0)
-- Dependencies: 232
-- Name: vouchers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vouchers_id_seq OWNED BY public.vouchers.id;


--
-- TOC entry 3353 (class 2604 OID 17836)
-- Name: carts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts ALTER COLUMN id SET DEFAULT nextval('public.carts_id_seq'::regclass);


--
-- TOC entry 3322 (class 2604 OID 17603)
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- TOC entry 3328 (class 2604 OID 17629)
-- Name: course_tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_tags ALTER COLUMN id SET DEFAULT nextval('public.course_tags_id_seq'::regclass);


--
-- TOC entry 3312 (class 2604 OID 17567)
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- TOC entry 3318 (class 2604 OID 17593)
-- Name: favorites id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites ALTER COLUMN id SET DEFAULT nextval('public.favorites_id_seq'::regclass);


--
-- TOC entry 3343 (class 2604 OID 17683)
-- Name: gifts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifts ALTER COLUMN id SET DEFAULT nextval('public.gifts_id_seq'::regclass);


--
-- TOC entry 3309 (class 2604 OID 17554)
-- Name: levels id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.levels ALTER COLUMN id SET DEFAULT nextval('public.levels_id_seq'::regclass);


--
-- TOC entry 3356 (class 2604 OID 17932)
-- Name: redeemables id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.redeemables ALTER COLUMN id SET DEFAULT nextval('public.redeemables_id_seq'::regclass);


--
-- TOC entry 3306 (class 2604 OID 17541)
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- TOC entry 3325 (class 2604 OID 17616)
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- TOC entry 3346 (class 2604 OID 17696)
-- Name: tracks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks ALTER COLUMN id SET DEFAULT nextval('public.tracks_id_seq'::regclass);


--
-- TOC entry 3334 (class 2604 OID 17650)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 3315 (class 2604 OID 17582)
-- Name: user_courses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_courses ALTER COLUMN id SET DEFAULT nextval('public.user_courses_id_seq'::regclass);


--
-- TOC entry 3350 (class 2604 OID 17708)
-- Name: user_gifts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts ALTER COLUMN id SET DEFAULT nextval('public.user_gifts_id_seq'::regclass);


--
-- TOC entry 3340 (class 2604 OID 17672)
-- Name: user_vouchers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers ALTER COLUMN id SET DEFAULT nextval('public.user_vouchers_id_seq'::regclass);


--
-- TOC entry 3303 (class 2604 OID 17524)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3337 (class 2604 OID 17659)
-- Name: vouchers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers ALTER COLUMN id SET DEFAULT nextval('public.vouchers_id_seq'::regclass);


--
-- TOC entry 3620 (class 0 OID 17833)
-- Dependencies: 243
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3600 (class 0 OID 17600)
-- Dependencies: 223
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories VALUES (1, 'Back End Development', '2023-01-19 12:48:09.047643', '2023-01-19 12:48:09.047643', NULL);
INSERT INTO public.categories VALUES (2, 'Front End Development', '2023-01-27 12:02:10.780753', '2023-01-27 12:02:10.780753', NULL);


--
-- TOC entry 3604 (class 0 OID 17626)
-- Dependencies: 227
-- Data for Name: course_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.course_tags VALUES (69, 26, 23, '2023-01-20 17:54:03.996733', '2023-01-20 17:54:03.996733', NULL);
INSERT INTO public.course_tags VALUES (70, 26, 29, '2023-01-20 17:54:03.996733', '2023-01-20 17:54:03.996733', NULL);
INSERT INTO public.course_tags VALUES (71, 27, 23, '2023-01-20 18:47:44.104583', '2023-01-20 18:47:44.104583', NULL);
INSERT INTO public.course_tags VALUES (72, 27, 30, '2023-01-20 18:47:44.104583', '2023-01-20 18:47:44.104583', NULL);
INSERT INTO public.course_tags VALUES (77, 31, 27, '2023-01-27 15:35:28.889027', '2023-01-27 15:35:28.889027', NULL);
INSERT INTO public.course_tags VALUES (78, 31, 28, '2023-01-27 15:35:28.889027', '2023-01-27 15:35:28.889027', NULL);
INSERT INTO public.course_tags VALUES (111, 32, 23, '2023-01-27 18:14:11.482654', '2023-01-27 18:14:11.482654', NULL);
INSERT INTO public.course_tags VALUES (112, 32, 25, '2023-01-27 18:14:11.482654', '2023-01-27 18:14:11.482654', NULL);
INSERT INTO public.course_tags VALUES (113, 33, 23, '2023-01-28 11:19:20.729086', '2023-01-28 11:19:20.729086', NULL);
INSERT INTO public.course_tags VALUES (114, 33, 29, '2023-01-28 11:19:20.729086', '2023-01-28 11:19:20.729086', NULL);
INSERT INTO public.course_tags VALUES (119, 34, 27, '2023-02-03 09:55:26.302942', '2023-02-03 09:55:26.302942', NULL);
INSERT INTO public.course_tags VALUES (120, 34, 30, '2023-02-03 09:55:26.302942', '2023-02-03 09:55:26.302942', NULL);
INSERT INTO public.course_tags VALUES (123, 34, 23, '2023-02-03 11:18:13.614922', '2023-02-03 11:18:13.614922', NULL);
INSERT INTO public.course_tags VALUES (124, 34, 25, '2023-02-03 11:18:13.614922', '2023-02-03 11:18:13.614922', NULL);
INSERT INTO public.course_tags VALUES (125, 34, 26, '2023-02-03 11:18:13.614922', '2023-02-03 11:18:13.614922', NULL);
INSERT INTO public.course_tags VALUES (126, 34, 28, '2023-02-03 11:18:13.614922', '2023-02-03 11:18:13.614922', NULL);
INSERT INTO public.course_tags VALUES (133, 35, 23, '2023-02-03 11:55:08.436911', '2023-02-03 11:55:08.436911', NULL);
INSERT INTO public.course_tags VALUES (134, 35, 27, '2023-02-03 11:55:08.436911', '2023-02-03 11:55:08.436911', NULL);
INSERT INTO public.course_tags VALUES (135, 35, 28, '2023-02-03 11:55:08.436911', '2023-02-03 11:55:08.436911', NULL);
INSERT INTO public.course_tags VALUES (136, 35, 29, '2023-02-03 11:55:08.436911', '2023-02-03 11:55:08.436911', NULL);
INSERT INTO public.course_tags VALUES (137, 35, 30, '2023-02-03 11:55:08.436911', '2023-02-03 11:55:08.436911', NULL);
INSERT INTO public.course_tags VALUES (138, 35, 25, '2023-02-03 11:55:08.436911', '2023-02-03 11:55:08.436911', NULL);
INSERT INTO public.course_tags VALUES (139, 36, 27, '2023-02-03 13:17:53.91707', '2023-02-03 13:17:53.91707', NULL);
INSERT INTO public.course_tags VALUES (140, 36, 25, '2023-02-03 13:17:53.91707', '2023-02-03 13:17:53.91707', NULL);
INSERT INTO public.course_tags VALUES (141, 37, 23, '2023-02-03 13:21:10.506024', '2023-02-03 13:21:10.506024', NULL);
INSERT INTO public.course_tags VALUES (142, 37, 29, '2023-02-03 13:21:10.506024', '2023-02-03 13:21:10.506024', NULL);
INSERT INTO public.course_tags VALUES (143, 38, 23, '2023-02-03 13:26:14.244527', '2023-02-03 13:26:14.244527', NULL);
INSERT INTO public.course_tags VALUES (144, 38, 29, '2023-02-03 13:26:14.244527', '2023-02-03 13:26:14.244527', NULL);
INSERT INTO public.course_tags VALUES (145, 39, 23, '2023-02-03 13:27:23.262769', '2023-02-03 13:27:23.262769', NULL);
INSERT INTO public.course_tags VALUES (67, 25, 23, '2023-01-20 17:53:05.935745', '2023-01-20 17:53:05.935745', NULL);
INSERT INTO public.course_tags VALUES (68, 25, 28, '2023-01-20 17:53:05.935745', '2023-01-20 17:53:05.935745', NULL);


--
-- TOC entry 3594 (class 0 OID 17564)
-- Dependencies: 217
-- Data for Name: courses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.courses VALUES (26, 1, 'Kafka', 'kafka', 'Materi kafka', '<p>Ini content</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1674816532/courses/kafka.png', 'Ike', 'publish', '2023-01-20 17:54:03.999045', '2023-01-27 17:48:52.948946', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1674816532/courses/kafka.png', 100000);
INSERT INTO public.courses VALUES (25, 1, 'PostgreSQL', 'postgresql', 'Materi PG', 'Ini content', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1674211985/courses/postgresql', 'Rudi', 'publish', '2023-01-20 17:53:05.937265', '2023-01-20 17:53:05.937265', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1674211985/courses/postgresql', 100000);
INSERT INTO public.courses VALUES (31, 2, 'Introduction to React', 'introduction-to-react', 'React is a JavaScript library for building user interfaces', '<h1>React is a JavaScript library for building user interfaces</h1>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1674808528/courses/introduction-to-react.png', 'Merisa Syafrina', 'publish', '2023-01-27 15:35:28.891985', '2023-01-27 18:11:51.511465', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1674808528/courses/introduction-to-react.png', 90000);
INSERT INTO public.courses VALUES (32, 1, 'Be A Back End Expert', 'be-a-back-end-expert', 'Be A Back End Expert', '<p>Be A Back End Expert</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1674818050/courses/be-a-back-end-expert.png', 'Javin Djapri', 'publish', '2023-01-27 18:14:11.483671', '2023-01-27 18:14:11.483671', '2023-01-27 18:26:29.860614', 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1674818050/courses/be-a-back-end-expert.png', 999000);
INSERT INTO public.courses VALUES (33, 1, 'RESTful API', 'restful-api', 'Tutorial restful api', '<p>REST is an acronym for&nbsp;REpresentational&nbsp;State&nbsp;Transfer and an architectural style for&nbsp;distributed hypermedia systems. Roy Fielding first presented it in 2000 in his famous&nbsp;<a href="https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm" rel="noopener noreferrer" target="_blank" style="color: var(--accent);">dissertation</a>.</p><p>Like other architectural styles, REST has its guiding principles and constraints. These principles must be satisfied if a service interface needs to be referred to as&nbsp;RESTful.</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1674879560/courses/restful-api.png', 'Rezza Aldy', 'publish', '2023-01-28 11:19:20.730983', '2023-01-28 11:19:20.730983', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1674879560/courses/restful-api.png', 95000);
INSERT INTO public.courses VALUES (27, 1, 'Database Tuning', 'database-tuning', 'Materi database tuning', '<p><strong>Database tuning</strong>&nbsp;describes a group of activities used to optimize and homogenize the performance of a&nbsp;<a href="https://en.wikipedia.org/wiki/Database" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(6, 69, 173);">database</a>. It usually overlaps with&nbsp;<a href="https://en.wikipedia.org/wiki/Query_language" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(6, 69, 173);">query</a>&nbsp;tuning, but refers to design of the database files, selection of the&nbsp;<a href="https://en.wikipedia.org/wiki/Database_management_system" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(6, 69, 173);">database management system</a>&nbsp;(DBMS) application, and configuration of the database''s environment (<a href="https://en.wikipedia.org/wiki/Operating_system" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(6, 69, 173);">operating system</a>,&nbsp;<a href="https://en.wikipedia.org/wiki/CPU" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(6, 69, 173);">CPU</a>, etc.).</p><p>Database tuning aims to maximize use of system resources to perform work as efficiently and rapidly as possible. Most systems are designed to manage their use of system resources, but there is still much room to improve their efficiency by customizing their settings and configuration for the database and the DBMS.</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1674879560/courses/database-tuning', 'Ike Nurjanah', 'publish', '2023-01-20 18:47:44.1078', '2023-01-28 11:31:41.233581', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1674879560/courses/database-tuning', 1000000);
INSERT INTO public.courses VALUES (38, 1, 'Oracle Database', 'oracle-database', 'Oracle Database (commonly referred to as Oracle DBMS, Oracle Autonomous Database, or simply as Oracle) is a multi-model[4] database management system ', '<p><strong>Oracle Database</strong>&nbsp;(commonly referred to as&nbsp;<strong>Oracle DBMS</strong>,&nbsp;<strong>Oracle Autonomous Database</strong>, or simply as&nbsp;<strong>Oracle</strong>) is a&nbsp;<a href="https://en.wikipedia.org/wiki/Multi-model_database" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">multi-model</a><a href="https://en.wikipedia.org/wiki/Oracle_Database#cite_note-4" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);"><sup>[4]</sup></a>&nbsp;<a href="https://en.wikipedia.org/wiki/Database_management_system" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">database management system</a>&nbsp;produced and marketed by&nbsp;<a href="https://en.wikipedia.org/wiki/Oracle_Corporation" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">Oracle Corporation</a>.</p><p>It is a database commonly used for running&nbsp;<a href="https://en.wikipedia.org/wiki/Online_transaction_processing" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">online transaction processing</a>&nbsp;(OLTP),&nbsp;<a href="https://en.wikipedia.org/wiki/Data_warehouse" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">data warehousing</a>&nbsp;(DW) and mixed (OLTP &amp; DW) database workloads. Oracle Database is available by several service providers&nbsp;<a href="https://en.wikipedia.org/wiki/On-premises_software" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">on-prem</a>,&nbsp;<a href="https://en.wikipedia.org/wiki/Cloud_computing" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">on-cloud</a>, or as a hybrid cloud installation. It may be run on third party servers as well as on Oracle hardware (<a href="https://en.wikipedia.org/wiki/Oracle_Exadata" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">Exadata</a>&nbsp;on-prem, on&nbsp;<a href="https://en.wikipedia.org/wiki/Oracle_Cloud" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">Oracle Cloud</a>&nbsp;or at Cloud at Customer).<a href="https://en.wikipedia.org/wiki/Oracle_Database#cite_note-5" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);"><sup>[5]</sup></a></p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1675405573/courses/oracle-database.png', 'Javin Djapri', 'publish', '2023-02-03 13:26:14.245517', '2023-02-03 13:26:14.245517', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1675405573/courses/oracle-database.png', 5000);
INSERT INTO public.courses VALUES (34, 2, 'Cara Menjadi Marco Loen', 'cara-menjadi-marco-loen', 'Marco Loen adalah seorang pemuda asal Riau', '<p>Marco Loen adalah seorang pemuda asal Riau</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1675392925/courses/cara-menjadi-marco-loen.png', 'Marco Loen', 'publish', '2023-02-03 09:55:26.304432', '2023-02-03 11:18:13.624107', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1675392925/courses/cara-menjadi-marco-loen.png', 970000);
INSERT INTO public.courses VALUES (35, 1, 'Cara Menjadi Javin Djapri', 'cara-menjadi-javin-djapri', 'Javin Djapri adalah seorang pemuda dari Kalideres', '<p>Javin Djapri adalah seorang pemuda dari Kalideres yang setiap hari berangkat kerja naik Transjakarta</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1675400107/courses/cara-menjadi-javin-djapri.png', 'Javin Djapri', 'publish', '2023-02-03 11:55:08.440304', '2023-02-03 11:55:08.440304', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1675400107/courses/cara-menjadi-javin-djapri.png', 90000);
INSERT INTO public.courses VALUES (37, 1, 'MongoDB', 'mongodb', 'Be A Back End Expert', '<p><strong style="color: rgb(32, 33, 34);">MongoDB</strong><span style="color: rgb(32, 33, 34);">&nbsp;is a&nbsp;</span><a href="https://en.wikipedia.org/wiki/Source-available" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">source-available</a><span style="color: rgb(32, 33, 34);">&nbsp;</span><a href="https://en.wikipedia.org/wiki/Cross-platform" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">cross-platform</a><span style="color: rgb(32, 33, 34);">&nbsp;</span><a href="https://en.wikipedia.org/wiki/Document-oriented_database" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">document-oriented database</a><span style="color: rgb(32, 33, 34);">&nbsp;program. Classified as a&nbsp;</span><a href="https://en.wikipedia.org/wiki/NoSQL" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">NoSQL</a><span style="color: rgb(32, 33, 34);">&nbsp;database program, MongoDB uses&nbsp;</span><a href="https://en.wikipedia.org/wiki/JSON" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">JSON</a><span style="color: rgb(32, 33, 34);">-like documents with optional&nbsp;</span><a href="https://en.wikipedia.org/wiki/Database_schema" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">schemas</a><span style="color: rgb(32, 33, 34);">. MongoDB is developed by&nbsp;</span><a href="https://en.wikipedia.org/wiki/MongoDB_Inc." rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">MongoDB Inc.</a><span style="color: rgb(32, 33, 34);">&nbsp;and licensed under the&nbsp;</span><a href="https://en.wikipedia.org/wiki/Server_Side_Public_License" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204);">Server Side Public License</a><span style="color: rgb(32, 33, 34);">&nbsp;(SSPL) which is deemed non-free by several distributions.</span></p><p><br></p><p><span class="ql-cursor">﻿</span>10gen software company began developing MongoDB in 2007 as a component of a planned&nbsp;<a href="https://en.wikipedia.org/wiki/Platform_as_a_service" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">platform as a service</a>&nbsp;product. In 2009, the company shifted to an open-source development model, with the company offering commercial support and other services. In 2013, 10gen changed its name to MongoDB Inc.<a href="https://en.wikipedia.org/wiki/MongoDB#cite_note-gigaom-rename-5" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);"><sup>[5]</sup></a></p><p>On October 20, 2017, MongoDB became a publicly traded company, listed on NASDAQ as MDB with an IPO price of $24 per share.<a href="https://en.wikipedia.org/wiki/MongoDB#cite_note-6" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);"><sup>[6]</sup></a></p><p>MongoDB is a global company with US headquarters in New York City, USA and International headquarters in Dublin, Ireland.</p><p>On October 30, 2019, MongoDB teamed up with&nbsp;<a href="https://en.wikipedia.org/wiki/Alibaba_Cloud" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);">Alibaba Cloud</a>, who will offer its customers a MongoDB-as-a-service solution. Customers can use the managed offering from BABA''s global data centers.<a href="https://en.wikipedia.org/wiki/MongoDB#cite_note-7" rel="noopener noreferrer" target="_blank" style="background-color: initial; color: rgb(51, 102, 204);"><sup>[7]</sup></a></p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1675405270/courses/mongodb.png', 'Marco Loen', 'publish', '2023-02-03 13:21:10.506433', '2023-02-03 13:21:10.506433', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1675405270/courses/mongodb.png', 9000);
INSERT INTO public.courses VALUES (36, 2, 'Merisa Syafrina 101', 'merisa-syafrina-101', 'Impassioned to gain more experiences in Front End Development and is an enthusiast in UI/UX design. Experienced working in a team, open to any ideas a', '<p><span style="color: rgba(0, 0, 0, 0.9);">Impassioned to gain more experiences in Front End Development and is an enthusiast in UI/UX design. Experienced working in a team, open to any ideas and eager to find solutions to overcome problems.</span></p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1675405073/courses/merisa-syafrina-101.png', 'Merisa Syafrina', 'publish', '2023-02-03 13:17:53.918687', '2023-02-03 15:56:53.474233', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1675405073/courses/merisa-syafrina-101.png', 5000000);
INSERT INTO public.courses VALUES (39, 1, 'RabbitMQ', 'rabbitmq', 'RabbitMQ is an open-source message-broker software (sometimes called message-oriented middleware) that originally implemented the Advanced Message Que', '<p><strong>RabbitMQ</strong>&nbsp;is an open-source&nbsp;<a href="https://en.wikipedia.org/wiki/Message_broker" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">message-broker</a>&nbsp;software (sometimes called&nbsp;<a href="https://en.wikipedia.org/wiki/Message-oriented_middleware" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">message-oriented middleware</a>) that originally implemented the&nbsp;<a href="https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">Advanced Message Queuing Protocol</a>&nbsp;(AMQP) and has since been extended with a&nbsp;<a href="https://en.wikipedia.org/wiki/Plug-in_(computing)" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">plug-in architecture</a>&nbsp;to support&nbsp;<a href="https://en.wikipedia.org/wiki/Streaming_Text_Oriented_Messaging_Protocol" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">Streaming Text Oriented Messaging Protocol</a>&nbsp;(STOMP),&nbsp;<a href="https://en.wikipedia.org/wiki/MQ_Telemetry_Transport" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">MQ Telemetry Transport</a>&nbsp;(MQTT), and other protocols.<a href="https://en.wikipedia.org/wiki/RabbitMQ#cite_note-1" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;"><sup>[1]</sup></a></p><p>Written in&nbsp;<a href="https://en.wikipedia.org/wiki/Erlang_(programming_language)" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">Erlang</a>, the RabbitMQ server is built on the&nbsp;<a href="https://en.wikipedia.org/wiki/Open_Telecom_Platform" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">Open Telecom Platform</a>&nbsp;framework for clustering and failover. Client libraries to interface with the broker are available for all major programming languages. The source code is released under the&nbsp;<a href="https://en.wikipedia.org/wiki/Mozilla_Public_License" rel="noopener noreferrer" target="_blank" style="color: rgb(51, 102, 204); background-color: initial;">Mozilla Public License</a>.</p>', 'https://res.cloudinary.com/dz70b8vjw/image/upload/v1675405642/courses/rabbitmq.png', 'Marco Loen', 'publish', '2023-02-03 13:27:23.263047', '2023-02-03 14:20:08.936091', NULL, 'https://res.cloudinary.com/dz70b8vjw/image/upload/w_auto,h_300,c_fill,g_auto,f_auto/v1675405642/courses/rabbitmq.png', 10000);


--
-- TOC entry 3598 (class 0 OID 17590)
-- Dependencies: 221
-- Data for Name: favorites; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.favorites VALUES (19, 7, 26, '2023-01-23 12:40:35.407917', '2023-01-23 12:40:35.408076', '2023-01-23 12:40:35.408076', NULL);
INSERT INTO public.favorites VALUES (22, 1, 25, '2023-01-24 10:05:12.196641', '2023-01-24 10:05:12.196731', '2023-01-24 10:05:12.196731', NULL);
INSERT INTO public.favorites VALUES (31, 7, 31, '2023-02-01 15:37:42.194887', '2023-02-01 15:37:42.195006', '2023-02-01 15:37:42.195006', NULL);
INSERT INTO public.favorites VALUES (55, 13, 27, '2023-02-02 14:16:30.40271', '2023-02-02 14:16:30.402893', '2023-02-02 14:16:30.402893', NULL);


--
-- TOC entry 3614 (class 0 OID 17680)
-- Dependencies: 237
-- Data for Name: gifts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3606 (class 0 OID 17635)
-- Dependencies: 229
-- Data for Name: invoices; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.invoices VALUES ('c2e46f7e-928e-4926-8875-de1264773e1c', 4, 'completed', 100000, '2023-01-24 21:24:59.049405', NULL, '2023-01-24 21:24:47.941958', '2023-01-24 21:25:47.864651', NULL, 0, 100000, NULL);
INSERT INTO public.invoices VALUES ('c2a9baf0-8283-4ffb-a5ae-b9caaf4ae5ff', 4, 'completed', 100000, '2023-01-24 19:45:42.065795', NULL, '2023-01-24 19:41:40.708713', '2023-01-24 19:45:57.019399', NULL, 0, 100000, NULL);
INSERT INTO public.invoices VALUES ('d97d6bb8-1340-45af-b83f-f9fb56603310', 7, 'completed', 175000, '2023-01-24 21:24:59.049', 1, '2023-01-24 14:11:35.643572', '2023-01-24 17:21:18.769343', NULL, 0, 200000, 1);
INSERT INTO public.invoices VALUES ('e1ad65e7-fd5d-40bb-9e78-a540c30595df', 13, 'completed', 1200000, '2023-01-30 16:50:29.147749', NULL, '2023-01-30 13:03:57.892611', '2023-01-30 16:50:49.728228', NULL, 0, 1200000, NULL);
INSERT INTO public.invoices VALUES ('55e1e04d-e5d8-4e5e-850c-8c0b4073aa37', 13, 'waiting_payment', 185000, NULL, NULL, '2023-02-02 15:14:13.977687', '2023-02-02 15:14:13.977687', NULL, 0, 185000, NULL);
INSERT INTO public.invoices VALUES ('89bfb51d-5e51-4ff4-a3f6-ecb807f9e5d3', 7, 'completed', 1160000, '2023-02-01 18:15:19.551882', 1, '2023-02-01 10:54:12.898389', '2023-02-03 11:42:01.890904', NULL, 0, 1185000, 5);
INSERT INTO public.invoices VALUES ('05264e56-f856-4046-a886-aebb3ffcd3d7', 14, 'waiting_payment', 1385000, NULL, NULL, '2023-02-03 14:31:14.696458', '2023-02-03 14:31:14.696458', NULL, 0, 1385000, NULL);
INSERT INTO public.invoices VALUES ('40fcf8a1-0844-4a43-b3e3-b5548dde3daf', 14, 'waiting_confirmation', 5000, '2023-02-03 15:45:17.928529', NULL, '2023-02-03 15:44:35.637814', '2023-02-03 15:45:17.928611', NULL, 0, 5000, NULL);
INSERT INTO public.invoices VALUES ('d32f892b-82f3-4e15-80c5-1a7432f69a13', 14, 'waiting_confirmation', 5000000, '2023-02-04 16:53:13.264591', NULL, '2023-02-04 16:52:50.826214', '2023-02-04 16:53:13.264879', NULL, 0, 5000000, NULL);


--
-- TOC entry 3592 (class 0 OID 17551)
-- Dependencies: 215
-- Data for Name: levels; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.levels VALUES (1, 'newbie', 0, '2023-01-17 17:34:16.519485', '2023-01-17 17:34:16.519485', NULL, 0, 'https://media.tenor.com/jwHkGGFNoH8AAAAi/shiny-charmander-pokemon.gif', 0);
INSERT INTO public.levels VALUES (2, 'junior', 0.05, '2023-01-25 11:47:04.816099', '2023-01-25 11:47:04.816099', NULL, 10, 'https://64.media.tumblr.com/tumblr_ma4fpfD6Tu1rfjowdo1_500.gif', 0);
INSERT INTO public.levels VALUES (3, 'senior', 0.1, '2023-01-25 11:47:04.816099', '2023-01-25 11:47:04.816099', NULL, 20, 'https://64.media.tumblr.com/4d40efa93e8cfe522fff1d81c311f49f/tumblr_mj4eheib8x1s3bc1no1_500.gifv', 1);
INSERT INTO public.levels VALUES (4, 'master', 0.2, '2023-01-25 11:47:04.816099', '2023-01-25 11:47:04.816099', NULL, 50, 'https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/abae6e28-7eb5-4262-8145-4dc4ac179c1d/d91l46h-b390d769-85c8-4608-ac79-2bce10bc9748.gif?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcL2FiYWU2ZTI4LTdlYjUtNDI2Mi04MTQ1LTRkYzRhYzE3OWMxZFwvZDkxbDQ2aC1iMzkwZDc2OS04NWM4LTQ2MDgtYWM3OS0yYmNlMTBiYzk3NDguZ2lmIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.damFVhL_HpN-XCxxG2vaOvmOp4jZkCLFYchcrW2dVjE', 2);


--
-- TOC entry 3622 (class 0 OID 17929)
-- Dependencies: 245
-- Data for Name: redeemables; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.redeemables VALUES (1, 9, 0, '2023-01-25 11:21:45.8777', '2023-01-25 11:21:45.8777', NULL);
INSERT INTO public.redeemables VALUES (2, 4, 0, '2023-01-25 11:24:00.773574', '2023-01-25 11:24:00.773574', NULL);
INSERT INTO public.redeemables VALUES (3, 1, 0, '2023-01-25 11:24:00.773574', '2023-01-25 11:24:00.773574', NULL);
INSERT INTO public.redeemables VALUES (4, 7, 0, '2023-01-25 11:24:00.773574', '2023-01-25 11:24:00.773574', NULL);
INSERT INTO public.redeemables VALUES (5, 13, 0, '2023-01-29 19:35:27.178661', '2023-01-29 19:35:27.178661', NULL);
INSERT INTO public.redeemables VALUES (6, 14, 0, '2023-02-02 09:17:05.84478', '2023-02-02 09:17:05.84478', NULL);


--
-- TOC entry 3590 (class 0 OID 17538)
-- Dependencies: 213
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.roles VALUES (1, 'admin', '2023-01-17 17:34:14.580294', '2023-01-17 17:34:14.580294', NULL);
INSERT INTO public.roles VALUES (2, 'user', '2023-01-17 21:35:13.764704', '2023-01-17 21:35:13.764704', NULL);


--
-- TOC entry 3602 (class 0 OID 17613)
-- Dependencies: 225
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.tags VALUES (23, 'back end', '2023-01-19 13:38:32.239696', '2023-01-19 13:38:32.239696', NULL);
INSERT INTO public.tags VALUES (25, 'master', '2023-01-19 15:37:22.323784', '2023-01-19 15:37:22.323784', NULL);
INSERT INTO public.tags VALUES (26, 'back', '2023-01-19 17:02:41.828935', '2023-01-19 17:02:41.828935', NULL);
INSERT INTO public.tags VALUES (27, 'front end', '2023-01-19 18:15:05.580769', '2023-01-19 18:15:05.580769', NULL);
INSERT INTO public.tags VALUES (28, 'newbie', '2023-01-20 15:44:07.947247', '2023-01-20 15:44:07.947247', NULL);
INSERT INTO public.tags VALUES (29, 'intermediate', '2023-01-20 17:54:03.998184', '2023-01-20 17:54:03.998184', NULL);
INSERT INTO public.tags VALUES (30, 'Expert', '2023-01-20 18:47:44.106918', '2023-01-20 18:47:44.106918', NULL);


--
-- TOC entry 3616 (class 0 OID 17693)
-- Dependencies: 239
-- Data for Name: tracks; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3608 (class 0 OID 17647)
-- Dependencies: 231
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions VALUES (20, 'd97d6bb8-1340-45af-b83f-f9fb56603310', 26, 100000, '2023-01-24 14:11:35.646285', '2023-01-24 14:11:35.646285', NULL);
INSERT INTO public.transactions VALUES (19, 'd97d6bb8-1340-45af-b83f-f9fb56603310', 25, 100000, '2023-01-24 14:11:35.646', '2023-01-24 14:11:35.646285', NULL);
INSERT INTO public.transactions VALUES (21, 'c2a9baf0-8283-4ffb-a5ae-b9caaf4ae5ff', 25, 100000, '2023-01-24 19:41:40.709', '2023-01-24 19:41:40.709225', NULL);
INSERT INTO public.transactions VALUES (22, 'c2e46f7e-928e-4926-8875-de1264773e1c', 27, 100000, '2023-01-24 21:24:47.942952', '2023-01-24 21:24:47.942952', NULL);
INSERT INTO public.transactions VALUES (23, 'e1ad65e7-fd5d-40bb-9e78-a540c30595df', 26, 100000, '2023-01-30 13:03:57.911111', '2023-01-30 13:03:57.911111', NULL);
INSERT INTO public.transactions VALUES (24, 'e1ad65e7-fd5d-40bb-9e78-a540c30595df', 25, 100000, '2023-01-30 13:03:57.911111', '2023-01-30 13:03:57.911111', NULL);
INSERT INTO public.transactions VALUES (25, 'e1ad65e7-fd5d-40bb-9e78-a540c30595df', 27, 1000000, '2023-01-30 13:03:57.911111', '2023-01-30 13:03:57.911111', NULL);
INSERT INTO public.transactions VALUES (26, '89bfb51d-5e51-4ff4-a3f6-ecb807f9e5d3', 31, 90000, '2023-02-01 10:54:12.911807', '2023-02-01 10:54:12.911807', NULL);
INSERT INTO public.transactions VALUES (27, '89bfb51d-5e51-4ff4-a3f6-ecb807f9e5d3', 33, 95000, '2023-02-01 10:54:12.911807', '2023-02-01 10:54:12.911807', NULL);
INSERT INTO public.transactions VALUES (28, '89bfb51d-5e51-4ff4-a3f6-ecb807f9e5d3', 27, 1000000, '2023-02-01 10:54:12.911807', '2023-02-01 10:54:12.911807', NULL);
INSERT INTO public.transactions VALUES (29, '55e1e04d-e5d8-4e5e-850c-8c0b4073aa37', 31, 90000, '2023-02-02 15:14:13.990849', '2023-02-02 15:14:13.990849', NULL);
INSERT INTO public.transactions VALUES (30, '55e1e04d-e5d8-4e5e-850c-8c0b4073aa37', 33, 95000, '2023-02-02 15:14:13.990849', '2023-02-02 15:14:13.990849', NULL);
INSERT INTO public.transactions VALUES (31, '05264e56-f856-4046-a886-aebb3ffcd3d7', 26, 100000, '2023-02-03 14:31:14.707402', '2023-02-03 14:31:14.707402', NULL);
INSERT INTO public.transactions VALUES (32, '05264e56-f856-4046-a886-aebb3ffcd3d7', 25, 100000, '2023-02-03 14:31:14.707402', '2023-02-03 14:31:14.707402', NULL);
INSERT INTO public.transactions VALUES (33, '05264e56-f856-4046-a886-aebb3ffcd3d7', 31, 90000, '2023-02-03 14:31:14.707402', '2023-02-03 14:31:14.707402', NULL);
INSERT INTO public.transactions VALUES (34, '05264e56-f856-4046-a886-aebb3ffcd3d7', 33, 95000, '2023-02-03 14:31:14.707402', '2023-02-03 14:31:14.707402', NULL);
INSERT INTO public.transactions VALUES (35, '05264e56-f856-4046-a886-aebb3ffcd3d7', 27, 1000000, '2023-02-03 14:31:14.707402', '2023-02-03 14:31:14.707402', NULL);
INSERT INTO public.transactions VALUES (36, '40fcf8a1-0844-4a43-b3e3-b5548dde3daf', 38, 5000, '2023-02-03 15:44:35.639631', '2023-02-03 15:44:35.639631', NULL);
INSERT INTO public.transactions VALUES (37, 'd32f892b-82f3-4e15-80c5-1a7432f69a13', 36, 5000000, '2023-02-04 16:52:50.828826', '2023-02-04 16:52:50.828826', NULL);


--
-- TOC entry 3596 (class 0 OID 17579)
-- Dependencies: 219
-- Data for Name: user_courses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_courses VALUES (10, 7, 26, 'in_progress', '2023-01-24 17:21:18.771535', '2023-01-24 17:21:18.771535', NULL);
INSERT INTO public.user_courses VALUES (11, 4, 25, 'in_progress', '2023-01-24 19:45:57.020465', '2023-01-24 19:45:57.020465', NULL);
INSERT INTO public.user_courses VALUES (12, 4, 27, 'in_progress', '2023-01-24 21:25:47.865406', '2023-01-24 21:25:47.865406', NULL);
INSERT INTO public.user_courses VALUES (9, 7, 25, 'completed', '2023-01-24 17:21:18.771535', '2023-01-31 16:53:44.871721', NULL);
INSERT INTO public.user_courses VALUES (15, 13, 27, 'completed', '2023-01-30 16:50:49.728741', '2023-02-02 14:16:15.965938', NULL);
INSERT INTO public.user_courses VALUES (14, 13, 25, 'completed', '2023-01-30 16:50:49.728741', '2023-02-02 14:28:02.210301', NULL);
INSERT INTO public.user_courses VALUES (13, 13, 26, 'completed', '2023-01-30 16:50:49.728741', '2023-02-02 14:28:15.312047', NULL);
INSERT INTO public.user_courses VALUES (16, 7, 31, 'in_progress', '2023-02-03 11:42:01.89141', '2023-02-03 11:42:01.89141', NULL);
INSERT INTO public.user_courses VALUES (17, 7, 33, 'in_progress', '2023-02-03 11:42:01.89141', '2023-02-03 11:42:01.89141', NULL);
INSERT INTO public.user_courses VALUES (18, 7, 27, 'in_progress', '2023-02-03 11:42:01.89141', '2023-02-03 11:42:01.89141', NULL);


--
-- TOC entry 3618 (class 0 OID 17705)
-- Dependencies: 241
-- Data for Name: user_gifts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3612 (class 0 OID 17669)
-- Dependencies: 235
-- Data for Name: user_vouchers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_vouchers VALUES (1, 7, 1, '2023-01-30 19:46:39.933856', true, '2023-01-23 19:46:39.933856', '2023-01-24 14:11:35.648587', NULL);
INSERT INTO public.user_vouchers VALUES (3, 7, 1, '2023-01-17 22:57:07.34174', false, '2023-01-24 22:57:07.34174', '2023-01-24 22:57:07.34174', NULL);
INSERT INTO public.user_vouchers VALUES (4, 7, 1, '2023-01-31 22:57:27.349541', false, '2023-01-24 22:57:27.349541', '2023-01-24 22:57:27.349541', NULL);
INSERT INTO public.user_vouchers VALUES (5, 7, 1, '2023-02-27 22:57:46.838', true, '2023-01-24 22:57:46.838192', '2023-02-01 10:54:12.913943', NULL);
INSERT INTO public.user_vouchers VALUES (6, 4, 3, '2023-03-03 11:42:01.903032', false, '2023-02-03 11:42:01.903048', '2023-02-03 11:42:01.903048', NULL);


--
-- TOC entry 3588 (class 0 OID 17521)
-- Dependencies: 211
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (4, 'acong@mail.com', '$2a$04$J.ryjplDRmhTKhhDpAQtlOPIEoXvtpfJwyg3pDy7I5EHCaGwfcgE6', 2, 'acong', 'acong bahlul', 'cipulir', '089274628174', 1, 'e04q3l5', NULL, '2023-01-18 15:40:55.175074', '2023-01-18 15:40:55.175074', NULL);
INSERT INTO public.users VALUES (7, 'rachelvenya@mail.com', '$2a$04$eDO/PknM.elF2s3fczrPauK33lsRf1vYn6SCJsQPPA9hw3M/RrRea', 2, 'rachelvenya', 'Rachel Venya', 'cipulir', '088977955978', 1, 'Asx8kZ4', 'e04q3l5', '2023-01-18 16:05:51.177429', '2023-01-18 16:05:51.177429', NULL);
INSERT INTO public.users VALUES (1, 'andra@gmail.com', '$2a$12$D45Zr2uvcy/zvNH.yfQ8UuctMkAIswM9ITfv3HfOon1iSN9DKj.HK', 1, 'andra', 'Andra Adhiatma Nugraha', 'kebayoran lama', '082139154623', 1, 'XBBG357', NULL, '2023-01-17 17:37:00.70602', '2023-01-20 21:57:24.459152', NULL);
INSERT INTO public.users VALUES (9, 'okin@mail.com', '$2a$04$faxoL0y4jyiLX7QgX02fO.KIBRWUYIcIwk74Lj7bD3onzq9P.w1N6', 2, 'okin', 'Okin Aja', 'cipulir', '081234567890', 1, 'T3ZX8pQ', NULL, '2023-01-25 11:21:45.869146', '2023-01-25 11:21:45.869146', NULL);
INSERT INTO public.users VALUES (14, 'rezza@mail.com', '$2a$04$9X9O1Eks1ZEWfVc0l8vW/e1A/i.42i8bm317jhaTNnVbpUOXbIvy6', 2, 'wanjaldy', 'Rezza Aldy Sofyan', 'Jatiwaringin', '087284618374', 1, 'UXcslYa', NULL, '2023-02-02 09:17:05.843929', '2023-02-03 15:44:02.011203', NULL);
INSERT INTO public.users VALUES (13, 'guntur@mail.com', '$2a$04$Fy1mzQIJA3s75O2znhTIve32EZK9bXPd2krER8x.RpxBorNyRy5LC', 2, 'guntur', 'Gunturrrrrr', 'Gandaria Selatan', '089719283746', 1, 'v5abmkN', NULL, '2023-01-29 19:35:27.177237', '2023-02-02 08:57:49.494707', NULL);


--
-- TOC entry 3610 (class 0 OID 17656)
-- Dependencies: 233
-- Data for Name: vouchers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.vouchers VALUES (1, 'Discount 25k', 'HEMAT25', '2023-01-23 18:07:33.294766', '2023-01-23 18:07:33.294766', NULL, 25000, 200000);
INSERT INTO public.vouchers VALUES (2, 'Discount 50k', 'HEMAT50', '2023-01-23 18:07:33.294766', '2023-01-23 18:07:33.294766', NULL, 50000, 500000);
INSERT INTO public.vouchers VALUES (3, 'Discount 100k', 'HEMAT100', '2023-01-23 18:07:33.294766', '2023-01-23 18:07:33.294766', NULL, 100000, 1000000);


--
-- TOC entry 3648 (class 0 OID 0)
-- Dependencies: 242
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.carts_id_seq', 32, true);


--
-- TOC entry 3649 (class 0 OID 0)
-- Dependencies: 222
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 2, true);


--
-- TOC entry 3650 (class 0 OID 0)
-- Dependencies: 226
-- Name: course_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.course_tags_id_seq', 151, true);


--
-- TOC entry 3651 (class 0 OID 0)
-- Dependencies: 216
-- Name: courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.courses_id_seq', 39, true);


--
-- TOC entry 3652 (class 0 OID 0)
-- Dependencies: 220
-- Name: favorites_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.favorites_id_seq', 62, true);


--
-- TOC entry 3653 (class 0 OID 0)
-- Dependencies: 236
-- Name: gifts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gifts_id_seq', 1, false);


--
-- TOC entry 3654 (class 0 OID 0)
-- Dependencies: 228
-- Name: invoices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoices_id_seq', 16, true);


--
-- TOC entry 3655 (class 0 OID 0)
-- Dependencies: 214
-- Name: levels_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.levels_id_seq', 5, true);


--
-- TOC entry 3656 (class 0 OID 0)
-- Dependencies: 244
-- Name: redeemables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.redeemables_id_seq', 6, true);


--
-- TOC entry 3657 (class 0 OID 0)
-- Dependencies: 212
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 2, true);


--
-- TOC entry 3658 (class 0 OID 0)
-- Dependencies: 224
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tags_id_seq', 30, true);


--
-- TOC entry 3659 (class 0 OID 0)
-- Dependencies: 238
-- Name: tracks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tracks_id_seq', 1, false);


--
-- TOC entry 3660 (class 0 OID 0)
-- Dependencies: 230
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 37, true);


--
-- TOC entry 3661 (class 0 OID 0)
-- Dependencies: 218
-- Name: user_courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_courses_id_seq', 18, true);


--
-- TOC entry 3662 (class 0 OID 0)
-- Dependencies: 240
-- Name: user_gifts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_gifts_id_seq', 1, false);


--
-- TOC entry 3663 (class 0 OID 0)
-- Dependencies: 234
-- Name: user_vouchers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_vouchers_id_seq', 6, true);


--
-- TOC entry 3664 (class 0 OID 0)
-- Dependencies: 210
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 14, true);


--
-- TOC entry 3665 (class 0 OID 0)
-- Dependencies: 232
-- Name: vouchers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vouchers_id_seq', 3, true);


--
-- TOC entry 3420 (class 2606 OID 17840)
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (id);


--
-- TOC entry 3422 (class 2606 OID 17852)
-- Name: carts carts_un; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_un UNIQUE (user_id, course_id);


--
-- TOC entry 3390 (class 2606 OID 17611)
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- TOC entry 3392 (class 2606 OID 17609)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 3398 (class 2606 OID 17633)
-- Name: course_tags course_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_tags
    ADD CONSTRAINT course_tags_pkey PRIMARY KEY (id);


--
-- TOC entry 3400 (class 2606 OID 17821)
-- Name: course_tags course_tags_un; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_tags
    ADD CONSTRAINT course_tags_un UNIQUE (course_id, tag_id);


--
-- TOC entry 3378 (class 2606 OID 17573)
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- TOC entry 3380 (class 2606 OID 17577)
-- Name: courses courses_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_slug_key UNIQUE (slug);


--
-- TOC entry 3382 (class 2606 OID 17575)
-- Name: courses courses_title_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_title_key UNIQUE (title);


--
-- TOC entry 3386 (class 2606 OID 17598)
-- Name: favorites favorites_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_pkey PRIMARY KEY (id);


--
-- TOC entry 3388 (class 2606 OID 17831)
-- Name: favorites favorites_un; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_un UNIQUE (user_id, course_id);


--
-- TOC entry 3412 (class 2606 OID 17691)
-- Name: gifts gifts_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_name_key UNIQUE (name);


--
-- TOC entry 3414 (class 2606 OID 17689)
-- Name: gifts gifts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_pkey PRIMARY KEY (id);


--
-- TOC entry 3402 (class 2606 OID 17911)
-- Name: invoices invoices_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_pkey PRIMARY KEY (id);


--
-- TOC entry 3374 (class 2606 OID 17562)
-- Name: levels levels_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_name_key UNIQUE (name);


--
-- TOC entry 3376 (class 2606 OID 17560)
-- Name: levels levels_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_pkey PRIMARY KEY (id);


--
-- TOC entry 3424 (class 2606 OID 17936)
-- Name: redeemables redeemables_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.redeemables
    ADD CONSTRAINT redeemables_pkey PRIMARY KEY (id);


--
-- TOC entry 3370 (class 2606 OID 17549)
-- Name: roles roles_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


--
-- TOC entry 3372 (class 2606 OID 17547)
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- TOC entry 3394 (class 2606 OID 17624)
-- Name: tags tags_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_name_key UNIQUE (name);


--
-- TOC entry 3396 (class 2606 OID 17622)
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- TOC entry 3416 (class 2606 OID 17703)
-- Name: tracks tracks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT tracks_pkey PRIMARY KEY (id);


--
-- TOC entry 3404 (class 2606 OID 17654)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3384 (class 2606 OID 17588)
-- Name: user_courses user_courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_courses
    ADD CONSTRAINT user_courses_pkey PRIMARY KEY (id);


--
-- TOC entry 3418 (class 2606 OID 17712)
-- Name: user_gifts user_gifts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT user_gifts_pkey PRIMARY KEY (id);


--
-- TOC entry 3410 (class 2606 OID 17678)
-- Name: user_vouchers user_vouchers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers
    ADD CONSTRAINT user_vouchers_pkey PRIMARY KEY (id);


--
-- TOC entry 3360 (class 2606 OID 17532)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 3362 (class 2606 OID 17823)
-- Name: users users_phone_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_no_key UNIQUE (phone_no);


--
-- TOC entry 3364 (class 2606 OID 17530)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3366 (class 2606 OID 17536)
-- Name: users users_referral_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_referral_key UNIQUE (referral);


--
-- TOC entry 3368 (class 2606 OID 17534)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 3406 (class 2606 OID 17667)
-- Name: vouchers vouchers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT vouchers_name_key UNIQUE (name);


--
-- TOC entry 3408 (class 2606 OID 17665)
-- Name: vouchers vouchers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT vouchers_pkey PRIMARY KEY (id);


--
-- TOC entry 3445 (class 2606 OID 17846)
-- Name: carts carts_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3446 (class 2606 OID 17841)
-- Name: carts carts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3432 (class 2606 OID 17748)
-- Name: course_tags course_tags_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_tags
    ADD CONSTRAINT course_tags_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3433 (class 2606 OID 17753)
-- Name: course_tags course_tags_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_tags
    ADD CONSTRAINT course_tags_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- TOC entry 3427 (class 2606 OID 17723)
-- Name: courses courses_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- TOC entry 3430 (class 2606 OID 17743)
-- Name: favorites favorites_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3431 (class 2606 OID 17738)
-- Name: favorites favorites_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3434 (class 2606 OID 17860)
-- Name: invoices invoices_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_fk FOREIGN KEY (user_voucher_id) REFERENCES public.user_vouchers(id);


--
-- TOC entry 3435 (class 2606 OID 17758)
-- Name: invoices invoices_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3436 (class 2606 OID 17763)
-- Name: invoices invoices_voucher_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_voucher_id_fkey FOREIGN KEY (voucher_id) REFERENCES public.vouchers(id);


--
-- TOC entry 3447 (class 2606 OID 17937)
-- Name: redeemables redeemables_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.redeemables
    ADD CONSTRAINT redeemables_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3441 (class 2606 OID 17788)
-- Name: tracks tracks_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT tracks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3437 (class 2606 OID 17773)
-- Name: transactions transactions_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3438 (class 2606 OID 17923)
-- Name: transactions transactions_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_fk FOREIGN KEY (invoice_id) REFERENCES public.invoices(id);


--
-- TOC entry 3428 (class 2606 OID 17733)
-- Name: user_courses user_courses_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_courses
    ADD CONSTRAINT user_courses_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3429 (class 2606 OID 17728)
-- Name: user_courses user_courses_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_courses
    ADD CONSTRAINT user_courses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3442 (class 2606 OID 17798)
-- Name: user_gifts user_gifts_gift_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT user_gifts_gift_id_fkey FOREIGN KEY (gift_id) REFERENCES public.gifts(id);


--
-- TOC entry 3443 (class 2606 OID 17803)
-- Name: user_gifts user_gifts_track_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT user_gifts_track_id_fkey FOREIGN KEY (track_id) REFERENCES public.tracks(id);


--
-- TOC entry 3444 (class 2606 OID 17793)
-- Name: user_gifts user_gifts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_gifts
    ADD CONSTRAINT user_gifts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3439 (class 2606 OID 17778)
-- Name: user_vouchers user_vouchers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers
    ADD CONSTRAINT user_vouchers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3440 (class 2606 OID 17783)
-- Name: user_vouchers user_vouchers_voucher_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_vouchers
    ADD CONSTRAINT user_vouchers_voucher_id_fkey FOREIGN KEY (voucher_id) REFERENCES public.vouchers(id);


--
-- TOC entry 3425 (class 2606 OID 17718)
-- Name: users users_level_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_level_id_fkey FOREIGN KEY (level_id) REFERENCES public.levels(id);


--
-- TOC entry 3426 (class 2606 OID 17713)
-- Name: users users_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- TOC entry 3629 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-02-05 10:38:17 WIB

--
-- PostgreSQL database dump complete
--

