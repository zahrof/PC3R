mtype{rouge, orange, vert,INDETERMINEE}
chan observe = [0] of { mytype, bool}

active proctype feu(){
    bool clignotant = false; 
    mytype couleur = INDETERMINEE; 
}

active proctype observateur(){
    mytype precedent = INDETERMINEE
    mytype courant = INDETERMINEE 
    bool clignote = false 
    do::observe ? couleur, clignote -> 
        if :: atomic {couleur==orange -> assert(precedent != vert); precedent = orange};  
        if :: atomic {couleur==rouge  -> assert(precedent != orange); precedent = rouge};
        if :: atomic {couleur==vert -> assert(precedent != rouge); precedent = vert};
}

initial: 
    couleur = orange; 
    clignotant = true; 
    observe ! couleur, clignotant
    if 
        :: true -> clignotant = false; goto red
        :: true -> goto initial 
    fi 

red: 
atomic{
    couleur = red; 
    observe ! couleur, clignotant}
    if 
        :: true -> goto red; 
        :: true -> goto green; 
        :: true -> goto fault; 
    fi 

green: 
atomic{
    couleur = vert; 
    observe ! couleur, clignotant}
    if 
        :: true -> goto green 
        :: true -> goto orange 
        :: true -> goto fault 
    fi

orange: 
atomic {
    couleur = orange; 
    observe ! couleur, clignotant}
    if 
        :: true -> goto orange 
        :: true -> goto red 
        :: true -> goto fault 
    fi 

fault : 
    clignotant = true 

fault_loop :
atomic{
    couleur = orange ; 
    observe ! couleur, clignotant }
    if 
        :: true -> goto fault_loop
    fi 