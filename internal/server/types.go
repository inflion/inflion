// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package server

type HasuraProjectAddedEventRoot struct {
	Id    string                  `json:"id"`
	Event HasuraProjectAddedEvent `json:"event"`
}

type HasuraProjectAddedEvent struct {
	SessionVariables                HasuraSessionVariables `json:"session_variables"`
	HasuraProjectAddedEventDataRoot `json:"data"`
}

type HasuraProjectAddedEventDataRoot struct {
	New HasuraProjectAddedEventData `json:"new"`
}

type HasuraProjectAddedEventData struct {
	Id     int64  `json:"id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

type HasuraInvitationEventRoot struct {
	Id    string                `json:"id"`
	Event HasuraInvitationEvent `json:"event"`
}

type HasuraInvitationEvent struct {
	SessionVariables                 HasuraSessionVariables `json:"session_variables"`
	HasuraProjectInvitationEventRoot `json:"data"`
}

type HasuraSessionVariables struct {
	XHasuraRole   string `json:"x-hasura-role"`
	XHasuraUserId string `json:"x-hasura-user-id"`
}

type HasuraProjectInvitationEventRoot struct {
	New HasuraProjectInvitationEventData `json:"new"`
}

type HasuraProjectInvitationEventData struct {
	MailAddress   string `json:"mail_address"`
	ProjectId     int    `json:"project_id"`
	InviteeUserId string `json:"invitee_user_id"`
	Token         string `json:"token"`
}

type ConfirmationActionRoot struct {
	SessionVariables  HasuraSessionVariables `json:"session_variables"`
	ConfirmationInput `json:"input"`
}

type ConfirmationInput struct {
	Input ConfirmationToken `json:"input"`
}

type ConfirmationToken struct {
	Token string `json:"token"`
}
