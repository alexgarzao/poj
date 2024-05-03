program function_call_with_one_param_string;

function myfunction(name: string) : string;
begin
    myfunction := 'Hello ' + name + '!';
end;

begin
    writeln('Hello from main!');
    writeln(myfunction('Alex'));
end.