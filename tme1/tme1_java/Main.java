class Main 
{ 

	public static void main(String[] args) 
	{ 
		final String[] produits = {"pomme", "poire", "orange", "kiwi", "banane"};
		final int nbre_Producteurs = 5; 
		final int nbre_Consommateur = 4; 
		final int cible_production = 3;
		Counter compteur = new Counter(cible_production * nbre_Producteurs); 
		final int taille_tapis = 1000; 
		Tapis tapis = new Tapis(taille_tapis); 


		for (int i=0; i<nbre_Producteurs; i++) 
		{ 
			Thread object = new Thread(new Producteurs(produits[i], cible_production,tapis)); 
			object.start(); 
		} 

		for (int i=0; i<nbre_Consommateur; i++) 
		{ 
			Thread object = new Thread(new Consommateurs(i, compteur,tapis)); 
			object.start(); 
		} 


		while(compteur.myCount>0); 
	} 
} 
