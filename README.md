# Install nset-cli

## Install Golang
CentOS
```
rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO
curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo
yum install golang
```
Debian
```
version=1.16.4
wget https://golang.org/dl/go${verison}.linux-amd64.tar.gz
mkdir -pv ~/go/{tmp,src}
tar zxf go${verison}.linux-amd64.tar.gz -C ~/go/tmp/
mv ~/go/tmp/go  ~/go/go${verison}
ln -s ~/go/go${verison} /opt/go/goroot
if [ `grep 'go/goroot/bin' ~/.bashrc -c` -eq 0 ]; then
  echo 'export PATH=$PATH:~/go/goroot/bin' | tee -a ~/.bashrc
  . ~/.bashrc
fi
```

## Golang Env
```
if [ `grep 'GO111MODULE' ~/.bashrc -c` -eq 0 ]; then
  echo "export GO111MODULE=off GOPATH=~/go" >> ~/.bash_profile
  . ~/.bashrc
fi
```

## Install nset-cli
```
go get github.com/spf13/cobra
go get golang.org/x/crypto/ssh
go get github.com/pkg/sftp
go get gopkg.in/yaml.v2

cd ~/go/src
git clone https://github.com/kevinu2/nset-cli
cd nset-cli
make
make install
```
