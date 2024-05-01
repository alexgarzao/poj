program function_call_with_two_params;

function addvalues(value1, value2: integer) : integer;
begin
    addvalues := value1 + value2;
end;

var
    xpto: integer;

begin
    writeln('2+4=', addvalues(2, 4));
end.