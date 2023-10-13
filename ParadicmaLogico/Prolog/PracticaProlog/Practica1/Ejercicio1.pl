sumlist([], 0).  % La suma de una lista vac�a es 0.

sumlist([X|Xs], S) :-  % Para una lista [X|Xs], la suma es X m�s la suma de Xs.
    sumlist(Xs, S1),  % Calcula la suma de Xs.
    S is X + S1.      % Calcula S como X m�s la suma de Xs.
