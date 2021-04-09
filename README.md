# Install nset-cli

## Install Golang
`rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO`<br />
`curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo`<br />
`yum install golang`

## Golang Env
`mkdir -p ~/go/src`<br />
`echo "export GO111MODULE=off GOPATH=~/go" >> ~/.bash_profile`<br />
`source ~/.bash_profile`

## Install nset-cli
`go get github.com/spf13/cobra`<br />
`go get golang.org/x/crypto/ssh`<br />
`go get github.com/pkg/sftp`<br />
`go get gopkg.in/yaml.v2`<br />

`cd ~/go/src`<br />
`git clone https://github.com/kevinu2/nset-cli`<br />
`cd nset-cli`<br />
`make`<br />
`make install`
