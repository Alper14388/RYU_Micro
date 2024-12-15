# ryu_forwarder.py
import json
import requests
from ryu.base import app_manager
from ryu.controller import ofp_event
from ryu.controller.handler import MAIN_DISPATCHER
from ryu.controller.handler import set_ev_cls
from ryu.ofproto import ofproto_v1_0

# Go mikroservislerinin adresleri
GO_MICROSERVICES = {
    "packetin": "http://127.0.0.1:8090/packetin",
    "topology": "http://127.0.0.1:8091/topology",
    
}

class Forwarder(app_manager.RyuApp):
    OFP_VERSIONS = [ofproto_v1_0.OFP_VERSION]

    def __init__(self, *args, **kwargs):
        super(Forwarder, self).__init__(*args, **kwargs)

    @set_ev_cls(ofp_event.EventOFPPacketIn, MAIN_DISPATCHER)
    def _packet_in_handler(self, ev):
   
        msg = ev.msg
        datapath = msg.datapath
        dpid = datapath.id
        packet = msg.to_jsondict()
        packet['dpid'] = dpid
        print('Packet ', packet)
        # Packet-In verisi
        """packet_data = {
            "buffer_id": msg.buffer_id,
            "data": msg.data.hex(),
            "in_port": in_port,
            "reason": msg.reason,
            "total_len": msg.total_len,
            "dpid": datapath.id,
            "action": action
        }"""
        print(packet)
        self.forward_to_go("packetin", packet)

    def topology_discovery(self):
     
        self.forward_to_go("topology", {})

    def forward_to_go(self, service, payload):
     
        url = GO_MICROSERVICES.get(service)
        if not url:
            self.logger.error("Service %s not found in GO_MICROSERVICES.", service)
            return

        try:
            response = requests.post(url, json=payload)
            response.raise_for_status()
            self.logger.info("Data forwarded to %s successfully: %s", service, response.status_code)
        except requests.exceptions.RequestException as e:
            self.logger.error("Error forwarding data to %s: %s", service, str(e))
