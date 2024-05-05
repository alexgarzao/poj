program NestedFors;
var
  i: integer;
  j: integer;
begin
  for i := 1 to 4 do
  begin
    for j := 1 to 4 do
    begin
      writeln ('i=', i, ' j=', j);
    end;
  end;
end.
