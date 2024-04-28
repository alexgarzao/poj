program procedure_call_wo_params;

procedure myprocedure;
begin
    writeln('Hello from myprocedure!');
end;

begin
    writeln('Hello from main!');
    myprocedure;
end.