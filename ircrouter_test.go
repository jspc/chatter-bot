package main

import (
    "testing"
)

func TestRouter_normaliseCommand(t *testing.T) {
    type fields struct {
        Channel string
        Nick    string
        User    string
    }
    type args struct {
        cmd string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   string
    }{
        {"public channel, sent to user", fields{"#test-channel", "test-nick", "control-user"}, args{"test-nick: command"}, "command"},
        {"private chat, sent with username", fields{"control-user", "test-nick", "control-user"}, args{"test-nick: command"}, "command"},
        {"private chat, sent without username", fields{"control-user", "test-nick", "control-user"}, args{"command"}, "command"},
        {"public channel, sent to user, extra space", fields{"control-user", "test-nick", "control-user"}, args{"test-nick:  command"}, "command"},
        {"public channel, sent to user, not enough space", fields{"control-user", "test-nick", "control-user"}, args{"test-nick:command"}, "command"},
    }
    for _, tt := range tests {
        r := &Router{
            Channel: tt.fields.Channel,
            Nick:    tt.fields.Nick,
            User:    tt.fields.User,
        }
        if got := r.normaliseCommand(tt.args.cmd); got != tt.want {
            t.Errorf("%q. Router.normaliseCommand() = %v, want %v", tt.name, got, tt.want)
        }  else {
            t.Logf("%q. Router.isValid() => success", tt.name)
        }
    }
}

func TestRouter_isValid(t *testing.T) {
    type fields struct {
        Channel string
        Nick    string
        User    string
    }
    type args struct {
        requestor string
        cmd       string
    }

    tests := []struct {
        name   string
        fields fields
        args   args
        want   bool
    }{
        {"safe command, invalid user", fields{"#test-channel", "test-nick", "control-user"}, args{"a-user", "uptime"}, true},
        {"safe command, valid user", fields{"#test-channel", "test-nick", "control-user"}, args{"control-user", "uptime"}, true},

        {"unsafe command, invalid user", fields{"#test-channel", "test-nick", "control-user"}, args{"a-user", "some-command"}, false},
        {"unsafe command, valid user", fields{"#test-channel", "test-nick", "control-user"}, args{"control-user", "some-command"}, true},
    }
    for _, tt := range tests {
        r := &Router{
            Channel: tt.fields.Channel,
            Nick:    tt.fields.Nick,
            User:    tt.fields.User,
        }
        if got := r.isValid(tt.args.requestor, tt.args.cmd); got != tt.want {
            t.Errorf("%q. Router.isValid() = %v, want %v", tt.name, got, tt.want)
        } else {
            t.Logf("%q. Router.isValid() => success", tt.name)
        }
    }
}
