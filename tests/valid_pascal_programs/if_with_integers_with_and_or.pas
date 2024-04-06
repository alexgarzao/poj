program IfWithInteger;
begin
  {* Simple conditional operator. *}
  write('2>1 must be true: ');
  if (2 > 1) then
    writeln('true')
  else
    writeln('false');

  write('1>2 must be false: ');
  if (1 > 2) then
    writeln('true')
  else
    writeln('false');

  {* AND operator. *}
  write('111>222 and 222>333 must be false: ');
  if ( 111 > 222) and ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 and 222<333 must be true: ');
  if ( 111 < 222) and ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111>222 and 222<333 must be false: ');
  if ( 111 > 222) and ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 and 222>333 must be false: ');
  if ( 111 < 222) and ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');
  
  {* OR operator. *}
  write('111>222 or 222>333 must be false: ');
  if ( 111 > 222) or ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 or 222<333 must be true: ');
  if ( 111 < 222) or ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111>222 or 222<333 must be true: ');
  if ( 111 > 222) or ( 222 < 333 ) then
    writeln('true')
  else
    writeln('false');

  write('111<222 or 222>333 must be true: ');
  if ( 111 < 222) or ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');
end.
