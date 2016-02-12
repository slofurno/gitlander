## Usage

Download the gitlander cli

```
curl -L https://raw.githubusercontent.com/slofurno/gitlander/master/gitlander.sh -O
chmod +x gitlander.sh
```

Find an existing ssh key or generate a new one. 

```
ssh-keygen -t rsa -C "your_email@example.com"
```

Create an account with your public key
    
```
./gitlander.sh register <username> <public key>
./gitlander.sh register slofurno ~/.ssh/mygit.pub
```

Add gitlander to your ssh config with your username and private key 

```
Host gitlander
HostName gitlander.com
User slofurno
IdentityFile ~/.ssh/mygit
```

Create your first gitlander repo 

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

View the contents of your repo on gitlander.com
