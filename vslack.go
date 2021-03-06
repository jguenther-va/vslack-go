package vslack

import (
	"errors"
)

// Interface is a VSlack interface
//
//go:generate mockery --inpackage --name=Interface
type Interface interface {
	SetIncomingWebhookURI(h string) *VSlack
	SetChannel(c string) *VSlack
	SetUsername(u string) *VSlack
	SetIconEmoji(i string) *VSlack
	SetMessage(m string) *VSlack
	Send() error
	SendAsync(e chan error)
	SetAttachments(a ...Attachment) *VSlack
	SetLinkNames(linkNames bool) *VSlack
	validate() error
}

// VSlack a structure holding data for the slack message
type VSlack struct {
	IncomingWebhookURI string       `json:"-"`
	Message            string       `json:"text,omitempty"`
	Username           string       `json:"username"`
	IconEmoji          string       `json:"icon_emoji,omitempty"`
	Channel            string       `json:"channel"`
	Attachments        []Attachment `json:"attachments,omitempty"`
	LinkNames          bool         `json:"link_names,omitempty"`
}

// NewVSlack returns a new instance of VSlack
func NewVSlack(incomingwebHookURI string) *VSlack {
	return &VSlack{IncomingWebhookURI: incomingwebHookURI}
}

// NewVSlackAttachment returns an instance of a new VSlack attachment
func NewVSlackAttachment() Attachment {
	return Attachment{}
}

// SetIncomingwebHookURI sets the incoming web hook
func (v *VSlack) SetIncomingWebhookURI(h string) *VSlack {
	v.IncomingWebhookURI = h
	return v
}

// SetChannel sets the channel
func (v *VSlack) SetChannel(c string) *VSlack {
	v.Channel = c
	return v
}

// SetUsername sets the username
func (v *VSlack) SetUsername(u string) *VSlack {
	v.Username = u
	return v
}

// SetIconEmoji sets the emoji for the icon
func (v *VSlack) SetIconEmoji(i string) *VSlack {
	v.IconEmoji = i
	return v
}

// SetMessage sets the message
func (v *VSlack) SetMessage(m string) *VSlack {
	v.Message = m
	return v
}

// SetAttachments takes in attachments
func (v *VSlack) SetAttachments(a ...Attachment) *VSlack {
	v.Attachments = append(v.Attachments, a...)
	return v
}

// SetLinkNames takes in attachments
func (v *VSlack) SetLinkNames(linkNames bool) *VSlack {
	v.LinkNames = linkNames
	return v
}

// Send the message
func (v *VSlack) Send() error {
	if err := v.validate(); err != nil {
		return err
	}
	return v.send()
}

// SendAsync sends the message asynchronously
func (v *VSlack) SendAsync(e chan error) {
	if err := v.validate(); err != nil {
		e <- err
	}
	v.send()
}

func (v *VSlack) validate() error {
	if v.IncomingWebhookURI == "" {
		return errors.New("VSlack needs an incoming web hook")
	}
	if v.Message == "" && len(v.Attachments) == 0 {
		return errors.New("VSlack needs a message, or attachments")
	}
	return nil
}
