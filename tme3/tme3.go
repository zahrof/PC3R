package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const NB_TRAVAILLEURS = 10000

type paquet struct {
	arrivee string
	depart  string
	arret   int
}

func serveurDeCalcul(serveurCalChan chan paquet) {
	for {
		p := <-serveurCalChan
		fmt.Println("paquet: arrivee " + p.arrivee + " depart " + p.depart)
		adate, _ := time.Parse("15:04:05", p.arrivee)
		ddate, _ := time.Parse("15:04:05", p.depart)
		duree := int(ddate.Sub(adate).Seconds())
		paquetTransforme := paquet{arrivee: p.arrivee, depart: p.depart, arret: duree}
		serveurCalChan <- paquetTransforme

	}
}

type requete struct {
	arg    paquet
	retour chan paquet
}

func travailleur(lectureChan chan string, serveurCalChan chan requete, retourServeur chan requete, chanelRedacteur chan requete) {
	for i := 0; i < NB_TRAVAILLEURS; i++ {
		str := <-lectureChan // j'attends sur lectureChan
		go func(s string) {
			res := strings.Split(s, ",")
			paquet := paquet{arrivee: res[1], depart: res[2], arret: 0}
			serveurCalChan <- requete{arg: paquet, retour: make(chan paquet)}
			p := <-retourServeur // struct sera la requete?
			chanelRedacteur <- p
		}(str)
	}
}

func lecteur(lignePourTravailleurs chan string) {
	file, err := os.Open("/home/zahrof/Documents/master/semestre2/pc3r/export_gtfs_voyages/test.txt")
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

func redacteur(chanelRedacteur chan paquet) {

	// faire un select pour voir s'il doit arrêter ou continuer
	compteur := 0
	for {
		p := <-chanelRedacteur
		compteur = p.arret + compteur
		//fmt.Println("compteur " + compteur)
	}

}

func main() {
	//argTime, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	//fmt.Println(argTime)
	// }
	ligne := make(chan string)
	paquets := make(chan requete)
	paquetsToServer := make(chan requete)
	paquetsToRedactor := make(chan requete)
	go func() { lecteur(ligne) }()
	go func() { travailleur(ligne, paquets, paquetsToServer, paquetsToRedactor) }()
	time.Sleep(3000 * time.Second)

}
