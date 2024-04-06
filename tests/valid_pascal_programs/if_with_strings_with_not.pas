program IfWithStrings;
begin
  {* NOT operator. *}
  write('not(aaa>bbb) must be true: ');
  if not ( 'aaa' > 'bbb' ) then
    writeln('true')
  else
    writeln('false');

  {* NOT and AND operator. *}
  write('not(aaa>bbb and bbb>ccc) must be true: ');
  if not (( 'aaa' > 'bbb' ) and ( 'bbb' > 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('not(aaa<bbb and bbb<ccc) must be false: ');
  if not(( 'aaa' < 'bbb' ) and ( 'bbb' < 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('not(aaa>bbb and bbb<ccc) must be true: ');
  if not (( 'aaa' > 'bbb' ) and ( 'bbb' < 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('not(aaa<bbb and bbb>ccc) must be true: ');
  if not(( 'aaa' < 'bbb' ) and ( 'bbb' > 'ccc' )) then
    writeln('true')
  else
    writeln('false');
  
  {* NOT and OR operator. *}
  write('not(aaa>bbb or bbb>ccc) must be true: ');
  if not(( 'aaa' > 'bbb' ) or ( 'bbb' > 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('not(aaa<bbb or bbb<ccc) must be false: ');
  if not(( 'aaa' < 'bbb' ) or ( 'bbb' < 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('not(aaa>bbb or bbb<ccc) must be false : ');
  if not(( 'aaa' > 'bbb' ) or ( 'bbb' < 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('not(aaa<bbb or bbb>ccc) must be false: ');
  if not(( 'aaa' < 'bbb' ) or ( 'bbb' > 'ccc' )) then
    writeln('true')
  else
    writeln('false');

  write('(aaa<bbb) and not (bbb>ccc) must be true: ');
  if ( 'aaa' < 'bbb' ) and not ( 'bbb' > 'ccc' ) then
    writeln('true')
  else
    writeln('false');

end.
