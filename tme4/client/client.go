package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	st "./structures" // contient la structure Personne
	tr "./travaux"
)

var ADRESSE string = "localhost"                           // adresse de base pour la Partie 2

var FICHIER_SOURCE string = "client/conseillers-municipaux.txt" // fichier dans lequel piocher des personnes
var TAILLE_SOURCE int = 450000                             // inferieure au nombre de lignes du fichier, pour prendre une ligne au hasard
var TAILLE_G int = 5                                       // taille du tampon des gestionnaires
var NB_G int = 2                                           // nombre de gestionnaires
var NB_P int = 2                                           // nombre de producteurs
var NB_O int = 4                                           // nombre d'ouvriers
var NB_PD int = 2                                          // nombre de producteurs distants pour la Partie 2

var pers_vide = st.Personne{Nom: "", Prenom: "", Age: 0, Sexe: "M"} // une personne vide

type message_lec struct {
	contenu int
	retour chan string
}

// paquet de personne, sur lequel on peut travailler, implemente l'interface personne_int
type personne_emp struct {
	statut string
	ligne int
	aFaire []func(st.Personne) st.Personne
	st.Personne
	lecture chan message_lec
}

type message_proxy struct{
	retour chan string
	methode string
	identifiant int
}

// paquet de personne distante, pour la Partie 2, implemente l'interface personne_int
type personne_dist struct {
	identifiant int
	proxy chan message_proxy
}

// interface des personnes manipulees par les ouvriers, les
type personne_int interface {
	initialise()          // appelle sur une personne vide de statut V, remplit les champs de la personne et passe son statut à R
	travaille()           // appelle sur une personne de statut R, travaille une fois sur la personne et passe son statut à C s'il n'y a plus de travail a faire
	vers_string() string  // convertit la personne en string
	donne_statut() string // renvoie V, R ou C
}

// fabrique une personne à partir d'une ligne du fichier des conseillers municipaux
// à changer si un autre fichier est utilisé
func personne_de_ligne(l string) st.Personne {
	separateur := regexp.MustCompile("\u0009") // oui, les donnees sont separees par des tabulations ... merci la Republique Francaise
	separation := separateur.Split(l, -1)
	naiss, _ := time.Parse("2/1/2006", separation[7])
	a1, _, _ := time.Now().Date()
	a2, _, _ := naiss.Date()
	agec := a1 - a2
	return st.Personne{Nom: separation[4], Prenom: separation[5], Sexe: separation[6], Age: agec}
}

// *** METHODES DE L'INTERFACE personne_int POUR LES PAQUETS DE PERSONNES ***

func (p *personne_emp) initialise() {
	ret := make(chan string)
	p.lecture <- message_lec{contenu:p.ligne, retour:ret}
	ligne := <- ret

	p.Personne = personne_de_ligne(ligne)

	for i := 0; i < rand.Intn(6)+1; i++{
		p.aFaire= append(p.aFaire,tr.UnTravail())
	}
	p.statut="R"
}

func (p *personne_emp) travaille() {
	p.Personne = p.aFaire[0](p.Personne)
	p.aFaire = p.aFaire[1:]
	if len(p.aFaire) == 0{
		p.statut = "C"
	}
}

func (p *personne_emp) vers_string() string {
	var add string
	if p.Sexe == "F" {
		add = "Madame "
	}else {
		add = "Monsieur"
	}
	return fmt.Sprint(add, p.Prenom, " ",p.Nom, " : ", p.Age, " ans. ")
}

func (p *personne_emp) donne_statut() string {
	return p.statut
}
// *** METHODES DE L'INTERFACE personne_int POUR LES PAQUETS DE PERSONNES DISTANTES (PARTIE 2) ***
// ces méthodes doivent appeler le proxy (aucun calcul direct)

func (p personne_dist) initialise() {
	local := make(chan string)
	mess := message_proxy{methode:"initialise", retour:local, identifiant:p.identifiant}
	p.proxy <- mess
	<-local
}

