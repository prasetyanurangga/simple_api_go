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
	"crypto/tls"
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
	go runWithHttps2()
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

func runWithHttps2(){

	certPair1, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}

	tlsConfig := new(tls.Config)
	tlsConfig.NextProtos = []string{"http/1.1"}
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.PreferServerCipherSuites = true

	tlsConfig.Certificates = []tls.Certificate{
		certPair1, /** add other certificates here **/
	}
	tlsConfig.BuildNameToCertificate()

	tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
	tlsConfig.CurvePreferences = []tls.CurveID{
		tls.CurveP521,
		tls.CurveP384,
		tls.CurveP256,
	}
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	}

	mux := new(http.ServeMux)

	mux.HandleFunc("/get", getuser)
	mux.HandleFunc("/post", insertuser)
	mux.HandleFunc("/update", updateuser)
	mux.HandleFunc("/delete", deleteuser)

	var handler http.Handler = mux

	handler = middlewareCheckMethod(handler)

	server := &http.Server{
		Addr : ":9009",
		Handler : handler,
		TLSConfig : tlsConfig,
	}

	server.ListenAndServeTLS("","")

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