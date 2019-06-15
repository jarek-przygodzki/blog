/*
 * Useful to monitor process creation when additional tools cannot be installed.
 *
 * It depends only on .NET framework (csc compiler is part of standard installation, somewhere in C:\Windows\Microsoft.NET\Framework64\)
 * 
 * csc procmon_wmi.cs
 */
 
using System;
using System.Management;

class ProcessMonitor
{
    static public void Main(String[] args)
    {
        var processStartEvent = 
            new ManagementEventWatcher("SELECT * FROM Win32_ProcessStartTrace");
        var processStopEvent = 
            new ManagementEventWatcher("SELECT * FROM Win32_ProcessStopTrace");

        processStartEvent.EventArrived += 
            new EventArrivedEventHandler(
                delegate (object sender, EventArrivedEventArgs e)
        {
            var processName = e.NewEvent.Properties["ProcessName"].Value;
            var processId = e.NewEvent.Properties["ProcessID"].Value;

            Console.WriteLine("{0} Process started. Name: {1} | PID: {2}", 
                DateTime.Now, processName, processId);
        });

        processStopEvent.EventArrived += 
            new EventArrivedEventHandler(
                delegate (object sender, EventArrivedEventArgs e)
        {
            var processName = e.NewEvent.Properties["ProcessName"].Value;
            var processId = e.NewEvent.Properties["ProcessID"].Value;

            Console.WriteLine("{0} Process stopped. Name: {1} | PID: {2}", 
                DateTime.Now, processName, processId);
        });

        processStartEvent.Start();
        processStopEvent.Start();

        Console.ReadKey();
    }
}