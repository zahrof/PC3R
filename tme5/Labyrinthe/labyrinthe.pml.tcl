wm title . "scenario"
wm geometry . 320x600+650+100
canvas .c -width 800 -height 800 \
	-scrollregion {0c -1c 30c 100c} \
	-xscrollcommand ".hscroll set" \
	-yscrollcommand ".vscroll set" \
	-bg white -relief raised -bd 2
scrollbar .vscroll -relief sunken  -command ".c yview"
scrollbar .hscroll -relief sunken -orient horiz  -command ".c xview"
pack append . \
	.vscroll {right filly} \
	.hscroll {bottom fillx} \
	.c {top expand fill}
.c yview moveto 0
# ProcLine[1] stays at 0 (Used 0 nobox 0)
.c create rectangle 86 0 198 20 -fill black
# ProcLine[1] stays at 0 (Used 0 nobox 0)
.c create rectangle 84 -2 196 18 -fill ivory
.c create text 140 8 -text "0:labyrinthe"
.c create text 70 32 -fill #eef -text "1"
.c create line 140 32 140 32 -fill #eef -dash {6 4}
.c create line 140 36 140 20 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 0 to 1 (Used 1 nobox 0)
# ProcLine[1] stays at 1 (Used 1 nobox 1)
.c create rectangle 31 22 249 42 -fill white -width 0
.c create text 140 32 -text "Début à la case 0_4. "
.c create text 70 56 -fill #eef -text "3"
.c create line 140 56 140 56 -fill #eef -dash {6 4}
.c create line 140 48 140 44 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 1 to 3 (Used 1 nobox 1)
# ProcLine[1] stays at 3 (Used 1 nobox 1)
.c create rectangle 31 46 249 66 -fill white -width 0
.c create text 140 56 -text "Haut vers la case 1_4. "
.c create text 70 80 -fill #eef -text "5"
.c create line 140 80 140 80 -fill #eef -dash {6 4}
.c create line 140 72 140 68 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 3 to 5 (Used 1 nobox 1)
# ProcLine[1] stays at 5 (Used 1 nobox 1)
.c create rectangle 20 70 260 90 -fill white -width 0
.c create text 140 80 -text "Gauche vers la case 1_3. "
.c create text 70 104 -fill #eef -text "7"
.c create line 140 104 140 104 -fill #eef -dash {6 4}
.c create line 140 96 140 92 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 5 to 7 (Used 1 nobox 1)
# ProcLine[1] stays at 7 (Used 1 nobox 1)
.c create rectangle 20 94 260 114 -fill white -width 0
.c create text 140 104 -text "Gauche vers la case 1_2. "
.c create text 70 128 -fill #eef -text "9"
.c create line 140 128 140 128 -fill #eef -dash {6 4}
.c create line 140 120 140 116 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 7 to 9 (Used 1 nobox 1)
# ProcLine[1] stays at 9 (Used 1 nobox 1)
.c create rectangle 20 118 260 138 -fill white -width 0
.c create text 140 128 -text "Gauche vers la case 1_1. "
.c create text 70 152 -fill #eef -text "11"
.c create line 140 152 140 152 -fill #eef -dash {6 4}
.c create line 140 144 140 140 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 9 to 11 (Used 1 nobox 1)
# ProcLine[1] stays at 11 (Used 1 nobox 1)
.c create rectangle 31 142 249 162 -fill white -width 0
.c create text 140 152 -text "Haut vers la case 2_1. "
.c create text 70 176 -fill #eef -text "13"
.c create line 140 176 140 176 -fill #eef -dash {6 4}
.c create line 140 168 140 164 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 11 to 13 (Used 1 nobox 1)
# ProcLine[1] stays at 13 (Used 1 nobox 1)
.c create rectangle 31 166 249 186 -fill white -width 0
.c create text 140 176 -text "Haut vers la case 3_1. "
.c create text 70 200 -fill #eef -text "15"
.c create line 140 200 140 200 -fill #eef -dash {6 4}
.c create line 140 192 140 188 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 13 to 15 (Used 1 nobox 1)
# ProcLine[1] stays at 15 (Used 1 nobox 1)
.c create rectangle 20 190 260 210 -fill white -width 0
.c create text 140 200 -text "Droite vers la case 3_2. "
.c create text 70 224 -fill #eef -text "17"
.c create line 140 224 140 224 -fill #eef -dash {6 4}
.c create line 140 216 140 212 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 15 to 17 (Used 1 nobox 1)
# ProcLine[1] stays at 17 (Used 1 nobox 1)
.c create rectangle 20 214 260 234 -fill white -width 0
.c create text 140 224 -text "Droite vers la case 3_3. "
.c create text 70 248 -fill #eef -text "19"
.c create line 140 248 140 248 -fill #eef -dash {6 4}
.c create line 140 240 140 236 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 17 to 19 (Used 1 nobox 1)
# ProcLine[1] stays at 19 (Used 1 nobox 1)
.c create rectangle 31 238 249 258 -fill white -width 0
.c create text 140 248 -text "Haut vers la case 4_3. "
.c create text 70 272 -fill #eef -text "21"
.c create line 140 272 140 272 -fill #eef -dash {6 4}
.c create line 140 264 140 260 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 19 to 21 (Used 1 nobox 1)
# ProcLine[1] stays at 21 (Used 1 nobox 1)
.c create rectangle 20 262 260 282 -fill white -width 0
.c create text 140 272 -text "Gauche vers la case 4_2. "
.c create text 70 296 -fill #eef -text "23"
.c create line 140 296 140 296 -fill #eef -dash {6 4}
.c create line 140 288 140 284 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 21 to 23 (Used 1 nobox 1)
# ProcLine[1] stays at 23 (Used 1 nobox 1)
.c create rectangle 20 286 260 306 -fill white -width 0
.c create text 140 296 -text "Gauche vers la case 4_1. "
.c create text 70 320 -fill #eef -text "25"
.c create line 140 320 140 320 -fill #eef -dash {6 4}
.c create line 140 312 140 308 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 23 to 25 (Used 1 nobox 1)
# ProcLine[1] stays at 25 (Used 1 nobox 1)
.c create rectangle 20 310 260 330 -fill white -width 0
.c create text 140 320 -text "Gauche vers la case 4_0. "
.c create text 70 344 -fill #eef -text "27"
.c create line 140 344 140 344 -fill #eef -dash {6 4}
.c create line 140 336 140 332 -fill lightgrey -tags grid -width 1 
.c lower grid
# ProcLine[1] from 25 to 27 (Used 1 nobox 1)
# ProcLine[1] stays at 27 (Used 1 nobox 1)
.c create rectangle 103 334 177 354 -fill white -width 0
.c create text 140 344 -text "Sortie. "
.c lower grid
.c raise mesg
