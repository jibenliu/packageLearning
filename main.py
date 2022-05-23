#!/usr/bin/env python3
# encoding: utf-8

import dbus
import dbus.service
from dbus.mainloop.glib import DBusGMainLoop
import gobject

MSG_OBJ_PATH = "/com/example/msg"
MSG_INTERFACE_URI = "com.example.msg"

TIMEFORMAT = "%H:%M:%S"


class Msg(dbus.service.Object):
    def __init__(self,bus,object_path):
        dbus.service.Object.__init__(self,bus,object_path)
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


if __name__ == "__main__":
    DBusGMainLoop(set_as_default=True)
    bus = dbus.SessionBus()
    aMsg = Msg(bus,MSG_OBJ_PATH)
    loop = gobject.MainLoop()
    loop.run()