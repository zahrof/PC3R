public class Tapis {
	int sz; 
	int begin; 
	int sizeMax; 
	Paquet[] tab; 

	Tapis(int sizeMax){
		this.sizeMax=sizeMax;
		this.sz=0; 
		this.begin=0; 
		this.tab= new Paquet[sizeMax]; 
	}

	public  int getSz() {
		return sz;
	}

	public boolean empty() {
		return this.sz==0; 
	}

	public boolean full() {
		return this.sz==this.sizeMax; 
	}

	public Paquet defiler(Counter compt) throws InterruptedException {

		synchronized(this) {
			while(this.empty()) this.wait();
			if(this.full()) this.notify();
			Paquet p = this.tab[this.begin]; 
			this.sz--; 
			this.begin=(this.begin+1)%this.sizeMax; 
			compt.decrement(); 
			return p; 
		}
	}

	public void enfiler(Paquet p) throws InterruptedException {
		synchronized(this) {
			while(this.full()) this.wait();
			if(this.empty()) this.notify();
			this.tab[(this.begin+this.sz)%this.sizeMax]=p;
			this.sz++; 
		}
	}
}
