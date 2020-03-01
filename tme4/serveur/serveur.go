package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	st "../client/structures"
	tr "./travaux"
)

var m = make(map[int]personne_serv)
var ADRESSE = "localhost"

var pers_vide = st.Personne{Nom: "", Prenom: "", Age: 0, Sexe: "M"}

// type d'un paquet de personne stocke sur le serveur, n'implemente pas forcement personne_int (qui n'existe pas ici)
type personne_serv struct {
	statut string
	aFaire []func(st.Personne) st.Personne
	st.Personne
}

// cree une nouvelle personne_serv, est appelé depuis le client, par le proxy, au moment ou un producteur distant
// produit une personne_dist
func creer(id int) *personne_serv{
	fmt.Println("dans créer")
	np :=  pers_vide
	nt := make([]func (st.Personne) st.Personne, 0)
	npe := personne_serv{statut: "V", aFaire:nt, Personne :np}
	fmt.Println("je sors")
	fmt.Println(&npe)
	return &npe
}

// Méthodes sur les personne_serv, on peut recopier des méthodes des personne_emp du client
// l'initialisation peut être fait de maniere plus simple que sur le client
// (par exemple en initialisant toujours à la meme personne plutôt qu'en lisant un fichier)
func (p *personne_serv) initialise() {

	p.Personne = st.Personne{Nom: "Dupond", Prenom: "Paul", Sexe: "M", Age: 22}

	for i := 0; i < rand.Intn(6)+1; i++{
		p.aFaire= append(p.aFaire,tr.UnTravail())
	}
	p.statut="R"
}

func (p *personne_serv) travaille() {
	p.Personne = p.aFaire[0](p.Personne)
	p.aFaire = p.aFaire[1:]
	if len(p.aFaire) == 0{
		p.statut = "C"
	}
}

func (p *personne_serv) vers_string() string {
	var add string
	if p.Sexe == "F" {
		add = "Madame "
	}else {
		add = "Monsieur"
	}
	return fmt.Sprint(add, p.Prenom, " ",p.Nom, " : ", p.Age, " ans. ")
}

func (p *personne_serv) donne_statut() string {
	return p.statut
}

// Goroutine qui maintient une table d'association entre identifiant et personne_serv
// il est contacté par les goroutine de gestion avec un nom de methode et un identifiant
// et il appelle la méthode correspondante de la personne_serv correspondante
func mainteneur(nom_methode string, id int, res chan string){
	fmt.Println(nom_methode)
	ps := m[id]
	switch nom_methode {
	case "creer":
		fmt.Println("holaaaa")
		m[id] = *creer(id)
		fmt.Println(m[id])
		res <- "OK"
	case "initialise":
		ps.initialise()
		res <-  "OK"
	case "travaile":
		ps.travaille()
		res <-  "OK"
	case "vers_string":
		res <- ps.vers_string()
	case "donne_statut":
		res <- ps.donne_statut()
	}

}

// Goroutine de gestion des connections
// elle attend sur la socketi un message content un nom de methode et un identifiant et
//appelle le mainteneur avec ces arguments
// elle recupere le resultat du mainteneur et l'envoie sur la socket, puis ferme la socket
func gere_connection(conn net.Conn) {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil{
			fmt.Println(err)
			return
		}
		message=strings.TrimSuffix(message, "\n")
		fmt.Println(message)
		res := strings.Split(message, ",")
		id, _ := strconv.Atoi(res[0])
		resChan := make(chan string)
		go func(methode string, ident int, res chan string){mainteneur(methode,ident, res )}(res[1],id,resChan)
		messageRes := <- resChan
		fmt.Println(messageRes)
		conn.Write([]byte(messageRes+"\n"))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Format: client <port>")
		return
	}
	port, _ := strconv.Atoi(os.Args[1]) // doit être le meme port que le client
	addr := ADRESSE + ":" + fmt.Sprint(port)
	// A FAIRE: creer les canaux necessaires, lancer un mainteneur
	ln, _ := net.Listen("tcp", addr) // ecoute sur l'internet electronique
	fmt.Println("Ecoute sur", addr)
	for {

		conn, err := ln.Accept() // recoit une connection, cree une socket
		if err != nil{
			fmt.Println("mauvais port !! ")
			return
		}
		fmt.Println("Accepte une connection.")
		go gere_connection(conn) // passe la connection a une routine de gestion des connections

	}

}
