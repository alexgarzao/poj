program InvalidExample;
begin
  if ( 111 < 'a' ) then
    writeln('true');
  
  if ( 'a' + 2 < 5 ) then
    writeln('true');

  if ( 'a' - 'b' < 'c' ) then
    writeln('true');
end.
