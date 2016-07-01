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

package handlers

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/golang/glog"
	"github.com/tcolgate/hugot"
)

func init() {
	hugot.AddCommandHandler(New())
}

type ping struct {
}

func New() hugot.CommandHandler {
	return &ping{}
}

func (*ping) Describe() (string, string) {
	return "ping", "replies Pong to any ping"
}

func (*ping) Command(ctx context.Context, w hugot.ResponseWriter, m *hugot.Message) error {
	glog.Info("Got ping ", *m)
	if err := m.Parse(); err != nil {
		return err
	}

	w.SetChannel("bottest")
	fmt.Fprintf(w, "PONG!")

	return nil
}
