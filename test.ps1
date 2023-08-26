If (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
  $arguments = "& '" + $myinvocation.mycommand.definition + "'"
  Start-Process powershell -Verb runAs -ArgumentList $arguments
  Break
}

$remoteport = "127.0.0.1"
$found = $remoteport -match '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}';

if ($found) {
  $remoteport = $matches[0];
}
else {
  Write-Output "IP address could not be found";
  exit;
}

$ports=@(80,443,8000,9000,9001,9002,9003,9009,9998,27500,28910,29900,29901,29920 );


for ($i = 0; $i -lt $ports.length; $i++) {
  $port = $ports[$i];
  Invoke-Expression "netsh interface portproxy delete v4tov4 listenport=$port";
  Invoke-Expression "netsh advfirewall firewall delete rule name=$port";

  Invoke-Expression "netsh interface portproxy add v4tov4 listenport=$port connectport=$port connectaddress=$remoteport";
  Invoke-Expression "netsh advfirewall firewall add rule name=$port dir=in action=allow protocol=TCP localport=$port";
}

Invoke-Expression "netsh interface portproxy delete v4tov4 listenport=27900";
Invoke-Expression "netsh advfirewall firewall delete rule name=27900";

Invoke-Expression "netsh interface portproxy add v4tov4 listenport=27900 connectport=27900 connectaddress=$remoteport";
Invoke-Expression "netsh interface portproxy add v4tov4 listenport=27901 connectport=27901 connectaddress=$remoteport";
Invoke-Expression "netsh advfirewall firewall add rule name=27900 dir=in action=allow protocol=UDP localport=27900";
Invoke-Expression "netsh advfirewall firewall add rule name=27901 dir=in action=allow protocol=UDP localport=27901";

Invoke-Expression "netsh interface portproxy show v4tov4";
