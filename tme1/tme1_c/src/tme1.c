/*
 ============================================================================
 Name        : TME1.c
 Author      : 
 Version     :
 Copyright   : Your copyright notice
 Description : Hello World in C, Ansi-style
 ============================================================================
 */

#include<stdio.h>
#include<string.h>
#include<pthread.h>
#include<stdlib.h>
#include<unistd.h>


char * produits[5]={"pomme", "poire", "orange", "kiwi", "banane"};
int PROD_NB_THREAD=sizeof(produits)/sizeof(char*);

int CONS_NB_THREAD=3;
int TAILLE_TAPIS = 15;
int CIBLE_PRODUCTION=2;

typedef struct{
	char * nom;
}paquet;

typedef struct{
	paquet ** tab;
	size_t allocsize;
	size_t begin;
	size_t sz;
	pthread_mutex_t lock;
	pthread_cond_t  cv_full;
	pthread_cond_t  cv_empty ;
}tapis;

char * newcopy(const char * src){
	size_t n = strlen(src);
	char * dest = malloc((strlen(src)+1)*sizeof(char));
	for(size_t i=0; i <=n; ++i) dest[i]=src[i];
	return dest;
}

void mkpaquet(paquet* p, char * str){
	p->nom = newcopy(str);
}

void free_paquet(paquet *p){
	if (p !=NULL){
	free(p->nom);
	free(p);
	}
}

void mktapis(size_t maxsize, tapis * t){
	t->allocsize= maxsize;
	t->begin=0;
	t->sz=0;

	t->tab=malloc(maxsize*sizeof(paquet));
	if (pthread_mutex_init(&t->lock, NULL) != 0)
		printf("\n mutex init failed\n");

	if (pthread_cond_init(&t->cv_empty, NULL) != 0)
			printf("\n mutex init failed\n");

	if (pthread_cond_init(&t->cv_full, NULL) != 0)
			printf("\n mutex init failed\n");

}

//pas besoin de relocker le mutex car on utilise cette fonction uniquement
//dans enfiler et defiler lesquels ils ont déjà locker le mutex
//dans le cas contraire on devrait effectivement protéger l'accès
int empty(tapis * t){
	return (t->sz==0);
}

int full(tapis * t){
	return (t->sz ==t->allocsize);
}

size_t size(tapis * t){
	pthread_mutex_lock(&t->lock);
	size_t res = t->sz;
	pthread_mutex_unlock(&t->lock);
	return res;
}

void printTapis(tapis *t){
	if(t->sz==0)printf("Le tapis est vide\n");
	if(t->begin==t->sz) printf("%zueme élément du tapis: %s\n",t->begin, t->tab[t->begin]->nom);
	else{
		for (int i=t->begin; i<t->sz; i++)
			printf("%deme élément du tapis: %s\n",i, t->tab[i]->nom);
	}
	printf("\n");
}

char * stringConcat(char * src1, char * src2){
	int lengthSrc1 = strlen(src1);
	int lengthSrc2 = strlen(src2);
	char * res = malloc ((lengthSrc1+lengthSrc2+1)*sizeof(char));
	for (int j=0; j < lengthSrc1; j++)
		res [j] = src1[j];
	for (int j = 0; src2[j] != '\0'; ++j, ++lengthSrc1)
		res [lengthSrc1] = src2[j];
	res[lengthSrc1]='\0';
	return res;
}
int digitsInANumber(int i){
	int count =0;
	while(i != 0)
	{
		count++;
		i /= 10;
	}
	return count;
}

char * intToString(int i){
	char * res = malloc((digitsInANumber(i)+1)*sizeof(char));
	sprintf(res, "%d", i);
	return res;
}

typedef struct{
	char * nom_de_produit;
	int cible_production;
	int production_actuel;
	tapis * tapis;
}thProd;

void delThProd(thProd * t){
	if(t != NULL){
		free(t->nom_de_produit);
		free(t);
	}
}

typedef struct{
	int id;
	int * compteur;
	tapis * tapis;
}thCons;

