package imp

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"path"
	"time"
)

type SCPResult struct {
	Host    string
	Success bool
	Result  string
}

func SftpConn (user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func PscpCopy (localFile, remoteDir string, sshHosts []string) {
	var scpResult SCPResult
	/*
		for _, host := range sshHosts {
			scpResult = ScpCopy(localFile, remoteDir, host)
			fmt.Println(scpResult.Success)
		}
	*/
	ch := make(chan SCPResult, len(sshHosts))
	for  _, host := range sshHosts {
		scpResult = ScpCopy(localFile, remoteDir, host)
		ch <- scpResult
	}
	close(ch)
	for scpResult = range ch {
		if scpResult.Success {
			fmt.Println(scpResult.Host + ":", scpResult.Success)
		} else {
			fmt.Println(scpResult.Host + ":", scpResult.Result)
		}
	}
	/* */
	return
}

func ScpCopy(localFile, remoteDir, host string) SCPResult {
	var scpResult SCPResult
	sftpClient, err := SftpConn(SSH_USER, SSH_PASS, host, 22)
	scpResult.Host = host
	if err != nil {
		scpResult.Success = false
		scpResult.Result = fmt.Sprintf("<%s>", err.Error())
		return scpResult
	}
	defer sftpClient.Close()
	srcFile, err := os.Open(localFile)
	if err != nil {
		scpResult.Success = false
		scpResult.Result = fmt.Sprintf("<%s>", err.Error())
		return scpResult
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFile)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		scpResult.Success = false
		scpResult.Result = fmt.Sprintf("<%s>", err.Error())
		return scpResult
	}
	defer dstFile.Close()

	buf := make([]byte, 102400)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
	}
	scpResult.Success = true
	scpResult.Result = remoteFileName + ", done!"
	return scpResult
}