program IfWithStrings;
begin
  {* Simple conditional operator. *}
  write('b>a must be true: ');
  if ('b' > 'a') then
    writeln('true')
  else
    writeln('false');

  write('a>b must be false: ');
  if ('a' > 'b') then
    writeln('true')
  else
    writeln('false');

  {* AND operator. *}
  write('aaa>bbb and bbb>ccc must be false: ');
  if ( 'aaa' > 'bbb') and ( 'bbb' > 'ccc' ) then
    writeln('true')
  else
    writeln('false');

  write('aaa<bbb and bbb<ccc must be true: ');
  if ( 'aaa' < 'bbb') and ( 'bbb' < 'ccc' ) then
    writeln('true')
  else
    writeln('false');

  write('aaa>bbb and bbb<ccc must be false: ');
  if ( 'aaa' > 'bbb') and ( 'bbb' < 'ccc' ) then
    writeln('true')
  else
    writeln('false');

  write('aaa<bbb and bbb>ccc must be false: ');
  if ( 'aaa' < 'bbb') and ( 'bbb' > 'ccc' ) then
    writeln('true')
  else
    writeln('false');
  
  {* OR operator. *}
  write('aaa>bbb or bbb>ccc must be false: ');
  if ( 'aaa' > 'bbb') or ( 'bbb' > 'ccc' ) then
    writeln('true')
  else
    writeln('false');

  write('aaa<bbb or bbb<ccc must be true: ');
  if ( 'aaa' < 'bbb') or ( 'bbb' < 'ccc' ) then
    writeln('true')
  else
    writeln('false');

  write('aaa>bbb or bbb<ccc must be true: ');
  if ( 'aaa' > 'bbb') or ( 'bbb' < 'ccc' ) then
    writeln('true')
  else
    writeln('false');

  write('aaa<bbb or bbb>ccc must be true: ');
  if ( 'aaa' < 'bbb') or ( 'bbb' > 'ccc' ) then
    writeln('true')
  else
    writeln('false');
end.
