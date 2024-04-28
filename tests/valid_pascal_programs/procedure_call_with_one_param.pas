program procedure_call_with_one_param;

procedure myprocedure(name: string);
begin
    writeln('Hello ' + name + '!')
end;

begin
    writeln('Hello from main!');
    myprocedure('Alex');
end.