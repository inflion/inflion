CREATE FUNCTION public.rand() RETURNS text
    LANGUAGE sql
    AS $$
SELECT array_to_string(array(select substr('abcdefghikjlmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789',((random()*(36-1)+1)::integer),1) from generate_series(1,24)),'');
$$;
CREATE TABLE public.action (
    id bigint NOT NULL,
    user_id character varying(255) NOT NULL,
    project_id bigint NOT NULL,
    body jsonb DEFAULT jsonb_build_array() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    name character varying NOT NULL
);
CREATE SEQUENCE public.action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.action_id_seq OWNED BY public.action.id;
CREATE TABLE public.aws_account (
    id bigint NOT NULL,
    user_id character varying(255) NOT NULL,
    project_id bigint NOT NULL,
    account_id character varying(255) NOT NULL,
    role_name character varying(255) NOT NULL,
    external_id character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.aws_account_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.aws_account_id_seq OWNED BY public.aws_account.id;
CREATE TABLE public.instance (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    instance_id character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    project_id bigint NOT NULL,
    status character varying(255)
);
CREATE TABLE public.instance_at_service (
    id bigint NOT NULL,
    service_id bigint NOT NULL,
    instance_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.instance_at_service_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.instance_at_service_id_seq OWNED BY public.instance_at_service.id;
CREATE SEQUENCE public.instance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.instance_id_seq OWNED BY public.instance.id;
CREATE SEQUENCE public.instance_project_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.instance_project_id_seq OWNED BY public.instance.project_id;
CREATE TABLE public.notification_rule (
    id bigint NOT NULL,
    project_id bigint NOT NULL,
    rule_name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    rules jsonb DEFAULT jsonb_build_array() NOT NULL
);
CREATE SEQUENCE public.notification_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.notification_rule_id_seq OWNED BY public.notification_rule.id;
CREATE TABLE public.organization (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    unique_name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.organization_id_seq OWNED BY public.organization.id;
CREATE TABLE public.project (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    user_id character varying(255) NOT NULL,
    description text
);
CREATE TABLE public.project_collaborator (
    user_id character varying(255) NOT NULL,
    project_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.project_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.project_id_seq OWNED BY public.project.id;
CREATE TABLE public.project_in_organization (
    project_id bigint NOT NULL,
    organization_id bigint NOT NULL
);
CREATE TABLE public.project_invitation (
    id bigint NOT NULL,
    invitee_user_id character varying(255) NOT NULL,
    project_id bigint NOT NULL,
    mail_address character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    token character varying DEFAULT public.rand() NOT NULL,
    confirmed boolean DEFAULT false NOT NULL
);
CREATE SEQUENCE public.project_invitation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.project_invitation_id_seq OWNED BY public.project_invitation.id;
CREATE TABLE public.service (
    id bigint NOT NULL,
    project_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.service_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.service_id_seq OWNED BY public.service.id;
CREATE TABLE public.slack_webhook (
    id bigint NOT NULL,
    user_id character varying(255) NOT NULL,
    project_id bigint NOT NULL,
    webhook_url character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    channel character varying NOT NULL,
    name character varying NOT NULL
);
CREATE SEQUENCE public.slack_webhook_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.slack_webhook_id_seq OWNED BY public.slack_webhook.id;
CREATE TABLE public.user_account (
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    id character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
ALTER TABLE ONLY public.action ALTER COLUMN id SET DEFAULT nextval('public.action_id_seq'::regclass);
ALTER TABLE ONLY public.aws_account ALTER COLUMN id SET DEFAULT nextval('public.aws_account_id_seq'::regclass);
ALTER TABLE ONLY public.instance ALTER COLUMN id SET DEFAULT nextval('public.instance_id_seq'::regclass);
ALTER TABLE ONLY public.instance ALTER COLUMN project_id SET DEFAULT nextval('public.instance_project_id_seq'::regclass);
ALTER TABLE ONLY public.instance_at_service ALTER COLUMN id SET DEFAULT nextval('public.instance_at_service_id_seq'::regclass);
ALTER TABLE ONLY public.notification_rule ALTER COLUMN id SET DEFAULT nextval('public.notification_rule_id_seq'::regclass);
ALTER TABLE ONLY public.organization ALTER COLUMN id SET DEFAULT nextval('public.organization_id_seq'::regclass);
ALTER TABLE ONLY public.project ALTER COLUMN id SET DEFAULT nextval('public.project_id_seq'::regclass);
ALTER TABLE ONLY public.project_invitation ALTER COLUMN id SET DEFAULT nextval('public.project_invitation_id_seq'::regclass);
ALTER TABLE ONLY public.service ALTER COLUMN id SET DEFAULT nextval('public.service_id_seq'::regclass);
ALTER TABLE ONLY public.slack_webhook ALTER COLUMN id SET DEFAULT nextval('public.slack_webhook_id_seq'::regclass);
ALTER TABLE ONLY public.action
    ADD CONSTRAINT action_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.aws_account
    ADD CONSTRAINT aws_account_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.instance_at_service
    ADD CONSTRAINT instance_at_service_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.instance_at_service
    ADD CONSTRAINT instance_at_service_service_id_instance_id_key UNIQUE (service_id, instance_id);
ALTER TABLE ONLY public.instance
    ADD CONSTRAINT instance_instance_id_key UNIQUE (instance_id);
ALTER TABLE ONLY public.instance
    ADD CONSTRAINT instance_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.notification_rule
    ADD CONSTRAINT notification_rule_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.project_in_organization
    ADD CONSTRAINT project_in_organization_pkey PRIMARY KEY (project_id, organization_id);
ALTER TABLE ONLY public.project_invitation
    ADD CONSTRAINT project_invitation_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.project_invitation
    ADD CONSTRAINT project_invitation_token_key UNIQUE (token);
ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_project_id_name_key UNIQUE (project_id, name);
ALTER TABLE ONLY public.slack_webhook
    ADD CONSTRAINT slack_webhook_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.slack_webhook
    ADD CONSTRAINT slack_webhook_project_id_channel_key UNIQUE (project_id, channel);
ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_email_key UNIQUE (email);
ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_username_key UNIQUE (username);
ALTER TABLE ONLY public.action
    ADD CONSTRAINT action_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.action
    ADD CONSTRAINT action_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user_account(id);
ALTER TABLE ONLY public.aws_account
    ADD CONSTRAINT aws_account_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.aws_account
    ADD CONSTRAINT aws_account_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user_account(id);
ALTER TABLE ONLY public.instance_at_service
    ADD CONSTRAINT instance_at_service_instance_id_fkey FOREIGN KEY (instance_id) REFERENCES public.instance(id);
ALTER TABLE ONLY public.instance_at_service
    ADD CONSTRAINT instance_at_service_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.service(id);
ALTER TABLE ONLY public.instance
    ADD CONSTRAINT instance_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.notification_rule
    ADD CONSTRAINT notification_rule_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.project_collaborator
    ADD CONSTRAINT project_collaborator_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.project_collaborator
    ADD CONSTRAINT project_collaborator_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user_account(id);
ALTER TABLE ONLY public.project_in_organization
    ADD CONSTRAINT project_in_organization_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id);
ALTER TABLE ONLY public.project_in_organization
    ADD CONSTRAINT project_in_organization_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.project_invitation
    ADD CONSTRAINT project_invitation_invitee_user_id_fkey FOREIGN KEY (invitee_user_id) REFERENCES public.user_account(id);
ALTER TABLE ONLY public.project_invitation
    ADD CONSTRAINT project_invitation_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user_account(id);
ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.slack_webhook
    ADD CONSTRAINT slack_webhook_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
ALTER TABLE ONLY public.slack_webhook
    ADD CONSTRAINT slack_webhook_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user_account(id);
