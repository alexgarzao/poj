program IfWithInteger;
begin
  {* Simple conditional operator. *}
  write('2>1? ');
  if (2 > 1) then
    writeln('true')
  else
    writeln('false');

  write('1>2? ');
  if (1 > 2) then
    writeln('true')
  else
    writeln('false');

  {* AND operator. *}
  write('111>222 and 222>333? ');
  if ( 111 > 222) and ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 and 222<333? ');
  if ( 111 < 222) and ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111>222 and 222<333? ');
  if ( 111 > 222) and ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 and 222>333? ');
  if ( 111 < 222) and ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');
  
  {* OR operator. *}
  write('111>222 or 222>333? ');
  if ( 111 > 222) or ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 or 222<333? ');
  if ( 111 < 222) or ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111>222 or 222<333? ');
  if ( 111 > 222) or ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 or 222>333? ');
  if ( 111 < 222) or ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');
end.
