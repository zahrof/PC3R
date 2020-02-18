#include "Threads_structures_and_routines.h"

#include<stdlib.h>
#include<unistd.h>

void fin_routine(void * args){
	thFin * thFin =args;
  ft_thread_await(*thFin->fin);
}

void prodWork(void * args){
	thProd * prod = args;
	while (prod->cible_production!=prod->production_actuel){
		paquet * p = malloc(sizeof(paquet));
		mkpaquet(p, stringConcat(prod->nom_de_produit,intToString(prod->production_actuel)));
		enfiler("producteur", prod->tapis, p);
		prod->production_actuel++;
		fprintf (prod->journalProducteur, "JOURNAL DE PRODUCTION: enfilage de %s \n", p->nom);
		printf("JOURNAL DE PRODUCTION: enfilage de %s \n", p->nom);

		ft_thread_cooperate();
	}
	free(prod);
}

void consWork(void * args){
	thCons * cons = (thCons*)args;
	while((*(cons->compteur))>0){
		paquet * p =defiler("consommateur", cons->tapis);
		fprintf (cons->journal_consommateur, "JOURNAL DE CONSOMMATION:C%d mange %s\n", cons->id, p->nom);
		printf("JOURNAL DE CONSOMMATION:C%d mange %s\n", cons->id, p->nom);
		(*cons->compteur)--;
		ft_thread_cooperate();
		free_paquet(p);
	}
	free(cons);
	ft_thread_generate(*cons->fin);
}


void messWork(void * args){
	thMessager * messagers = (thMessager*)args;
	ft_scheduler_t sched = ft_thread_scheduler();
	ft_thread_unlink();

	while((*(messagers->compteur))>0){
		ft_thread_link(*messagers->sched_pro);
		paquet * p = defiler("messager ",messagers->tapis_production);
		ft_thread_unlink();
		fprintf (messagers->journal_voyage, "JOURNAL DE VOYAGE:voyage par %d de %s\n", messagers->id, p->nom);
		printf("JOURNAL DE VOYAGE:voyage par %d de %s\n", messagers->id, p->nom);
		ft_thread_link(*messagers->sched_cons);
		enfiler("messager", messagers->tapis_consommation,p);
		ft_thread_unlink();
	}
	ft_thread_link(sched);
	free(messagers);
}



