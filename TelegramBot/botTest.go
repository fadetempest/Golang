package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bot-api/telegram"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

const token = "######"
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "######"
	dbname = "avito_db"
)

func main(){
	getInformation()
}

func getInformation(){
	urls:=[3]string{"https://www.avito.ru/buzuluk/kvartiry/prodam-ASgBAgICAUSSA8YQ?cd=1&f=ASgBAQICAUSSA8YQAUCQvg0Ulq41&s=104","https://www.avito.ru/buzuluk/doma_dachi_kottedzhi/prodam/dom-ASgBAQICAUSUA9AQAUDYCBTOWQ?cd=1&s=104&user=1","https://www.avito.ru/buzuluk/zemelnye_uchastki/prodam-ASgBAgICAUSWA9oQ?cd=1&f=ASgBAgECAUSWA9oQAUW4Exh7ImZyb20iOm51bGwsInRvIjoxNDQ1OX0&s=104&user=1"}
	client:= http.Client{Timeout: 15*time.Second}
	for _, url:=range urls{
		req,err:=client.Get(url)
		if err != nil{
			log.Fatal("Error while sending the request")
		}
		if req.StatusCode != 200{
			log.Fatalf("Wrong status code: %d, %s", req.StatusCode, req.Status)
		}

		doc, err:= goquery.NewDocumentFromReader(req.Body)
		if err != nil{
			log.Fatal("Reading error")
		}
		doc.Find(".iva-item-content-UnQQ4").Each(findAll)
	}
}

func findAll(i int, s *goquery.Selection){
	link, _ := s.Find("a").Attr("href")
	if checkLink(link){
		sendMess(link)
	}
}

func sendMess(link string){
	api:=telegram.New(token)
	editedLink:="Новое объявление!\n\n"+"www.avito.ru"+link+"\n"
	msg:= telegram.NewMessage(-306908388,editedLink)
	ctx:= context.Background()
	_, err:= api.SendMessage(ctx, msg)
	if err!= nil{
		log.Fatal("Sending message error")
	}
	time.Sleep(10*time.Second)
}

func checkLink(link string) bool{
	sqcon:=fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host,port,user,password,dbname)
	db, err:= sql.Open("postgres",sqcon)
	if err != nil{
		log.Fatal("Error while open DB")
	}

	defer db.Close()

	e:= db.Ping()
	if e != nil{
		log.Fatal(e)
	}

	var check bool = false

	searchQuery:= `SELECT link FROM links_avito WHERE link=$1`
	if db.QueryRow(searchQuery, link).Scan(&link) == sql.ErrNoRows {
		insertValue := `INSERT INTO links_avito (link) VALUES ($1)`
		_, er := db.Exec(insertValue, link)
		if er != nil {
			log.Fatal("Error while inserting value into DB")
		}
		check = true
	} 
	
	return check
}
