package main

import(
	"net/http"
	"simple_api_go/utils"
	"simple_api_go/repository"
	"simple_api_go/model"
	"log"
	"fmt"
	"context"
	"encoding/json"
	"strconv"
)

const port string = "9000"

//GetUser.....
func getuser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		//ddd
		user, err := repository.GetAll(ctx)

		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseJSON(w, user, http.StatusOK)
		return
	}
}

//insert user
func insertuser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		if r.Header.Get("Content-Type") != "application/json"{
			http.Error(w,"Gunakan context-type aplicatopn json", http.StatusBadRequest)
		}

		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		var usr model.User

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil{
			log.Fatal(er)
			return
		}

		if er := repository.Insert(ctx, usr); er != nil{
			log.Fatal(er)
			return
		}

		res := map[string]string{
			"status" : "successfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

func updateuser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		if r.Header.Get("Content-Type") != "application/json"{
			http.Error(w,"Gunakan context-type aplicatopn json", http.StatusBadRequest)
		}

		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		var usr model.User

		if er := json.NewDecoder(r.Body).Decode(&usr); er != nil{
			log.Fatal(er)
			return
		}

		if er := repository.Update(ctx, usr); er != nil{
			log.Fatal(er)
			return
		}

		res := map[string]string{
			"status" : "successfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

func deleteuser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		if r.Header.Get("Content-Type") != "application/json"{
			http.Error(w,"Gunakan context-type aplicatopn json", http.StatusBadRequest)
		}

		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		var usr model.User

		id := r.URL.Query().Get("id")



		if id == ""{
			utils.ResponseJSON(w, "ID Tidak BOleh kosong", http.StatusBadRequest)
		}

		usr.ID, _ = strconv.Atoi(id)

		if err := repository.Delete(ctx, usr); err != nil{
			utils.ResponseJSON(w, "Error", http.StatusBadRequest)
		}



		res := map[string]string{
			"status" : "successfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

func main(){

	db, err := utils.MySQL()

	if err != nil{
		log.Fatal(err)
	}

	eb := db.Ping()
	if eb != nil{
		panic(eb.Error())
	}

	fmt.Println("Success")

	http.HandleFunc("/get", getuser)
	http.HandleFunc("/post", insertuser)
	http.HandleFunc("/update", updateuser)
	http.HandleFunc("/delete", deleteuser)

	server := &http.Server{
		Addr : ":"+port,
	}

	server.ListenAndServe()
}