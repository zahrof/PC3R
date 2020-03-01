package travaux

import (
	"math/rand"

	st "../../client/structures"
)

// *** LISTES DE FONCTION DE TRAVAIL DE Personne DANS Personne DU SERVEUR ***
// Essayer de trouver des fonctions *différentes* de celles du client


func f1(p st.Personne) st.Personne {
	return  st.Personne{Age:4000, Nom:p.Nom, Prenom:p.Prenom+"TOOOOOOO", Sexe:p.Sexe}
}

func f2(p st.Personne) st.Personne {
	return  st.Personne{Age:p.Age, Nom:"EHY", Prenom:p.Prenom, Sexe:p.Sexe}
}

func f3(p st.Personne) st.Personne {
	return  st.Personne{Age:p.Age, Nom:p.Nom, Prenom:p.Prenom+"-TAAAA"+p.Prenom, Sexe:p.Sexe}
}

func f4(p st.Personne) st.Personne {
	return  st.Personne{Age:p.Age, Nom:p.Nom+"TACHE 4", Prenom:p.Prenom, Sexe:p.Sexe}
}

func UnTravail() func(st.Personne) st.Personne {
	tableau := make([]func(st.Personne) st.Personne, 0)
	tableau = append(tableau, func(p st.Personne) st.Personne { return f1(p) })
	tableau = append(tableau, func(p st.Personne) st.Personne { return f2(p) })
	tableau = append(tableau, func(p st.Personne) st.Personne { return f3(p) })
	tableau = append(tableau, func(p st.Personne) st.Personne { return f4(p) })
	i := rand.Intn(len(tableau))
	return tableau[i]
}
