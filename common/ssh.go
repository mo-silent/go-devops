package common

import (
	"golang.org/x/crypto/ssh"
	"os"
)

type SSH struct {
	Addr string
	User string
}

func (s *SSH) ExecuteWithPasswd(passwd, cmd string) ([]byte, error) {
	// ssh config
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connection ssh
	conn, err := ssh.Dial("tcp", s.Addr, config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// execute command
	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *SSH) ExecuteWithKeyFile(file, cmd string) ([]byte, error) {
	// SSH connection configuration
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			publicKeyFile(file),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略 HostKey 验证
	}

	// establish an SSH connection
	client, err := ssh.Dial("tcp", s.Addr, config)
	if err != nil {
		return nil, err
	}

	// create a new SSH session
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// execute the command and get the output
	output, err := session.Output(cmd)
	if err != nil {
		return nil, err
	}

	// return result
	return output, nil
}

// publicKeyFile Use SSH key files for authentication
func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}
