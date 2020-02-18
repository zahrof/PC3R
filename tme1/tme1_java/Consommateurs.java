
public 
class Consommateurs implements Runnable{
	int id; 
	Counter compteur; 
	Tapis tapis; 

	Consommateurs(int id, Counter compteur,  Tapis tapis){
		this.id=id; 
		this.compteur=compteur; 
		this.tapis =tapis; 
	}

	@Override
	public void run() {
		while(this.compteur.myCount>0) {
			try {
				Paquet p = this.tapis.defiler(this.compteur);
				System.out.println("C"+this.id+" mange "+p.getPaquet()); 
				// le décrement du compteur se fait à l'intérieur de défiler 
			} catch (InterruptedException e) {
				System.out.print("Problème lors du défilage"); 
				e.printStackTrace();
			} 
		}
	}
}