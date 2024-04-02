program ForExample;
var
  i: integer;
begin
  for i := 10 to 19 do
  begin
    writeln ('i=', i);
  end;

  for i := 10 to 2 do
  begin
    writeln ('dont write i=', i);
  end;
end.
