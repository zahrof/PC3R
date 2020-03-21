mtype{ROUGE, ORANGE, VERT, INDETERMINEE}; 
chan observe = [0] of {mtype,bool}
active proctype feu (){
    bool clignotant = false;
    mtype couleur = INDETERMINEE 

    initial:
        couleur = ORANGE;
        clignotant = true;
        observe!couleur, clignotant
        if
            ::true -> clignotant = false ; goto red
            ::true -> goto initial
        fi
    
    red:
        
        atomic{couleur = ROUGE;
        observe! couleur, clignotant;} 
        if
            ::true -> goto red;
            ::true -> goto green;
            ::true -> goto fault;
        fi
    green:
        atomic{couleur = VERT;
        observe! couleur, clignotant;} 
        if
            ::true -> goto green;
            ::true -> goto orange;
            ::true -> goto fault;
        fi
    orange:
        atomic{couleur = ORANGE ;
        observe! couleur, clignotant;}
        
        if
            ::true -> goto orange;
            ::true -> goto red;
            ::true -> goto fault;
        fi
    fault:
        clignotant = true 
    fault_loop:
        couleur = ORANGE 
        if
            ::true -> goto fault_loop
        fi
}
active proctype observateur (){
    mtype precedant = INDETERMINEE
    mtype courant = INDETERMINEE
    bool clignote = false
do:: observe ? courant, clignote
    if 
     :: atomic{courant ==ORANGE -> assert(precedant != ROUGE );precedant = ORANGE};
     :: atomic{courant ==VERT -> assert(precedant != ORANGE );precedant = VERT};
     :: atomic{courant ==ROUGE -> assert(precedant != VERT);precedant = ROUGE};
     :: atomic {clignote -> assert (courant == ORANGE )}
    fi
    
    
od  
}



