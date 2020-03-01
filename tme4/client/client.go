package main

import (
"bufio"
"fmt"
"log"
"math/rand"
"os"
"regexp"
"strconv"
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

// Partie 2: contacté par les méthodes de personne_dist, le proxy appelle la méthode à travers le réseau et récupère le résultat
// il doit utiliser une connection TCP sur le port donné en ligne de commande
func proxy() {
// A FAIRE
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
func ouvrier( toCollector chan personne_int, uW chan requete) {
for{
//paquet := <- fromGest
	req := <- uW
	go func ( r requete ){
		switch r.personne.donne_statut() {
		case "V":
			r.personne.initialise()
			fmt.Println("i've come back from initializing the packet : " + r.personne.vers_string())
			r.retourW <- r.personne

			fmt.Println("i've sent the packet " + r.personne.vers_string())
		case "R":
			fmt.Println("i'm working the packet : " + r.personne.vers_string())
			r.personne.travaille()
			r.retourW <- r.personne

		case "C":
			fmt.Println("i'm finishing the packet : " + r.personne.vers_string())
			toCollector <- r.personne
		}
	}(req)



/*fmt.Println("i've received the packet : "+ paquet.vers_string())
switch paquet.donne_statut(){
case "V":
	paquet.initialise()
	fmt.Println("i've come back from initializing the packet : "+ paquet.vers_string())

	fmt.Println("potatoe : "+ paquet.vers_string())
	fromWorker <- paquet



	fmt.Println("i've sent the packet "+ paquet.vers_string())
case "R":
	fmt.Println("i'm working the packet : "+ paquet.vers_string())
	paquet.travaille()
//	ack <- 1
	fromWorker <-paquet

case "C":
	fmt.Println("i'm finishing the packet : "+ paquet.vers_string())
	toCollector <- paquet
//	ack <- 1

}*/
}

}

// Partie 1: les producteurs cree des personne_int implementees par des personne_emp initialement vides,
// de statut V mais contenant un numéro de ligne (pour etre initialisee depuis le fichier texte)
// la personne est passée aux gestionnaires
func producteur(lire chan message_lec, uP chan requete) {
for {
	req := <-uP
	go func ( r requete ){
		np := pers_vide
		nt := make([]func (st.Personne) st.Personne, 0)
		npe := personne_emp{statut: "V", ligne:rand.Intn(TAILLE_SOURCE), aFaire:nt, Personne :np, lecture:lire}
		r.retourP <- personne_int(&npe)
	}(req)

}
}

// Partie 2: les producteurs distants cree des personne_int implementees par des personne_dist qui contiennent un identifiant unique
// utilisé pour retrouver l'object sur le serveur
// la creation sur le client d'une personne_dist doit declencher la creation sur le serveur d'une "vraie" personne, initialement vide, de statut V
func producteur_distant(enfiler chan personne_int, port int, proxer chan message_proxy, frais chan int) {
for{
n:= <- frais
fmt.Println("Porducteur Distant crée identifiant", n)
np := personne_dist{identifiant: n, proxy: proxer}
local := make (chan string)
proxer <- message_proxy{identifiant:n , methode:"creer", retour: local}
<- local
enfiler <- np
}
}
type requete struct{
personne personne_int
retourW  chan personne_int
retourP chan personne_int
}

// Partie 1: les gestionnaires recoivent des personne_int des producteurs et des ouvriers et maintiennent chacun une file de personne_int
// ils les passent aux ouvriers quand ils sont disponibles
// ATTENTION: la famine des ouvriers doit être évitée: si les producteurs inondent les gestionnaires de paquets, les ouvrier ne pourront
// plus rendre les paquets surlesquels ils travaillent pour en prendre des autres
func gestionnaire(uP chan requete, uW chan requete){
var queue []personne_int
turnProd := 0
turnWork := 0

rP := make(chan personne_int)
rW := make(chan personne_int)


	npe := personne_emp{}

	uP <- requete{personne: &npe,retourW:rW, retourP:rP}
	for {
		select {
		case work := <-rW:
			if turnProd>turnWork {
				fmt.Println("worker")
					queue = append(queue, work)
				turnWork = turnWork+1
			}

		case prod := <-rP:
			if turnProd == turnWork {
				fmt.Println("prod")
				queue = append(queue, prod)
				turnProd = turnProd+1
			}
		}
		//fmt.Println("popo")
		if len(queue)>0 {
			toSend := queue[0]
			queue = queue[1:]
			uW <- requete{personne: toSend, retourW: rW, retourP: rP}
		}
		if len(queue)<TAILLE_G {
			npe := personne_emp{}
			uP <- requete{personne: &npe, retourW: rW, retourP: rP}
		}

		//toWorker <- toSend
//fmt.Println("queue "+ queue[0].vers_string())
//time.Sleep(time.Duration(5000) * time.Millisecond)
}
}

// Partie 1: le collecteur recoit des personne_int dont le statut est c, il les collecte dans un journal
// quand il recoit un signal de fin du temps, il imprime son journal.
func collecteur(toCollector chan personne_int, quit chan int) {
var journal string
for {
select{
	case p := <- toCollector:
		journal = journal + p.vers_string()+ " \n"
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
//port, _ := strconv.Atoi(os.Args[1]) // utile pour la partie 2
millis, _ := strconv.Atoi(os.Args[1]) // duree du timeout
fintemps := make(chan int)
// A FAIRE
// creer les canaux
/*	url := make (chan message_lec)
fromGest := make (chan personne_int)*/
toCollector := make (chan personne_int)
lire := make (chan message_lec)

	uP := make(chan requete)
	uW := make(chan requete)
/*	quit := make(chan int)*/

// lancer les goroutines (parties 1 et 2): 1 lecteur, 1 collecteur, des producteurs,
//des gestionnaires, des ouvriers
go func() { lecteur(lire)}()
go func() { collecteur(toCollector,fintemps)}()
for i := 0 ; i <NB_O ; i++ {
go func() { ouvrier(toCollector, uW ) }()
}
for i := 0 ; i <NB_P ; i++ {
go func() { producteur(lire, uP) }()
}
for i := 0 ; i <NB_G ; i++ {
go func() { gestionnaire(uP, uW) }()
}

fmt.Println(millis)
// lancer les goroutines (partie 2): des producteurs distants, un proxy
time.Sleep(time.Duration(millis) * time.Millisecond)
fmt.Println("je sors")
fintemps <- 0
<-fintemps
}
