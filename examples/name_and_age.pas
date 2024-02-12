program NameAndAge;
var
  MyName: String;
  MyAge : Byte;
begin
  Write('What is your name? '); Readln(MyName);
  Write('How old are you? '); Readln(MyAge);
  Writeln;
  Writeln('Hello ', MyName);
  Writeln('You are ', MyAge, ' years old');
end.