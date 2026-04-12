#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from imp import reload
import subprocess
import configparser
from itertools import chain
import re
import os
import sys
import json
import base64
from jinja2 import FileSystemLoader,Environment

def check_sysargs():
    if len(sys.argv) <= 1:
      print("Input Error,expect jinja2 template file")
      exit(1)

def get_template_files():
    configFile=""
    valueFiles=[]
    for arg in sys.argv[1:]:
        if arg.strip() == "config.j2":
            configFile = "config.j2"
            continue
        valueFiles.append(arg)
    
    return (configFile, valueFiles)

def get_dict_from_cmd(command):
    # r=os.popen(command)
    # configStr=json.dumps(r.readlines()[0]) 
    # return json.loads(json.loads(configStr))
    return json.loads(command)

def get_custom_config():
    filename = os.path.join(os.path.abspath('.'), "config")

    parser = configparser.ConfigParser()
    with open(filename) as lines:
        lines = chain(("[top]",), lines)  # This line does the trick.
        parser.readfp(lines)

    items = parser.items()
    for item in items:
        (key, value) = item
        print("key: ", key, " value: ", value)

    # with open("foo.conf") as stream:
    #     stream = StringIO("[top]" + stream.read())  # This line does the trick.
    #     parser.readfp(stream)
    # items = config.defaults().items()

    variables=dict()
            
    return json.loads(json.dumps(variables))

def get_config():
    # configmap_cmd='kubectl get cm env -o jsonpath={".data"}'
    configmapCmd='''
    {"nacos_ip":"10.6.109.201",
     "nacos_port":"30122",
     "nacos_user":"nacos",
     "registry_hostname":"intranet.uihcloud.registry", 
     "registry_port": 5000,
     "k8s_worker_ip":"127.0.0.1",
     "ingress_ip":"10.6.109.251",
     "ingress_port":30122
     }
    '''
    configmap=get_dict_from_cmd(configmapCmd)
    # secret_cmd='kubectl get secret env -o jsonpath={".data"}'
    secretCmd='''{"nacos_passwd":"TkAybzIyYWMwc144Cg=="}'''
    secretmap=get_dict_from_cmd(secretCmd)
    for k,v in secretmap.items():
        secretmap[k]=base64.b64decode(v)[:-1] #delete  end '\n'
    #set all env key,value
    configmap.update(secretmap)
    #add self_def_config
    # config.update(self_def_config)
    return configmap

def render_config(valueFiles, data):
    for file in valueFiles:
       write_file(file, data)

def render_custom_config(configFile, data):
    write_file(configFile, data)


def write_file(templateFile, data):
    templateEnv = Environment(loader=FileSystemLoader(os.path.abspath('.')))
    templateConfig = templateEnv.get_template(templateFile).render(**data)
    (filename, _ ) = os.path.splitext(templateFile)

    with open(filename,'w', encoding='UTF-8') as file:
        file.write(templateConfig)
    
    return templateConfig


def main():
    reload(sys)
   # sys.setdefaultencoding('utf8')
    check_sysargs()

    # get tempalte file
    (configFile, valuesFiles) = get_template_files()
    conf = get_config()

    # render custom config
    render_custom_config(configFile, conf)

    customConf = get_custom_config()

    allConf = json.loads(json.dumps(dict()))
    allConf.update(customConf)
    allConf.update(conf)

    # render internal config
    render_config(valuesFiles, allConf)


if __name__ == "__main__":
    main()