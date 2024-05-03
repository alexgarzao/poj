program global_and_local_vars_3;

var
    myglobal1: integer;
    myglobal2: integer;
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

function mulvalues(value1, value3: integer) : integer;
var
    mylocal1: integer;
    mylocal3: integer;
begin
    mylocal1 := value1 * value3;
    mylocal3 := mylocal1 + addedValue;
    mulvalues := mylocal3;
end;

begin
    addedValue := 5;
    myglobal1 := addvalues(2, 4);
    myglobal2 := mulvalues(6, 8);

    writeln('addedValue=', addedValue);
    writeln('2+4 + addedValue=', myglobal1);
    writeln('6*8 + addedValue=', myglobal2);
end.
