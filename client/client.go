package client

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type SSHProxmoxClient struct {
	client *ssh.Client
}

func NewSSHProxmoxClient(user, password, address string) (*SSHProxmoxClient, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		return nil, err
	}

	return &SSHProxmoxClient{client: client}, nil
}

func (c *SSHProxmoxClient) RunCommand(cmd string) (string, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	err = session.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w\nStderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func (c *SSHProxmoxClient) Close() {
	c.client.Close()
}
