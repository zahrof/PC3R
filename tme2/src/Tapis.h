/*
 * Tapis.h
 *
 *  Created on: Feb 16, 2020
 *      Author: zahrof
 */

#ifndef TAPIS_H_
#define TAPIS_H_
#include "paquet.h"

typedef struct{
	char * nom;
	paquet ** tab;
	size_t allocsize;
	size_t begin;
	size_t sz;
	ft_event_t * cv;
}tapis;

void mktapis(size_t maxsize, tapis * t, ft_event_t * cv, char * str);

int empty(tapis * t);

int full(tapis * t);

size_t size(tapis * t);

void printTapis(tapis *t);

paquet * defiler(char * str, tapis * t );

void enfiler(char * str, tapis *t, paquet * p);


#endif /* TAPIS_H_ */
