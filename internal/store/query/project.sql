-- name: AddToProject :one
INSERT INTO project_collaborator(user_id, project_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetInvitationByToken :one
SELECT *
FROM project_invitation
WHERE
    token = $1;

-- name: CountProjectCollaboratorByUserId :one
SELECT count(*)
FROM project_collaborator
WHERE
    user_id = $1
    AND
    project_id = $2
;

-- name: ConfirmInvitation :one
UPDATE project_invitation SET confirmed = true WHERE token = $1 RETURNING *;