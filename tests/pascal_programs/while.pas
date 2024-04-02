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

  i := 20;
  while i <= 10 do
  begin
    writeln ('dont write i=', i);
    i := i + 1;
  end;

end.
