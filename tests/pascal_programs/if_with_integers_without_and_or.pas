program SubIntegers;
begin
  {* Check > if then sentence (else true). *}
  write('111>222? ');
  if ( 111 > 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check > if then sentence (then true). *}
  write('333>222? ');
  if ( 333 > 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check > if sentence (then false). *}
  writeln('2>4? ');
  if ( 2 > 4 ) then
    writeln('true');

  {* Check > if sentence (then true). *}
  write('4>2? ');
  if ( 4 > 2 ) then
    writeln('true');

  {* Check < if then sentence (then true). *}
  write('111<222? ');
  if ( 111 < 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check < if then sentence (else true). *}
  write('222<111? ');
  if ( 222 < 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (then true). *}
  write('222>=111? ');
  if ( 222 >= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (else true). *}
  write('111>=222? ');
  if ( 111 >= 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (then true). *}
  write('111>=111? ');
  if ( 111 >= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (then true). *}
  write('111<=222? ');
  if ( 111 <= 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (else true). *}
  write('222<=111? ');
  if ( 222 <= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (then true). *}
  write('111<=111? ');
  if ( 111 <= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check = if then sentence (then true). *}
  write('222=222? ');
  if ( 222 = 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check = if then sentence (else true). *}
  write('111=222? ');
  if ( 111 = 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <> if then sentence (then true). *}
  write('111<>222? ');
  if ( 111 <> 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <> if then sentence (else true). *}
  write('111<>111? ');
  if ( 111 <> 111 ) then
    writeln('true')
  else
    writeln('false');

end.
