aplanar([], []).
aplanar([H|T], L2) :-
    is_list(H),
    aplanar(H, FlatH),
    aplanar(T, Rest),
    append(FlatH, Rest, L2).
aplanar([H|T], [H|L2]) :-
    \+ is_list(H),
    aplanar(T, L2).


is_list(X) :- is_list(X, []).
is_list([], _).
is_list([_|_], _).
