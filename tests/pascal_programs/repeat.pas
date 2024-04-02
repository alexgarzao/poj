program RepeatExample;
var
  i: integer;
begin
  i := 10;
  repeat
    writeln ('i=', i);
    i := i + 1;
  until i = 20;
end.