paquet * defiler(tapis * t, int * compt){
	pthread_mutex_lock(&t->lock);
	while(empty(t)) pthread_cond_wait(&t->cv_empty , &t->lock);
	if(full(t)) pthread_cond_signal(&t->cv_full);
	paquet * p = t->tab[t->begin];
	t->sz--;
	t->begin=(t->begin+1)%t->allocsize;
	(*compt)--;
	pthread_mutex_unlock(&t->lock);
	return p;
}

void enfiler(tapis *t, paquet * p){
	pthread_mutex_lock(&t->lock);
	while(full(t)) pthread_cond_wait(&t->cv_full , &t->lock);
	if(empty(t)) pthread_cond_signal(&t->cv_empty);
	t->tab[(t->begin+t->sz)%t->allocsize]=p;
	t->sz++;
	pthread_mutex_unlock(&t->lock);
}

void* prodWork(void * args){
	thProd * prod = args;

	while (prod->cible_production!=prod->production_actuel){
		paquet * p = malloc(sizeof(paquet));
		mkpaquet(p, stringConcat(prod->nom_de_produit,intToString(prod->production_actuel)));
		enfiler(prod->tapis, p);
		prod->production_actuel++;
	}
	free(prod);
	return 0;
}

void* consWork(void * args){
	thCons * cons = (thCons*)args;

	while((*(cons->compteur))>0){
		paquet * p =defiler(cons->tapis,cons->compteur);
		printf("C%d mange %s\n", cons->id, p->nom);
		free_paquet(p);
	}
	free(cons);
	return 0;
}


int main(void) {

	/*** INITIALISATION DES VARIABLES ***/
	pthread_t tidPROD[PROD_NB_THREAD];
	pthread_t tidCONS[CONS_NB_THREAD];
	int prodCreat =0;
	int consCreat =0;
	int terminaisonsPROD = 0;
	int terminaisonsCONS = 0;
	int err;
	tapis tapis;
	mktapis(TAILLE_TAPIS,&tapis); // free
	int compteur = CIBLE_PRODUCTION*PROD_NB_THREAD;
	pthread_mutex_t  comptLock;
	if (pthread_mutex_init(&comptLock, NULL) != 0)
			printf("\n mutex init failed\n");


	/*** CREATION DES PRODUCTEURS ***/
	while(prodCreat < PROD_NB_THREAD){
		thProd *args = malloc(sizeof (thProd));
		args->nom_de_produit=produits[prodCreat];
		args->production_actuel=0;
		args->cible_production= CIBLE_PRODUCTION;
		args->tapis=&tapis;
		err = pthread_create(&(tidPROD[prodCreat]), NULL, prodWork, args);
		if (err != 0)
			printf("\ncan't create thread :[%s]", strerror(err));
		prodCreat++;
	}

	/*** CREATION DES CONSOMMATEURS ***/
	while(consCreat < CONS_NB_THREAD){
		thCons * args = malloc(sizeof (thCons));
		args->compteur=&compteur;
		args->id=consCreat;
		args->tapis=&tapis;
		err = pthread_create(&(tidCONS[consCreat]), NULL, consWork, args);
		if (err != 0)
			printf("\ncan't create thread :[%s]", strerror(err));
		consCreat++;
	}

	/*** TERMINAISON DES PRODUCTEURS ***/
	while(terminaisonsPROD < PROD_NB_THREAD){
		pthread_join(tidPROD[terminaisonsPROD],NULL);
		terminaisonsPROD++;
	}

	/*** TERMINAISON DES CONSOMMATEURS ***/
	while(terminaisonsCONS < CONS_NB_THREAD){
		pthread_join(tidCONS[terminaisonsCONS],NULL);
		terminaisonsCONS++;
	}

	/*** DESTRUCTION DU MUTEX ***/
	free(tapis.tab);
	pthread_cond_destroy(&tapis.cv_full);
	pthread_mutex_destroy(&tapis.lock);
	return 0;
}
