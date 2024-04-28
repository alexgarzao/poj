program procedure_call_with_two_param;

procedure myprocedure(msg, name: string);
begin
    writeln(msg + ' ' + name + '!')
end;

begin
    writeln('Hello from main!');
    myprocedure('Hello', 'Alex');
end.