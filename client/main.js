function onLoaded() {
    var source = new EventSource("/sse/serveUpdateddata");
    //its been a minute since ive had to put semicolons
    source.onmessage = function(event) {
      console.log("OnMessage called: ");
      console.dir(event);
      var cpuinfo = JSON.parse(event.data)
      var cpuinfo2 = cpuinfo["cpumem"]
      document.getElementById("cpu").innerHTML = cpuinfo2["CPU"];
      document.getElementById("usr").innerHTML = cpuinfo2["%usr"];
      document.getElementById("nice").innerHTML = cpuinfo2["%nice"];
      document.getElementById("sys").innerHTML = cpuinfo2["%sys"];
      document.getElementById("iowait").innerHTML = cpuinfo2["%iowait"];
      document.getElementById("irq").innerHTML = cpuinfo2["%irq"];
      document.getElementById("soft").innerHTML = cpuinfo2["%soft"];
      document.getElementById("steal").innerHTML = cpuinfo2["%steal"];
      document.getElementById("guest").innerHTML = cpuinfo2["%guest"];
      document.getElementById("gnice").innerHTML = cpuinfo2["%gnice"];
      document.getElementById("idle").innerHTML = cpuinfo2["%idle"];
    };
  }