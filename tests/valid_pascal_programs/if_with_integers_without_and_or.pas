program IfWithInteger;
begin
  {* Check > if then sentence (else true). *}
  write('111>222  must be false: ');
  if ( 111 > 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check > if then sentence (then true). *}
  write('333>222 must be true: ');
  if ( 333 > 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check > if sentence (then false). *}
  writeln('2>4 must be false: ');
  if ( 2 > 4 ) then
    writeln('true');

  {* Check > if sentence (then true). *}
  write('4>2 must be true: ');
  if ( 4 > 2 ) then
    writeln('true');

  {* Check < if then sentence (then true). *}
  write('111<222 must be true: ');
  if ( 111 < 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check < if then sentence (else true). *}
  write('222<111 must be false: ');
  if ( 222 < 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (then true). *}
  write('222>=111 must be true: ');
  if ( 222 >= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (else true). *}
  write('111>=222  must be false: ');
  if ( 111 >= 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (then true). *}
  write('111>=111 must be true: ');
  if ( 111 >= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (then true). *}
  write('111<=222 must be true: ');
  if ( 111 <= 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (else true). *}
  write('222<=111 must be false: ');
  if ( 222 <= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (then true). *}
  write('111<=111 must be true: ');
  if ( 111 <= 111 ) then
    writeln('true')
  else
    writeln('false');

  {* Check = if then sentence (then true). *}
  write('222=222 must be true: ');
  if ( 222 = 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check = if then sentence (else true). *}
  write('111=222 must be false: ');
  if ( 111 = 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <> if then sentence (then true). *}
  write('111<>222 must be true: ');
  if ( 111 <> 222 ) then
    writeln('true')
  else
    writeln('false');

  {* Check <> if then sentence (else true). *}
  write('111<>111 must be false: ');
  if ( 111 <> 111 ) then
    writeln('true')
  else
    writeln('false');
end.
