# -*- coding: utf-8 -*-
import sys
import yaml
import time
import json

def load_config(config_file):
    with open(config_file, 'r') as f:
        config = yaml.load(f, Loader=yaml.FullLoader)
    return config

#import sys
#import yaml2
#import time
#import json
#def load_config(config_file):
#    f = open(config_file)
#    config = yaml2.load(f, Loader=yaml2.FullLoader)
#    f.close()
#    return config
