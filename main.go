package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

type fileddo struct {
	Name     string
	Size     int64
	Modified string
	IsDir    bool
}

func initRepo(w http.ResponseWriter, req *http.Request) {

	qs := req.URL.Query()
	user := qs.Get("name")
	repo := qs.Get("repo")

	if user == "" || repo == "" {
		return
	}

	cmd := exec.Command("sh", "./gitlander-init", user, repo)
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
	}

}

func registerUser(w http.ResponseWriter, req *http.Request) {

	qs := req.URL.Query()
	user := qs.Get("name")

	if user == "" {
		return
	}

	publickey, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd := exec.Command("sh", "./gitlander-useradd", user, string(publickey))
	_, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
	}

}

func getDir(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	user := qs.Get("user")
	repo := qs.Get("repo")

	if user == "" || repo == "" {
		return
	}

	path := "gitlander/" + user + "/" + repo

	dir, _ := os.Open(path)
	defer dir.Close()

	contents, err := dir.Readdir(0)

	ddos := make([]fileddo, len(contents))

	for i, content := range contents {
		ddos[i] = fileddo{
			Name:     content.Name(),
			Size:     content.Size(),
			Modified: content.ModTime().String(),
			IsDir:    content.IsDir(),
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	j, err := json.Marshal(ddos)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(j)

}

func getRepos(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	user := qs.Get("user")

	if user == "" {
		return
	}
	fmt.Println(user)

	path := "gitlander/" + user
	fmt.Println(path)

	repo, err := os.Open(path)
	defer repo.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dirs, err := repo.Readdirnames(0)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	j, err := json.Marshal(dirs)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(j)

}

func getUsers(w http.ResponseWriter, req *http.Request) {
	dir, _ := os.Open("gitlander")
	defer dir.Close()

	dirs, err := dir.Readdirnames(0)

	if err != nil {
		fmt.Println(err.Error())
	}

	for d := range dirs {
		fmt.Println(d)
	}

	j, err := json.Marshal(dirs)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(j)
}

func main() {

	http.HandleFunc("/api/repo", initRepo)
	http.HandleFunc("/api/user", registerUser)
	http.HandleFunc("/api/dir", getDir)
	http.HandleFunc("/api/repos", getRepos)
	http.HandleFunc("/api/users", getUsers)

	http.Handle("/", http.FileServer(http.Dir("gitlander")))
	http.ListenAndServe(":80", nil)

}
