program local_vars;

function addvalues(value1, value2: integer) : integer;
var
    total: integer;
begin
    total := value1 + value2;
    addvalues := total;
end;

begin
    writeln(addvalues(2, 4));
end.
