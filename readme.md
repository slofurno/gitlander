1. find an existing ssh key or generate a new one. 

  ```
  ssh-keygen -t rsa -C "your_email@example.com"
  ```

2. use the gitlander alpha client to register your account
    
  ```
  ./gitlander.sh register <username> <~/path/to/ssh/key.pub>
  eg ./gitlander.sh register slofurno ~/.ssh/mygit.pub
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

5. clone your repo locally

  ```
  git clone <username>@gitlander:<repo>.git
  eg git clone slofurno@gitlander:finaltest.git
  ```

6. add something and make your first commit

7. view your repo on gitlander.com