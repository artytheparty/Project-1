function onLoaded() {
    var source = new EventSource("/sse/serveUpdateddata");
    //its been a minute since ive had to put semicolons
    source.onmessage = function(event) {
      console.log("OnMessage called: ");
      console.dir(event);
      var cpuinfo = JSON.parse(event.data)
      var cpuusage = cpuinfo["CPUUSAGE"]
      var sysinfo = cpuinfo["SYSINFO"]
      var lscpu = cpuinfo["LSCPU"]
      var cpumem = cpuinfo["CPUMEM"]
      document.getElementById("cpu").innerHTML = cpuusage["cpumem"]["CPU"];
      document.getElementById("usr").innerHTML = cpuusage["cpumem"]["%usr"];
      document.getElementById("nice").innerHTML = cpuusage["cpumem"]["%nice"];
      document.getElementById("sys").innerHTML = cpuusage["cpumem"]["%sys"];
      document.getElementById("iowait").innerHTML = cpuusage["cpumem"]["%iowait"];
      document.getElementById("irq").innerHTML = cpuusage["cpumem"]["%irq"];
      document.getElementById("soft").innerHTML = cpuusage["cpumem"]["%soft"];
      document.getElementById("steal").innerHTML = cpuusage["cpumem"]["%steal"];
      document.getElementById("guest").innerHTML = cpuusage["cpumem"]["%guest"];
      document.getElementById("gnice").innerHTML = cpuusage["cpumem"]["%gnice"];
      document.getElementById("idle").innerHTML = cpuusage["cpumem"]["%idle"];
      document.getElementById("sysuser").innerHTML = sysinfo["SystemUser"];
      document.getElementById("syskernel").innerHTML = sysinfo["SystemKernel"];
      document.getElementById("syskernelR").innerHTML = sysinfo["SystemKernelRelease"];
      document.getElementById("syskernelV").innerHTML = sysinfo["SystemKernelVersion"];
      document.getElementById("sysarch").innerHTML = sysinfo["SystemArch"];
      document.getElementById("sysproc").innerHTML = sysinfo["SystemProcessor"];
      document.getElementById("syshardw").innerHTML = sysinfo["SystemHardwarePlatform"];
      document.getElementById("sysOS").innerHTML = sysinfo["SystemOS"];
      document.getElementById("arch").innerHTML = lscpu["Architechture"];
      document.getElementById("cpuopmode").innerHTML = lscpu["CPUopmode"];
      document.getElementById("cpus").innerHTML = lscpu["CPUs"];
      document.getElementById("threadspercore").innerHTML = lscpu["ThreadsPerCore"];
      document.getElementById("vendorid").innerHTML = lscpu["VendorID"];
      document.getElementById("modename").innerHTML = lscpu["ModelName"];
      document.getElementById("cpumhz").innerHTML = lscpu["CPUMHz"];
      document.getElementById("cpumaxmhz").innerHTML = lscpu["CPUmaxMHz"];
      document.getElementById("cpuminmhz").innerHTML = lscpu["CPUminMHz"];
      document.getElementById("virt").innerHTML = lscpu["Virtualization"];
    };
  }