program NestedRepeats;
var
  i: integer;
  j: integer;
begin

  i := 1;
  repeat

    j := 1;
    repeat
      writeln ('i=', i, ' j=', j);
      j := j + 1;
    until j >= 5;

    i := i + 1;
  until i >= 5;
end.
