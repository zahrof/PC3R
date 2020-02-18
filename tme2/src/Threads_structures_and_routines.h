/*
 * Threadsstructuresandroutines.h
 *
 *  Created on: Feb 16, 2020
 *      Author: zahrof
 */
#include "Tapis.h"
#include<stdio.h>
typedef struct{
	ft_event_t * fin;
}thFin;

typedef struct{
	int id;
	ft_scheduler_t * sched_pro;
	ft_scheduler_t * sched_cons;
	int * compteur;
	tapis * tapis_production;
	tapis * tapis_consommation;
	FILE * journal_voyage;
}thMessager;

typedef struct{
	int id;
	ft_scheduler_t * sched_pro;
	char * nom_de_produit;
	int cible_production;
	int production_actuel;
	tapis * tapis;
	FILE * journalProducteur;
}thProd;

typedef struct{
	ft_scheduler_t * sched_cons;
	int id;
	int * compteur;
	tapis * tapis;
	ft_event_t * fin;
	FILE * journal_consommateur;
}thCons;



void prodWork(void * args);

void consWork(void * args);

void messWork(void * args);

void fin_routine(void * args);
