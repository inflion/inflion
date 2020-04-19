CREATE TABLE public.metrics (
    "time" timestamp with time zone NOT NULL,
    instance_id character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    value double precision
);
CREATE INDEX metrics_time_idx ON public.metrics USING btree ("time" DESC);
CREATE INDEX metrics_type_time_idx ON public.metrics USING btree (type, "time" DESC);

SELECT create_hypertable('metrics', 'time');
