partir([], _, [], []).
partir([X|Rest], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Rest, Umbral, Menores, Mayores).
partir([X|Rest], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Rest, Umbral, Menores, Mayores).
