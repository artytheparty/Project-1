function onLoaded(){
    var source = new EventSource("/sse/serveUpdateddata");
    //its been a minute since ive had to put semicolons
    source.onmessage= function(event){
        console.log("OnMessage called: ");
        console.dir(event);
        document.getElementById("counter").innerHTML = event.data;
    }
}