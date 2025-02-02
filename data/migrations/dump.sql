--
-- PostgreSQL database dump
--

-- Dumped from database version 13.18 (Debian 13.18-1.pgdg120+1)
-- Dumped by pg_dump version 13.18 (Debian 13.18-1.pgdg120+1)

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
-- Name: comments; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.comments (
    id integer NOT NULL,
    author_id integer NOT NULL,
    post_id integer NOT NULL,
    parent_id integer,
    text text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.comments OWNER TO root;

--
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comments_id_seq OWNER TO root;

--
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.posts (
    id integer NOT NULL,
    author_id integer NOT NULL,
    text text NOT NULL,
    photo text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.posts OWNER TO root;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO root;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.refresh_tokens (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.refresh_tokens OWNER TO root;

--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.refresh_tokens_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.refresh_tokens_id_seq OWNER TO root;

--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.refresh_tokens_id_seq OWNED BY public.refresh_tokens.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: refresh_tokens id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.refresh_tokens ALTER COLUMN id SET DEFAULT nextval('public.refresh_tokens_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.comments (id, author_id, post_id, parent_id, text, created_at) FROM stdin;
1	6	2	\N	да вроде нет... мб помню что ли... Но слушай тебе понравится и может добавят тебя в избранное	2025-02-02 01:02:56.937913
2	6	1	\N	блин посмотри внимательнее пожалуйста....... эээээй Арсений... Посмотри внимательнее пожалуйста, не пострадал ли	2025-02-02 01:03:06.778189
3	6	4	\N	Пытаюсь шевелить конечностями.. 0_0...90мм.. 90мм....80мм.....1,	2025-02-02 01:03:35.012655
4	6	3	\N	Да нет вроде все нормально вроде вроде... ниче страшного же!!!!!!0_0!!!!!!!1	2025-02-02 01:03:42.555899
5	6	10	\N	Покушали......Вот это я бы хотела заказать..Хотя блинчеки уже достала уже наверное...........Ни	2025-02-02 01:04:32.095591
6	6	9	\N	0_о дааааа понимаю тебя тоже..._._ ыхазхвзахваз	2025-02-02 01:04:39.845247
7	6	8	\N	ну если мой есть то я могу попробовать......... короче они озадачились моим расчетам, и посоветовали сходить в	2025-02-02 01:04:48.838584
8	6	7	\N	Ты хочешь сидеть дома а? Или наоборот хочешь посидеть отдохнуть? Давай выбирай места поуютнее например скамейки	2025-02-02 01:04:55.997477
9	6	6	\N	И яяя тебя люблюююю сильно сильно очень сильно больше чем когда ты чувствуешь нехватку чего то мощного и	2025-02-02 01:05:03.674365
10	6	5	\N	вечером конечно пойду же......... малыш.... надеюсь не сильно задерживаться я ведь тоже хочу спать дальше спать а мама	2025-02-02 01:05:11.005416
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.posts (id, author_id, text, photo, created_at) FROM stdin;
1	2	Стикеры рисую и играю все еще... [EPSYP) сижу каракули читаю.........	\N	2025-02-02 00:36:50.357125
2	3	Я думала ты на афтобсуе уже сидишь ждешь чего нибудь?.._._.._	\N	2025-02-02 01:02:07.964519
3	1	Отпросился что бы тебе на почту отвечать почаще..??... Не знаю блинб....... не отвечают	\N	2025-02-02 01:02:54.180895
4	4	Пытаюсь шевелиться ааааюца..._._.........0_...0,9	\N	2025-02-02 01:03:04.063877
8	4	Ну ты же можешь попросить меня купить энерготик на двак дешевле..??... ээээ хз...	\N	2025-02-02 01:03:50.924794
10	5	Нада в покер поиграть что ли... Или кальянчек.....Поесть зашли.......Морож	\N	2025-02-02 01:03:59.880197
6	1	Ну я хочу что бы ты знал что я тебя любит! Я знаю!!!Ахаххавхыхах	\N	2025-02-02 01:03:39.832067
9	5	Я просто слюни пускала когда он рядом сидела жопу вытирала...0_0.7 Прости пожалуйста	\N	2025-02-02 01:03:55.255016
5	2	я думала ты следом напишешь что плохо себя чувствуешь и придешь домой..?... ээээ ладно,	\N	2025-02-02 01:03:32.254026
7	3	Ну я хочу что бы ты посидел вот так посидеть сегодня вечером. Я устала просто ужасно да, захотелось выбраться из	\N	2025-02-02 01:03:46.377309
11	6	Я просто еду на такси с коллегой по работе. Он такой привет заява отдает..0_0 Что	\N	2025-02-02 01:04:29.318672
12	6	ну я вроде выспалась после сна во сне еще сильнее чем вчера..0_0 Надеюсь нет... Температура спала	\N	2025-02-02 01:04:36.859537
13	6	Иду в душ наверное уже... Ноги две минуты валяются просто..._._.(SOP]	\N	2025-02-02 01:04:46.044973
14	6	Ну можно попросить тебя прибраться там внизу что ли... мб....... ээ ну даа	\N	2025-02-02 01:04:53.231555
15	6	Пытаюсь записаться к лору на рентген плеча. Бля уже второй раз.. сижу жду благодарностей от	\N	2025-02-02 01:05:00.946635
16	6	Блин забыла уже как мы начали Гарри поттера смотреть... Был целый день вчера был.. Ужасно хочу спать....	\N	2025-02-02 01:05:08.295677
17	6	Я все равно быстрее теюярюсеньки рисовать начну завтра утром..!!!0_.5 мб	\N	2025-02-02 01:05:15.123135
18	6	А если серьезно то надо выбрать фотку подходящую под настроение... Мне нравится такое.. [sEp] Ну	\N	2025-02-02 01:05:22.675685
\.


--
-- Data for Name: refresh_tokens; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.refresh_tokens (id, user_id, token, created_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.users (id, username, password_hash, created_at) FROM stdin;
1	МихаилGPT	a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3	2025-02-02 00:36:41.863899
2	АртурGPT	a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3	2025-02-02 00:36:41.863899
3	РомаGPT	a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3	2025-02-02 00:36:41.863899
4	РусланGPT	a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3	2025-02-02 00:36:41.863899
5	СеняGPT	a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3	2025-02-02 00:36:41.863899
6	ЕваGPT	a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3	2025-02-02 00:36:41.863899
\.


--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.comments_id_seq', 10, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.posts_id_seq', 18, true);


--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.refresh_tokens_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.users_id_seq', 6, true);


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_user_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_user_id_key UNIQUE (user_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: comments comments_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: comments comments_parent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.comments(id) ON DELETE CASCADE;


--
-- Name: comments comments_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: posts posts_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: refresh_tokens refresh_tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

