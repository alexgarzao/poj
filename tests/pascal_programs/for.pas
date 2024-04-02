program ForExample;
var
  i: integer;
begin
  for i := 10 to 19 do
  begin
    writeln ('inc i=', i);
  end;

  for i := 10 to 2 do
  begin
    writeln ('dont write i=', i);
  end;

  for i := 10 downto 5 do
  begin
    writeln ('dec i=', i);
  end;
end.
