package imp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	SSH_USER = "root"
	SSH_PASS = "kevin@0311"
	SSH_PORT = 22
)

type SSHResult struct {
	Host    string
	Success bool
	Result  string
}

func SshConn(user, password, host, key string, port int, cipherList []string) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		session      *ssh.Session
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	if key == "" {
		auth = append(auth, ssh.Password(password))
	} else {
		pemBytes, err := ioutil.ReadFile(key)
		if err != nil {
			return nil, err
		}

		var signer ssh.Signer
		if password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(password))
		}
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}

func PsshRun(commandName string, sshHosts []string) {
	var sshResult SSHResult
	ciphers := []string{}
	/*
		for _, host := range sshHosts {
			sshResult = SshRun(commandName, SSH_USER, SSH_PASS, host, "", SSH_PORT, ciphers)
			fmt.Println(sshResult)
		}
	*/
	ch := make(chan SSHResult, len(sshHosts))
	for _, host := range sshHosts {
		sshResult = SshRun(commandName, SSH_USER, SSH_PASS, host, "", SSH_PORT, ciphers)
		ch <- sshResult
	}
	close(ch)
	for sshResult = range ch {
		if sshResult.Result != "" {
			fmt.Println(sshResult.Host + ":")
			res := strings.Split(sshResult.Result, "\n")
			for n, line := range res {
				if (n + 1) < len(res) {
					fmt.Println(" ", line)
				}
			}
		}
	}
	return
}

func SshRun(commandName, user, password, host, key string, port int, ciphers []string) SSHResult {
	var cmdResult SSHResult
	cmdResult.Host = host
	session, err := SshConn(user, password, host, key, port, ciphers)
	if err != nil {
		cmdResult.Success = false
		cmdResult.Result = fmt.Sprintf("<%s>", err.Error())
		return cmdResult
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(commandName)
	cmdResult.Success = true
	cmdResult.Result = stdoutBuf.String()
	return cmdResult
}
