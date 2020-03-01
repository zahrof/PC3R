package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	. "time"
)

const NB_TRAVAILLEURS = 2

type paquet struct {
	arrivee string
	depart  string
	arret   int // temps d'arrêt en secondes
}

func serveurDeCalcul(serveurCalChan chan requete) {
	for {
		req := <-serveurCalChan

		go func (r requete){
			adate, _ := time.Parse("15:04:05", r.arg.arrivee)
			ddate, _ := time.Parse("15:04:05", r.arg.depart)
			duree := int(ddate.Sub(adate).Seconds())
			r.arg.arret=duree
			res:=r.arg
			r.retour <-res
		}(req)
	/*	fmt.Println("paquet: arrivee " + p.arrivee + " depart " + p.depart)
		adate, _ := time.Parse("15:04:05", p.arrivee)
		ddate, _ := time.Parse("15:04:05", p.depart)
		duree := int(ddate.Sub(adate).Seconds())
		paquetTransforme := paquet{arrivee: p.arrivee, depart: p.depart, arret: duree}
		serveurCalChan <- paquetTransforme*/

	}
}

type requete struct {
	arg    paquet
	retour chan paquet
}

func travailleur(lignePourTravailleurs chan string, serveurCalChan chan requete, redacteurChan chan paquet  ){
//chan string, serveurCalChan chan requete, retourServeur chan requete, chanelRedacteur chan requete) {
	for i := 0; i < NB_TRAVAILLEURS; i++ {
		str := <-lignePourTravailleurs // j'attends sur lectureChan
		go func(s string) {
			res := strings.Split(s, ",")
			p := paquet{arrivee: res[1], depart: res[2], arret: 0}

			r:= make(chan paquet)
			serveurCalChan <- requete{arg: p, retour: r}
			paquetRes := <-r
			//fmt.Println("envoye:" ,p, " recu ", paquetRes)
		//	p := <-retourServeur // struct sera la requete?
			redacteurChan <- paquetRes
		}(str)
	}
}

func lecteur(lignePourTravailleurs chan string, file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		lignePourTravailleurs <- str
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func redacteur(redacteurChan chan paquet, quit chan int, mainChan chan string) {
	times := 0
	compteurArret := 0
	for {
		select{
			case p := <-redacteurChan:
				compteurArret = p.arret + compteurArret
				times++
			case <- quit:
				moy := compteurArret /times
				fmt.Println(strconv.Itoa(compteurArret))
				fmt.Println(strconv.Itoa(times))
				s2 := strconv.Itoa(moy)
				mainChan <- s2
				fmt.Println("quit")
				return
		}
	}

}

func main() {
	ligne := make(chan string)
	url := make(chan requete)
	redacteurChan := make(chan paquet)
	quit := make (chan int)
	mainChan := make (chan string)
	fmt.Println(" args"+ os.Args[2])
	go func() { lecteur(ligne,os.Args[2]) }()
	for i := 0 ; i <5 ; i++ {
		go func() { travailleur(ligne, url,redacteurChan) }()
	}
	go func() { serveurDeCalcul(url) }()
	go func() { redacteur(redacteurChan, quit, mainChan) }()

	i, _ := strconv.Atoi(os.Args[1])
	Sleep(Second*Duration(*flag.Int("time",i,"an int")))
	quit <- 1
	compteur := <- mainChan
	fmt.Println("FIN moyenne: "+ compteur)

}