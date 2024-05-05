program NestedIfs;
begin
  if (1 > 2) then
    if (2 > 3 ) then
      writeln('1 > 2 and 2 > 3')
    else
      writeln('1 > 2 and 2 <= 3')
  else
    writeln('1 <= 2');

  if (1 < 2) then
    if (2 < 3 ) then
      writeln('1 < 2 and 2 < 3')
    else
      writeln('1 < 2 and 2 >= 3')
  else
    writeln('1 >= 2');

  if (1 < 2) then
    if (2 > 3 ) then
      writeln('1 < 2 and 2 > 3')
    else
      writeln('1 < 2 and 2 <= 3')
  else
    writeln('1 >= 2');

  if (1 > 2) then
    if (2 > 3 ) then
      writeln('1 > 2 and 2 > 3')
    else
      writeln('1 > 2 and 2 <= 3')
  else
    if (1 < 2) then
      if (2 < 3 ) then
        writeln('1 < 2 and 2 < 3')
      else
        writeln('1 < 2 and 2 >= 3');
end.
