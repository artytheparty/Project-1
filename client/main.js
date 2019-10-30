function onLoaded() {
  var source = new EventSource("/sse/serveUpdateddata");
  //its been a minute since ive had to put semicolons
  source.onmessage = function (event) {
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
    document.getElementById("tmem").innerHTML = cpumem["TotalMEM"];
    document.getElementById("umem").innerHTML = cpumem["UsedMEM"];
    document.getElementById("fmem").innerHTML = cpumem["FreeMEM"];
    document.getElementById("cmem").innerHTML = cpumem["CacheMEM"];
    
  };
}
function addRow(tableID){
  let tableRef = document.getElementById(tableID)
  for (var i = 0; i < cpumem["Processes"].length; i++) {
    let newRow = tableRef.insertRow(i+1);
    let newCell1 = newRow.insertCell(0);
    let newCell2 = newRow.insertCell(1);
    let newCell3 = newRow.insertCell(2);
    let newCell4 = newRow.insertCell(3);
    let newCell5 = newRow.insertCell(4);
    let newCell6 = newRow.insertCell(5);
    let newCell7 = newRow.insertCell(6);
    let newCell8 = newRow.insertCell(7);
    let newCell9 = newRow.insertCell(8);
    let newCell10 = newRow.insertCell(9);
    let newCell11 = newRow.insertCell(10);
    let newCell12 = newRow.insertCell(11);
    newCell1 = cpumem["Processes"][i].PID;
    newCell2 = cpumem["Processes"][i].User;
    newCell3 = cpumem["Processes"][k].PR;
    newCell4 = cpumem["Processes"][k].NI;
    newCell5 = cpumem["Processes"][k].VIRT;
    newCell6 = cpumem["Processes"][k].RES;
    newCell7 = cpumem["Processes"][k].SHR;
    newCell8 = cpumem["Processes"][k].S;
    newCell9 = cpumem["Processes"][k].CPU;
    newCell10 = cpumem["Processes"][k].MEM;
    newCell11 = cpumem["Processes"][k].TIME;
    newCell12 = cpumem["Processes"][k].Command;
  }
};
addRow("populateinhere");
