using System;
using System.Collections.Generic;
using System.Linq;
using System.IO.Ports;

namespace MyApp
{
    //dotnet run #wont work
    //dotnet build
    //dotnet bin/Debug/net6.0/cs-serial.dll
    public class Program
    {
        public static void Main(string[] args)
        {
            Console.WriteLine("Starting...");
            var sp = new SerialPort();
            sp.PortName = "/dev/ttyUSB0";
            sp.ReadTimeout = 800;
            sp.Open();
            while (true)
            {
                sp.Write(new byte[] { 1, 5, 16, 9, 255, 0, 88, 248 }, 0, 8);
                Read(sp, 8);
                System.Threading.Thread.Sleep(100);
                sp.Write(new byte[] { 1, 5, 16, 9, 0, 0, 25, 8 }, 0, 8);
                Read(sp, 8);
                System.Threading.Thread.Sleep(100);
            }
        }

        private static void Read(SerialPort sp, int nn)
        {
            for (var i = 0; i < nn; i++)
            {
                // var b = sp.ReadByte();
                // System.Console.WriteLine("<{0} {1}", i, b);
            }
        }
    }
}