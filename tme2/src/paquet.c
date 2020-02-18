/*
 * paquet.cpp
 *
 *  Created on: Feb 16, 2020
 *      Author: zahrof
 */

#include "paquet.h"

void mkpaquet(paquet* p, char * str){ p->nom = newcopy(str); }

void free_paquet(paquet *p){
	if (p !=NULL){
		free(p->nom);
		free(p);
	}
}



