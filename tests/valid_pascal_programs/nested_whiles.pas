program NestedWhiles;
var
  i: integer;
  j: integer;
begin
  i := 1;
  while i <= 4 do
  begin
    j := 1;
    while j <= 4 do
    begin
      writeln ('i=', i, ' j=', j);
      j := j + 1;
    end;

    i := i + 1;
  end;
end.
