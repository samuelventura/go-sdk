//https://developercenter.robotstudio.com/api/pcsdk/articles/Manual/Using-the-PC-SDK/Messaging-domain.html#what-can-be-sent-in-a-message
//ControlPanel/Configuration/Controller/Task
//RmqType = Remote
//RmqMode = Synchronous
using ABB.Robotics.Controllers;
using ABB.Robotics.Controllers.Discovery;
using ABB.Robotics.Controllers.Messaging;
using System.Text;


Console.WriteLine("Scanning...");
var s = new NetworkScanner();
s.Scan();
var ci = s.Controllers[0];
Console.WriteLine("Connecting...");
Console.WriteLine("Name {0}", ci.Name);
Console.WriteLine("IP {0}", ci.IPAddress);
Console.WriteLine("Availability {0}", ci.Availability);
var c = new Controller(ci);
var q = c.Ipc.GetQueue("RMQ_T_ROB1");

Action<string> send = str =>
{
    //important NOT to reuse the IpcMessage
    Console.WriteLine("Sending {0}", str);
    var m = new IpcMessage();
    m.SetData(UTF8Encoding.UTF8.GetBytes(str));
    q.Send(m);
};

//third message gets discarded and halts the program
//https://www.youtube.com/watch?v=wPCBtdSS-AA
//corner path failure = corner zone converted to stop point
//use fine instead of zXXX in RAPID MoveAbsJ
while (true)
{
    send("robjoint;[90,0,0,0,0,0]");
    send("robjoint;[0,0,0,0,0,0]");
    Console.ReadKey();
}

c.Dispose();
