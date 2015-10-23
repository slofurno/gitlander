1. first make your ssh key and copy the public key to where youd like to work

2. run the gitlander alpha client to register your account

```
./gitlander.sh register <username> <location of your public key>
eg ./gitlander.sh register slofurno mygit.pub
```
3. make an entry in your ssh config with the username you picked, and the location of your private key

```
Host gitlander
HostName gitlander.com
User slofurno
IdentityFile ~/.ssh/mygit
```
4. now that your account is setup, you can init as many bare repos as you want

```
./gitlander.sh init <your username> <your repo name>
eg ./gitlander.sh init slofurno finaltest
```
5. now clone your repo locally

```
git clone <username>@gitlander:<repo>.git
eg git clone slofurno@gitlander:finaltest.git
```

6. now add something and make your first commit

7. you can view your repo on gitlander.com