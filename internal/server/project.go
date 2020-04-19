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

import (
	"encoding/json"
	"github.com/inflion/inflion/internal/store"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type projectHandler struct {
	querier store.Querier
}

func newProjectHandler(querier store.Querier) projectHandler {
	return projectHandler{querier: querier}
}

func (i *projectHandler) added(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg HasuraProjectAddedEventRoot
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	params := msg.Event.HasuraProjectAddedEventDataRoot.New

	_, err = i.querier.AddToProject(r.Context(), store.AddToProjectParams{
		ProjectID: params.Id,
		UserID:    params.UserId,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(output)
}

func (i *projectHandler) sendInvitation(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg HasuraInvitationEventRoot
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sendInvitationMail(msg.Event.HasuraProjectInvitationEventRoot.New)

	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(output)
}

func (i *projectHandler) confirmInvitation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}

	var msg ConfirmationActionRoot
	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println(msg.Input.Token)

	invitation, err := i.querier.GetInvitationByToken(r.Context(), msg.Input.Token)
	if err != nil {
		log.Println(err)
		output := []byte("{\"result\": false}")
		_, _ = w.Write(output)
		return
	}

	_, err = i.querier.AddToProject(r.Context(), store.AddToProjectParams{
		UserID:    msg.SessionVariables.XHasuraUserId,
		ProjectID: invitation.ProjectID,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		output := []byte("{\"result\": false}")
		_, _ = w.Write(output)
		return
	}

	_, _ = i.querier.ConfirmInvitation(r.Context(), msg.Input.Token)

	output := []byte("{\"result\": true}")
	_, _ = w.Write(output)
}

func sendInvitationMail(event HasuraProjectInvitationEventData) {
	var url string

	server := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	domain := os.Getenv("DOMAIN")

	if domain == "localhost" {
		port := os.Getenv("WEB_EXTERNAL_PORT")
		url = "http://" + domain + ":" + port
	} else {
		url = "https://" + domain
	}

	msg := []byte("To: " + event.MailAddress + "\r\n" +
		"Subject: Inflion invitation\r\n" +
		"\r\n" +
		"This is the invitation mail from the inflion.\r\n" +
		"\r\n" +
		url + "/projectHandler/invitation/confirm/" + event.Token + "\r\n")

	err := smtp.SendMail(
		server,
		nil,
		"system@inflion.com",
		[]string{event.MailAddress},
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}
}
