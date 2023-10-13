subconj([], _).
subconj([H|T], Set) :-
    member(H, Set),
    subconj(T, Set).
