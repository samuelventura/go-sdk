#https://github.com/robotics/open_abb
#https://github.com/robotics/open_abb/wiki/Python-Control
#enable port 5000 in windows firewall.cpl/advanced
#10.77.3.211
#in server.mod/main
#PERS string ipController:= "0.0.0.0"; !local IP for testing in simulation
#python3 socket send requies str.encode()
#see attached datasheet for arm cartesian ranges

import abb
import time
import logging

logging.basicConfig(level=logging.DEBUG)
R = abb.Robot(ip='127.0.0.1')
R.set_joints([0,0,0,0,0,0])
#in millimeters
#workobject cartesian+quaternion
R.set_cartesian([[200,0,200], [0,0,1,0]])
#R.set_joints([0,0,0,0,-90,0])
time.sleep(1)
R.close()
