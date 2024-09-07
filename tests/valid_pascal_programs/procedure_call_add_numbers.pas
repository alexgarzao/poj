program procedure_call_add_numbers;

procedure add(value1, value2: integer);
begin
    writeln(value1 + value2);
end;

begin
    writeln('Hello from main!');
    add(4, 6);
end.
