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

package main

import (
	"flag"

	"golang.org/x/net/context"

	// Add some handlers
	"github.com/golang/glog"
	"github.com/tcolgate/hugot"
	"github.com/tcolgate/hugot/adapters/irc"
	"github.com/tcolgate/hugot/handlers/ping"
	"github.com/tcolgate/hugot/handlers/tableflip"
	irce "github.com/thoj/go-ircevent"
)

var (
	nick    = flag.String("nick", "hugot-test", "Bot nick")
	user    = flag.String("irc.user", "hugot-test", "IRC username")
	pass    = flag.String("irc.pass", "", "IRC password")
	server  = flag.String("irc.server", "chat.freenode.net:6697", "Server to connect to")
	ircchan = flag.String("irc.channel", "#hugottest", "Channel to listen in")
	useSSL  = flag.Bool("irc.usessl", true, "Use SSL to connect")
)

func main() {
	flag.Parse()

	c := irce.IRC(*nick, *user)
	if c == nil {
		glog.Fatal("could not create IRC event instance")
	}
	c.UseTLS = *useSSL
	c.Password = *pass

	err := c.Connect(*server)
	if err != nil {
		glog.Fatalf("could not connecto server", err)
	}
	c.Join(*ircchan)
	defer c.Quit()

	a, err := irc.New(c)
	if err != nil {
		glog.Fatal("could not create irc adapter")
	}

	hugot.Add(tableflip.New())
	hugot.Add(ping.New())

	go hugot.ListenAndServe(context.Background(), a, nil)

	c.Loop()

}
