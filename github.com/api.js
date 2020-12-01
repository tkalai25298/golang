var XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;
var request = new XMLHttpRequest();
request.open("GET","http://localhost:8200/health");
request.send();
request.onload = () =>{
    console.log(request);
}
