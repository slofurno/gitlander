## Usage

clone this repo or download a copy of gitlander.sh

find an existing ssh key or generate a new one. 

```
ssh-keygen -t rsa -C "your_email@example.com"
```

use the gitlander alpha client to register your account
    
```
./gitlander.sh register <username> <public key>
./gitlander.sh register slofurno ~/.ssh/mygit.pub
```

make an entry in your ssh config with the username you picked, and the location of your private key

```
Host gitlander
HostName gitlander.com
User slofurno
IdentityFile ~/.ssh/mygit
```

now that your account is setup, you can create as many repos as you want

```
./gitlander.sh init <username> <repo name>
./gitlander.sh init slofurno finaltest
```

Push or clone your new repo over ssh

```
git remote add origin gitlander:finaltest.git
git push -u origin master

git clone gitlander:finaltest.git
```

Publicly clone over http

```
git clone http://gitlander.com/slofurno/finaltest.git
```

view the contents of your repo on gitlander.com
