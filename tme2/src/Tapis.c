/*
 * Tapis.c
 *
 *  Created on: Feb 16, 2020
 *      Author: zahrof
 */

#include "Tapis.h"


void mktapis(size_t maxsize, tapis * t, ft_event_t * cv, char * str){
	t->nom=newcopy(str);
	t->allocsize= maxsize;
	t->begin=0;
	t->sz=0;
	t->cv=cv;
	t->tab=malloc(maxsize*sizeof(paquet));
}

int empty(tapis * t){ return (t->sz==0); }

int full(tapis * t){ return (t->sz ==t->allocsize); }

size_t size(tapis * t){
	size_t res = t->sz;
	return res;
}

paquet * defiler(char * str, tapis * t){
	while(empty(t)){
		ft_thread_await(*t->cv);
		ft_thread_cooperate();
	}
	if(full(t))ft_thread_generate(*t->cv);
	paquet * p = t->tab[t->begin];
	t->sz--;
	t->begin=(t->begin+1)%t->allocsize;
	return p;
}

void enfiler(char * str, tapis *t, paquet * p){
	while(full(t)){
		ft_thread_await(*t->cv);
		ft_thread_cooperate();
	}
	if(empty(t)) ft_scheduler_broadcast(*t->cv);
	t->tab[(t->begin+t->sz)%t->allocsize]=p;
	t->sz++;
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




