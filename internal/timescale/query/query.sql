-- name: GetCpuUtilization :many
SELECT *
FROM metrics
ORDER BY time
LIMIT 100;

-- name: InsertCpuUtilization :one
INSERT INTO metrics(time, instance_id, type, value)
VALUES ($1, $2, $3, $4)
RETURNING time
;

-- name: UpsertCpuUtilization :one
INSERT INTO metrics(time, instance_id, type, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING
RETURNING time;

-- name: CountCpuUtilization :one
SELECT count(time)
FROM metrics
WHERE time = $1
  AND type = $2
  AND instance_id = $3;


-- name: Average :one
SELECT time_bucket('10 minutes', time) AS period,
       instance_id,
       avg(value)
FROM metrics
WHERE instance_id = $1
GROUP BY period, instance_id
LIMIT 1;