func (p personne_dist) travaille() {
	local := make (chan string)
	mess := message_proxy{methode:"travaille", retour: local , identifiant:p.identifiant}
	p.proxy <- mess
	<- local
}

func (p personne_dist) vers_string() string {
	local := make(chan string)
	mess := message_proxy{methode:"vers_string", retour: local, identifiant:p.identifiant}
	p.proxy <- mess
	return <- local
}

func (p personne_dist) donne_statut() string {
	local := make(chan string)
	mess := message_proxy{methode: "donne_statut", retour: local, identifiant:p.identifiant}
	p.proxy <- mess
	s:= <- local
	return s
}


// *** CODE DES GOROUTINES DU SYSTEME ***

// Partie 2: contacté par les méthodes de personne_dist, le proxy appelle la méthode
//à travers le réseau et récupère le résultat
// il doit utiliser une connection TCP sur le port donné en ligne de commande
func proxy(port string, u chan message_proxy) {
	add := ADRESSE + ":"+ port
	conn, _ := net.Dial("tcp", add)

	for {
		m_proxy := <- u
		requete := strconv.Itoa(m_proxy.identifiant)+ ","+ m_proxy.methode+ "\n"
		fmt.Fprint(conn, fmt.Sprintf(requete))
		reponse,_ := bufio.NewReader(conn).ReadString('\n')
		reponse = strings.TrimSuffix(reponse, "\n")
		fmt.Println("j'ai reç ma m_proxy de  de : "+ m_proxy.methode)
		fmt.Println("réponse "+ reponse)
		m_proxy.retour <- reponse
	}
	conn.Close()
}

// Partie 1 : contacté par la méthode initialise() de personne_emp, récupère une ligne donnée dans le fichier source
func lecteur(url chan message_lec) {
	for{
		m := <-url
		fmt.Println("Lecteur contacté pour ligne", m.contenu)
		file, err := os.Open(FICHIER_SOURCE)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		_ = scanner.Scan()
		for i := 0; i < m.contenu; i++ {
			_ = scanner.Scan()
		}
		resultat := scanner.Scan()
		if resultat == false {
			log.Fatal(err)
		}else{
			m.retour <- scanner.Text()
		}
		file.Close()

	}

}

// Partie 1: récupèrent des personne_int depuis les gestionnaires, font une opération dépendant de donne_statut()
// Si le statut est V, ils initialise le paquet de personne puis le repasse aux gestionnaires
// Si le statut est R, ils travaille une fois sur le paquet puis le repasse aux gestionnaires
// Si le statut est C, ils passent le paquet au collecteur
func ouvrier( fromGest chan personne_int, toGest chan personne_int, toCollector chan personne_int) {
	for{
		paquet := <- fromGest
		switch paquet.donne_statut() {
		case "V":
			paquet.initialise()
			toGest <- paquet
		case "R":
			paquet.travaille()
			toGest <- paquet
		case "C":
			toCollector <- paquet
		}
	}

}

// Partie 1: les producteurs cree des personne_int implementees par des personne_emp initialement vides,
// de statut V mais contenant un numéro de ligne (pour etre initialisee depuis le fichier texte)
// la personne est passée aux gestionnaires
func producteur(lire chan message_lec, prodChan chan personne_int) {
	for {
		nt := make([]func (st.Personne) st.Personne, 0)
		npe := personne_emp{statut: "V", ligne:rand.Intn(TAILLE_SOURCE), aFaire:nt, Personne :pers_vide, lecture:lire}
		prodChan <- personne_int(&npe)

	}
}

// Partie 2: les producteurs distants cree des personne_int implementees par des personne_dist qui contiennent un identifiant unique
// utilisé pour retrouver l'object sur le serveur
// la creation sur le client d'une personne_dist doit declencher la creation sur le serveur d'une "vraie" personne, initialement vide, de statut V
func producteur_distant(prodChan chan personne_int,  proxer chan message_proxy,frais chan int) {
	for{
		n := <- frais
		np := personne_dist{identifiant: n, proxy: proxer}
		local := make (chan string)
		proxer <- message_proxy{identifiant:n , methode:"creer", retour: local}
		<- local
		prodChan <- np
	}
}

