      *    This program is designed to plot and optionally zoom the 
      *    Mandlebrot set in COBOL in the terminal window
      *    
      *    To compile, you can use GNU COBOL or Microfocus COBOL
      *
       IDENTIFICATION DIVISION.
           PROGRAM-ID. VANCISE-MANDLEBROT.
           AUTHOR. Nick VanCise.
       ENVIRONMENT DIVISION.

       DATA DIVISION.
       WORKING-STORAGE SECTION.  
           01 LIMITERS.
      *        STATIC LIMITS
               03 SCREEN-X PIC 99 VALUE 24.
               03 SCREEN-Y PIC 99 VALUE 64.
               03 N PIC 999 VALUE 100.
               03 THRESH PIC 999V99 VALUE 2.
           01 COUNTERS.
      *        COUNTER VARIABLES
               03 ITER-COUNT PIC 999 VALUE 0.
               03 X-CNT PIC 99 VALUE 0.
               03 Y-CNT PIC 99 VALUE 0.
           
           01 MANDLE-STORAGE.
      *        STANDARD MANDLEBROT STORAGE FIELDS
               03 X-STORE PIC S9V9(6) VALUE 0.
               03 Y-STORE PIC S9V9(6) VALUE 0.
               03 X-NEXT PIC S9V9(6) VALUE 0.
               03 Y-NEXT PIC S9V9(6) VALUE 0.
               03 X-NSQ PIC S9V9(6) VALUE 0.
               03 Y-NSQ PIC S9V9(6) VALUE 0.
               03 T-SQRT PIC S9V9(6) VALUE 0.
               03 T-TEST-TOT PIC S9V9(6) VALUE 0.
               03 DELTA-X PIC 9V9(6) VALUE 0.
               03 DELTA-Y PIC 9V9(6) VALUE 0.
               03 USER-X-C1 PIC S9V9(6) VALUE 0.
               03 USER-Y-C2 PIC S9V9(6) VALUE 0.
               03 PLOT-CHAR PIC X.
               03 START-POS-X PIC S9V9(6) VALUE -2.0.
               03 START-POS-Y PIC S9V9(6) VALUE 2.0.

           01 USER-VARS.
      *        USER DEFINED RECT FIELDS
               03 USER-X-MIN PIC S9V9(6) VALUE -2.
               03 USER-Y-MIN PIC S9V9(6) VALUE -2.
               03 USER-X-MAX PIC S9V9(6) VALUE 2.
               03 USER-Y-MAX PIC S9V9(6) VALUE 2.
               03 USER-RECT-W  PIC S9V9(6) VALUE 0.
               03 USER-RECT-H  PIC S9V9(6) VALUE 0.
               03 HAS-SET-USER-VALS PIC 9 VALUE 0.
               03 USER-ZOOM PIC X VALUE "N".

       PROCEDURE DIVISION.
      *    MAIN LOGIC
       START-MANDLE.
           PERFORM SETUP.
           DISPLAY " ".
           PERFORM COMPUTE-LOOP
           VARYING X-CNT FROM 0 BY 1 UNTIL X-CNT > SCREEN-X.
           DISPLAY " ".
           DISPLAY "ZOOM AGAIN? (Y/N) ".
           ACCEPT USER-ZOOM.
           IF USER-ZOOM = "Y" OR USER-ZOOM = "y"
              GO TO START-MANDLE
           END-IF.
       END-MANDLE.
           STOP RUN.
               
       SETUP SECTION.
      *    COMPUTE INITIAL SCREEN DELTAS
           DIVIDE 4.0 BY SCREEN-X GIVING DELTA-X.
           DIVIDE 4.0 BY SCREEN-Y GIVING DELTA-Y.

      *    USER START POSITION LOGIC BELOW
           DISPLAY "ENTER YOUR START AND END RECT COORDINATES".
           DISPLAY "ENTER NOTHING FOR DEFAULTS OF -2.0 AND -2.0".
           DISPLAY "ENTER RECT LOWER LHS MIN POINT".
           DISPLAY "ENTER START X: ".
           ACCEPT USER-X-MIN.
           DISPLAY "ENTER START Y: ".
           ACCEPT USER-Y-MIN.
           DISPLAY "ENTER RECT UPPER RHS MAX POINT".
           DISPLAY "ENTER END X: ".
           ACCEPT USER-X-MAX.
           DISPLAY "ENTER END Y: ".
           ACCEPT USER-Y-MAX.

           IF NOT (USER-X-MIN > USER-X-MAX OR 
              USER-Y-MIN > USER-Y-MAX OR
              USER-Y-MIN < -2.0 OR USER-X-MIN < -2.0 OR 
              USER-X-MAX > 2.0 OR USER-Y-MAX > 2.0) AND
              (USER-X-MAX <> 0 OR USER-Y-MAX <> 0 OR USER-X-MIN <> 0 OR 
              USER-Y-MIN <> 0)
               MOVE 1 TO HAS-SET-USER-VALS
           ELSE
               DISPLAY "INVALID COORDINATES. USING DEFAULTS"
               MOVE 2.0 TO USER-X-MAX USER-Y-MAX
               MOVE -2.0 TO USER-X-MIN USER-Y-MIN
               MOVE 0 TO HAS-SET-USER-VALS
           END-IF.

      *    DO POINT AND RECTANGLE CONVERSION IF APPLICABLE
           IF HAS-SET-USER-VALS = 1
               DISPLAY "CALCULATING RECTANGLE FROM USER COORDINATES"
               DISPLAY "("USER-X-MIN","USER-Y-MIN")" "("USER-X-MAX","USE
      -        R-Y-MAX")"
               SUBTRACT USER-X-MIN FROM USER-X-MAX GIVING USER-RECT-W
               SUBTRACT USER-Y-MIN FROM USER-Y-MAX GIVING USER-RECT-H               
      
      *        SET START POSITION
               COMPUTE START-POS-X = (USER-X-MIN / SCREEN-X) * (USER-X-M
      -        AX - USER-X-MIN) + USER-X-MIN
               COMPUTE START-POS-Y = (USER-Y-MIN / (-SCREEN-Y)) * (USER-
      -        Y-MAX - USER-Y-MIN) + USER-Y-MAX

      *        RECALCULATE DELTA X/Y BASED ON NEW RECT
               COMPUTE DELTA-X = (USER-X-MAX - USER-X-MIN) / SCREEN-X
               COMPUTE DELTA-Y = (USER-Y-MAX - USER-Y-MIN) / SCREEN-Y

               DISPLAY DELTA-X ", " DELTA-Y

               MOVE 0 TO HAS-SET-USER-VALS 
           END-IF.

       END-SETUP.
           EXIT.
      *    
       COMPUTE-LOOP SECTION.
      *    THIS IS A DOUBLY NESTED LOOP, TRAVERSING THE 'SCREEN'
      *    ONE BLOCK BY ONE BLOCK
           PERFORM VARYING Y-CNT FROM 0 BY 1 UNTIL Y-CNT > SCREEN-Y
               MOVE SPACES TO PLOT-CHAR
      *        compute user coordinates, user coordinates are b/w 2 and
      *        -2, coordinates are translated using below
               COMPUTE USER-X-C1 = START-POS-X + (X-CNT * DELTA-X)
               COMPUTE USER-Y-C2 = START-POS-Y - (Y-CNT * DELTA-Y)

               MOVE ZERO TO X-STORE Y-STORE
               
               PERFORM THRESH-TEST VARYING ITER-COUNT FROM 0 BY 1
               UNTIL ITER-COUNT IS GREATER THAN N 
               OR PLOT-CHAR IS EQUAL TO '#'

               DISPLAY PLOT-CHAR WITH NO ADVANCING
           END-PERFORM.
           DISPLAY " ".
       END-COMPUTE-LOOP.
           EXIT.

       THRESH-TEST SECTION.
      *    compute next real value
           COMPUTE X-NEXT = (X-STORE**2) - (Y-STORE**2) + USER-X-C1.
      *    compute next imaginary value     
           COMPUTE Y-NEXT = 2.0 * (X-STORE * Y-STORE) + USER-Y-C2. 
           
      *    PERFORM CONVERGENCE TEST   
           COMPUTE T-TEST-TOT = X-NEXT**2 + Y-NEXT**2.
      *    testing convergence using sqrt(sum of squares)  
           MOVE FUNCTION SQRT (T-TEST-TOT) TO T-SQRT.     
           IF T-SQRT > THRESH
      *        diverging here
               MOVE '#' TO PLOT-CHAR
           END-IF.
      *    store values for next compute round
           MOVE X-NEXT TO X-STORE.
           MOVE Y-NEXT TO Y-STORE.

       END-THRESH-TEST.
           EXIT.
