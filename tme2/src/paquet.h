/*
 * paquet.h
 *
 *  Created on: Feb 16, 2020
 *      Author: zahrof
 */

#ifndef PAQUET_H_
#define PAQUET_H_
#include "auxiliar_methods.h"

typedef struct{
	char * nom;
}paquet;

void mkpaquet(paquet* p, char * str);

void free_paquet(paquet *p);



#endif /* PAQUET_H_ */
