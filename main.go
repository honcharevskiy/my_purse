package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

const db_file string = "./foo.db"
const create_db string = "create table if not exists balances (id integer not null primary key AUTOINCREMENT, created_time int64, response text);"

var ErrIDNotFound = fmt.Errorf("ID not found")

type Activities struct {
	db *sql.DB
}

type LastResponse struct {
	id         int64
	saved_time int64
	response   string
}

func main() {

	activities, err := InitActivities()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	http.HandleFunc("/monobank/personal", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		result, err := databaseOrAPIResponse(activities)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(result)

	})
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("secrets.env")

	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}

	return os.Getenv(key)

}

func databaseOrAPIResponse(activities *Activities) (clientInfo, error) {
	provider := monobankProvider{
		apiToken: goDotEnvVariable("MONOBANK_TOKEN"),
		client:   http.Client{},
	}

	last_response, err := activities.LastResponse()
	if err == ErrIDNotFound {
		log.Println("Try to get information from API")
		personalInformation, err := provider.clientInfo()
		if err != nil {
			log.Fatal(err)
		}
		savedId, err := activities.Insert(personalInformation)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response saved: %d", savedId)
		return personalInformation, nil

	} else {
		result := clientInfo{}
		err := json.Unmarshal([]byte(last_response.response), &result)
		if err != nil {
			return clientInfo{}, err
		}
		return result, nil
	}
}

func InitActivities() (*Activities, error) {
	log.Println("Init database")
	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create_db); err != nil {
		return nil, err
	}
	return &Activities{
		db: db,
	}, nil
}

func (c *Activities) Insert(bankInfo clientInfo) (int, error) {
	serializedResponse, err := json.Marshal(bankInfo)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	now := time.Now().Unix()
	res, err := c.db.Exec("INSERT INTO balances VALUES(NULL,?,?);", now, serializedResponse)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (c *Activities) LastResponse() (LastResponse, error) {
	row := c.db.QueryRow("select id, created_time, response from balances order by created_time desc limit 1")

	last_response := LastResponse{}
	var err error
	if err = row.Scan(&last_response.id, &last_response.saved_time, &last_response.response); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return LastResponse{}, ErrIDNotFound
	}

	return last_response, err
}
