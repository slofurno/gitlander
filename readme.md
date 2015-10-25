## Usage

clone this repo or download a copy of gitlander.sh

find an existing ssh key or generate a new one. 

```
ssh-keygen -t rsa -C "your_email@example.com"
```

use the gitlander alpha client to register your account
    
```
./gitlander.sh register <username> <~/path/to/ssh/key.pub>
./gitlander.sh register slofurno ~/.ssh/mygit.pub
```

make an entry in your ssh config with the username you picked, and the location of your private key

```
Host gitlander
HostName gitlander.com
User slofurno
IdentityFile ~/.ssh/mygit
```

now that your account is setup, you can init as many bare repos as you want

```
./gitlander.sh init <your username> <your repo name>
./gitlander.sh init slofurno finaltest
```

clone your repo locally

```
git clone <username>@gitlander:<repo>.git
git clone slofurno@gitlander:finaltest.git
```

add something and make your first commit

view your repo on gitlander.com