#!/usr/bin/env python3
# encoding: utf-8

"""
@time: 18-4-23 上午11:54
"""
import sys

import dbus
import dbus.service
from PyQt5.QtWidgets import QWidget, QApplication

from dbus.mainloop.pyqt5 import DBusQtMainLoop


class Window(QWidget):
    def __init__(self):
        super().__init__()
        # loop = DBusGMainLoop()

        DBusQtMainLoop(set_as_default=True)
        sessionbus = dbus.SessionBus()
        sessionbus.add_signal_receiver(signal_name="PropertiesChanged",
                                       dbus_interface="org.freedesktop.DBus.Properties",
                                       path="/com/deepin/utcloud/Daemon",
                                       bus_name="com.deepin.utcloud.Daemon",
                                       handler_function=self.ut_login_data)

    def ut_login_data(self, sender, id_, data):
        print("-" * 10, "LoginChanged", "-" * 10)
        print(sender)
        print(id_)
        print(data)


if __name__ == '__main__':
    app = QApplication(sys.argv)
    win = Window()
    win.show()
    app.exec_()
