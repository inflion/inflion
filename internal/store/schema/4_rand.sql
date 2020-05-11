-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION public.rand()
    RETURNS text
    LANGUAGE sql
AS
$function$

SELECT array_to_string(array(select substr('abcdefghikjlmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789',
                                           ((random() * (36 - 1) + 1)::integer), 1)
                             from generate_series(1, 24)), '');

$function$
-- +migrate StatementEnd
