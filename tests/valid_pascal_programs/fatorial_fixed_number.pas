program fatorial;

var numero : integer;

function fatorial(n : integer) : integer;
begin
    if n<0 then fatorial := 0
    else begin
        if n<=1 then fatorial := 1
        else fatorial := n * fatorial(n-1);
    end;
end;

begin
    numero := 5;
    writeln('O fatorial de ', numero, ' e: ', fatorial(numero));
end.