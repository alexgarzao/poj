program function_call_wo_params;

function myfunction : string;
begin
    myfunction := 'Hello from myfunction!';
end;

begin
    writeln('Hello from main!');
    writeln(myfunction());
end.