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
	"github.com/spf13/viper"
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

	go runWithHttp()
	runWithHttps()
	
}

func runWithHttp(){
	mux := new(http.ServeMux)

	mux.HandleFunc("/get", getuser)
	mux.HandleFunc("/post", insertuser)
	mux.HandleFunc("/update", updateuser)
	mux.HandleFunc("/delete", deleteuser)

	var handler http.Handler = mux

	handler = middlewareCheckMethod(handler)



	server := &http.Server{
		Addr : ":"+viper.GetString("server.port"),
		Handler : handler,
	}

	server.ListenAndServe()
}

func runWithHttps(){
	mux := new(http.ServeMux)

	mux.HandleFunc("/get", getuser)
	mux.HandleFunc("/post", insertuser)
	mux.HandleFunc("/update", updateuser)
	mux.HandleFunc("/delete", deleteuser)

	var handler http.Handler = mux

	handler = middlewareCheckMethod(handler)




	http.ListenAndServeTLS(":"+viper.GetString("server.portTls"),"server.crt","server.key",handler)
}

func middlewareCheckMethod(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			checkMethod := (r.Method == "GET") || (r.Method == "POST")
			if !checkMethod{
				w.Write([]byte("Method Harus GET atau POST"))
			}

			next.ServeHTTP(w,r)
		})
}