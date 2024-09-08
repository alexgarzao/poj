program WhileExample;
var
  i: integer;
begin
  i := 10;
  while i < 20 do
  begin
    writeln ('i=', i);
    i := i + 1;
  end;
end.
