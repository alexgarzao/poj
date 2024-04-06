program InvalidExample;
begin
  if ( 111 < 222) and ( 'a' < 'b' ) then
    writeln('true')
  else
    writeln('false');

  if ( 111 < 'a' ) then
    writeln('true');
  
  if ( 'a' + 2 < 5 ) then
    writeln('true');

  if ( 'a' - 'b' < 'c' ) then
    writeln('true');
end.
