sub_cadenas(_, [], []).
sub_cadenas(Subcadena, [String|Rest], Filtradas) :-
    sub_string(String, _, _, _, Subcadena),
    sub_cadenas(Subcadena, Rest, FiltradasRest),
    Filtradas = [String|FiltradasRest].
sub_cadenas(Subcadena, [_|Rest], Filtradas) :-
    sub_cadenas(Subcadena, Rest, Filtradas).
sub_cadenas(Subcadena, Lista, Filtradas) :-
    atom(Subcadena),
    is_list(Lista),
    sub_cadenas(Subcadena, Lista, Filtradas).
