package common

import (
	"reflect"
	"testing"
)

func TestSSH_ExecuteWithKeyFile(t *testing.T) {
	type fields struct {
		Addr string
		User string
	}
	type args struct {
		file string
		cmd  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test2",
			fields: fields{
				Addr: "10.12.41.53:22",
				User: "ec2-user",
			},
			args: args{
				file: "E:\\ssh\\devops-ssh-key.key",
				cmd:  "pwd",
			},
			want: []byte{47, 104, 111, 109, 101, 47, 101, 99, 50, 45, 117, 115, 101, 114, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SSH{
				Addr: tt.fields.Addr,
				User: tt.fields.User,
			}
			got, err := s.ExecuteWithKeyFile(tt.args.file, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteWithKeyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecuteWithKeyFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSSH_ExecuteWithPasswd(t *testing.T) {
	type fields struct {
		Addr string
		User string
	}
	type args struct {
		passwd string
		cmd    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Addr: "10.65.2.57:22",
				User: "root",
			},
			args: args{
				passwd: "123456",
				cmd:    "pwd",
			},
			want: []byte{47, 114, 111, 111, 116, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SSH{
				Addr: tt.fields.Addr,
				User: tt.fields.User,
			}
			got, err := s.ExecuteWithPasswd(tt.args.passwd, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteWithPasswd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecuteWithPasswd() got = %v, want %v", got, tt.want)
			}
		})
	}
}
