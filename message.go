// Copyright (c) 2016 Tristan Colgate-McFarlane
//
// This file is part of hugot.
//
// hugot is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// hugot is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with hugot.  If not, see <http://www.gnu.org/licenses/>.

package hugot

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mattn/go-shellwords"
	"github.com/nlopes/slack"
)

type Attachment slack.Attachment

type Message struct {
	Adapter Adapter // The to use for sending

	Event   *slack.MessageEvent
	From    string
	Channel string

	Text        string
	Attachments []Attachment

	Private bool
	ToBot   bool

	*flag.FlagSet
}

var ErrBadCLI = errors.New("coul not process as command line")

func (m *Message) Reply(txt string) *Message {
	out := *m
	out.Text = txt

	if !m.Private && m.ToBot {
		out.Text = fmt.Sprintf("@%s: %s", m.From, txt)
	}

	out.Event = nil
	out.From = ""

	return &out
}

func (m *Message) Replyf(s string, is ...interface{}) *Message {
	return m.Reply(fmt.Sprintf(s, is...))
}

func (m *Message) Parse() error {
	args, err := shellwords.Parse(m.Text)
	if err != nil {
		return err
	}

	return m.FlagSet.Parse(args[1:])
}
