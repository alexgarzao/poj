program global_and_local_vars_2;

var
    myglobal1: integer;
    addedValue: integer;

function addvalues(value1, value2: integer) : integer;
var
    mylocal1: integer;
    mylocal2: integer;
begin
    mylocal1 := value1 + value2;
    mylocal2 := mylocal1 + addedValue;
    addvalues := mylocal2;
end;

begin
    addedValue := 5;
    myglobal1 := addvalues(2, 4);
    writeln('addedValue=', addedValue);
    writeln('2+4 + addedValue=', myglobal1);
end.
