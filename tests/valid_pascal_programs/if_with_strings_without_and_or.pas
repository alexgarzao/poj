program IfWithStrings;
begin
  {* Check > if then sentence (else true). *}
  write('aaa>bbb  must be false: ');
  if ( 'aaa' > 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check > if then sentence (then true). *}
  write('ccc>bbb must be true: ');
  if ( 'ccc' > 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check > if sentence (then false). *}
  writeln('b>d must be false: ');
  if ( 'b' > 'd' ) then
    writeln('true');

  {* Check > if sentence (then true). *}
  write('d>b must be true: ');
  if ( 'd' > 'b' ) then
    writeln('true');

  {* Check < if then sentence (then true). *}
  write('aaa<bbb must be true: ');
  if ( 'aaa' < 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check < if then sentence (else true). *}
  write('bbb<aaa must be false: ');
  if ( 'bbb' < 'aaa' ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (then true). *}
  write('bbb>=aaa must be true: ');
  if ( 'bbb' >= 'aaa' ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (else true). *}
  write('aaa>=bbb  must be false: ');
  if ( 'aaa' >= 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check >= if then sentence (then true). *}
  write('aaa>=aaa must be true: ');
  if ( 'aaa' >= 'aaa' ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (then true). *}
  write('aaa<=bbb must be true: ');
  if ( 'aaa' <= 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (else true). *}
  write('bbb<=aaa must be false: ');
  if ( 'bbb' <= 'aaa' ) then
    writeln('true')
  else
    writeln('false');

  {* Check <= if then sentence (then true). *}
  write('aaa<=aaa must be true: ');
  if ( 'aaa' <= 'aaa' ) then
    writeln('true')
  else
    writeln('false');

  {* Check = if then sentence (then true). *}
  write('bbb=bbb must be true: ');
  if ( 'bbb' = 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check = if then sentence (else true). *}
  write('aaa=bbb must be false: ');
  if ( 'aaa' = 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check <> if then sentence (then true). *}
  write('aaa<>bbb must be true: ');
  if ( 'aaa' <> 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* Check <> if then sentence (else true). *}
  write('aaa<>aaa must be false: ');
  if ( 'aaa' <> 'aaa' ) then
    writeln('true')
  else
    writeln('false');
end.
