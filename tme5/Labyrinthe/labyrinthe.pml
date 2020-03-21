active proctype labyrinthe (){
   // int x = false;
    //int y = INDETERMINEE 

    initial:
        if 
            ::true -> printf("Début à la case 0_4. "); goto Case_0_4
        fi

    Case_2_0 : 
        if 
            ::true -> printf("Droite vers la case 2_1. "); goto Case_2_1 
        fi

    Case_3_0: 
        if 
            ::true -> printf("Haut vers la case 1_4. "); goto Case_4_0 
        fi

    Case_4_0: 
        if 
            ::true -> printf("Sortie. "); assert false 
            ::true -> printf("Droite vers la case 4_1. "); goto Case_4_1
            ::true -> printf("Bas vers la case 3_0. "); goto Case_3_0 
        fi


    Case_0_1: 
        if 
            ::true -> printf("Haut vers la case 1_1. ");  goto Case_1_1
            ::true -> printf("Droite vers la case 0_2. "); goto Case_0_2 
        fi

    Case_1_1: 
        if 
            ::true -> printf("Haut vers la case 2_1. "); goto Case_2_1
            ::true -> printf("Droite vers la case 1_2. "); goto Case_1_2 
            ::true -> printf("Bas vers la case 0_1. "); goto Case_0_1
        fi

    Case_2_1: 
        if 
            ::true -> printf("Haut vers la case 3_1. "); goto Case_3_1
            ::true -> printf("Gauche vers la case 2_0. "); goto Case_2_0 
            ::true -> printf("Bas vers la case 1_1. "); goto Case_1_1
        fi

    Case_3_1: 
        if 
            ::true -> printf("Droite vers la case 3_2. "); goto Case_3_2
            ::true -> printf("Bas vers la case 2_1. "); goto Case_2_1 
        fi

    Case_4_1: 
        if 
            ::true -> printf("Gauche vers la case 4_0. "); goto Case_4_0
            ::true -> printf("Droite vers la case 4_2. "); goto Case_4_2
        fi


    Case_0_2: 
        if 
            ::true -> printf("Gauche vers la case 0_1. "); goto Case_0_1
            ::true -> printf("Droite vers la case 0_3. ");  goto Case_0_3
        fi

    Case_1_2: 
        if 
            ::true ->printf("Gauche vers la case 1_1. ");  goto Case_1_1
            ::true ->printf("Droite vers la case 1_3. "); goto Case_1_3
        fi

    Case_3_2: 
        if 
            ::true -> printf("Gauche vers la case 3_1. "); goto Case_3_1
            ::true -> printf("Droite vers la case 3_3. "); goto Case_3_3
        fi

    Case_4_2: 
        if 
            ::true -> printf("Gauche vers la case 4_1. "); goto Case_4_1
            ::true -> printf("Droite vers la case 4_3. "); goto Case_4_3
        fi


    Case_0_3: 
        if 
            ::true -> printf("Gauche vers la case 0_2. "); goto Case_0_2
        fi

    Case_1_3: 
        if 
            ::true -> printf("Gauche vers la case 1_2. "); goto Case_1_2
            ::true -> printf("Droite vers la case 1_4. "); goto Case_1_4
        fi

    Case_3_3: 
        if 
            ::true -> printf("Gauche vers la case 3_2. "); goto Case_3_2 
            ::true -> printf("Haut vers la case 4_3. "); goto Case_4_3
        fi

    Case_4_3: 
        if 
            ::true -> printf("Bas vers la case 3_3. "); goto Case_3_3
            ::true -> printf("Gauche vers la case 4_2. "); goto Case_4_2
        fi

    Case_0_4: 
        if 
            ::true -> printf("Haut vers la case 1_4. "); goto Case_1_4 
        fi

    Case_1_4:  
        if 
            ::true -> printf("Bas vers la case 0_4. "); goto Case_0_4
            ::true -> printf("Gauche vers la case 1_3. "); goto Case_1_3 
            ::true -> printf("Haut vers la case 2_4. "); goto Case_2_4
        fi

    Case_2_4:
        if 
            ::true -> printf("Haut vers la case 3_4. "); goto Case_3_4
            ::true -> printf("Bas vers la case 1_4. "); goto Case_1_4 
        fi

    Case_3_4:
        if 
            ::true -> printf("Haut vers la case 4_4. "); goto Case_4_4
            ::true -> printf("Bas vers la case 2_4. "); goto Case_2_4 
        fi

    Case_4_4: 
        if 
            ::true -> printf("Bas vers la case 3_4. "); goto Case_3_4 
        fi
}