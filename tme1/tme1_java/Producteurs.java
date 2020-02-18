
public class Producteurs implements Runnable{
	String nom_produit; 
	int cible_production; 
	int production_actuel; 
	Tapis tapis; 

	Producteurs(String nom_produit, int cible_production, Tapis tapis){
		this.nom_produit=nom_produit; 
		this.cible_production=cible_production; 
		this.production_actuel=0; 
		this.tapis=tapis; 
	}

	@Override
	public void run() {
		while(this.cible_production!=this.production_actuel) {
			Paquet p = new Paquet(this.nom_produit+" "+this.production_actuel); 
			try {
				this.tapis.enfiler(p);
			} catch (InterruptedException e) {
				System.out.print("Problème lors de l'enfilage"); 
				e.printStackTrace();
			}
			this.production_actuel++; 
		}

	}
}

