program NameAndAge;
var
  myname: string;
  myage: integer;
begin
  write('What is your name? '); readln(myname);
  write('How old are you? '); readln(myage);
  writeln;
  writeln('Hello ', myname);
  writeln('You are ', myage, ' years old');
end.