program procedure_call_with_mixed_params;

procedure addvalues(msg: string; value1, value2: integer);
begin
    writeln(msg, value1 + value2)
end;

begin
    addvalues('2+4=', 2, 4);
end.