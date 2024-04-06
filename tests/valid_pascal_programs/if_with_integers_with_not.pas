program IfWithInteger;
begin
  {* NOT operator. *}
  write('not(111>222) must be true: ');
  if not ( 111 > 222) then
    writeln('true')
  else
    writeln('false');

  {* NOT and AND operator. *}
  write('not(111>222 and 222>333) must be true: ');
  if not (( 111 > 222) and ( 222 > 333 )) then
    writeln('true')
  else
    writeln('false');

  write('not(111<222 and 222<333) must be false: ');
  if not(( 111 < 222) and ( 222 < 333 )) then
    writeln('true')
  else
    writeln('false');

  write('not(111>222 and 222<333) must be true: ');
  if not (( 111 > 222) and ( 222 < 333 )) then
    writeln('true')
  else
    writeln('false');

  write('not(111<222 and 222>333) must be true: ');
  if not(( 111 < 222) and ( 222 > 333 )) then
    writeln('true')
  else
    writeln('false');
  
  {* NOT and OR operator. *}
  write('not(111>222 or 222>333) must be true: ');
  if not(( 111 > 222) or ( 222 > 333 )) then
    writeln('true')
  else
    writeln('false');

  write('not(111<222 or 222<333) must be false: ');
  if not(( 111 < 222) or ( 222 < 333 )) then
    writeln('true')
  else
    writeln('false');

  write('not(111>222 or 222<333) must be false : ');
  if not(( 111 > 222) or ( 222 < 333 )) then
    writeln('true')
  else
    writeln('false');

  write('not(111<222 or 222>333) must be false: ');
  if not(( 111 < 222) or ( 222 > 333 )) then
    writeln('true')
  else
    writeln('false');

  write('(111<222) and not (222>333) must be true: ');
  if ( 111 < 222) and not ( 222 > 333 ) then
    writeln('true')
  else
    writeln('false');

end.