// Partie 1: les gestionnaires recoivent des personne_int des producteurs et des ouvriers et maintiennent chacun une file de personne_int
// ils les passent aux ouvriers quand ils sont disponibles
// ATTENTION: la famine des ouvriers doit être évitée: si les producteurs inondent les gestionnaires de paquets, les ouvrier ne pourront
// plus rendre les paquets surlesquels ils travaillent pour en prendre des autres
func gestionnaire(fromProducer chan personne_int, toOuvr chan personne_int, fromOuv chan personne_int){
	queue := make([]personne_int, 0)
	for {
		switch len(queue) {
		case TAILLE_G:
			toOuvr <- queue[0]
			queue = queue[1:]
		case 0:
			select {
			case np := <-fromOuv:
				queue = append(queue, np)
			case np := <-fromProducer:
				queue = append(queue, np)
			}

		default:
			if len(queue) < TAILLE_G/2 {
				select {
				case np := <-fromOuv:
					queue = append(queue, np)
				case np := <-fromProducer:
					queue = append(queue, np)
				case toOuvr <- queue[0]:
					queue = queue[1:]
				}
			} else {
				select {
				case np := <-fromOuv:
					queue = append(queue, np)
				case toOuvr <- queue[0]:
					queue = queue[1:]
				}
			}
		}
	}
}

// Partie 1: le collecteur recoit des personne_int dont le statut est c, il les collecte dans un journal
// quand il recoit un signal de fin du temps, il imprime son journal.
func collecteur(toCollector chan personne_int, quit chan int) {
	var journal string
	for {
		select{
		case p := <- toCollector:
			journal +=  p.vers_string()+ " \n"
		case <- quit:
			fmt.Println("Collecteur: J'ai réçu le signal d'arrêt")
			fmt.Println("Journal: "+ journal)
			quit <- 0
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // graine pour l'aleatoire
	if len(os.Args) < 2 {
		fmt.Println("Format: client <port> <millisecondes d'attente>")
		return
	}
	millis, _ := strconv.Atoi(os.Args[2]) // duree du timeout
	fintemps := make(chan int)
	prodGest := make(chan personne_int)
	gestOuvr := make(chan personne_int)
	gestOuvr2 := make(chan personne_int)
	ouvrCollect := make(chan personne_int)
	lire := make(chan message_lec)
	proxer := make(chan message_proxy)
	frais := make(chan int)

	// lancer les goroutines (parties 1 et 2): 1 lecteur, 1 collecteur, des producteurs,
	//des gestionnaires, des ouvriers
	go func() { lecteur(lire)}()
	go func() { collecteur(ouvrCollect,fintemps)}()
	for i := 0 ; i <NB_O ; i++ {
		go func() { ouvrier(gestOuvr,gestOuvr2, ouvrCollect) }()
	}
	for i := 0 ; i <NB_PD ; i++ {
		go func() { producteur(lire, prodGest) }()
	}
	for i := 0 ; i <NB_G ; i++ {
		go func() { gestionnaire(prodGest,gestOuvr, gestOuvr2) }()
	}

	fmt.Println(millis)
	// lancer les goroutines (partie 2): des producteurs distants, un proxy
	go func(port string) { proxy(port, proxer) }(os.Args[1])
	go func(frais chan int ){ generator(frais)}(frais)

	for i := 0 ; i <NB_PD ; i++ {
			go func() {producteur_distant(prodGest, proxer,frais) }()
	}
	time.Sleep(time.Duration(millis) * time.Millisecond)
	fintemps <- 0
	<-fintemps
}

func generator(frais chan int) {
	compteur := 0
	for{
		frais <- compteur
		compteur++
	}
}
