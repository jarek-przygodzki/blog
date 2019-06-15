#!/usr/bin/env python

#
# Requries WMI and pywin32 modules. Install with 
# ```
# pip install wmi
# pip install pywin32
# ````
# 

import wmi, logging

logging.basicConfig(format='%(asctime)s - %(message)s', level=logging.INFO)

c = wmi.WMI()

process_watcher = c.Win32_Process.watch_for("creation")

while True:
    new_process = process_watcher()
    processName = new_process.Caption
    processId   = new_process.ProcessId
    logging.info("Process started. Name: %s | PID: %d", processName, processId)
