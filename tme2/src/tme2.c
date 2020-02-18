/*
 ============================================================================
 Name        : TME1.c
 Author      : 
 Version     :
 Copyright   : Your copyright notice
 Description : Hello World in C, Ansi-style
 ============================================================================
 */


#include<string.h>
#include<pthread.h>
#include<stdlib.h>
#include<stdio.h>
#include "Threads_structures_and_routines.h"


char * produits[5]={"pomme", "poire", "orange", "kiwi", "banane"};
int PROD_NB_THREAD=sizeof(produits)/sizeof(char*);
int CONS_NB_THREAD=5;
int MESS_NB_THREAD=5;
int TAILLE_TAPIS = 5;
int CIBLE_PRODUCTION=1;



int main(void) {
	/*** INITIALISATION DES VARIABLES ***/
	ft_scheduler_t sched_pro = ft_scheduler_create();
	ft_scheduler_t sched_cons = ft_scheduler_create();
	ft_scheduler_t sched_mess = ft_scheduler_create();
	ft_event_t  cv_prod = ft_event_create(sched_pro);
	ft_event_t  cv_cons = ft_event_create(sched_cons);
	tapis tapisProd;
	mktapis(TAILLE_TAPIS,&tapisProd, &cv_prod,"prod");
	tapis tapisCons;
	mktapis(TAILLE_TAPIS,&tapisCons,&cv_cons,"cons");
	pthread_mutex_t  comptLock;
	pthread_cond_t  cv_compt ;
	if (pthread_mutex_init(&comptLock, NULL) != 0)
				printf("\n mutex init failed\n");
	if (pthread_cond_init(&cv_compt, NULL) != 0)
				printf("\n mutex init failed\n");

	ft_event_t  fin;
		fin = ft_event_create(sched_cons);

	int prodCreat =0;
	int consCreat =0;
	int messCreat =0;
	int compteur = CIBLE_PRODUCTION*PROD_NB_THREAD;
	FILE * journal_producteur = fopen("journal_producteurs.txt","w");
	FILE * journal_consommateur = fopen("journal_consommateur.txt","w");
	FILE * journal_voyage = fopen("journal_voyage.txt","w") ;


	/*** CREATION DES PRODUCTEURS ***/
	while(prodCreat < PROD_NB_THREAD){
		thProd *args = malloc(sizeof (thProd));
		args->id=prodCreat;
		args->nom_de_produit=produits[prodCreat];
		args->production_actuel=0;
		args->cible_production= CIBLE_PRODUCTION;
		args->tapis=&tapisProd;
		args->sched_pro=&sched_pro;
		args->journalProducteur = journal_producteur;
		ft_thread_create(sched_pro,prodWork,NULL,(void *)args);
		prodCreat++;
	}

	/*** CREATION DES MESSAGERS ***/
	while(messCreat < MESS_NB_THREAD){
		//puts("dans les messagers");
		thMessager * args = malloc(sizeof (thMessager));
		args->compteur=&compteur;
		args->id=messCreat;
		args->tapis_consommation=&tapisCons;
		args->tapis_production=&tapisProd;
		args->sched_pro=&sched_pro;
		args->sched_cons=&sched_cons;
		args->journal_voyage = journal_voyage;
		ft_thread_create(sched_mess,messWork,NULL,(void *)args);
		messCreat++;
	}

	/*** CREATION DES CONSOMMATEURS ***/
	while(consCreat < CONS_NB_THREAD){
		thCons * args = malloc(sizeof (thCons));
		args->sched_cons=&sched_cons;
		args->compteur=&compteur;
		args->id=consCreat;
		args->tapis=&tapisCons;
		args->fin=&fin;
		args->journal_consommateur= journal_consommateur;
		ft_thread_create(sched_cons,consWork,NULL,(void *)args);
		consCreat++;
	}

	thFin * args = malloc(sizeof (thFin));
	args->fin=&fin;
	 ft_thread_t finThread = ft_thread_create(sched_cons, fin_routine, NULL, (void *)args);
	 ft_scheduler_start(sched_pro);
	 ft_scheduler_start(sched_mess);
	 ft_scheduler_start(sched_cons);
	pthread_join(ft_pthread(finThread),NULL);
	 fclose (journal_producteur);
	return 0;
}
