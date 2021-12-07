#python -m pip install pyserial
import serial
import time

print("starting...")
s = serial.Serial("/dev/ttyUSB0", 9600, timeout=0.4)
while True:
    s.write([1, 5, 16, 9, 255, 0, 88, 248])
    s.flush()
    time.sleep(0.1)
    n = s.read(8)
    print("n", str(n))
    s.write([1, 5, 16, 9, 0, 0, 25, 8])
    s.flush()
    time.sleep(0.1)
    n = s.read(8)
    print("n", str(n))
